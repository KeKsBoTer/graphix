package game

import (
	"dotcookie.me/graphix/graphics"
	"fmt"
	"log"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var texture graphics.Texture
var camera *graphics.Camera

var windowWidth, windowHeight int

var batch *graphics.SpriteBatch

type TestScreen struct {
}

func (screen TestScreen) Show() {

	windowWidth = graphics.App.Graphics.GetWidth()
	windowHeight = graphics.App.Graphics.GetHeight()

	graphics.App.Input.AddMouseListener(screen)
	graphics.App.Input.AddKeyListener(screen)

	fmt.Println("Create")
	// Load the texture
	tex, err := graphics.LoadTexture("square.png")
	if err != nil {
		log.Fatalln(err)
	}
	texture = *tex

	batch = graphics.NewSpriteBatch()

	camera = graphics.NewCamera(float32(windowWidth), float32(windowHeight))

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

}

func (screen TestScreen) Render(delta float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	/*mx, my := 0,0//window.GetCursorPos()
	if window.GetMouseButton(glfw.MouseButton1) == glfw.Press {
		angleX += (mx - lmx) * elapsed
		angleY += (my - lmy) * elapsed
	}
	lmx, lmy = mx, my*/
	//model = model.Mul4(mgl32.HomogRotate3D(float32(-angleY), mgl32.Vec3{0, 0, 1}))

	// Render
	batch.Begin()
	batch.SetProjectionMatrix(*camera.GetProjection())
	batch.SetTransformationMatrix(*camera.GetView())
	batch.DrawTexture(texture,0,100,100,100)
	batch.End()
}

func (screen TestScreen) Dispose() {
	fmt.Println("Dispose")
}

func (screen TestScreen) Resize(width, height int) {
	camera.SetViewport(float32(width),float32(height))
	batch.SetProjectionMatrix(*camera.GetProjection())
}

func (screen TestScreen) KeyPressed(key glfw.Key) {
	fmt.Println("Pressed:", key)
}

func (screen TestScreen) KeyReleased(key glfw.Key) {
	fmt.Println("Released:", key)
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