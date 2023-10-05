package config

import (
	"strings"

	"github.com/gosimple/slug"
	"github.com/joeshaw/envdecode"
	"github.com/samber/do"
)

type AppConfig struct {
	Name  string `env:"APP_NAME,default=dptsi-base-go"`
	Env   string `env:"APP_ENV,default=production"`
	Key   string `env:"APP_KEY"`
	Debug bool   `env:"APP_DEBUG,default=false"`
	URL   string `env:"APP_URL,default=http://localhost"`
}

type CorsConfig struct {
	Paths          []string `env:"CORS_PATHS,default=*"`
	AllowedMethods []string `env:"CORS_ALLOWED_METHODS,default=*"`
	AllowedOrigins []string `env:"CORS_ALLOWED_ORIGINS,default=*"`
	AllowedHeaders []string `env:"CORS_ALLOWED_HEADERS"`
	ExposedHeaders []string `env:"CORS_EXPOSED_HEADERS"`
	MaxAge         int      `env:"CORS_MAX_AGE,default=0"`
	SupportCred    bool     `env:"CORS_SUPPORT_CREDENTIALS,default=false"`
}

type HTTPConfig struct {
	Port   int  `env:"HTTP_PORT,default=8080"`
	Secure bool `env:"HTTP_SECURE,default=false"`
}

type SessionConfig struct {
	Lifetime   int    `env:"SESSION_LIFETIME,default=7200"`
	CookieName string `env:"SESSION_NAME,default=base-go"`
	CookiePath string `env:"SESSION_PATH,default=/"`
	Domain     string `env:"SESSION_DOMAIN,default=localhost"`
	Secure     bool   `env:"SESSION_SECURE_COOKIE,default=false"`

	// Firestore session adapter
	FirestoreProjectID  string `env:"SESSION_FIRESTORE_PROJECT_ID"`
	FirestoreCollection string `env:"SESSION_FIRESTORE_COLLECTION,default=sessions"`
}

type Config interface {
	App() AppConfig
	Cors() CorsConfig
	HTTP() HTTPConfig
	Session() SessionConfig
}

type ConfigImpl struct {
	app     AppConfig
	cors    CorsConfig
	http    HTTPConfig
	session SessionConfig
}

func (c ConfigImpl) App() AppConfig {
	return c.app
}

func (c ConfigImpl) Cors() CorsConfig {
	return c.cors
}

func (c ConfigImpl) HTTP() HTTPConfig {
	return c.http
}

func (c ConfigImpl) Session() SessionConfig {
	return c.session
}

func NewConfig(i *do.Injector) (Config, error) {
	var app AppConfig
	err := envdecode.StrictDecode(&app)
	if err != nil {
		return nil, err
	}

	var cors CorsConfig
	err = envdecode.StrictDecode(&cors)
	if err != nil {
		return nil, err
	}

	var http HTTPConfig
	err = envdecode.StrictDecode(&http)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	var session SessionConfig
	err = envdecode.StrictDecode(&session)
	if err != nil {
		return nil, err
	}
	if session.CookieName == "" || session.CookieName == "base-go" {
		name := slug.Make(app.Name)
		name = strings.ReplaceAll(name, "-", "_") + "_session"
		session.CookieName = name
	}

	return &ConfigImpl{app, cors, http, session}, err
}

func init() {
	do.Provide[Config](do.DefaultInjector, NewConfig)
}
