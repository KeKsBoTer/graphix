package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/KeKsBoTer/graphix/graphics"
	"time"
)

type UIScreen struct{}

var uiHolder = struct {
	animation *graphics.Animation
	camera    *graphics.Camera

	windowWidth, windowHeight int

	shape    *graphics.ShapeRenderer
	renderer *graphics.TiledMapRenderer

	renderCount int

	printFPS  bool
	stateTime float64
}{}

func (screen UIScreen) Show() {

	windowWidth := graphics.App.Graphics.GetWidth()
	windowHeight := graphics.App.Graphics.GetHeight()

	uiHolder.camera = graphics.NewCamera(float32(windowWidth), float32(windowHeight))

	uiHolder.shape = graphics.NewShapeRenderer()

	uiHolder.printFPS = true
	go func() {
		for ; ; {
			if uiHolder.printFPS {
				fmt.Println(uiHolder.renderCount)
			}
			uiHolder.renderCount = 0
			time.Sleep(time.Second)
		}
	}()
}

func (screen UIScreen) Render(delta float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.0, 1.0, 1.0, 1.0)
	uiHolder.stateTime += delta


	uiHolder.shape.Begin()
	uiHolder.shape.SetProjectionMatrix(*uiHolder.camera.GetProjection())
	uiHolder.shape.SetTransformationMatrix(*uiHolder.camera.GetView())

	uiHolder.shape.DrawRectangle(0,0,100,100)
	uiHolder.shape.End()

	uiHolder.renderCount++
}

func (screen UIScreen) Dispose() {
}

func (screen UIScreen) Resize(width, height int32) {
}
