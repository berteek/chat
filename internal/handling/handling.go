package handling

import (
    "net/http"

    "github.com/berteek/chat/internal/logging"

    ws "github.com/gorilla/websocket"
)

type messageJSON struct {
    Content string `json:"content"`
}

var upgrader = ws.Upgrader{}

func MakeEcho(logger logging.Logger) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        c, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            logger.Errorw("unable to upgrade connection!", "error", err.Error())
            return
        }
        defer c.Close()

        for {
            message := messageJSON{}
            err := c.ReadJSON(&message)
            if err != nil {
                logger.Errorw("unable to read message!", "error", err.Error())
                break
            }

            logger.Infow("new message received.", "message", message)

            err = c.WriteMessage(ws.TextMessage, []byte("message received."))
            if err != nil {
                logger.Errorw("unable to send message!", "error", err.Error())
                break
            }
        }
    }
}
