package pave

import (
	"context"
	"fmt"

	"encore.dev/rlog"
	tbTypes "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"

	"encore.app/internal/helpers"
	"encore.app/internal/storage"
)

//encore:api public path=/transfer/get/:transferID
func (s *Service) GetTransfer(ctx context.Context, transferID string) (*TransferResponse, error) {
	log := rlog.With(
		"transferID", transferID,
	)

	transfer, err := storage.GetTransfer(ctx, helpers.Uint128(transferID), s.db)
	if err != nil {
		log.Error("Error voiding transfer", "error", err)
		return nil, err
	}

	return &TransferResponse{Message: fmt.Sprintf("Found transfer"), Transfer: tbTransferToTransfer(transfer)}, nil
}

func tbTransferToTransfer(transfer tbTypes.Transfer) Transfer {
	return Transfer{
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
}
