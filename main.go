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
	Host         string
	SQLiteDBName string
	DB           db.MongoConfig
}

func processFlags() *config {
	cfg := &config{}

	flag.StringVar(&cfg.Host, "listen", "localhost:3000", "HTTP listen spec")
	flag.StringVar(&cfg.SQLiteDBName, "sqlite", "storage.sqlite", "SQLite database name")
	flag.StringVar(&cfg.DB.Host, "host", "localhost:12017", "Mongo serve host")
	flag.StringVar(&cfg.DB.Username, "user", "user", "Mongo user")
	flag.StringVar(&cfg.DB.Password, "password", "", "Mongo password")
	flag.Parse()
	return cfg
}

func initializeRoutes() { // определение роута главной страницы
	router.GET("/", ui.InitIndexPage(m))
	router.GET("/products", ui.InitProductsPage(m))
	router.GET("/about", ui.AboutPage)
	router.GET("/contacts", ui.ContactsPage)
}

func main() {
	fmt.Println(os.Args)
	cfg := processFlags()
	fmt.Println(cfg)

	db, err := db.InitSqlite(cfg.SQLiteDBName)
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
