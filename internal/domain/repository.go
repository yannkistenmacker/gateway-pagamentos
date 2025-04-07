package domain

type AccountRepository interface {
	Save(account *Account)
	FindByAPIKey(apikey string) (*Account, error)
	FindByID(id string) (*Account, error)
	Update(account *Account) error
}
