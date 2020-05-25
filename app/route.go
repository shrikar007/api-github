package app

import (
	"github-integration/app/database"
	"github-integration/app/handler"
	"net/http"
)

func (a *App) setRouters(data database.Database) {
	a.Get("/profile/{username}",data.GetProfile )
	a.Get("/repositories/{username}",data.GetRepo )
}

func (a *App) GetProfile(w http.ResponseWriter, r *http.Request) {
	handler.GetProfile(a.DB,w,r)
}

func (a *App) GetRepo(w http.ResponseWriter, r *http.Request) {
	handler.GetRepo(a.DB,w,r)
}

