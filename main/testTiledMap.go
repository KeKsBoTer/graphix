package main

import (
	"fmt"
	"log"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/KeKsBoTer/graphix/graphics"
	"time"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var animation *graphics.Animation
var camera *graphics.Camera

var windowWidth, windowHeight int

var batch *graphics.SpriteBatch
var renderer *graphics.TiledMapRenderer

var renderCount int

var printFPS = false

type TestScreen struct {
}

func (screen TestScreen) Show() {

	windowWidth = graphics.App.Graphics.GetWidth()
	windowHeight = graphics.App.Graphics.GetHeight()

	graphics.App.Input.AddMouseListener(screen)
	graphics.App.Input.AddKeyListener(screen)
	graphics.App.Input.AddMouseWheelListener(screen)

	fmt.Println("Create")
	tex2, err := graphics.LoadTexture("square.png")
	if err != nil {
		log.Fatalln(err)
	}
	animation = graphics.NewAnimation(tex2.Split(32, 32)[0], graphics.LoopPingPong, 1)

	batch = graphics.NewSpriteBatch()

	camera = graphics.NewCamera(float32(windowWidth), float32(windowHeight))

	tiledMap, err := graphics.LoadMap("testdata/orthogonal-outside.tmx")
	if err != nil {
		log.Fatalln(err)
	}
	renderer = graphics.NewTiledRenderer(*tiledMap, 1)
	go func() {
		for ; ; {
			if printFPS {
				fmt.Println(renderCount)
			}
			renderCount = 0
			time.Sleep(time.Second)
		}
	}()
}

var stateTime float64 = 0

func (screen TestScreen) Render(delta float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	stateTime += delta

	// Render
	batch.Begin()
	batch.SetProjectionMatrix(camera.GetProjection())
	batch.SetTransformationMatrix(camera.GetView())

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	renderer.SetFrustumC(camera)
	renderer.Render(batch)

	batch.End()
	renderCount++
}

func (screen TestScreen) Dispose() {
	fmt.Println("Dispose")
	r := animation.GetRegion(0)
	r.GetTexture().Dispose()
}

func (screen TestScreen) Resize(width, height int32) {
	fmt.Println("Resize", width, height)
	camera.SetViewport(float32(width), float32(height))
	batch.SetProjectionMatrix(camera.GetProjection())
}

func (screen TestScreen) KeyPressed(key glfw.Key) {
	switch key {
	case 93:
		camera.SetZoom(camera.GetZoom() + 0.1)
		camera.Update()
		break
	case 47:
		camera.SetZoom(camera.GetZoom() - 0.1)
		camera.Update()
		break
	case glfw.KeyD:
		printFPS = !printFPS
		break
	case glfw.KeyF:
		//graphics.App.Graphics.ToggleFullScreen(!graphics.App.Graphics.IsFullScreen())
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
