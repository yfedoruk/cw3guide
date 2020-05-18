package main

import "github.com/yfedoruck/cw3guide/pkg/web"

type App struct {
	server *web.Server
}

func (a *App) Init() {
	a.server = web.NewServer()
}

func (a *App) Run() {
	a.server.Start()
}
