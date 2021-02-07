package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID    string `json:"id" valid:"uuid"`
	Name  string `json:"name" valid:"notnull"`
	Email string `json:"email" valid:"notnull"`
}

func (user *User) IsValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(nome string, email string) (*User, error) {
	user := User{
		Name:  nome,
		Email: email,
	}
	user.ID = uuid.NewV4().String()

	err := user.IsValid()

	if err != nil {
		return nil, err
	}

	return &user, nil
}
