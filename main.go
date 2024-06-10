package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/crypto/app"
	"github.com/ptdrpg/crypto/controller"
	"github.com/ptdrpg/crypto/repository"
	"github.com/ptdrpg/crypto/router"
)

func main() {
	stream := gin.Default()
	app.DBconnexion()
	db := app.DB
	repo := repository.NewRepository(db)
	controller := controller.NewController(db, repo)
	r := router.NewRouter(stream, controller)
	r.RouteConnexion()

	r.R.Run(":4400")
}
