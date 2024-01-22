package main

import (
	"github.com/bhanusumanth/go/cryptograsp/api"
)

func main() {
	getCurrencyData("BTC")
	getCurrencyData("ETH")
	getCurrencyData("SHIB")
	getCurrencyData("SOL")
}

func getCurrencyData(c string) {
	res, err := api.GetRate(c)
	res.PrintRate(err)
}
