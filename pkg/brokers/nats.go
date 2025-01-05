package brokers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NatsBroker struct{
  StreamName string
  Stream jetstream.JetStream 
}

func NewNatsBroker(stream string) *NatsBroker{
  nc, err := nats.Connect(os.Getenv("NATS_BROKER_URL"))
  if err != nil{
    log.Fatalf("couldn't connect to nats: %s", err);
  }
  js, _ := jetstream.New(nc)
  ctx := context.Background()
  _, err = js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{Name: stream, Subjects: []string{"ws.>"}, Retention: jetstream.WorkQueuePolicy})
  if err != nil{
    log.Fatalf("error creating/updating stream: %s", err)
  }
  return &NatsBroker{Stream: js, StreamName: stream}
}

func (self *NatsBroker) Produce(subject string, data []byte){
  self.Produce(subject, data) 
}

func (self *NatsBroker) Consume(subject string, ch chan Message){
  ctx := context.Background()
  c, err := self.Stream.CreateOrUpdateConsumer(ctx, self.StreamName, jetstream.ConsumerConfig{
    Durable: strings.ReplaceAll(fmt.Sprintf("%s-consumer", subject), ".", "-"),
  })
  if err != nil{
    log.Fatalf("error creating consumer: %s", err)
  }
  cons, _ := c.Consume(func(msg jetstream.Msg){
    msg.Ack()
    log.Printf("Receved message: %v", msg.Data())
    ch <- Message{
      Subject: subject,
      Data: msg.Data(),
    }
  })
  defer cons.Stop()
}
