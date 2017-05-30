package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SealTV/handmade-shope/daemon"
	"github.com/SealTV/handmade-shope/db"
	"github.com/SealTV/handmade-shope/model"
	"github.com/gin-gonic/gin"
)

var assetsPath string
var router *gin.Engine
var m *model.Model

func processFlags() *daemon.Config {
	cfg := &daemon.Config{}

	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")
	flag.StringVar(&cfg.Db.Host, "host", "localhost:12017", "Mongo serve host")
	flag.StringVar(&cfg.Db.Username, "user", "user", "Mongo user")
	flag.StringVar(&cfg.Db.Password, "password", "", "Mongo password")
	// flag.StringVar(&cfg.Db.ConnectString, "db-connect", "host=/var/run/postgresql dbname=gowebapp sslmode=disable", "DB Connect String")
	flag.StringVar(&assetsPath, "assets-path", "assets", "Path to assets dir")
	flag.Parse()
	return cfg
}

func setupHttpAssets(cfg *daemon.Config) {
	log.Printf("Assets served from %q.", assetsPath)
	cfg.UI.Assets = http.Dir(assetsPath)
}

func initializeRoutes() { // определение роута главной страницы
	router.GET("/", showIndexPage)
}

// Define the route for the index page and display the index.html template
// To start with, we'll use an inline route handler. Later on, we'll create
// standalone functions that will be used as route handlers.
func showIndexPage(c *gin.Context) {
	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Home Page",
		},
	)
}

func main() {
	fmt.Println(os.Args)
	cfg := processFlags()
	fmt.Println(cfg)

	// setupHttpAssets(cfg)
	// if err := daemon.Run(cfg); err != nil {
	// 	log.Printf("Error in main(): %v", err)
	// }

	db, err := db.InitMongoDb(cfg.Db)
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
	router.Run(cfg.ListenSpec)
}
