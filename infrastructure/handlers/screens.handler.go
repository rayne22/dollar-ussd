package handlers

import (
	_interface "dollar-ussd/domain/interface"
	"dollar-ussd/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var HandleScreens = func(c *gin.Context) {

	screen := _interface.Screen{}
	//

	selection := _interface.InputDetail{
		Selection: c.PostForm("selection"),
		Screen:    c.PostForm("screen"),
		Type:      c.PostForm("input_type"),
	}

	fullResp, err := screen.SelectionProcessing(selection)
	if err != nil {
		newError := utils.ErrorJson(err.Error(), "Bad Request")
		utils.ErrorHandler(c.Writer, http.StatusBadRequest, err, newError)
		log.Println(err)
		return
	}

	// Render the HTML template with the updated session state
	c.HTML(http.StatusOK, "index", _interface.InputDetail{
		Screen: fullResp.ScreenID,
		Type:   fullResp.InputType,
		Title:  fullResp.Details,
	})
}
