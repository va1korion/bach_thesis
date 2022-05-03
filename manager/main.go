package main

import (
	api "bach_thesis/api/go_api"
	"bach_thesis/manager/manager"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	mongocli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := mongocli.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	manager.Coll = mongocli.Database("thesis").Collection("Workers")

	s := grpc.NewServer()
	srv := &manager.GRPCServer{}
	api.RegisterManagerServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
