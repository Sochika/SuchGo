package config

import (
	"log"
	"text/template"
)

// Hold the main Application Config
type SystemConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
