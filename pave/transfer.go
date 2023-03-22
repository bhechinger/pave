package pave

import (
	"context"
	"fmt"

	"encore.dev/rlog"

	"encore.app/internal/helpers"
	"encore.app/internal/storage"
)

//encore:api public path=/transfer/get/:transferID
func (s *Service) GetTransfer(ctx context.Context, transferID string) (*Response, error) {
	log := rlog.With(
		"transferID", transferID,
	)

	transfer, err := storage.GetTransfer(ctx, helpers.Uint128(transferID), s.db)
	if err != nil {
		log.Error("Error voiding transfer", "error", err)
		return nil, err
	}

	ret := Transfer{
		ID:              transfer.ID.String(),
		DebitAccountID:  transfer.DebitAccountID.String(),
		CreditAccountID: transfer.CreditAccountID.String(),
		UserData:        transfer.UserData.String(),
		Reserved:        transfer.Reserved.String(),
		PendingID:       transfer.PendingID.String(),
		Timeout:         transfer.Timeout,
		Ledger:          transfer.Ledger,
		Code:            transfer.Code,
		Flags:           transfer.Flags,
		Amount:          transfer.Amount,
		Timestamp:       transfer.Timestamp,
	}

	return &Response{Message: fmt.Sprintf("Found transfer"), Transfer: ret}, nil
}
