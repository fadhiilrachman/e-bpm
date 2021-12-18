package controllers

import (
	"log"

	"github.com/fadhiilrachman/e-bpm/graph/model"
)

func CreateNewUser(user *model.User) error {
	err := dbConnect.Insert(user)
	if err != nil {
		panic(err)
	}

	log.Printf("User created")

	return nil
}