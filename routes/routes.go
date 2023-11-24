package routes

import (
	"dollar-ussd/infrastructure/controllers"
	"dollar-ussd/infrastructure/handlers"
	"github.com/gin-gonic/gin"
)

// Serve keeps gin framework instance
type Serve struct {
	Engine *gin.Engine
}

func (e *Serve) Routes() *Serve {

	// Loads all the templates
	e.Engine.LoadHTMLGlob("simulator/templates/layouts/*")

	// Home Route
	home := e.Engine.Group("")
	{
		home.GET("api", controllers.HomeController)
		home.GET("", controllers.IndexController)
	}

	//
	scrn := e.Engine.Group("screen")
	{
		scrn.POST("", controllers.CreateScreenController)
		scrn.GET("", controllers.FetchAllScreenController)
	}

	req := e.Engine.Group("input")
	{
		req.POST("", handlers.HandleScreens)
		//scrn.GET("", controllers.FetchAllScreenController)
	}

	e.Engine.Static("/assets", "./simulator/static/assets")

	return e
}
