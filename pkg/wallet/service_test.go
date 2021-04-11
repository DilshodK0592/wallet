package wallet

import "testing"

func TestService_FindAccountByID_Found(t *testing.T) {
	
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