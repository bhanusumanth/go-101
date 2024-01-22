package models

import "fmt"

type Rate struct {
	Currency string
	Price    float64
}

func (r Rate) PrintRate(err error) {
	if err == nil {
		fmt.Printf("\n Currency : %s \t Price %f", r.Currency, r.Price)
	} else {
		fmt.Printf("\n We got an Error : %v", err)
	}
}
