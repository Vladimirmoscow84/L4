package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"L4.5/internal/handlers"
	"L4.5/internal/service"
	"github.com/wb-go/wbf/config"
	"github.com/wb-go/wbf/ginext"

	_ "net/http/pprof" // "/debug/pprof"
)

func Run() {
	cfg := config.New()
	err := cfg.LoadEnvFiles(".env")
	if err != nil {
		log.Fatalf("[app] error of loading cfg: %v", err)
	}
	cfg.EnableEnv("")
	serverAddr := cfg.GetString("SERVER_ADDRESS")
	if serverAddr == "" {
		log.Fatal("[app] SERVER_ADDRESS is empty")
	}

	srv := service.New()
	log.Println("[app] service initialized successfully")

	engine := ginext.New("debug")
	router, err := handlers.New(engine, srv)
	if err != nil {
		log.Fatalf("[app] failed to create router: %v", err)
	}
	log.Println("[app] router initialized successfully")
	router.Routes()

	// graceful shutdown
	go func() {
		log.Printf("[app] server running on %s", serverAddr)
		err = engine.Run(serverAddr)
		if err != nil {
			log.Fatalf("[app] server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[app] shutting down gracefully")

}
