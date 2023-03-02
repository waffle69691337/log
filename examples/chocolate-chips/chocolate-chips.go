package main

import (
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	logger := log.New(log.WithLogTimestamp(), log.WithLogTimeFormat(time.Kitchen),
		log.WithLogCaller(), log.WithLogPrefix("Baking 🍪 "))

	logger.SetReportTimestamp(false)
	logger.SetReportCaller(false)
	logger.SetLevel(log.DebugLevel)
	logger.Debug("Preparing batch 2...") // DEBUG baking 🍪: Preparing batch 2...}

	batch2 := logger.With("batch", 2, "chocolateChips", true)
	batch2.Debug("Adding chocolate chips")
}
