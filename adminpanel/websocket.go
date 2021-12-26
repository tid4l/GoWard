package adminpanel

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// Struct to pass data to HTML template for dynamic render on webpage.
var WebData struct {
	RequestCounter int
	OnlineHosts    int
	TotalHosts     int
	LastLogin      string
	RequestHost    string
}

// Writes the updated WebData to the websocket connection.
func WsWriter(ws *websocket.Conn) {
	ticker := time.NewTicker(4 * time.Second)
	defer func() {
		ticker.Stop()
		ws.Close()
	}()
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprint(WebData))); err != nil {
				return
			}
		}
	}
}
