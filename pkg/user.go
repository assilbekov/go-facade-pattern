package pkg

type User struct {
	Name string
	Card *Card
}

func (u *User) Balance() float64 {
	return u.Card.Balance
}

func (u *User) Pay(amount float64) {
	u.Card.Balance -= amount
}
