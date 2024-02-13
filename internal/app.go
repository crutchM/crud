package internal

import db2 "crud/internal/lib/db"

type App struct {
	config string
}

func (app *App) Run(cfg string) {
	db := db2.DB{}

	var routes interface{}
}
