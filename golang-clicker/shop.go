package main

import (
	_ "math"

	"github.com/hajimehoshi/ebiten/v2"
)

type shopItem struct {
	itemName         string
	currentlyOwned   float64
	currentPrice     float64
	priceMultiplier  float64
	upgradeIncrement float64
	upgradeKey       ebiten.Key
}

func (item *shopItem) purchaseItem() {
	if ebiten.IsKeyPressed(item.upgradeKey) && (number >= item.currentPrice) {
		number -= item.currentPrice
		item.currentPrice *= item.priceMultiplier
		item.currentlyOwned += item.upgradeIncrement
	}
}

func (item shopItem) itemInformation() string {

	info := "\n" + item.itemName + "($" + numberFormatting(item.currentPrice, 0) + "): " + numberFormatting(item.currentlyOwned, 2)

	return info

}
