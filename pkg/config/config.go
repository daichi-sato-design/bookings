package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig アプリケーション構成を保持する
type AppConfig struct{
	UseCache		bool
	TemplateCache	map[string]*template.Template
	InfoLog			*log.Logger
	InProduction	bool
	Session			*scs.SessionManager
}