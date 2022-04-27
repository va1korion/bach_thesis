package manager

import (
	"bach_thesis/api"
	"context"
	"io"
)

type WorkerInstance struct {
	Uuid   string
	Status string
}

type Workers struct {
	Worker []WorkerInstance
}

type GRPCServer struct {
	api.UnimplementedManagerServer
}

func (s *GRPCServer) GetStatus(ctx context.Context, worker *api.Worker) (*api.Status, error) {

	// No feature was found, return an unnamed feature
	return nil, nil
}

func (s *GRPCServer) StreamTS(stream api.Manager_StreamTSServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := serialize(in.Location)
		// look for notes to be sent to client
		for _, note := range s.routeNotes[key] {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}

func (s *GRPCServer) GetVideo(source *api.Source, server api.Manager_GetVideoServer) error {
	for _, feature := range s.savedFeatures {
		if inRange(feature.Location, rect) {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *GRPCServer) mustEmbedUnimplementedManagerServer() {
	//TODO implement me
	panic("implement me")
}
