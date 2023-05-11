package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

/*
var testButt ebitButton = ebitButton{
	height: 110,
	width:  110,
	//xpos   float64
	//ypos   float64

	//fillColor    image.RGBA
	//boarderColor image.RGBA

	text: "string",

	isPressed: false,
}
*/

var number float64 = 0

var cpsUpper float64 = 0
var cpsUpperPrice float64 = 20
var cpsUpperIncrement float64 = 1

var cpsMultiplier float64 = 1
var cpsMultiplierPrice float64 = 5000
var cpsMultiplierIncrement float64 = 0.5

var cpsToThePower float64 = 1
var cpsToThePowerPrice float64 = 800000
var cpsToThePowerIncrement float64 = 0.02

var rebirths float64 = 0
var rebirthPrice float64 = 1000000000

var currentCps float64 = 0
var playerInformation string = ""
var creditCheck bool = false

var img *ebiten.Image

/* var img2 *ebiten.Image */

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("winner.png")
	/*img2 = ebiten.NewImage(testButt.width, testButt.height) */

	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 50)
	op.GeoM.Scale(0.4, 0.4)

	screen.Fill(color.RGBA{0, 0, 0, 0xff})
	/*img2.Fill(color.RGBA{110, 100, 0, 0xff})*/
	/*screen.DrawImage(img2, op)*/

	playerInformation =
		"Current clicks: " + numberFormatting(number) +
			"\nCurrent cps: " + strconv.FormatFloat((math.Pow(((cpsUpper)*(cpsMultiplier)), cpsToThePower))+((math.Pow(((cpsUpper)*(cpsMultiplier)), cpsToThePower))*(rebirths)), 'f', 1, 64) +
			"\n\n[q] Current cps+ ($" + numberFormatting(cpsUpperPrice) + "): " + numberFormatting(cpsUpper) +
			"\n[w] Current cps* ($" + numberFormatting(cpsMultiplierPrice) + "): " + numberFormatting(cpsMultiplier) +
			"\n[e] Current cps^ ($" + numberFormatting(cpsToThePowerPrice) + "): " + numberFormatting(cpsToThePower) +
			"\n[r] Rebirths cps+cps* ($" + numberFormatting(rebirthPrice) + "): " + numberFormatting(rebirths) +
			"\n\n\n\n\n\n           [s] Win! ($" + numberFormatting(1000000000000) + ")"

	ebitenutil.DebugPrint(screen, playerInformation)

	if creditCheck {
		screen.Fill(color.RGBA{0, 0, 0, 0xff})
		screen.DrawImage(img, op)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	clickIncrement()
	makeCpsGoUp()
	debugControls() // :>
	cpsIncrement()

	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Golang clicker game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func clickIncrement() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		number = number + 1
	}

}

func cpsIncrement() {
	number = number + ((math.Pow(((cpsUpper)*(cpsMultiplier)), cpsToThePower))+((math.Pow(((cpsUpper)*(cpsMultiplier)), cpsToThePower))*(rebirths)))/60
}

func rebirthing() {
	number = 5000 * rebirths

	cpsUpper = 0
	cpsUpperPrice = 20
	cpsMultiplier = 1
	cpsMultiplierPrice = 5000
	cpsToThePower = 1
	cpsToThePowerPrice = 800000
}

func makeCpsGoUp() {
	if (ebiten.IsKeyPressed(ebiten.KeyQ)) && (number >= cpsUpperPrice) {
		number = number - cpsUpperPrice
		cpsUpperPrice = cpsUpperPrice * 1.005
		cpsUpper = cpsUpper + cpsUpperIncrement
	} else if (ebiten.IsKeyPressed(ebiten.KeyW)) && (number >= cpsMultiplierPrice) {
		number = number - cpsMultiplierPrice
		cpsMultiplierPrice = cpsMultiplierPrice * 1.03
		cpsMultiplier = cpsMultiplier + cpsMultiplierIncrement
	} else if (ebiten.IsKeyPressed(ebiten.KeyE)) && (number >= cpsToThePowerPrice) {
		number = number - cpsToThePowerPrice
		cpsToThePowerPrice = math.Pow(cpsToThePowerPrice, 1.02)
		cpsToThePower = cpsToThePower + cpsToThePowerIncrement
	} else if (ebiten.IsKeyPressed(ebiten.KeyR)) && (number >= rebirthPrice) {
		rebirths = rebirths + 0.5
		rebirthing()

	} else if (inpututil.IsKeyJustPressed(ebiten.KeyS)) && (number >= 1000000000000) {
		creditCheck = true
	}
}

func debugControls() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		cpsUpperPrice = cpsUpperPrice * 1.005
		cpsUpper = cpsUpper + cpsUpperIncrement
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF2) {
		cpsMultiplierPrice = cpsMultiplierPrice * 1.03
		cpsMultiplier = cpsMultiplier + cpsMultiplierIncrement
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		cpsToThePowerPrice = math.Pow(cpsToThePowerPrice, 1.02)
		cpsToThePower = cpsToThePower + cpsToThePowerIncrement
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF4) {
		rebirths = rebirths + 0.5
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		creditCheck = true
	}
}

func numberFormatting(input float64) string {
	// tl;dr
	// 2.0e10 = 2*10^10 = 20000000000
	// 100000 = 1*10^(numOfZeros) = 1x10^6 = 1e6
	// but if less than 100000, don't do anything

	if input > 10000000000 {
		var convertedInput string = strconv.FormatFloat(input, 'f', 0, 64)
		var digits int = len(convertedInput) - 1
		return firstN(convertedInput, 1) + "e" + strconv.Itoa(digits)
	} else {
		return strconv.FormatFloat(input, 'f', 1, 64)
	}

}

func firstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

/*
func saving() {

}
*/
