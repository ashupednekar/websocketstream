package stream

import (
	"fmt"

	"github.com/ashupednekar/websocketstream/pkg/brokers"
	"github.com/gorilla/websocket"
)

func RecvServiceMessages(conn *websocket.Conn, broker brokers.Broker, service string, user string){
  ch := make(chan brokers.Message)
  go broker.Consume(fmt.Sprintf("ws.recv.%s.%s", service, user), ch)
  for msg := range(ch){
    conn.WriteMessage(websocket.BinaryMessage, msg.Data)
  }
}

