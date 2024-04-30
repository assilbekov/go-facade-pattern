package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (b *Bank) CheckBalance(cardNumber string) error {
	fmt.Println("Checking balance for card", cardNumber)
	time.Sleep(time.Millisecond * 300)

	for _, c := range b.Cards {
		if c.Name != cardNumber {
			continue
		}
		if c.Balance <= 0 {
			return errors.New(fmt.Sprintf("card %s has negative balance", cardNumber))
		}

		fmt.Println("Card", cardNumber, "has balance", c.Balance)
		return nil
	}

	return fmt.Errorf("card %s not found", cardNumber)
}
