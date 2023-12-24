package main

import (
	"context"

	"halyard.bot/config"
	"halyard.bot/log"
)

func main() {
	ctx := context.Background()
	ctx = config.ContextWithConfig(ctx)
	ctx = log.ContextWithLogger(ctx)

	telegram.startBot(ctx)

}
