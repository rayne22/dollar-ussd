package controllers

import (
	"dollar-ussd/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HomeController handler
func HomeController(c *gin.Context) {
	// Holds welcome message
	mess := "Welcome to the system built by the GODS. !!!"

	//start := routines.StartRoutine()
	//fmt.Println(start)

	// Creates object
	obj := map[string]interface{}{
		"message": mess,
	}

	// Formats Json Object
	resp := utils.JsonResp(obj)

	// Returns Response
	utils.JsonHandler(c.Writer, http.StatusOK, resp)
}
