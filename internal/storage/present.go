package storage

import (
	"context"
	"fmt"

	"encore.dev/rlog"
	tb "github.com/tigerbeetledb/tigerbeetle-go"
	tbTypes "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"

	"encore.app/internal/helpers"
)

func Present(ctx context.Context, pendingID string, db tb.Client) error {
	randID := helpers.GetRandID()

	log := rlog.With(
		"id", randID,
		"pendingID", pendingID,
	)

	transfer := tbTypes.Transfer{
		ID:              randID,
		PendingID:       helpers.Uint128(pendingID),
		DebitAccountID:  tbTypes.Uint128{},
		CreditAccountID: tbTypes.Uint128{},
		UserData:        helpers.Uint128("2"),
		Reserved:        tbTypes.Uint128{},
		Timeout:         0,
		Ledger:          1,
		Code:            0,
		Flags:           tbTypes.TransferFlags{PostPendingTransfer: true}.ToUint16(),
		Amount:          0,
		Timestamp:       0,
	}

	err := createTransfer(ctx, transfer, db)
	if err != nil {
		log.Error("Error posting transfer", "error", err)
		return fmt.Errorf("error posting transfer: %s", err)
	}

	return nil
}
