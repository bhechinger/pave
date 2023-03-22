package pave

type Account struct {
	ID             string
	UserData       string
	Ledger         uint32
	Code           uint16
	Flags          uint16
	DebitsPending  uint64
	DebitsPosted   uint64
	CreditsPending uint64
	CreditsPosted  uint64
	Timestamp      uint64
}

type Transfer struct {
	ID              string
	DebitAccountID  string
	CreditAccountID string
	UserData        string
	Reserved        string
	PendingID       string
	Timeout         uint64
	Ledger          uint32
	Code            uint16
	Flags           uint16
	Amount          uint64
	Timestamp       uint64
}

type Response struct {
	Transfer   Transfer
	Message    string
	TransferID string
	Accounts   []Account
}
