package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	combined                      mgl32.Mat4
	projection, view              mgl32.Mat4
	near, far                     float32
	position                      mgl32.Vec3
	viewportWidth, viewportHeight float32
	zoom                          float32
	up, direction                 mgl32.Vec3
}

func (c *Camera) Update() {
	c.projection = mgl32.Ortho(
		-c.viewportWidth/2, c.viewportWidth/2,
		-c.viewportHeight/2, c.viewportHeight/2,
		c.near, c.far,
	).Mul4(mgl32.Scale3D(c.zoom, c.zoom, 1)) // zoom in/out
	c.view = mgl32.LookAtV(
		c.position,
		c.direction,
		c.up,
	)
}

func (c *Camera) SetViewport(width, height float32) {
	c.viewportWidth = width
	c.viewportHeight = height
	c.Update()
}
func NewCamera(viewportWidth, viewportHeight float32) *Camera {
	c := Camera{
		far:            1,
		near:           -1,
		viewportWidth:  viewportWidth,
		viewportHeight: viewportHeight,
		zoom:           1,
		position:       mgl32.Vec3{},
		up:             mgl32.Vec3{0, 1, 0},
		direction:      mgl32.Vec3{0, 0, -1},
	}
	c.Update()
	return &c
}

func (c *Camera) GetProjection() *mgl32.Mat4 {
	return &c.projection
}

func (c *Camera) GetView() *mgl32.Mat4 {
	return &c.view
}

func (c *Camera) SetPosition(x, y float32) {
	c.position[0], c.position[1] = x, y
}

func (c *Camera) Translate(x, y float32) {
	c.position[0], c.position[1] = c.position[0]+x, c.position[1]+y
}
