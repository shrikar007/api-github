package handler

import (
	"encoding/json"
	"github-integration/app/common"
	"github-integration/app/model"
	"github.com/jinzhu/gorm"
	"net/http"
)
 
func GetProfile(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	pro := model.Profile{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pro); err != nil {
		common.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&pro).Error; err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondJSON(w, http.StatusCreated, pro)

}
 
func GetRepo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	pro := model.Repository{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pro); err != nil {
		common.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&pro).Error; err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondJSON(w, http.StatusCreated, pro)


}
 
