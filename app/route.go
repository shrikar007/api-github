package app

import (
	"github-integration/app/database"
	"github-integration/app/handler"
	"net/http"
)

func (a *App) setRouters(data database.Database) {
	a.Post("/profile",data.GetProfile )
	a.Post("/repositories",data.GetRepo )
}

func (a *App) GetProfile(w http.ResponseWriter, r *http.Request) {

	handler.GetProfile(a.DB,w,r)
}

func (a *App) GetRepo(w http.ResponseWriter, r *http.Request) {
	handler.GetRepo(a.DB,w,r)
}

