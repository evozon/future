package collector

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

var (
	Address = "localhost:8111"
)

type Server struct {
	UnimplementedCollectorServer

	srv      *grpc.Server
	listener net.Listener

	summary map[string]*SummaryData
}

func NewServer() (*Server, error) {
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		return nil, err
	}
	srv := &Server{
		srv:      grpc.NewServer(),
		listener: lis,
		summary:  make(map[string]*SummaryData),
	}

	RegisterCollectorServer(srv.srv, srv)

	return srv, nil
}

func (s *Server) Start() error {
	return s.srv.Serve(s.listener)
}

func (s *Server) Stop() {
	s.listener.Close()
}

func (s *Server) Push(_ context.Context, req *PushRequest) (*PushResponse, error) {
	s.summary[req.Command] = &SummaryData{
		Output: req.Output,
		Status: req.Status,
	}

	return &PushResponse{}, nil
}

func (s *Server) Summary(_ context.Context, _ *SummaryRequest) (*SummaryResponse, error) {
	return &SummaryResponse{
		Summary: s.summary,
	}, nil
}

func (s *Server) Shutdown(_ context.Context, _ *ShutdownRequest) (*Void, error) {
	s.srv.Stop()
	s.listener.Close()

	return &Void{}, nil
}
