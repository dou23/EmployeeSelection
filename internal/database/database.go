package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var databaseEngine *gorm.DB

func newDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("EmployeeSelection.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database," + err.Error())
	}
	return db
}

func setNewEngine() {
	databaseEngine = newDatabase()
}

func init() {
	if databaseEngine == nil {
		setNewEngine()
	}
}

func GetDatabase() *gorm.DB {
	return databaseEngine
}
