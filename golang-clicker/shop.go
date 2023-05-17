package main

import (
	_ "math"
)

type shopItem struct {
	owned    float64
	price    float64
	strength float64
}

/*
func (shopItem shopItem) test() {
	if number >= shopItem.price {
		number = number - shopItem.price
		shopItem.price = math.Pow(shopItem.price, 1.02)
		cpsToThePower = cpsToThePower + cpsToThePowerIncrement
	}
}
*/
