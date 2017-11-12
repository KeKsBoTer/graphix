package main

import (
	"dotcookie.me/graphix/graphics"
	"dotcookie.me/graphix/game"
)

func main() {
	graphics.DesktopApplication(graphics.WindowConfig{
		Width:      800,
		Height:     600,
		Title:      "test",
		Resizeable: true,
		Vsync:true,
		Samples:0,
		Fullscreen:false,
	}, new(game.TestScreen))
}
