package config

import (
	"html/template"
	"log"
)

// AppConifg holds the application config
type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
