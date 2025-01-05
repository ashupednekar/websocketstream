package stream

import (
	"log"

	"github.com/ashupednekar/websocketstream/pkg/brokers"
	"github.com/gorilla/websocket"
)

func RecvClientMessages (conn *websocket.Conn, broker brokers.Broker){
  go func(){
    for {
      _, message, err := conn.ReadMessage()
      if err != nil{
        log.Printf("error reading client message: %s\n", err)
        break
      }
      log.Printf("received message from client: %s\n", message)
      broker.Produce("ws.send.svc.user", message)
    }
  }()
}
