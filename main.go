package main

import (
	// api adalah directory root project go yang kita buat
    "api/models" // memanggil package models pada directory models
    "api/routes"
)

func main() {

    db := models.SetupDB()
    db.AutoMigrate(&models.Task{})

    r := routes.SetupRoutes(db)
    r.Run()
}