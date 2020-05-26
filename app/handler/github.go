package handler

import (
	"context"
	"github-integration/app/common"
	"github-integration/app/model"
	"github.com/google/go-github/v31/github"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
	"net/http"
)
 
func GetProfile(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	ctx:=context.Background()
	client := github.NewClient(nil)
	user,_,_ :=client.Users.Get(ctx,vars["username"])
	pro := model.Profile{Name: *user.Name}
	if err := db.Save(&pro).Error; err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		logrus.WithError(err).Error("unable read config file")
		return
	}
	common.RespondJSON(w, http.StatusCreated, pro)
	logrus.WithField("Profile",pro).Print("fetched github user full name and stored in database")
}
 
func GetRepo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	ctx:=context.Background()
	vars:=mux.Vars(r)
	var insertrepo []interface{}
	client := github.NewClient(nil)
	opt := &github.RepositoryListOptions{ }
	repos, _, _ := client.Repositories.List(ctx,vars["username"],opt)
	for _, repo := range repos {
		insertrepo=append(insertrepo,model.Repository{Repo_name: *repo.Name,Repo_url: *repo.CloneURL})
	}
	err := gormbulk.BulkInsert(db, insertrepo, 3000)
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		logrus.WithError(err).Error("unable read config file")
		return
	}
	logrus.WithField("Profile",vars["username"]).Print("fetched all repositories of this profile and stored in database")

}
 
