package controllers

import (
	"log"

	"github.com/fadhiilrachman/e-bpm/graph/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)


func CreateUserTable(db *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&model.User{}, options)
	if err != nil {
		panic(err)
	}

	log.Printf("User table created")

	return nil
}

// initiate

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}