package gui

import "github.com/KeKsBoTer/graphix/graphics"

type Image struct {
	x, y, width, height float32 //Todo make generic struct with methods
	region              graphics.TextureRegion
}

func NewImage(region graphics.TextureRegion) *Image {
	tw, th := float32(region.GetTexture().GetWidth()), float32(region.GetTexture().GetHeight())
	rw, rh := region.GetRegionSize()
	return &Image{
		region: region,
		x:      0,
		y:      0,
		width:  rw * tw,
		height: rh * th,
	}
}

func (c *Image) Update(delta float64) {
}

func (c *Image) GetPos() (float32, float32) {
	return c.x, c.y
}

func (c *Image) GetSize() (float32, float32) {
	return c.width, c.height
}

func (c *Image) GetBounds() (float32, float32, float32, float32) {
	return c.x, c.y, c.width, c.height
}

func (c *Image) Render(batch *graphics.SpriteBatch) {
	x, y, width, height := c.GetBounds()
	batch.DrawRegion(c.region, x, y, width, height)
}

func (c *Image) Dispose() {
}

func (c *Image) SetPos(x, y float32) {
	c.x, c.y = x, y
}

func (c *Image) SetSize(width, height float32) {
	c.width, c.height = width, height
}
