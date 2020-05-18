package env

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func BasePath() string {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		log.Panic("Caller error")
	}
	env := filepath.Dir(b)
	pkg := filepath.Dir(env)
	app := filepath.Dir(pkg)
	return app
}

var port *string

func Port() string {
	if port != nil {
		return *port
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return port
}
