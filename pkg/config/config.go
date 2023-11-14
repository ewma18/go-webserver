package config

import "html/template"

// AppConfig holds all application configs
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}

var app *AppConfig = new(AppConfig)

func GetAppConfig() *AppConfig {
	return app
}
