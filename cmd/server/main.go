package main

import (

	// Configurations
	"github.com/BogdanStaziyev/shop-test/config"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/app"
)

func main() {
	// Initialize configuration
	conf := config.GetConfiguration()

	// Run application
	app.Run(conf)
}
