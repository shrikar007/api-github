package model
 
import (
	"github.com/jinzhu/gorm"

)
 
type Profile struct {
	Id     int       `gorm:"unique;not null;PRIMARY_KEY;AUTO_INCREMENT"`
	Name   string    `json:"name"`

}
type Repository struct {
	Id     int       `gorm:"unique;not null;PRIMARY_KEY;AUTO_INCREMENT"`
	Repo_name string `json:"repo"`
	Repo_url  string  `json:"repo_url"`
}
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Profile{},&Repository{})
	return db
}

