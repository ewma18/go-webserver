package renders

import (
	"fmt"
	"go-sample-webserver/pkg/config"
	"html/template"
	"io"
	"path/filepath"

	"github.com/gofiber/fiber/v2/log"
)

func RenderHtmlTemplate(writer io.Writer, fileName string) error {

	parsedTemplate, err := getTemplate(fileName)
	if err != nil {
		return err
	}

	return parsedTemplate.Execute(writer, nil)
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
	template, err := loadTemplateFromFile("./resources/templates/pages/"+fileName, layouts)
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

func SetupPageTemplates(config *config.AppConfig) {
	if config.UseCache {
		templateCache, err := preLoadTemplates()
		if err != nil {
			log.Fatal("cannot create template cache", err)
		}
		config.TemplateCache = templateCache
	}
}

func preLoadTemplates() (map[string]*template.Template, error) {

	log.Info("Building Templates cache...")

	var templateCache = map[string]*template.Template{}

	pages, err := filepath.Glob("./resources/templates/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	if len(pages) == 0 {
		return nil, fmt.Errorf("no templates found on folder ./resources/templates/pages")
	}

	layouts, err := getLayoutFiles()
	if err != nil {
		return nil, err
	}

	log.Infof("Found %d pages, %d layouts", len(pages), len(layouts))

	for _, page := range pages {
		fileName := filepath.Base(page)

		template, err := loadTemplateFromFile(page, layouts)
		if err != nil {
			return nil, err
		}
		templateCache[fileName] = template
	}

	log.Info("Building Templates Cache Done!")
	return templateCache, nil
}

func getLayoutFiles() ([]string, error) {
	layouts, err := filepath.Glob("./resources/templates/layouts/*.tmpl")
	return layouts, err
}

func loadTemplateFromFile(file string, layouts []string) (*template.Template, error) {
	fileName := filepath.Base(file)

	log.Infof("Building template for %s ...", fileName)
	template, err := template.New(fileName).ParseFiles(file)

	if err != nil {
		return nil, err
	}

	log.Infof("Building template for %s. Adding layouts...", fileName)
	if len(layouts) > 0 {
		template, err = template.ParseFiles(layouts...)
		if err != nil {
			return nil, err
		}
	}

	log.Infof("Building template for %s Done", fileName)
	return template, nil
}
