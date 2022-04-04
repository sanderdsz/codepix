package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type PixKeyRepository interface {
	Register(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	FindAccount(account *Account) error
	AddAccount(account *Account) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("Invalid kind type")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("Invalid status type")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewPixKey(kind string, key string, accountID string, account *Account, status string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:      kind,
		Key:       key,
		AccountID: accountID,
		Account:   account,
		Status:    status,
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}