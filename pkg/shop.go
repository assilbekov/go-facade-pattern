package pkg

import (
	"errors"
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

	if u.Card.Balance < p.Price {
		return errors.New("not enough money")
	}

	u.Card.Balance -= p.Price
	s.Products = append(s.Products, *p)

	return nil
}
