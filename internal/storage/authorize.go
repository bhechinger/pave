package storage

import (
	"context"
	"fmt"

	"encore.dev/rlog"
	tb "github.com/tigerbeetledb/tigerbeetle-go"
	tbTypes "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"

	"encore.app/internal/helpers"
)

type AuthorizeOpts struct {
	TransferID  tbTypes.Uint128
	PendingID   string
	Code        uint16
	FromAccount string
	ToAccount   string
	Amount      uint64
}

func Authorize(ctx context.Context, opts AuthorizeOpts, db tb.Client) error {
	log := rlog.With(
		"TransferID", opts.TransferID,
		"PendingID", opts.PendingID,
	)

	transfer := tbTypes.Transfer{
		ID:              opts.TransferID,
		PendingID:       tbTypes.Uint128{},
		DebitAccountID:  helpers.Uint128(opts.FromAccount),
		CreditAccountID: helpers.Uint128(opts.ToAccount),
		UserData:        helpers.Uint128("2"),
		Reserved:        tbTypes.Uint128{},
		Timeout:         0,
		Ledger:          1,
		Code:            opts.Code,
		Flags:           tbTypes.TransferFlags{Pending: true}.ToUint16(),
		Amount:          opts.Amount,
		Timestamp:       0,
	}

	err := createTransfer(ctx, transfer, db)
	if err != nil {
		log.Error("Error authorizing transfer", "error", err)
		return fmt.Errorf("error authorizing transfer: %s", err)
	}

	return nil
}
