package pave

import (
	"context"
	"fmt"

	"encore.dev/rlog"

	"encore.app/internal/helpers"
	"encore.app/internal/storage"
)

type AuthorizeOpts struct {
	TransferID  string
	Code        uint16
	FromAccount string
	ToAccount   string
	Amount      uint64
}

//encore:api public method=POST path=/authorize
func (s *Service) Authorize(ctx context.Context, opts *AuthorizeOpts) (*TransferResponse, error) {
	log := rlog.With(
		"transferID", opts.TransferID,
		"code", opts.Code,
		"fromAccount", opts.FromAccount,
		"toAccount", opts.ToAccount,
		"amount", opts.Amount,
	)

	log.Debug("Authorize")

	transfer, err := storage.Authorize(ctx, storage.AuthorizeOpts{
		TransferID:  helpers.Uint128(opts.TransferID),
		Code:        opts.Code,
		FromAccount: opts.FromAccount,
		ToAccount:   opts.ToAccount,
		Amount:      opts.Amount,
	}, s.db)
	if err != nil {
		log.Error("Error transferring funds", "error", err)
		return nil, fmt.Errorf("error transferring funds: %s", err)
	}

	return &TransferResponse{Message: "Authorized transfer", Transfer: tbTransferToTransfer(transfer)}, nil
}
