package routes

import (
	"dollar-ussd/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

// Serve keeps gin framework instance
type Serve struct {
	Engine *gin.Engine
}

func (e *Serve) Routes() *Serve {

	// Handle the public endpoints
	//e.Engine.POST("/register", controllers.RegistrationController)
	//e.Engine.POST("/login", controllers.LoginController)

	// Home Route
	home := e.Engine.Group("")
	{
		home.GET("", controllers.HomeController)
	}

	return e
}
