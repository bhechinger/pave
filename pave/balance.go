package pave

import (
	"context"
	"fmt"

	"encore.dev/rlog"
	tbTypes "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"

	"encore.app/internal/helpers"
)

//encore:api public path=/balance/:id
func (s *Service) GetBalance(_ context.Context, id string) (*AccountResponse, error) {
	accounts, err := s.db.LookupAccounts([]tbTypes.Uint128{helpers.Uint128(id)})
	if err != nil {
		rlog.Error("could not fetch accounts", "error", err)
		return nil, fmt.Errorf("could not fetch accounts: %s", err)
	}

	if len(accounts) == 0 {
		return nil, fmt.Errorf("account not found: %s", id)
	}

	return &AccountResponse{Message: "Account info", Account: tbAccountToAccount(accounts[0])}, nil
}
