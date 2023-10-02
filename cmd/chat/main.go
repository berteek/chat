package main

import (
    "github.com/berteek/chat/internal/serving"
    "github.com/berteek/chat/internal/logging"
)

func main() {
    // create logger
    logger := logging.NewZap()
    // flush logger buffers
    defer logger.Flush()

    // run server with dependecies
    server := serving.NewServer(logger)
    server.Run()
}
