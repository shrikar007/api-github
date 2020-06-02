package handler

import (
	"context"
	"github-integration/app/common"
	"github-integration/app/model"
	"github.com/google/go-github/v31/github"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
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
