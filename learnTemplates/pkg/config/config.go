package config

import "html/template"

type AppConfig struct{
	AppCache map[string]*template.Template
}