package serve

import (
	"dollar-ussd/routes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
)

func Serve() {
	//Loads .env file
	//err := godotenvvault.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	f, err := os.OpenFile(""+
		"activity.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")

	// Sets the mode of the Web (Either Release or Debug mode)
	mode, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Println(err)
	}

	if mode == true {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Defines the Gin Gonic Engine
	router := gin.Default()

	router.Use()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// router.Use(cors.Default())

	// Instantiates Serve struct
	route := routes.Serve{}

	route.Engine = router

	_ = route.Routes()

	// Instantiates the Port the Web App runs on
	port := fmt.Sprintf("%s%s", ":", os.Getenv("PORT"))

	//router.LoadHTMLGlob("templates/**")

	// Starts listening and serving
	_ = route.Engine.Run(port)

}
