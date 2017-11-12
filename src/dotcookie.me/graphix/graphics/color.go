package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Color struct {
	R, G, B, A float32
}

func (c *Color) DecodeFloatRGBA() float32 {
	enc := mgl32.Vec4{c.R, c.G, c.B, c.A}
	kDecodeDot := mgl32.Vec4{1.0, 1 / 255.0, 1 / 65025.0, 1 / 160581375.0}
	return enc.Dot(kDecodeDot)
}