package wallet

import (
	"errors"
	"github.com/google/uuid"
	"github.com/DilshodK0592/wallet/pkg/types"
)

var (
	ErrPhoneregistered = errors.New("phone already registered")
	ErrAmountMustBePositive = errors.New("amount must be greater than zero")
	ErrAccountNotFound = errors.New("account not found")
	ErrNotEnoughBalance = errors.New("not enough balance")
	ErrPaymentNotFound = errors.New("payment not found")
)

type Service struct {
	nextAccountID int64
	accounts []*types.Account
	payments []*types.Payment
}


func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneregistered
		}
	}
	s.nextAccountID++
	 account := &types.Account{
		ID: s.nextAccountID,
		Phone: phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount <= 0 {
		return ErrAmountMustBePositive
	}
	var account *types.Account

	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}
	if account == nil {
		return ErrAccountNotFound
	}

	// зачисление средств пока не рассматриваем как платёж

	account.Balance += amount
	return nil
}

func (s *Service) Pay(accID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error) {
	if amount <= 0 {
		return nil, ErrAmountMustBePositive
	}
	var account *types.Account

	for _, acc := range s.accounts {
		if acc.ID == accID {
			account = acc
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}
	if account.Balance < amount {
		return nil, ErrNotEnoughBalance
	}
	account.Balance -= amount
	paymentID := uuid.New().String()
	payment := &types.Payment {
		ID: paymentID,
		AccountID: accID,
		Amount: amount,
		Category: category,
		Status: types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}

func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			return account, nil
		}
	}
	
	return  nil, ErrAccountNotFound
}


func (s *Service)Reject(paymentID string) error {
	var rejectPayment *types.Payment
	for _, payment := range s.payments {
		if payment.ID == paymentID {
			rejectPayment = payment
			break
		}
	}
	if rejectPayment == nil {
		return ErrPaymentNotFound
	}

	var rejectAccount *types.Account
	for _, account := range s.accounts {
		if account.ID == rejectPayment.AccountID {
			rejectAccount = account
			break
		}
	}
	if rejectAccount == nil {
		return ErrAccountNotFound
	}

	rejectPayment.Status = types.PaymentStatusFail
	rejectAccount.Balance += rejectPayment.Amount
	return nil

}


func (s *Service)FindPaymentByID(paymentID string) (*types.Payment, error) {
	for _, payment := range s.payments {
		if payment.ID == paymentID {
			return payment, nil
		}
	}
	return nil, ErrPaymentNotFound
}