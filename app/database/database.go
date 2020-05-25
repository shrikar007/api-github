package database

import "net/http"

type Database interface {
	GetProfile(w http.ResponseWriter, r *http.Request)
	GetRepo(w http.ResponseWriter, r *http.Request)
}