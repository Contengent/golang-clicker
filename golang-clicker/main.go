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

func init() {
	var err error
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

	playerInformation =
		"Current clicks: " + strconv.FormatFloat(number, 'f', 1, 64) +
			"\nCurrent cps: " + strconv.FormatFloat((math.Pow(((cpsUpper)*(cpsMultiplier)), cpsToThePower))+((math.Pow(((cpsUpper)*(cpsMultiplier)), cpsToThePower))*(rebirths)), 'f', 1, 64) +
			"\n\n[q] Current cps+ ($" + strconv.FormatFloat(cpsUpperPrice, 'f', 0, 64) + "): " + strconv.FormatFloat(cpsUpper, 'f', 0, 64) +
			"\n[w] Current cps* ($" + strconv.FormatFloat(cpsMultiplierPrice, 'f', 0, 64) + "): " + strconv.FormatFloat(cpsMultiplier, 'f', 1, 64) +
			"\n[e] Current cps^ ($" + strconv.FormatFloat(cpsToThePowerPrice, 'f', 0, 64) + "): " + strconv.FormatFloat(cpsToThePower, 'f', 2, 64) +
			"\n[r] Rebirths cps+cps* ($" + strconv.FormatFloat(rebirthPrice, 'f', 0, 64) + "): " + strconv.FormatFloat(rebirths, 'f', 2, 64) +
			"\n\n\n\n\n\n           [s] Win! ($999999999999999999)"

	ebitenutil.DebugPrint(screen, playerInformation)

	if creditCheck {
		screen.Fill(color.RGBA{0, 0, 0, 0xff})
		screen.DrawImage(img, op)
	}

	// ebitenutil.DebugPrint(screen, "Hello, World! I'm gonna lose my marbles\nasdf")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	clickIncrement()
	makeCpsGoUp()
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
	number = 5000

	cpsUpper = 0
	cpsUpperPrice = 20
	cpsMultiplier = 1
	cpsMultiplierPrice = 5000
	cpsToThePower = 111
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

	} else if (ebiten.IsKeyPressed(ebiten.KeyS)) && (number >= 999999999999999999) {
		creditCheck = true
	}
}
