package pave

import (
	"context"
	"fmt"

	"encore.dev/rlog"

	"encore.app/internal/storage"
)

//encore:api public path=/post/:pendingID
func (s *Service) Post(ctx context.Context, pendingID string) (*Response, error) {
	log := rlog.With(
		"pendingID", pendingID,
	)

	err := storage.Present(ctx, pendingID, s.db)
	if err != nil {
		log.Error("Error posting funds", "error", err)
		return nil, err
	}

	return &Response{Message: fmt.Sprintf("Posted funds for transfer: %s", pendingID)}, nil
}
