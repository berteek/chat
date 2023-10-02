package main

import (
    "os"
    "os/signal"
    "syscall"

    "github.com/berteek/chat/internal/serving"
    "github.com/berteek/chat/internal/logging"
)

func main() {
    // create logger
    logger := logging.NewZap()
    // flush logger buffers
    defer logger.Flush()

    // make new server with dependencies
    server := serving.NewServer(logger)
    // handle SIGINT and SIGTERM
    exit := make(chan os.Signal, 1)
    signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-exit
        server.Stop()
        os.Exit(0)
    }()
    // run server
    server.Run()
}
