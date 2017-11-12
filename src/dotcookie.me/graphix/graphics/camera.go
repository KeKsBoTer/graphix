package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	combined                      mgl32.Mat4
	projection, view              mgl32.Mat4
	near, far                     float32
	position, direction           mgl32.Vec3
	viewportWidth, viewportHeight float32
}

func (c *Camera) SetViewport(width, height float32) {
	c.viewportWidth = width
	c.viewportHeight = height
	c.projection = mgl32.Ortho2D(-width/2, width/2, -height/2, height/2)
}
func NewCamera(viewportWidth, viewportHeight float32) *Camera {
	c := Camera{
		near:           0.1,
		far:            10.0,
		viewportWidth:  viewportWidth,
		viewportHeight: viewportHeight,
	}
	c.position = mgl32.Vec3{0,0,0}
	c.SetViewport(viewportWidth,viewportHeight)
	//c.projection = mgl32.Perspective(mgl32.DegToRad(45.0), (c.viewportWidth)/(c.viewportHeight), c.near, c.far)
	c.view = mgl32.Ident4().Mul4(mgl32.Translate3D(c.position.X(),c.position.Y(),c.position.Z()))
	//c.view = mgl32.Translate3D(c.position.X(),c.position.Y(),c.position.Z())
	return &c
}

func (c *Camera) GetProjection() *mgl32.Mat4 {
	return &c.projection
}

func (c *Camera) GetView() *mgl32.Mat4 {
	return &c.view
}