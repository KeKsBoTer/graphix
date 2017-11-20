package graphics

import (
	"unsafe"
	"github.com/go-gl/mathgl/mgl32"
)

type Color struct {
	R, G, B, A float32
}

func (c *Color) Pack() float32 {
	r:= mgl32.Clamp(c.R,0,1)*255
	g:= mgl32.Clamp(c.G,0,1)*255
	b:= mgl32.Clamp(c.B,0,1)*255
	a:= mgl32.Clamp(c.A,0,1)*255
	data := [4]uint8{uint8(r),uint8(g),uint8(b),uint8(a)}
	return *(*float32)(unsafe.Pointer(&data))
}