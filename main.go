package main

import (
	"context"
	"fmt"
	"halyard.bot/config"
	"halyard.bot/log"
	"halyard.bot/telegram"
)

func main() {
	ctx := context.Background()
	ctx = config.ContextWithConfig(ctx)
	ctx = log.ContextWithLogger(ctx)
	fmt.Println("Config and Logger have been initialized successfully")

	telegram.StartBot(ctx)
}
