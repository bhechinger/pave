package pave

import (
	"context"
	"fmt"

	"encore.dev/rlog"
	tbTypes "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"

	"encore.app/internal/helpers"
)

//encore:api public path=/account/create/:id
func (s *Service) CreateAccount(_ context.Context, id string) (*AccountResponse, error) {
	account := tbTypes.Account{
		ID:             helpers.Uint128(id),
		UserData:       tbTypes.Uint128{},
		Reserved:       [48]uint8{},
		Ledger:         1,
		Code:           1,
		Flags:          0,
		DebitsPending:  0,
		DebitsPosted:   0,
		CreditsPending: 0,
		CreditsPosted:  0,
		Timestamp:      0,
	}

	accountsRes, err := s.db.CreateAccounts([]tbTypes.Account{account})
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

	return &AccountResponse{Message: fmt.Sprintf("Added account: %s", id), Account: tbAccountToAccount(account)}, nil
}

//encore:api public path=/account/get/:id
func (s *Service) GetAccount(ctx context.Context, id string) (*AccountResponse, error) {
	accounts, err := s.db.LookupAccounts([]tbTypes.Uint128{helpers.Uint128(id)})
	if err != nil {
		rlog.Error("could not fetch accounts", "error", err)
		return nil, fmt.Errorf("could not fetch accounts: %s", err)
	}

	if len(accounts) == 0 {
		return nil, fmt.Errorf("account not found: %s", id)
	}

	fmt.Println(accounts)

	return &AccountResponse{Message: "Account info", Account: tbAccountToAccount(accounts[0])}, nil
}

func tbAccountToAccount(account tbTypes.Account) Account {
	return Account{
		ID:             fmt.Sprintf("%v", account.ID),
		UserData:       fmt.Sprintf("%s", account.UserData),
		Ledger:         account.Ledger,
		Code:           account.Code,
		Flags:          account.Flags,
		DebitsPending:  account.DebitsPending,
		DebitsPosted:   account.DebitsPosted,
		CreditsPending: account.CreditsPending,
		CreditsPosted:  account.CreditsPosted,
		Timestamp:      account.Timestamp,
	}
}
