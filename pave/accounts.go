package pave

import (
	"context"
	"fmt"

	"encore.dev/rlog"
	tbTypes "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"

	"encore.app/internal/helpers"
)

//encore:api public path=/account/create/:id
func (s *Service) CreateAccount(ctx context.Context, id string) (*Response, error) {
	accountsRes, err := s.db.CreateAccounts([]tbTypes.Account{
		{
			ID:             helpers.Uint128(id),
			UserData:       tbTypes.Uint128{},
			Reserved:       [48]uint8{},
			Ledger:         1,
			Code:           718,
			Flags:          0,
			DebitsPending:  0,
			DebitsPosted:   0,
			CreditsPending: 0,
			CreditsPosted:  0,
			Timestamp:      0,
		},
	})
	if err != nil {
		rlog.Error("Error creating accounts: %s", err)
		return nil, fmt.Errorf("error creating accounts: %s", err)
	}

	for _, err := range accountsRes {
		if err.Result != tbTypes.AccountOK {
			rlog.Error("error creating account", "index", err.Index, "result", err.Result)
			return nil, fmt.Errorf("error from results: %s", err.Result.String())
		}
	}

	return &Response{Message: fmt.Sprintf("Added account: %s", id)}, nil
}

//encore:api public path=/account/get/:id
func (s *Service) GetAccount(ctx context.Context, id string) (*Response, error) {
	accounts, err := s.db.LookupAccounts([]tbTypes.Uint128{helpers.Uint128(id)})
	if err != nil {
		rlog.Error("could not fetch accounts", "error", err)
		return nil, fmt.Errorf("could not fetch accounts: %s", err)
	}

	if len(accounts) == 0 {
		return nil, fmt.Errorf("account not found: %s", id)
	}

	fmt.Println(accounts)

	eAccounts := make([]Account, len(accounts))

	for i := range accounts {
		eAccounts[i].ID = fmt.Sprintf("%v", accounts[i].ID)
		eAccounts[i].UserData = fmt.Sprintf("%s", accounts[i].UserData)
		eAccounts[i].Ledger = accounts[i].Ledger
		eAccounts[i].Code = accounts[i].Code
		eAccounts[i].Flags = accounts[i].Flags
		eAccounts[i].DebitsPending = accounts[i].DebitsPending
		eAccounts[i].DebitsPosted = accounts[i].DebitsPosted
		eAccounts[i].CreditsPending = accounts[i].CreditsPending
		eAccounts[i].CreditsPosted = accounts[i].CreditsPosted
		eAccounts[i].Timestamp = accounts[i].Timestamp
	}

	return &Response{Message: "Account info", Accounts: eAccounts}, nil
	//return &Response{Message: fmt.Sprintf("Account info: %v", accounts)}, nil
}
