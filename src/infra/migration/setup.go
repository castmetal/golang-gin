package main

import (
	"fmt"
	"log"

	_user "golang-gin/src/core/domains/user"
	_infra_db "golang-gin/src/infra/db"
)

func main() {
	runMigrations()
}

func runMigrations() {
	db, err := _infra_db.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Error on Database Connection: %v", err)
	}

	fmt.Println("Migrating database ...")

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	db.AutoMigrate(&_user.User{})

	fmt.Println("Migration completed ...")
}
