package main

import (
	"log"
	"net"

	"github.com/wizelineacademy/GoWorkshop/list/handlers"
	"github.com/wizelineacademy/GoWorkshop/proto/list"
	"github.com/wizelineacademy/GoWorkshop/shared"
	"google.golang.org/grpc"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	log.Print("[main] service started")

	shared.Init()

	srv := grpc.NewServer()
	list.RegisterListServer(srv, &handlers.Service{})
	srv.Serve(listener)
}
