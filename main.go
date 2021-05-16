package main

import (
	http2 "cs-ut-ee/build-it-project/pkg/handlers"
	"cs-ut-ee/build-it-project/pkg/repositories"
	"cs-ut-ee/build-it-project/pkg/services"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	pgUser   = os.Getenv("POSTGRES_USER")
	pgPass   = os.Getenv("POSTGRES_PASSWORD")
	pgDb     = os.Getenv("POSTGRES_DB")
	pgHost   = os.Getenv("POSTGRES_HOST")
	pgPort   = os.Getenv("POSTGRES_PORT")
	logLevel = os.Getenv("LOG_LEVEL")
	dsn      = ""
)

const (
	httpServicePort = 8080
	//postgreDbName      = "postgres"
	//postgresConnection = "dbname=" + postgreDbName + " host=postgres password=postgres user=postgres sslmode=disable port=5432"
	//dsns               = "postgres:postgres@tcp(postgres:5432)/postgres?charset=utf8mb4&parseTime=True&loc=Local"
	//	postgres:postgres@tcp(postgres:5432)/postgres?charset=utf8mb4&parseTime=True&loc=Local
)

func init() {

	log.Debug("Init called")
	if logLevel == "" {
		logLevel = "debug"
	}
	if pgUser == "" {
		log.Fatal("POSTGRES_USER not set")
	}
	if pgPass == "" {
		log.Fatal("POSTGRES_PASSWORD not set")
	}
	if pgDb == "" {
		log.Fatal("POSTGRES_DB not set")
	}

	if pgHost == "" {
		log.Fatal("POSTGRES_HOST not set")
	}
	if pgPort == "" {
		log.Fatal("POSTGRES_PORT not set")
	}

	dsn = buildDsn()

	log.Debug("DSN is: ", dsn)
}

func main() {

	level, err := log.ParseLevel(logLevel)

	if err != nil {
		panic(err)
	}
	log.SetLevel(level)
	log.Infof("Connecting with postgres %v", dsn)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Couldn't establish db connection with err: %v ", err)
		panic(err)
	}

	log.Info("Conneced with db")
	httpRouter := mux.NewRouter()
	plantHireRepo := repositories.NewPlantHireRepository(dbConn)
	purchaseOrderDriverService := services.NewPurchaseOrderDriverService()
	invoiceDriverService := services.NewInvoiceDriverService()
	plantHireService := services.NewPlantHireService(plantHireRepo, purchaseOrderDriverService)
	purchaseOrderRepo := repositories.NewPurchaseOrderRepository(dbConn)
	purchaseOrderService := services.NewPurchaseOrderService(purchaseOrderRepo, purchaseOrderDriverService)
	invoiceRepo := repositories.NewInvoiceRepository(dbConn)
	invoiceService := services.NewInvoiceService(invoiceRepo, purchaseOrderRepo, invoiceDriverService)

	httpHandler := http2.NewHTTPHandler(plantHireService, purchaseOrderService, invoiceService)
	httpHandler.RegisterRoutes(httpRouter)
	httpHandler.RegisterPORoutes(httpRouter)
	httpHandler.RegisterInvoiceRoutes(httpRouter)
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpServicePort),
		Handler: httpRouter,
	}
	log.Infof("Starting server at port ", httpServicePort)
	err = httpSrv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server")
	}

	log.Info("FINISHED starting server")

}

func buildDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", pgHost, pgUser, pgPass, pgDb, pgPort)
	//return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", pgUser, pgPass, pgUrl, pgDb)
}
