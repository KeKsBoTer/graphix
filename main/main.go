package main

import (
	"github.com/KeKsBoTer/graphix/graphics"
)

func main() {
	graphics.DesktopApplication(graphics.WindowConfig{
		Width:      800,
		Height:     600,
		Title:      "Testing environment",
		Resizeable: true,
		Vsync:true,
		Samples:0,
		Fullscreen:false,
	}, UIScreen{})
}
