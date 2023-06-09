package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"

	// Configuration
	"github.com/BogdanStaziyev/shop-test/config"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/v1"
	"github.com/BogdanStaziyev/shop-test/internal/database"
	"github.com/BogdanStaziyev/shop-test/internal/service"

	// External
	"github.com/BogdanStaziyev/shop-test/pkg/httpserver"
	"github.com/BogdanStaziyev/shop-test/pkg/logger"
	"github.com/BogdanStaziyev/shop-test/pkg/passwords"
	"github.com/BogdanStaziyev/shop-test/pkg/postgres"
)

func Run(conf config.Configuration) {
	l := logger.New(conf.LogLevel)

	// Start migrations
	if err := Migrate(conf); err != nil {
		l.Fatal(fmt.Errorf("unable to apply migrations: %s", err))
	}

	pgURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		conf.DatabaseUser, conf.DatabasePassword, conf.DatabaseHost, conf.DatabasePort, conf.DatabaseName,
	)

	// Connect to database
	pg, err := postgres.New(pgURL)
	if err != nil {
		l.Fatal(fmt.Errorf("unable to make postgreSQL connection: %s", err))
	}
	defer pg.Close()

	// Create password generator
	passGen := passwords.NewGeneratePasswordHash(conf.Cost)

	// Databases struct of db
	databases := service.Databases{
		Customer: database.NewCustomerRepo(pg),
	}

	// Services struct of all services
	services := v1.Services{
		Customer: service.NewCustomerService(databases.Customer, passGen),
	}

	// HTTP server start
	handler := chi.NewRouter()
	v1.Router(handler, services, l, conf.AdminName, conf.AdminPassword)
	server := httpserver.New(handler, httpserver.Port(conf.ServerPort))

	// Waiting signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Error("Signal interrupt error: " + s.String())
	case err = <-server.Notify():
		l.Error("Server notify", "err", err)
	}

	// Shutdown server
	err = server.Shutdown()
	if err != nil {
		l.Error("Server shutdown: ", "err", err)
	}
}
