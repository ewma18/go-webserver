package config

import (
	"html/template"
)

// AppConfig holds all application configs
type AppConfig struct {
	TemplateCache map[string]*template.Template
	InProduction  bool
}

var app *AppConfig = &AppConfig{
	InProduction: true,
}

func GetAppConfig() *AppConfig {
	return app
}
