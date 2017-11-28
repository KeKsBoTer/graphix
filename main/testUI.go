package main

import (
	"fmt"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/KeKsBoTer/graphix/graphics"
	"time"
	"log"
	"github.com/go-gl/glfw/v3.2/glfw"
	"io/ioutil"
)

type UIScreen struct{}

var uiHolder = struct {
	animation *graphics.Animation
	camera    *graphics.Camera

	windowWidth, windowHeight int

	shape       *graphics.ShapeRenderer
	spriteBatch *graphics.SpriteBatch
	font        *graphics.BitmapFont

	renderCount int

	printFPS  bool
	stateTime float64
}{}

func (screen UIScreen) Show() {

	windowWidth := graphics.App.Graphics.GetWidth()
	windowHeight := graphics.App.Graphics.GetHeight()

	graphics.App.Input.AddMouseListener(screen)
	graphics.App.Input.AddKeyListener(screen)
	graphics.App.Input.AddMouseWheelListener(screen)

	uiHolder.camera = graphics.NewCamera(float32(windowWidth), float32(windowHeight))
	uiHolder.camera.SetZoom(2.5)
	uiHolder.camera.Position()[0] = 390
	uiHolder.shape = graphics.NewShapeRenderer()

	btmFont, err := graphics.LoadBitmapFont("Fonts/arial.fnt")
	if err != nil {
		log.Fatalln(err)
	}
	uiHolder.font = btmFont
	uiHolder.spriteBatch = graphics.NewSpriteBatch()

	vert, err := ioutil.ReadFile("shader/font.vert")
	frag, err := ioutil.ReadFile("shader/font.frag")
	fontShader, err := graphics.NewShaderProgram(string(vert), string(frag))
	if err!=nil{
		log.Fatalln(err)
	}
	uiHolder.spriteBatch.SetShader(fontShader)

	uiHolder.printFPS = false
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
	gl.ClearColor(0.5, .5, 1.0, 1.0)
	uiHolder.stateTime += delta

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	uiHolder.spriteBatch.SetColor(graphics.Color{1, 0, 0, 1})

	uiHolder.spriteBatch.SetProjectionMatrix(*uiHolder.camera.GetProjection())
	uiHolder.spriteBatch.SetTransformationMatrix(*uiHolder.camera.GetView())
	uiHolder.shape.SetProjectionMatrix(*uiHolder.camera.GetProjection())
	uiHolder.shape.SetTransformationMatrix(*uiHolder.camera.GetView())

	uiHolder.shape.Begin()
	uiHolder.shape.DrawRectangle(0, 0, 1000, 1)
	uiHolder.shape.End()

	uiHolder.spriteBatch.Begin()
	uiHolder.font.Draw("golang is Awsome", uiHolder.spriteBatch)
	uiHolder.spriteBatch.End()

	uiHolder.renderCount++
}

func (screen UIScreen) Dispose() {
}

func (screen UIScreen) Resize(width, height int32) {
	uiHolder.camera.SetViewport(float32(width), float32(height))
}

func (screen UIScreen) KeyPressed(key glfw.Key) {
	switch key {
	case 93:
		uiHolder.camera.SetZoom(uiHolder.camera.GetZoom() + 0.5)
		uiHolder.camera.Update()
		break
	case 47:
		uiHolder.camera.SetZoom(uiHolder.camera.GetZoom() - 0.5)
		uiHolder.camera.Update()
		break
	case glfw.KeyD:
		uiHolder.printFPS = !uiHolder.printFPS
		break
	case glfw.KeyF:
		//graphics.App.Graphics.ToggleFullScreen(!graphics.App.Graphics.IsFullScreen())
		break
	}
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
	uiHolder.camera.Position()[0] -= float32(xOff)
	uiHolder.camera.Position()[1] += float32(yOff)
	uiHolder.camera.Update()
}
