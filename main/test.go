package main

import (
	"fmt"
	"log"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/KeKsBoTer/graphix/graphics"
)

var animation *graphics.Animation
var camera *graphics.Camera

var windowWidth, windowHeight int

var batch *graphics.SpriteBatch
var renderer *graphics.TiledMapRenderer

type TestScreen struct {
}

func (screen TestScreen) Show() {

	windowWidth = graphics.App.Graphics.GetWidth()
	windowHeight = graphics.App.Graphics.GetHeight()

	graphics.App.Input.AddMouseListener(screen)
	graphics.App.Input.AddKeyListener(screen)
	graphics.App.Input.AddMouseWheelListener(screen)

	fmt.Println("Create")
	tex2, err := graphics.LoadTexture("easing.png")
	if err != nil {
		log.Fatalln(err)
	}
	animation = graphics.NewAnimation(tex2.Split(32, 32)[0], graphics.LoopPingPong, 1)

	batch = graphics.NewSpriteBatch()

	camera = graphics.NewCamera(float32(windowWidth), float32(windowHeight))
	camera.SetZoom(2)
	camera.Update()

	tiledMap, err := graphics.LoadMap("testdata/orthogonal-outside.tmx")
	if err != nil {
		log.Fatalln(err)
	}
	renderer = graphics.NewTiledRenderer(*tiledMap, 1)
}

var stateTime float64 = 0

func (screen TestScreen) Render(delta float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.0, 1.0, 1.0, 1.0)
	stateTime += delta

	// Render
	batch.Begin()
	batch.SetProjectionMatrix(*camera.GetProjection())
	batch.SetTransformationMatrix(*camera.GetView())

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	renderer.Render(batch)

	batch.End()
}

func (screen TestScreen) Dispose() {
	fmt.Println("Dispose")
	r := animation.GetRegion(0)
	r.GetTexture().Dispose()
}

func (screen TestScreen) Resize(width, height int32) {
	fmt.Println("Resize", width, height)
	camera.SetViewport(float32(width), float32(height))
	batch.SetProjectionMatrix(*camera.GetProjection())
}

func (screen TestScreen) KeyPressed(key glfw.Key) {
	switch key {
	case 93:
		camera.SetZoom(camera.GetZoom()+0.1)
		camera.Update()
		break
	case 47:
		camera.SetZoom(camera.GetZoom()-0.1)
		camera.Update()
		break
	}
}

func (screen TestScreen) KeyReleased(key glfw.Key) {
}

func (screen TestScreen) MouseMoved(x, y int) bool {
	return true
}

func (screen TestScreen) MousePressed(x, y int, button glfw.MouseButton) bool {
	return true
}
func (screen TestScreen) MouseReleased(x, y int, button glfw.MouseButton) bool {
	return true
}

func (screen TestScreen) Scrolled(xOff, yOff float64) {
	camera.Position()[0] -= float32(xOff)
	camera.Position()[1] += float32(yOff)
	camera.Update()
}
