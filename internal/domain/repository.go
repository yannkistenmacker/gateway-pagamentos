package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(apikey string) (*Account, error)
	FindByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
}
