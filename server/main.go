package main

import (
	"bach_thesis/api"
	"bach_thesis/manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {
	creds, _ := credentials.NewServerTLSFromFile(certFile, keyFile)
	s := grpc.NewServer(grpc.Creds(creds))
	srv := &manager.GRPCServer{}
	Workers := make([]manager.WorkerInstance, 1)
	Workers[0] = manager.WorkerInstance{
		Uuid:   "some-cool-uuid",
		Status: "free",
	}
	api.RegisterManagerServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
