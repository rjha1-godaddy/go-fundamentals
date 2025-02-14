package main

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNegativeAmount    = errors.New("amount cannot be negative")
)

type Account struct {
	Balance float64
}

func (a *Account) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("withdrawal failed: %w", ErrNegativeAmount)
	}
	if amount > a.Balance {
		return fmt.Errorf("withdrawal failed: %w", ErrInsufficientFunds)
	}

	a.Balance -= amount
	return nil
}

func main() {
	acc := &Account{Balance: 100}

	transactions := []float64{50, 200, -10}

	for _, amount := range transactions {
		err := acc.Withdraw(amount)
		if err != nil {
			if errors.Is(err, ErrInsufficientFunds) {
				fmt.Println("Error: You don't have enough balance.")
			} else if errors.Is(err, ErrNegativeAmount) {
				fmt.Println("Error: Invalid amount entered.")
			} else {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Withdrawal successful! New Balance:", acc.Balance)
		}
	}
}
