package Graphix

import (
	"dotcookie.me/graphix/graphics"
	"dotcookie.me/vendor/github.com/go-gl/mathgl/mgl32"
	"fmt"
	"log"
	"dotcookie.me/vendor/github.com/go-gl/gl/v4.1-core/gl"
	"dotcookie.me/vendor/github.com/go-gl/glfw/v3.2/glfw"
	"dotcookie.me/graphix/io"
)

var texture graphics.Texture
var model mgl32.Mat4
var program *graphics.ShaderProgram
var vao uint32
var camera *graphics.Camera

var windowWidth, windowHeight int

type TestScreen struct {
}

func (screen TestScreen) Show() {

	windowWidth = graphics.App.Graphics.GetWidth()
	windowHeight = graphics.App.Graphics.GetHeight()

	graphics.App.Input.AddMouseListener(screen)
	graphics.App.Input.AddKeyListener(screen)

	fmt.Println("Create")
	// Load the texture
	tex, err := io.LoadTexture("square.png")
	if err != nil {
		log.Fatalln(err)
	}
	texture = *tex

	vertexShader = io.LoadFile("base.vert")
	fragmentShader = io.LoadFile("base.frag")

	// Configure the vertex and fragment shaders
	p, err := graphics.NewShaderProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}
	program =p

	camera = graphics.NewCamera(float32(windowWidth), float32(windowHeight))

	program.SetUniformMatrix4fv("projection", camera.GetProjection(), false)

	program.SetUniformMatrix4fv("camera", camera.GetView(), false)

	model = mgl32.Scale3D(100,100,1)
	program.SetUniformMatrix4fv("model", &model, false)

	program.SetUniform1i("tex", 0)

	program.BindFragDataLocation("outputColor", 0)

	// Configure the vertex data

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(planeVertices)*4, gl.Ptr(planeVertices), gl.STATIC_DRAW)
	//gl.BufferData(gl.ARRAY_BUFFER, len(planeVertices)*4, gl.Ptr(planeVertices), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(program.GetId(), gl.Str(graphics.GlString("vert"))))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program.GetId(), gl.Str(graphics.GlString("vertTexCoord"))))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

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
	gl.UseProgram(program.GetId())
	program.SetUniformMatrix4fv("model",&model,false)

	gl.BindVertexArray(vao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture.GetId())

	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
}

func (screen TestScreen) Dispose() {
	fmt.Println("Dispose")
}

func (screen TestScreen) Resize(width, height int) {
	camera.SetViewport(float32(width),float32(height))
	program.SetUniformMatrix4fv("projection",camera.GetProjection(),false)
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

var vertexShader string
var fragmentShader string

var planeVertices = []float32{
	-1.0, -1.0, 1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,

}
