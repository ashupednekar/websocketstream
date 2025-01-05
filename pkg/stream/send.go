package stream

import (
	"github.com/ashupednekar/websocketstream/pkg/brokers"
	"github.com/gorilla/websocket"
)

func RecvServiceMessages(conn *websocket.Conn, broker brokers.Broker){
  ch := make(chan brokers.Message)
  go broker.Consume("ws.recv.svc.user", ch)
  for msg := range(ch){
    conn.WriteMessage(websocket.BinaryMessage, msg.Data)
  }
}

