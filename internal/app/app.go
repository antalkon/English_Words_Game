// internal/app/app.go
package app

import (
	log2 "log"
	"log/slog"

	"github.com/antalkon/English_Words_Game/internal/transport/handler"
	"github.com/antalkon/English_Words_Game/internal/transport/router"
	"github.com/antalkon/English_Words_Game/pkg/config"
	"github.com/antalkon/English_Words_Game/pkg/connectDB"
)

func Main() {
	cfg := config.MustLoad() // get server configuration
	address := cfg.HTTPServer.Address

	// Setup logger
	log := SetupLogger(cfg.Env)
	log.Info("Starting app", slog.String("env", cfg.Env))

	// Initialize database
	connectDB.InitDB()
	connection, err := Db(cfg.Env, "development")
	if err != nil {
		log2.Fatalf("Ошибка при подключении к базе данных %v", err)
	}
	_ = connection
	// Setup GIN
	h := handler.NewHandler()
	r := router.SetupRouter(h)

	if address == "" {
		log2.Fatal("Failed to run server: Address is empty. \n pls. check config files")
	}

	if err := r.Run(address); err != nil {
		log2.Fatalf("Failed to run server: %s", err)
	}
}
