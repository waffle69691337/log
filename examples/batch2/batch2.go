package main

import (
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	logger := log.New(log.WithLogTimestamp(), log.WithLogTimeFormat(time.Kitchen),
		log.WithLogCaller(), log.WithLogPrefix("baking 🍪 ")).With("batch", 2, "chocolateChips", true)
	logger.SetReportTimestamp(false)
	logger.SetReportCaller(false)
	logger.SetLevel(log.DebugLevel)
	logger.Debug("Preparing batch 2...")
	logger.Debug("Adding chocolate chips")
}
