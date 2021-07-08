package config

import (
	"html/template"
	"log"
)

type AppConfig struct{
	AppCache map[string]*template.Template
	UseCache bool
	InfoLog *log.Logger
}