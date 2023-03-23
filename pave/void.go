package pave

import (
	"context"

	"encore.dev/rlog"

	"encore.app/internal/storage"
)

//encore:api public path=/void/:pendingID
func (s *Service) Void(ctx context.Context, pendingID string) (*VoidResponse, error) {
	log := rlog.With(
		"pendingID", pendingID,
	)

	err := storage.Void(ctx, pendingID, s.db)
	if err != nil {
		log.Error("Error voiding transfer", "error", err)
		return nil, err
	}

	return &VoidResponse{Message: "Voided transfer", PendingID: pendingID}, nil
}
