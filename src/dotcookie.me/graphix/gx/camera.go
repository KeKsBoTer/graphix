package gx

import (
	"github.com/go-gl/mathgl/mgl32"
	"fmt"
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
	c.projection = mgl32.Perspective(mgl32.DegToRad(45.0), (c.viewportWidth)/(c.viewportHeight), c.near, c.far)
	c.projection = mgl32.Ortho2D(width,0,height,0)
	//c.projection = c.projection.Transpose()
	fmt.Println(c.projection)
}

func ortho(width,height float32) mgl32.Mat4{
	near,far := float32(-1.0),float32(1.0)
	//x,y:=0.0,0.0
	left,right,top,bottom:=float32(0.0),width,height,float32(0.0)
	return mgl32.Mat4{	2./(right-left),0,0,0,
						0,2./(top-bottom),0,0,
						0,0,-2./(far-near),0,
						-(right+left)/(right-left),-(top+bottom)/(top-bottom),-(far+near)/(far-near),1}
}

func NewCamera(viewportWidth, viewportHeight float32) *Camera {
	c := Camera{
		near:           0.1,
		far:            10.0,
		viewportWidth:  viewportWidth,
		viewportHeight: viewportHeight,
	}
	c.position = mgl32.Vec3{0, 1, 0}
	c.SetViewport(viewportWidth,viewportHeight)
	//c.projection = mgl32.Perspective(mgl32.DegToRad(45.0), (c.viewportWidth)/(c.viewportHeight), c.near, c.far)

	fmt.Println(c.projection.Mul4x1(mgl32.Vec4{1,-1,0,1}))

	c.view = mgl32.LookAtV(
		c.position, //cam position
		mgl32.Vec3{0, 0, 0}, //center
		mgl32.Vec3{0, 1, 0}, //up
	)
	return &c
}

func (c *Camera) GetProjection() *mgl32.Mat4 {
	return &c.projection
}

func (c *Camera) GetView() *mgl32.Mat4 {
	return &c.view
}
