package renders

import (
	"fmt"
	"go-sample-webserver/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderHtmlTemplate(res http.ResponseWriter, fileName string) {

	parsedTemplate, err := getTemplate(fileName)
	if err != nil {
		handleError(res, err)
		return
	}

	err = parsedTemplate.Execute(res, nil)
	if err != nil {
		handleError(res, err)
		return
	}
}

func getTemplate(fileName string) (*template.Template, error) {
	config := config.GetAppConfig()
	if config.UseCache {
		return getCachedTemplate(fileName)
	}

	layouts, err := getLayoutFiles()
	if err != nil {
		return nil, err
	}
	template, err := loadTemplateFromFile("./templates/pages/"+fileName, layouts)
	return template, err

}

func getCachedTemplate(fileName string) (*template.Template, error) {
	config := config.GetAppConfig()
	parsedTemplate := config.TemplateCache[fileName]

	if parsedTemplate == nil {
		return nil, fmt.Errorf("unable to find parsed template in cache for file %s", fileName)
	}

	return parsedTemplate, nil
}

func PreLoadTemplates() (map[string]*template.Template, error) {

	log.Println("Building Templates cache...")

	var templateCache = map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	if len(pages) == 0 {
		return nil, fmt.Errorf("no templates found on folder ./templates/pages")
	}

	layouts, err := getLayoutFiles()
	if err != nil {
		return nil, err
	}

	log.Printf("Found %d pages, %d layouts", len(pages), len(layouts))

	for _, page := range pages {
		fileName := filepath.Base(page)

		template, err := loadTemplateFromFile(page, layouts)
		if err != nil {
			return nil, err
		}
		templateCache[fileName] = template
	}

	log.Println("Building Templates Cache Done!")
	return templateCache, nil
}

func getLayoutFiles() ([]string, error) {
	layouts, err := filepath.Glob("./templates/layouts/*.tmpl")
	return layouts, err
}

func loadTemplateFromFile(file string, layouts []string) (*template.Template, error) {
	fileName := filepath.Base(file)

	log.Printf("Building template for %s ...", fileName)
	template, err := template.New(fileName).ParseFiles(file)

	if err != nil {
		return nil, err
	}

	log.Printf("Building template for %s. Adding layouts...", fileName)
	if len(layouts) > 0 {
		template, err = template.ParseFiles(layouts...)
		if err != nil {
			return nil, err
		}
	}

	log.Printf("Building template for %s Done", fileName)
	return template, nil
}

func handleError(res http.ResponseWriter, err error) {
	res.WriteHeader(500)
	fmt.Fprintf(res, "Error processing template %v", err)
	log.Println("Error processing template ", err)
}
