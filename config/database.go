package config

import (
	"fmt"
	"loans-item-go/helper"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	cfg, err := LoadConfig()
	helper.PanicIfError(err)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)

	dialect := mysql.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})

	helper.PanicIfError(err)

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(1 * time.Minute)
	sqlDB.SetConnMaxIdleTime(30 * time.Second)

	return db
}
