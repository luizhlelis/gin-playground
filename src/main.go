package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luizhlelis/gin-playground/controllers"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}

	hostCommonName := os.Getenv("hostCommonName")
	serverPort := os.Getenv("hostCommonName")

	if hostCommonName == "" {
		log.Fatalf("Empty host common name")
	}

	log.Println(hostCommonName)

	router := gin.Default()

	// Routes first version
	v1 := router.Group("/v1")
	{
		v1.POST("/workedHours", controllers.PostWorkedHours)
	}

	log.Fatal(http.ListenAndServe(serverPort, router))

	//manager := autocert.Manager{
	//	Prompt:     autocert.AcceptTOS,
	//	HostPolicy: autocert.HostWhitelist(hostCommonName),
	//	Cache:      autocert.DirCache("/var/www/.cache"),
	//}

	//log.Fatal(autotls.RunWithManager(router, &manager))
}
