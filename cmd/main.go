package main

import "github.com/ashupednekar/websocketstream/pkg/server"

func main(){
  s := server.NewServer(8000)
  s.Start()
}
