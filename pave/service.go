package pave

import (
	"context"
	"log"

	tb "github.com/tigerbeetledb/tigerbeetle-go"
)

//encore:service
type Service struct {
	db tb.Client
}

func initService() (*Service, error) {
	client, err := tb.NewClient(0, []string{"3001", "3002", "3003"}, 1)
	if err != nil {
		log.Printf("Error creating tigerbeetle client: %s", err)
		return nil, err
	}

	return &Service{
		db: client,
	}, nil
}

func (s *Service) Shutdown(force context.Context) {
	s.db.Close()
}
