package main

// go get -u github.com/gofiber/fiber/v2
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql

import (
	"log"
	"src/bootstrap"
)

func main() {
	app, err := bootstrap.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
