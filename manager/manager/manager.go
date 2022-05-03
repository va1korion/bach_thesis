package manager

import (
	api "bach_thesis/api/go_api"

	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var Coll *mongo.Collection

type GRPCServer struct {
	api.UnimplementedManagerServer
}

func (s *GRPCServer) GetStatus(ctx context.Context, client *api.Client) (*api.Workers, error) {
	var result []bson.D
	var worker *api.Worker
	var workers []*api.Worker
	cursor, err := Coll.Find(context.TODO(), bson.D{})
	if err == mongo.ErrNoDocuments {
		log.Println("Nothing returned")
		return nil, nil
	}

	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Print("hello")
	}
	for _, entry := range result {
		bsonBytes, _ := bson.Marshal(entry)
		err := bson.Unmarshal(bsonBytes, &worker)
		if err != nil {
			return nil, err
		}
		workers = append(workers, worker)
	}
	return &api.Workers{Response: workers}, nil
}

func (s *GRPCServer) Join(ctx context.Context, worker *api.Worker) (*api.JoinResponse, error) {
	name := worker.Uri
	var result bson.M
	err := Coll.FindOne(context.TODO(), bson.D{{"uri", worker.Uri}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		_, err := Coll.InsertOne(context.TODO(), worker)
		if err != nil {
			log.Println(err)
		}
	}
	if err != nil {
		log.Println(err)
	}
	return &api.JoinResponse{Name: name}, nil
}

// weird side effect
func (s *GRPCServer) mustEmbedUnimplementedManagerServer() {
	panic("implement me")
}
