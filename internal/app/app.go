package app

import (
	"log"

	// configuration
	"github.com/BogdanStaziyev/shop-test/config"
)

func Run(conf config.Configuration) {
	// start migrations
	if err := Migrate(conf); err != nil {
		log.Fatal("Unable to apply migrations", err)
	}
}
