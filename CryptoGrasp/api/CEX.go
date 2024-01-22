package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bhanusumanth/go/cryptograsp/models"
)

const endpoint = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*models.Rate, error) {
	uppercaseCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(endpoint, uppercaseCurrency))
	var cexResponseObj CEXResponse
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		responseBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		// resJson := string(responseBytes)
		// fmt.Println("Response : ", resJson)
		err = json.Unmarshal(responseBytes, &cexResponseObj)
	} else {
		errStr := fmt.Sprintf("Status code received : %v", res.StatusCode)
		return nil, errors.New(errStr)
	}
	outputRate := models.Rate{Currency: currency, Price: cexResponseObj.Bid}
	return &outputRate, nil

}
