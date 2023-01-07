package summary

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	Port = ":8123"

	messages []*PushRequest
)

type (
	Server struct {
		UnimplementedSummaryServer
	}
)

func (s *Server) Push(ctx context.Context, message *PushRequest) (*PushResponse, error) {
	messages = append(messages, message)

	log.Printf("Received message: %v\n", message)

	return &PushResponse{
		Success: true,
	}, nil
}

func (s *Server) Pull(message *PullRequest, stream Summary_PullServer) error {
	log.Printf("Received message: %v\n", message)
	log.Printf("Responding with: %v\n", fmt.Sprintf("%v", messages))

	for {
		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream has ended")
		default:
			for _, m := range messages {
				if err := stream.SendMsg(m); err != nil {
					return err
				}
			}
		}
	}
}
