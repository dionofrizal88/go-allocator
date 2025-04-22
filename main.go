package main

import (
	"github.com/dionofrizal88/go-allocator/config"
	"github.com/dionofrizal88/go-allocator/interfaces/cmd"
	"github.com/dionofrizal88/go-allocator/pkg/db"
	"github.com/dionofrizal88/go-allocator/routes"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
)

// @host localhost:8080
// @schemes http
// main init the go allocator service.
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.GetConfig()

	redisConnection, errRedisConnection := db.NewRedisConnection(conf)
	if errRedisConnection != nil {
		log.Fatalf("Failed to connect to redis: %v", errRedisConnection)
	}

	// Init app
	app := cmd.NewCli()
	app.Action = func(c *cli.Context) error {
		// Init Router
		r := routes.
			NewRouter(
				routes.WithConfig(conf),
				routes.WithRedisDB(redisConnection),
			).
			Init()

		if conf.AppPort == "" {
			conf.AppPort = "8080"
		}

		log.Printf("Starting server on port %s", conf.AppPort)

		err := http.ListenAndServe(":"+conf.AppPort, r)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}

		return nil
	}

	// Init Cli
	cliCommands := cmd.NewCommand(
		conf,
		redisConnection,
	)
	app.Commands = cliCommands
	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("Failed to init CLI: %v", err)
	}
}
