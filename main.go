package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/SealTV/handmade-shope/db"
	"github.com/SealTV/handmade-shope/model"
	"github.com/SealTV/handmade-shope/ui"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var m *model.Model

type config struct {
	Host string
	DB   db.MongoConfig
}

func processFlags() *config {
	cfg := &config{}

	flag.StringVar(&cfg.Host, "listen", "localhost:3000", "HTTP listen spec")
	flag.StringVar(&cfg.DB.Host, "host", "localhost:12017", "Mongo serve host")
	flag.StringVar(&cfg.DB.Username, "user", "user", "Mongo user")
	flag.StringVar(&cfg.DB.Password, "password", "", "Mongo password")
	// flag.StringVar(&cfg.Db.ConnectString, "db-connect", "host=/var/run/postgresql dbname=gowebapp sslmode=disable", "DB Connect String")
	// flag.StringVar(&assetsPath, "assets-path", "assets", "Path to assets dir")
	flag.Parse()
	return cfg
}

func initializeRoutes() { // определение роута главной страницы
	router.GET("/", ui.InitIndexPage(m))
}

func main() {
	fmt.Println(os.Args)
	cfg := processFlags()
	fmt.Println(cfg)

	db, err := db.InitMongoDb(cfg.DB)
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
		log.Fatal(err)
	}
	defer db.Close()

	m = model.New(db)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	// Start serving the application
	router.Run(cfg.Host)
}
