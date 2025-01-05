package brokers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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
  ctx := context.Background()
  self.Stream.Publish(ctx, subject, data)
  defer ctx.Done()
}

func (self *NatsBroker) Consume(subject string, ch chan Message){
  ctx, _ := context.WithTimeout(context.Background(), time.Second * 300)
  defer ctx.Done()
  c, err := self.Stream.CreateOrUpdateConsumer(ctx, self.StreamName, jetstream.ConsumerConfig{
    Durable: strings.ReplaceAll(fmt.Sprintf("%s-consumer", subject), ".", "-"),
    FilterSubject: subject,
  })
  if err != nil{
    log.Fatalf("error creating consumer: %s", err)
  }
  consumer, err := c.Consume(func(msg jetstream.Msg){
    msg.Ack()
    log.Printf("Received message: %v", msg.Data())
    ch <- Message{
      Subject: subject,
      Data: msg.Data(),
    }
  })
  defer consumer.Stop()
  <-ctx.Done()
}
