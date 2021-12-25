package config

import (
    "log"
    "os"

    "github.com/go-pg/pg"
    "github.com/joho/godotenv"
    controller "github.com/fadhiilrachman/e-bpm/controllers"
)

func Connection() *pg.DB {
    // load env
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    // get envar
    DB_USER := os.Getenv("DB_USER")
    DB_PASS := os.Getenv("DB_PASS")
    DB_ADDR := os.Getenv("DB_ADDR")
    DB_NAME := os.Getenv("DB_NAME")

    option := &pg.Options{
        User:     DB_USER,
        Password: DB_PASS,
        Addr:     DB_ADDR,
        Database: DB_NAME,
    }

    var db *pg.DB = pg.Connect(option)

    if db == nil {
        log.Printf("Database %s not exist", DB_NAME)
    }

    log.Printf("Database %s is connected", DB_NAME)
    controller.CreateUserTable(db)

    return db
}