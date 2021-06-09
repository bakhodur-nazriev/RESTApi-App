package main

import (
	"restApi/models"
	"restApi/routers"
)

func main() {
	models.ConnectDatabase()
	routers.Routes()
}
