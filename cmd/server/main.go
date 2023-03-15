package main

import (

	// configurations
	"github.com/BogdanStaziyev/shop-test/config"

	// internal
	"github.com/BogdanStaziyev/shop-test/internal/app"
)

func main() {
	// Initialize configuration
	conf := config.GetConfiguration()

	// Run application
	app.Run(conf)
}
