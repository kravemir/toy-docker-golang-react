package main

import (
	"fmt"
)

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type (
	todoModel struct {
		gorm.Model
		Title   string
		Content string

		ListID uint
	}
	listModel struct {
		gorm.Model
		Name string

		Todos []todoModel `gorm:"foreignkey:ListID"`
	}
)

func openAndInitDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=db user=postgres dbname=postgres password=tralalala sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect the database: %v", err))
	}

	db.AutoMigrate(&todoModel{})
	db.AutoMigrate(&listModel{})

	return db
}
