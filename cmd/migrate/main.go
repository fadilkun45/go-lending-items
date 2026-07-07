package main

import (
	"loans-item-go/config"
	"log"
)

func main() {
	db := config.OpenConnection()
	config.RunMigration(db)
	log.Println("migration completed")
}
