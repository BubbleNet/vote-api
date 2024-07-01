package main

import (
	"context"
	"os"

	"github.com/BubbleNet/vote-api/internal/server"
)

func main() {
	server.Run(context.Background(), os.Getenv)
}
