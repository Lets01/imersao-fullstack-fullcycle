package model

import (
	"time"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string `json:"ownerName" valid:"notnull"`
	Bank      *Bank  `valid:"-"`
	Number    string `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func (account *Account) IsValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account {
		OwnerName: ownerName,
		Bank: bank,
		Number: number,
	}
	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.IsValid()

	if err != nil {
		return nil err
	}

	return &account, nil
}