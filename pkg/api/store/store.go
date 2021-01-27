package store

import (
	"github.com/jinzhu/gorm"
	"github.com/yoneyan/oit_kouzoka/pkg/api/core/task"
	"github.com/yoneyan/oit_kouzoka/pkg/api/tool/config"
	_ "gorm.io/driver/sqlite"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", config.Conf.DBPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	result := db.AutoMigrate(&task.Task{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
