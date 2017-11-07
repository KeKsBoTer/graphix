package main

import (
	"dotcookie.me/graphix/gx"
	"dotcookie.me/graphix/game"
)

func main() {
	gx.DesktopApplication(gx.WindowConfig{
		Width:      800,
		Height:     600,
		Title:      "test",
		Resizeable: true,
		Vsync:true,
		Samples:0,
		Fullscreen:false,
	}, new(game.TestScreen))
}
