package handlers

import (
	_interface "dollar-ussd/domain/interface"
	"dollar-ussd/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var HandleScreens = func(c *gin.Context) {

	screen := _interface.Screen{}
	//
	selection := _interface.InputDetail{}
	err := json.NewDecoder(c.Request.Body).Decode(&selection)
	if err != nil {
		newError := utils.ErrorJson(err.Error(), "Bad Request")
		utils.ErrorHandler(c.Writer, http.StatusBadRequest, err, newError)
		log.Println(err)
		return
	}

	fullResp, err := screen.SelectionProcessing(selection)
	if err != nil {
		newError := utils.ErrorJson(err.Error(), "Bad Request")
		utils.ErrorHandler(c.Writer, http.StatusBadRequest, err, newError)
		log.Println(err)
		return
	}

	// Render the HTML template with the updated session state
	c.HTML(http.StatusOK, "index", gin.H{
		"title": fullResp.Details,
	})
}
