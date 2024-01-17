package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bhanusumanth/go/cryptograsp/models"
)

const endpoint = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*models.Rate, error) {
	uppercaseCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(endpoint, uppercaseCurrency))
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n %v resposne %v : ", res, err)
	return nil, nil

}
