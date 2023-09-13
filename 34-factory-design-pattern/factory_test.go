package factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash) // common method

	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)

	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasnt correct")
	}

	// write logs (not errors)
	t.Log("LOG:", msg)

}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard) // common method

	if err != nil {
		t.Fatal("A payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(22.30)

	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasnt correct")
	}

	// write logs (not errors)
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodCreditCard(t *testing.T) {
	payment, err := GetPaymentMethod(CreditCard) // common method

	if err != nil {
		t.Fatal("A payment method of type 'CreditCard' must exist")
	}

	msg := payment.Pay(18.30)

	if !strings.Contains(msg, "paid using credit card") {
		t.Error("The credit card payment method message wasnt correct")
	}

	// write logs (not errors)
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T){
	_, err := GetPaymentMethod(50)

	if err == nil {
		t.Error("A payment method with ID 50 must return an error")
	}

	t.Log("LOG:", err)
}