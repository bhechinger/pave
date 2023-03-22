package pave

import (
	"context"
	"fmt"

	"encore.dev/rlog"
	tbTypes "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"

	"encore.app/internal/helpers"
)

//encore:api public path=/balance/:id
func (s *Service) GetBalance(_ context.Context, id string) (*Response, error) {
	accounts, err := s.db.LookupAccounts([]tbTypes.Uint128{helpers.Uint128(id)})
	if err != nil {
		rlog.Error("could not fetch accounts", "error", err)
		return nil, fmt.Errorf("could not fetch accounts: %s", err)
	}

	if len(accounts) == 0 {
		return nil, fmt.Errorf("account not found: %s", id)
	}

	eAccounts := make([]Account, 1)

	eAccounts[0].ID = fmt.Sprintf("%v", accounts[0].ID)
	eAccounts[0].UserData = fmt.Sprintf("%s", accounts[0].UserData)
	eAccounts[0].Ledger = accounts[0].Ledger
	eAccounts[0].Code = accounts[0].Code
	eAccounts[0].Flags = accounts[0].Flags
	eAccounts[0].DebitsPending = accounts[0].DebitsPending
	eAccounts[0].DebitsPosted = accounts[0].DebitsPosted
	eAccounts[0].CreditsPending = accounts[0].CreditsPending
	eAccounts[0].CreditsPosted = accounts[0].CreditsPosted
	eAccounts[0].Timestamp = accounts[0].Timestamp

	return &Response{Message: "Account info", Accounts: eAccounts}, nil
}
