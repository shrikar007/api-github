package model
 
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
 
type Profile struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`

}
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Profile{})
	return db
}
