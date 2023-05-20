package main

import (
	"math"
	_ "math"

	"github.com/hajimehoshi/ebiten/v2"
)

type shopItem struct {
	name             string
	currentlyOwned   float64
	currentPrice     float64
	priceMultiplier  float64
	upgradeIncrement float64
	upgradeKey       ebiten.Key
}

func (item *shopItem) purchaseItem(iter int, pM int) {
	for iter < pM && ebiten.IsKeyPressed(item.upgradeKey) {
		if number >= item.currentPrice {
			number -= item.currentPrice
			item.currentlyOwned += item.upgradeIncrement

			if item.name == "cpsToThePower" {
				item.currentPrice = math.Pow(item.currentPrice, item.priceMultiplier)
			} else {
				item.currentPrice *= item.priceMultiplier
			}
		}
		iter++
	}
}

func (item shopItem) itemInformation() string {

	info := "\n" + item.name + "($" + numberFormatting(item.currentPrice, 0) + "): " + numberFormatting(item.currentlyOwned, 2)

	return info

}
