package config

import "text/template"

type AppConfig struct {
	Templatecache map[string]*template.Template
	UseCache      bool
}
