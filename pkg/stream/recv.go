package stream

import (
	"fmt"
	"log"

	"github.com/ashupednekar/websocketstream/pkg/brokers"
	"github.com/gorilla/websocket"
)

func RecvClientMessages (conn *websocket.Conn, broker brokers.Broker, service string, user string){
  go func(){
    for {
      _, message, err := conn.ReadMessage()
      if err != nil{
        log.Printf("error reading client message: %s\n", err)
        break
      }
      log.Printf("received message from client: %s\n", message)
      log.Printf("producing to: %s", fmt.Sprintf("ws.send.%s.%s", service, user))
      broker.Produce(fmt.Sprintf("ws.send.%s.%s", service, user), message)
    }
  }()
}
