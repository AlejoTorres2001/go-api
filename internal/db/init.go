package db
import (
	"gorm.io/gorm"
	"github.com/AlejoTorres2001/go-gorm-restapi/pkg/models"
)

var Models = [] interface{}{}

func registerModels(models ... interface{}) {
	Models = append(Models, models...)
}

func InitDatabase(db *gorm.DB) {
	registerModels(
		&models.User{},
		&models.Task{},
	)
	db.AutoMigrate(Models...)
}