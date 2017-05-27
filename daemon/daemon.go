package daemon

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/SealTV/handmade-shope/db"
	"github.com/SealTV/handmade-shope/model"
	"github.com/SealTV/handmade-shope/ui"
)

type Config struct {
	ListenSpec string

	Db db.MongoConfig
	UI ui.Config
}

func Run(cfg *Config) error {
	log.Printf("Starting, HTTP on: %s\n", cfg.ListenSpec)

	db, err := db.InitMongoDb(cfg.Db)
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
		return err
	}
	defer db.Close()

	m := model.New(db)

	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Printf("Error creating listener: %v\n", err)
		return err
	}

	ui.Start(cfg.UI, m, l)

	waitForSignal()
	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch

	log.Printf("Got signal: %v, exiting.", s)
}
