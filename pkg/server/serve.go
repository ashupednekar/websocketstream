package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct{
  Port int
}

func NewServer(port int) *Server {
  return &Server{port}
}

func (s *Server) Start(){
  http.HandleFunc("/ws/", HandleWs)
  log.Printf("listening on port: %d", s.Port)
  if err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil); err != nil{
    log.Fatalf("Error starting server at port %d: %s\n", s.Port, err)
  }
}
