package serving

import (
    "github.com/berteek/chat/internal/logging"
    "github.com/berteek/chat/internal/handling"

    "net/http"

    _ "github.com/gorilla/websocket"
)

type Server struct {
    logger logging.Logger
}

func NewServer(logger logging.Logger) *Server {
    return &Server{logger: logger}
}

func (s *Server) Run() {
    s.startup()

    http.HandleFunc("/echo", handling.MakeEcho(s.logger))
    http.ListenAndServe("localhost:8000", nil)

    s.shutdown()
}

func (s *Server) startup() {
    s.logger.Info("server starting...")
    s.logger.Info("server started.")
}

func (s *Server) shutdown() {
    s.logger.Info("server shutting down...")
    s.logger.Info("server shutdown complete.")
}
