package db

import (
	"fmt"
	"myapp/config"
	"myapp/internal/domain/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB はDBを返します。
func GetDB() *gorm.DB {
	return db
}

func init() {
	c := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Name, c.DB.Params)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	if err := d.AutoMigrate(&entity.User{}); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	db = d
}
