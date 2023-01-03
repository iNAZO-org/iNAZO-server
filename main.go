package main

import (
	"karintou8710/iNAZO-server/config"
	"karintou8710/iNAZO-server/database"
	"karintou8710/iNAZO-server/models"
	"karintou8710/iNAZO-server/server"
)

func main() {
	config.Init("local")
	database.Init(&models.GradeDistribution{})

	if err := server.Init(); err != nil {
		panic(err)
	}
}
