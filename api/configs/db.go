package configs

import (
	"fmt"
	"os"

	"gamersnote.com/v1/dtos"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB 初期化された*gorm.DBを返します。初回はコネクションの確率とマイグレーションが行われます。
func GetDB() *gorm.DB {
	if db == nil {
		newDBConnection()
		dbMigration()
	}
	return db
}

// newDBConnection DBコネクションを確立する
func newDBConnection() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PORT"),
		"disable",
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to create db connection")
	}
}

// dbMigration マイグレーションを行う
func dbMigration() {
	err := db.AutoMigrate(
		&dtos.User{},
		&dtos.Article{},
		&dtos.Comment{},
	)
	if err != nil {
		panic("failed to migration")
	}
}
