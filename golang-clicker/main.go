package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
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
var cpsUpperPriceMulti float64 = 1.005
var cpsUpperIncrement float64 = 1

var cpsMultiplier float64 = 1
var cpsMultiplierPrice float64 = 5000
var cpsMultiplierPriceMulti float64 = 1.03
var cpsMultiplierIncrement float64 = 0.5

var cpsToThePower float64 = 1
var cpsToThePowerPrice float64 = 800000
var cpsToThePowerPriceMulti float64 = 1.02
var cpsToThePowerIncrement float64 = 0.03

var rebirths float64 = 0
var rebirthPrice float64 = 1000000000

var winPrice float64 = 10000000000000000000000000

var currentCps float64 = 0
var purchaseMulti float64 = 1
var i int = 0
var playerInformation string = ""
var creditCheck bool = false

var img *ebiten.Image

/* var img2 *ebiten.Image */
var (
	mplusNormalFont font.Face
)

func init() {
	var err error

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     36,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	img, _, err = ebitenutil.NewImageFromFile("winner.png")

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
		"Current clicks: " + numberFormatting(number, 1) +
			"\nCurrent cps: " + strconv.FormatFloat((math.Pow(((cpsUpper)*(cpsMultiplier)), cpsToThePower))+((math.Pow(((cpsUpper)*(cpsMultiplier)), cpsToThePower))*(rebirths)), 'f', 1, 64) +
			"\n[z-,x+]Shop multiplier: " + numberFormatting(purchaseMulti, 0) +
			"\n\n[q] Current cps+ ($" + numberFormatting(cpsUpperPrice, 0) + "): " + numberFormatting(cpsUpper, 0) +
			"\n[w] Current cps* ($" + numberFormatting(cpsMultiplierPrice, 0) + "): " + numberFormatting(cpsMultiplier, 1) +
			"\n[e] Current cps^ ($" + numberFormatting(cpsToThePowerPrice, 0) + "): " + numberFormatting(cpsToThePower, 2) +
			"\n[r] Rebirths cps+cps* ($" + numberFormatting(rebirthPrice, 0) + "): " + numberFormatting(rebirths, 2) +
			"\n\n\n\n\n\n                 [s] Win! ($" + numberFormatting(winPrice, 0) + ")"

	text.Draw(screen, playerInformation, mplusNormalFont, 0, 0, color.White)
	//ebitenutil.DebugPrint(screen, playerInformation)

	if creditCheck {
		screen.Fill(color.RGBA{0, 0, 0, 0xff})
		screen.DrawImage(img, op)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	clickIncrement()
	shopControls()
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
		number = number + 1 + (2 * rebirths)
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

func shopControls() {
	/* change this to a switch statement lol */
	i = 0

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		for i <= int(purchaseMulti) && (number >= cpsUpperPrice) {
			number = number - cpsUpperPrice
			cpsUpperPrice = cpsUpperPrice * cpsUpperPriceMulti
			cpsUpper = cpsUpper + cpsUpperIncrement
			i++
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		for i <= int(purchaseMulti) && (number >= cpsMultiplierPrice) {
			number = number - cpsMultiplierPrice
			cpsMultiplierPrice = cpsMultiplierPrice * cpsMultiplierPriceMulti
			cpsMultiplier = cpsMultiplier + cpsMultiplierIncrement
			i++
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		for i <= int(purchaseMulti) && (number >= cpsToThePowerPrice) {
			number = number - cpsToThePowerPrice
			cpsToThePowerPrice = math.Pow(cpsToThePowerPrice, cpsToThePowerPriceMulti)
			cpsToThePower = cpsToThePower + cpsToThePowerIncrement
			i++
		}
	}

	if (ebiten.IsKeyPressed(ebiten.KeyR)) && (number >= rebirthPrice) {
		rebirths = rebirths + 0.5
		rebirthing()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		purchaseMulti = purchaseMulti + 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyZ) && !(purchaseMulti <= 1) {
		purchaseMulti = purchaseMulti - 1
	}

	if (inpututil.IsKeyJustPressed(ebiten.KeyS)) && (number >= winPrice) {
		creditCheck = true
	}

}

func debugControls() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		cpsUpperPrice = cpsUpperPrice * cpsUpperPriceMulti
		cpsUpper = cpsUpper + cpsUpperIncrement
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF2) {
		cpsMultiplierPrice = cpsMultiplierPrice * cpsMultiplierPriceMulti
		cpsMultiplier = cpsMultiplier + cpsMultiplierIncrement
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		cpsToThePowerPrice = math.Pow(cpsToThePowerPrice, cpsToThePowerPriceMulti)
		cpsToThePower = cpsToThePower + cpsToThePowerIncrement
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF4) {
		rebirths = rebirths + 0.5
	} else if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		creditCheck = true
	}
}

func numberFormatting(input float64, precision int) string {
	if input > 10000000000 {
		var convertedInput string = strconv.FormatFloat(input, 'f', 0, 64)
		var digits int = len(convertedInput) - 1
		return firstN(convertedInput, 1) + "e" + strconv.Itoa(digits)
	} else {
		return strconv.FormatFloat(input, 'f', precision, 64)
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
