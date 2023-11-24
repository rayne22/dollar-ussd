package controllers

import (
	"dollar-ussd/domain/models/screens"
	"dollar-ussd/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var CreateScreenController = func(c *gin.Context) {

	screen := screens.Screen{}

	err := json.NewDecoder(c.Request.Body).Decode(&screen)
	fmt.Println(screen)
	if err != nil {

		newError := utils.ErrorJson(err.Error(), "Bad Request")
		utils.ErrorHandler(c.Writer, http.StatusBadRequest, err, newError)
		return
	}

	resp, err := screen.CreateScreen()
	if err != nil {

		newError := utils.ErrorJson(err.Error(), "Create MessageSchedule Error")
		utils.ErrorHandler(c.Writer, http.StatusConflict, err, newError)
		return
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s", c.Request.Host, c.Request.RequestURI))
	utils.JsonHandler(c.Writer, http.StatusCreated, resp)
}

var FetchAllScreenController = func(c *gin.Context) {

	screen := screens.Screen{}

	resp, err := screen.GetAllScreens()
	if err != nil {

		newError := utils.ErrorJson(err.Error(), "Fetch Screens Error")
		utils.ErrorHandler(c.Writer, http.StatusConflict, err, newError)
		return
	}

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s", c.Request.Host, c.Request.RequestURI))
	utils.JsonHandler(c.Writer, http.StatusCreated, resp)
}
