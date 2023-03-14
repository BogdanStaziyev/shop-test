package main

import (

	// configurations
	"github.com/BogdanStaziyev/shop-test/config"

	// internal
	"github.com/BogdanStaziyev/shop-test/internal/app"
)

func main() {
	// initialize configuration
	conf := config.GetConfiguration()

	// run application
	app.Run(conf)
}
