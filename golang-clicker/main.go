package main

import (
	"image/color"
	_ "image/png"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

var number int = 0
var minutePassing int = 0
var cpsUpper int = 0

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	screen.DrawImage(img, nil)
	ebitenutil.DebugPrint(screen, strconv.Itoa(number))
	// ebitenutil.DebugPrint(screen, "Hello, World! I'm gonna lose my marbles\nasdf")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	clickIncrement()
	makeCpsGoUp()
	if minutePassing >= 60 {
		cpsIncrement()
		minutePassing = 0
	}
	
	minutePassing = minutePassing + 1

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
	number = number + (1 * cpsUpper)
}

func makeCpsGoUp() {
	if (inpututil.IsKeyJustPressed(ebiten.KeyQ)) && (number >= 20) {
		number = number - 20
		cpsUpper = cpsUpper + 1
	}
}
