package account

import (
	"testing"
)

func TestDepositValidAmount(t *testing.T) {
	account := NewAccount(0.0, "Bruce Wayne")
	account.Deposit(200.0)
	expectedAccountBalance := 200.0
	actualAccountBalance := account.Balance()
	if expectedAccountBalance != actualAccountBalance {
		t.Fatalf("Failed to deposit money in the bank account, expected: %f, actual: %f", expectedAccountBalance, actualAccountBalance)
	}

}

func TestWithdrawValidAmount(t *testing.T) {
  account := NewAccount(200.0, "Bruce Wayne")
  account.Withdraw(50.00)
  expectedAccountBalance := 150.0
	actualAccountBalance := account.Balance()
  if expectedAccountBalance != actualAccountBalance {
		t.Fatalf("Failed to deposit money in the bank account, expected: %f, actual: %f", expectedAccountBalance, actualAccountBalance)
	}
}
