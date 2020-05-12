package services

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/oddx-team/odd-game-server/config"
	"github.com/oddx-team/odd-game-server/pb"
)

const Version = "1.0.0"

type ReadinessCheck func() bool

func DefaultReadinessCheck() bool {
	return true
}

type service struct {
	isReady        bool
	cfg            *config.Config
	readinessCheck ReadinessCheck
}

func New(config *config.Config) *service {
	return &service{
		isReady:        true,
		cfg:            config,
		readinessCheck: DefaultReadinessCheck,
	}
}

func (s *service) Version(context context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{
		Version: Version,
	}, nil
}

func (s *service) Liveness(context context.Context, req *pb.LivenessRequest) (*pb.LivenessResponse, error) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-osSignal:
		return nil, errors.New("Server is shutting shutdown")
	default:
		return &pb.LivenessResponse{
			Message: "OK",
		}, nil
	}
}

func (s *service) ToggleReadiness(context context.Context, req *pb.ToggleReadinessRequest) (*pb.ToggleReadinessResponse, error) {
	s.isReady = !s.isReady
	return &pb.ToggleReadinessResponse{
		Message: "OK",
	}, nil
}

func (s *service) Readiness(context context.Context, req *pb.ReadinessRequest) (*pb.ReadinessResponse, error) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-osSignal:
		return nil, errors.New("Server is shutting down")
	default:
		if s.readinessCheck() == false {
			return nil, errors.New("Server is not available")
		}

		if s.isReady {
			return &pb.ReadinessResponse{
				Message: "OK",
			}, nil
		}

		return nil, errors.New("Server isn't ready, status: toggle off")
	}
}
