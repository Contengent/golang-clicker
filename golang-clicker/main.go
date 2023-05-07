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
var cpsMultiplier float64 = 1
var cpsToThePower float64 = 1
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

	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})

	playerInformation =
		"Current clicks: " + strconv.FormatFloat(number, 'f', 1, 64) +
			"\nCurrent cps: " + strconv.FormatFloat(math.Pow(((1*cpsUpper)*(cpsMultiplier)), cpsToThePower), 'f', 1, 64) +
			"\n\n[q]Current cps+ ($20): " + strconv.FormatFloat(cpsUpper, 'f', 1, 64) +
			"\n[w]Current cps* ($5000): " + strconv.FormatFloat(cpsMultiplier, 'f', 1, 64) +
			"\n[e]Current cps^ ($800000): " + strconv.FormatFloat(cpsToThePower, 'f', 3, 64) +
			"\n\n\n\n\n\n            [r]Win! ($10000000000000000)"

	ebitenutil.DebugPrint(screen, playerInformation)

	if creditCheck {
		screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
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
	ebiten.SetWindowTitle("First project")
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
	number = number + (math.Pow(((1*cpsUpper)*(cpsMultiplier)), cpsToThePower))/60
}

func makeCpsGoUp() {
	if (ebiten.IsKeyPressed(ebiten.KeyQ)) && (number >= 20) {
		number = number - 20
		cpsUpper = cpsUpper + 1
	} else if (ebiten.IsKeyPressed(ebiten.KeyW)) && (number >= 5000) {
		number = number - 5000
		cpsMultiplier = cpsMultiplier + 0.5
	} else if (ebiten.IsKeyPressed(ebiten.KeyE)) && (number >= 800000) {
		number = number - 800000
		cpsToThePower = cpsToThePower + 0.002
	} else if (ebiten.IsKeyPressed(ebiten.KeyR)) && (number >= 10000000000000000) {
		number = number - 10000000000000000
		creditCheck = true
	}
}
