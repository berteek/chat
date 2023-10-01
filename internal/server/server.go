package server

import (
    "github.com/berteek/chat/internal/logging"
)

func Run(logger logging.Logger) {
    logger.Info("server starting...")
    logger.Info("server started.")

    logger.Warningw("this is a dummy warning!", "meaning of life", 42)

    logger.Info("server shutting down...")
    logger.Info("server shutdown complete.")
}
