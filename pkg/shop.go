package pkg

import (
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (s *Shop) Buy(p *Product, u *User) error {
	fmt.Println("Buying product", p.Name, "for", p.Price)
	time.Sleep(1 * time.Second)

	err := u.Card.CheckBalance()
	if err != nil {
		return err
	}

	if u.Balance() < p.Price {
		return fmt.Errorf("user %s has insufficient balance", u.Name)
	}

	u.Pay(p.Price)
	fmt.Println("Product", p.Name, "bought", "for", p.Price, "from", s.Name)

	return nil
}
