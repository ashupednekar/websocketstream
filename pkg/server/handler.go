package server

import (
	"log"
	"net/http"

	"github.com/ashupednekar/websocketstream/pkg/brokers"
	"github.com/ashupednekar/websocketstream/pkg/stream"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
}

func HandleWs(w http.ResponseWriter, r *http.Request){
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil{
    log.Fatalf("error upgrading websocket connection\n")
  }
  defer conn.Close()

  log.Println("Client connected")

  broker := brokers.NewBroker()
  stream.RecvClientMessages(conn, broker)
  stream.RecvServiceMessages(conn, broker)

}
