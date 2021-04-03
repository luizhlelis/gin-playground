package main

import (
	"log"
	"os"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luizhlelis/gin-playground/controllers"
	"golang.org/x/crypto/acme/autocert"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}

	hostCommonName := os.Getenv("hostCommonName")

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

	manager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(hostCommonName),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(router, &manager))
}
