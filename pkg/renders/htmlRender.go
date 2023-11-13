package renders

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

func RenderHtmlTemplate(res http.ResponseWriter, fileName string) {
	parsedTemplate, err := getOrSetCachedTemplate(fileName)
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

func getOrSetCachedTemplate(fileName string) (*template.Template, error) {
	parsedTemplate := templateCache[fileName]
	var err error

	if parsedTemplate == nil {
		log.Println("Template cache miss for ", fileName)
		log.Println("Creating and parsing template then adding to cache")
		parsedTemplate, err = template.ParseFiles(
			"./templates/"+fileName,
			"./templates/base.layout.tmpl",
		)

		if err != nil {
			log.Println("Error parsing template ", fileName, err)
			return nil, err
		}

		templateCache[fileName] = parsedTemplate
	} else {
		log.Println("Template cache hit for ", fileName)
	}
	return parsedTemplate, nil
}

func handleError(res http.ResponseWriter, err error) {
	res.WriteHeader(500)
	fmt.Fprintf(res, "Error parsing template %v", err)
	log.Println("Error parsing template ", err)
}
