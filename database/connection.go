package database

import (
	"github.com/StephenDsouza90/StudentApp/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GetConnection : ...
func GetConnection() *gorm.DB {
	var db *gorm.DB
	var err error

	databaseName := "database/StudentDatabase.db"
	db, err = gorm.Open("sqlite3", databaseName)
	utils.CheckErr(err)
	return db
}
