package config

import (
	"bookings/pkg/storage/mysql"
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	DB            *mysql.StorageRepository
}
