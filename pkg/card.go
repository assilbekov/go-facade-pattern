package pkg

import (
	"fmt"
	"time"
)

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (c *Card) CheckBalance() error {
	fmt.Println("Checking balance for card", c.Name, "with balance", c.Balance)
	time.Sleep(time.Millisecond * 500)
	return c.CheckBalance()
}
