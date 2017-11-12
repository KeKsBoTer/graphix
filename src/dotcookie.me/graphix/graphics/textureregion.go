package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
	"dotcookie.me/graphix/utils"
)

type TextureRegion struct {
	texture Texture
	u, v, // left upper corner
	u2, v2 float32 // right lower corner
}

func NewTextureRegion(texture Texture, x, y, width, height int32) *TextureRegion {
	region := TextureRegion{texture: texture}
	region.SetRegionA(x, y, width, height)
	return &region
}

func (t *TextureRegion) GetTexture() Texture {
	return t.texture
}

func (t *TextureRegion) SetTexture(texture Texture) {
	t.texture = texture
}

func (t *TextureRegion) GetRegionSize() (float32, float32) {
	return mgl32.Abs(t.u2 - t.u), mgl32.Abs(t.v2 - t.v)
}

func (t *TextureRegion) GetRegionPos() (float32, float32) {
	return t.u, t.v
}

func (t *TextureRegion) SetRegionR(u, v, u2, v2 float32) {
	t.u, t.v, t.u2, t.v2 = u, v, u2, v2
}

func (t *TextureRegion) SetRegionA(x, y, width, height int32) {
	textureWidth, textureHeight := t.texture.GetSize()
	t.SetRegionR(utils.Div(x, textureWidth), utils.Div(y, textureHeight), // u,v
		utils.Div(x+width, textureWidth), utils.Div(y+height, textureHeight)) // u2,v2
}

// TODO Setters,Getters,Scroll,Flip,Split