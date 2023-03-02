package main

import (
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	logger := log.New(log.WithLogTimestamp(), log.WithLogTimeFormat(time.Kitchen),
		log.WithLogCaller(), log.WithLogPrefix("Baking 🍪 "))
	logger.Info("Starting oven!", "degree", 375)
	time.Sleep(3 * time.Second)
	logger.Info("Finished baking")
}
