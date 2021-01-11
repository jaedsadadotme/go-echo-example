package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) Initialize() {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")) //Build connection string

	db, err := gorm.Open("mysql", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	h.DB = db
}
