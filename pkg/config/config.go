package config

import "html/template"

// AppConfig holds all application configs
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}

var app *AppConfig = &AppConfig{
	UseCache: true,
}

func GetAppConfig() *AppConfig {
	return app
}
