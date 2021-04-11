package wallet

import (
	"reflect"
	"testing"

	"github.com/DilshodK0592/wallet/pkg/types"
)


func TestService_FindAccountByID_success(t *testing.T) {
	
	svc := Service{}
	svc.RegisterAccount("+9920000001")
	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_FindAccountByID_NotFound(t *testing.T) {

	svc := Service{}
	svc.RegisterAccount("+9920000001")
	account, err := svc.FindAccountByID(2)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_Reject_Success(t *testing.T) {
	s := &Service{}

	phone := types.Phone("+992000000001")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("Reject(): can't register account, error = %v", err)
		return
	}

	err = s.Deposit(account.ID, 1_000_00)
	if err != nil {
		t.Errorf("Reject(): can't deposit account, error = %v", err)
		return
	}

	payment, err := s.Pay(account.ID, 1_000_00, "auto")
	if err != nil {
		t.Errorf("Reject(): can't create payment, error = %v", err)
		return
	}

	err = s.Reject(payment.ID)
	if err != nil {
		t.Errorf("Reject(): error = %v", err)
	}
}

func TestService_FindPaymentByID_success(t *testing.T) {
	s := &Service{}

	phone := types.Phone("+992000000001")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("FindPaymentByID(): can't register account, error = %v", err)
		return
	}

	err = s.Deposit(account.ID, 1_000_00)
	if err != nil {
		t.Errorf("FindPaymentByID(): can't deposit account, error = %v", err)
		return
	}

	payment, err := s.Pay(account.ID, 1_000_00, "auto")
	if err != nil {
		t.Errorf("FindPaymentByID(): can't create payment, error = %v", err)
		return
	}

	got, err := s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("FindPaymentByID(): error = %v", err)
		return
	}

	if !reflect.DeepEqual(payment, got) {
		t.Errorf("FindPaymentByID(): wrong payment returned = %v", err)
		return
	}
}