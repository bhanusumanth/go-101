package main

import (
	"fmt"
	"sync"

	"github.com/bhanusumanth/go/cryptograsp/api"
)

func main() {
	cryptoCurrencies := []string{"BTC", "ETH", "SHIB", "SOL", "DOGE", "USD", "MATIC", "XRP", "USDT", "BNB", "ADA",
		"AVAX", "TRX", "LINK", "TON", "LTC", "DAI", "ICP", "ATOM", "ETC", "OKB", "APT",
		"ARB", "KAS"}
	seqCalls(cryptoCurrencies)
	fmt.Printf("\n--------concurrent calls --------\n")
	var wg sync.WaitGroup
	for _, currency := range cryptoCurrencies {
		wg.Add(1)
		go func(c string) {
			getCurrencyData(c)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func seqCalls(currencies []string) {
	for _, currency := range currencies {
		getCurrencyData(currency)
	}
}

func getCurrencyData(c string) {
	res, err := api.GetRate(c)
	res.PrintRate(err)
}
