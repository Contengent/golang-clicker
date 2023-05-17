package main

import (
	_ "image/color"
)

type ebitButton struct {
	height int
	width  int

	/*
		xpos   float64
		ypos   float64

		fillColor    image.RGBA
		boarderColor image.RGBA
	*/

	text string

	isPressed bool
}
