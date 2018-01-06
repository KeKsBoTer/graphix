package main

import (
	"fmt"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/KeKsBoTer/graphix/graphics"
	"time"
	"log"
	"github.com/go-gl/glfw/v3.2/glfw"
	"math/rand"
	"github.com/KeKsBoTer/graphix/graphics/gui"
)

type UIScreen struct{}

var uiHolder = struct {
	windowWidth, windowHeight int

	renderCount int

	printFPS  bool
	stateTime float64

	fps int

	gui *gui.GUI
}{}

func (screen UIScreen) Show() {

	windowWidth := graphics.App.Graphics.GetWidth()
	windowHeight := graphics.App.Graphics.GetHeight()

	graphics.App.Input.AddMouseListener(screen)
	graphics.App.Input.AddKeyListener(screen)
	graphics.App.Input.AddMouseWheelListener(screen)

	uiHolder.printFPS = false
	go func() {
		for ; ; {
			if uiHolder.printFPS {
				fmt.Println(uiHolder.renderCount)
			}
			uiHolder.fps = uiHolder.renderCount
			uiHolder.renderCount = 0
			time.Sleep(time.Second)
		}
	}()

	test := make([]byte, 2000)
	for i := 0; i < len(test); i++ {
		test[i] = byte(rand.Intn(128))
		if i != 0 && i%100 == 0 {
			i++
			test[i] = '\n'
		}
	}

	uiHolder.gui = gui.NewGUI(windowWidth, windowHeight)
	tex, err := graphics.LoadTexture("morty.png")
	if err != nil {
		log.Fatal(err)
	}
	region := graphics.NewTextureRegion(tex, 0, 0, tex.GetWidth(), tex.GetHeight())
	img := gui.NewImage(*region)
	uiHolder.gui.AddComponent(img)
}

func (screen UIScreen) Render(delta float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(1, 1, 1.0, 1.0)
	uiHolder.stateTime += delta

	uiHolder.gui.Render(delta)
	uiHolder.renderCount++
}

func (screen UIScreen) Dispose() {
	uiHolder.gui.Dispose()
}

func (screen UIScreen) Resize(width, height int32) {
	uiHolder.gui.Resize(width, height)
}

func (screen UIScreen) KeyPressed(key glfw.Key) {
}

func (screen UIScreen) KeyReleased(key glfw.Key) {
}

func (screen UIScreen) MouseMoved(x, y int) bool {
	return true
}

func (screen UIScreen) MousePressed(x, y int, button glfw.MouseButton) bool {
	return true
}
func (screen UIScreen) MouseReleased(x, y int, button glfw.MouseButton) bool {
	return true
}

func (screen UIScreen) Scrolled(xOff, yOff float64) {
}
