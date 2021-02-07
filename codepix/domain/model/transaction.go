package model

import "errors"

const (
	TransactionPending string = "pending"
	TransactionCompleted string = "completed"
	TransactionError string = "error"
	TransactionConfirmed string= "confirmed"
)

type TransactionRepositoryInterface interface {
	RegisterKey(transaction *Transaction) error
	save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transaction struct {
	transactions []Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"name" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
}

func (t *Transaction) IsValid() error {
	_, err := govalidator.ValidateStruct(t)
	if t.Amount <= 0 {
		return errors.New("The amount must be grater than 0")
	}
	if t.Status != TransactionPending  && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("Invalid status for transaction")
	}

	if t.PixKeyTo.AccountID == t.AcountFrom.ID {
		return errors.New("The source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction {
		AccountFrom: accountFrom,
		Amount: amount,
		PixKeyTo: pixKeyTo,
		Status: TransactionPending,
		Description: description,
	}
	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.IsValid()

	if err != nil {
		return nil err
	}

	return &transaction, nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.Description = description
	err := t.isValid()
	return err
}