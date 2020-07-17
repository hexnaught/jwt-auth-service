package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/autotls"
	"github.com/hexnaught/jwt-auth-service/api"
	"github.com/hexnaught/jwt-auth-service/config"
	"github.com/hexnaught/jwt-auth-service/database"
	"golang.org/x/crypto/acme/autocert"
)

func main() {

	appConfig := config.LoadAppConfig()

	mongoDB := database.NewMongoHelper()
	defer mongoDB.Database.Client().Disconnect(context.TODO())

	router := api.SetUp(mongoDB, appConfig)

	// If we generate a cert manager without issue, we run with autotls
	// otherwise we fall back to running the service on :8080
	certManager, err := setupCertManager()
	if err != nil {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		log.Printf("App running on %s", fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), port))
		router.Run(
			fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), port),
		)
	} else {
		log.Printf("AUTOTLS Configured - Running App with TLS")
		log.Fatal(autotls.RunWithManager(router, certManager))
	}
}

// setupCertManager sets up a cert manager and returns it for use in gin
func setupCertManager() (*autocert.Manager, error) {
	envVarDomains := os.Getenv("TLS_DOMAINS")
	if os.Getenv("TLS_DOMAINS") == "" {
		return &autocert.Manager{}, errors.New("no domains specified when instantiating autocert manager")
	}
	domains := strings.Split(envVarDomains, ",")

	os.Mkdir("./certs", 0700)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domains...),
		Cache:      autocert.DirCache("certs"),
	}

	return &certManager, nil
}
