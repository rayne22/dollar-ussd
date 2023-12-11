package controllers

import (
	_interface "dollar-ussd/domain/interface"
	"dollar-ussd/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type Session struct {
//	gorm.Model
//	ID    string `gorm:"unique_index;not null"`
//	State string `gorm:"not null"`
//}

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

// IndexController Runs index page
var IndexController = func(c *gin.Context) {

	//sessionID := c.Query("session_id")
	// userInput := c.Query("user_input")

	//scr := _interface.Screen{}
	//test, err := scr.FetchScreens()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	//testData := test[0].Details
	menu := fmt.Sprintf("Welcome to MoMo FX. \n\nEnter the phone number (e.g 260****) to proceed")
	//c.HTML(http.StatusOK, "index", gin.H{
	//	"title": testData,
	// })
	//scr := _interface.Screen{}
	//test, err := scr.FetchScreens()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//testData := test[0].Details
	//	//c.HTML(http.StatusOK, "index", gin.H{
	//	//	"title": testData,
	//	//})

	//// Retrieve or create a session
	//session := getSession(sessionID)
	//
	//// Process user input and update the session state
	//nextState := processInput(session, userInput)
	//
	//// Update the session in the database
	//updateSession(session)

	// Render the HTML template with the updated session state
	c.HTML(http.StatusOK, "index", _interface.InputDetail{
		Screen: "main",
		Type:   "text",
		Title:  menu,
	},
	)

}

//func USSDHandler(c *gin.Context) {
//	sessionID := c.Query("session_id")
//	userInput := c.Query("user_input")
//
//	// Retrieve or create a session
//	session := getSession(sessionID)
//
//	// Process user input and update the session state
//	nextState := processInput(session, userInput)
//
//	// Update the session in the database
//	updateSession(session)
//
//	// Render the HTML template with the updated session state
//	c.HTML(http.StatusOK, "index.gohtml", gin.H{
//		"SessionID": sessionID,
//		"State":     nextState,
//	})
//}

//func getSession(sessionID string) *Session {
//	var session Session
//	configs.GetDB().FirstOrCreate(&session, Session{ID: sessionID})
//	return &session
//}
//
//func processInput(session *Session, userInput string) string {
//	// Add your logic to process user input and determine the next state
//	// In this example, just append user input to the session state
//	session.State += userInput
//	return session.State
//}
//
//func updateSession(session *Session) {
//	configs.GetDB().Save(session)
//}
