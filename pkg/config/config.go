package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConifg holds the application config
type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
