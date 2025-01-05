package brokers


type Message struct{
  Subject string
  Data []byte
}

type Broker interface{
  Produce(subject string, data []byte);
  Consume(subject string, ch chan Message)
}

func NewBroker() Broker {
  //if os.Getenv("PUBSUB_BROKER") == "nats"{
  return NewNatsBroker("websockets") 
}
