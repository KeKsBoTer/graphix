package main

import (
	"dotcookie.me/graphix/graphics"
	"dotcookie.me/graphix/game"
)

func main() {
	graphics.DesktopApplication(graphics.WindowConfig{
		Width:      800,
		Height:     600,
		Title:      "Testing environment",
		Resizeable: true,
		Vsync:false,
		Samples:0,
		Fullscreen:false,
	}, game.TestScreen{})
}
