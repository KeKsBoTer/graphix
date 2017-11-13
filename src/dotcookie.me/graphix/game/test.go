package game

import (
	"dotcookie.me/graphix/graphics"
	"fmt"
	"log"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var texture *graphics.Texture
var region *graphics.TextureRegion
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
	tex, err := graphics.LoadTexture("morty.png")
	if err != nil {
		log.Fatalln(err)
	}
	tex2, err := graphics.LoadTexture("square.png")
	if err != nil {
		log.Fatalln(err)
	}
	texture = tex2
	region = graphics.NewTextureRegion(tex, 50, 50, 300, 300)

	batch = graphics.NewSpriteBatch()

	camera = graphics.NewCamera(float32(windowWidth), float32(windowHeight))

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

}

var x, y float32 = 0, 0
var stateTime float64 = 0

func (screen TestScreen) Render(delta float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	stateTime += delta
	// Render
	batch.Begin()
	batch.SetProjectionMatrix(*camera.GetProjection())

	time := glfw.GetTime()

	border := 1 / 60.0

	batch.SetTransformationMatrix(*camera.GetView())
	for i := float64(0); i < 1; i++ {

		//x = float32(math.Sin(math.Pi*stateTime+i)) * 200
		//y = float32(math.Sin(math.Pi*stateTime+i)) * 200
		//batch.DrawTexture(*texture, 0, y, 100, 100)
		batch.DrawRegion(*region,0, 0, 100, 100)
		//batch.DrawRegion(*region, 100, 100, 100, 100) //TODO draws second image
		//batch.DrawTexture(texture, 0, 200, 100, 100)

		if glfw.GetTime()-time > border {
			fmt.Println(i)
			break
		}
	}
	batch.End()
}

func (screen TestScreen) Dispose() {
	fmt.Println("Dispose")
	texture.Dispose()
}

func (screen TestScreen) Resize(width, height int32) {
	fmt.Println("Resize",width,height)
	camera.SetViewport(float32(width), float32(height))
	//gl.Viewport(0, 0, width, height)
	batch.SetProjectionMatrix(*camera.GetProjection())
}

func (screen TestScreen) KeyPressed(key glfw.Key) {
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
