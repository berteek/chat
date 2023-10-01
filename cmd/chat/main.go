package main

import (
    "github.com/berteek/chat/internal/server"
    "github.com/berteek/chat/internal/logging"
)

func main() {
    logger := logging.NewZap()
    defer logger.Flush()
    server.Run(logger)
}
