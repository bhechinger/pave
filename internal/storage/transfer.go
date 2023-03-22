package storage

import (
	"context"
	"fmt"

	"encore.dev/rlog"
	tb "github.com/tigerbeetledb/tigerbeetle-go"
	tbTypes "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
)

func createTransfer(_ context.Context, transfer tbTypes.Transfer, db tb.Client) error {
	log := rlog.With(
		"id", transfer.ID,
	)

	transfersRes, err := db.CreateTransfers([]tbTypes.Transfer{transfer})
	if err != nil {
		log.Error("Error transferring funds", "error", err)
		return fmt.Errorf("error creating transfer: %s", err)
	}

	for _, err := range transfersRes {
		log.Error("error transferring funds", "index", err.Index, "result", err.Result)
		return fmt.Errorf("error from results: %s", err.Result.String())
	}

	return nil
}

func GetTransfer(_ context.Context, transferID tbTypes.Uint128, db tb.Client) (tbTypes.Transfer, error) {
	log := rlog.With("transferID", transferID)

	transfers, err := db.LookupTransfers([]tbTypes.Uint128{transferID})
	if err != nil {
		log.Error("Error getting transfers", "error", err)
		return tbTypes.Transfer{}, err
	}

	if len(transfers) == 0 {
		return tbTypes.Transfer{}, fmt.Errorf("transfer not found")
	}

	return transfers[0], nil
}
