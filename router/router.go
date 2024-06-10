package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/crypto/controller"
)

type Router struct {
	R *gin.Engine
	C *controller.Controller
}

func NewRouter(r *gin.Engine, c *controller.Controller) *Router {
	return &Router{
		R: r,
		C: c,
	}
}

func (r *Router) RouteConnexion() {
	apiR := r.R.Group("api")
	v1 := apiR.Group("v1")

	v1.GET("/", )
}