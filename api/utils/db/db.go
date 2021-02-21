package db

import (
	"fmt"
	"os"

	"gamersnote.com/v1/dtos"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBClient() *gorm.DB {
	// 接続
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PORT"),
		"disable",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to create db connection")
	}

	// マイグレーション
	err = db.AutoMigrate(
		&dtos.User{},
		&dtos.Article{},
		&dtos.Comment{},
		&dtos.Like{},
	)
	if err != nil {
		panic("failed to migration")
	}

	return db
}
