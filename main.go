package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SealTV/handmade-shope/daemon"
)

var assetsPath string

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

func main() {
	fmt.Println(os.Args)
	cfg := processFlags()
	fmt.Println(cfg)

	setupHttpAssets(cfg)
	if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
}
