package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
)

type TiledMapRenderer struct {
	tiledMap TiledMap
	scale    float32
	frustum  mgl32.Vec4
}

func NewTiledRenderer(tiledMap TiledMap, scale float32) *TiledMapRenderer {
	renderer := TiledMapRenderer{
		tiledMap: tiledMap,
		scale:    scale,
	}
	return &renderer
}

func (r *TiledMapRenderer) GetMap() *TiledMap {
	return &r.tiledMap
}

func (r *TiledMapRenderer) SetFrustum(x, y, width, height float32) {
	r.frustum[0] = x
	r.frustum[1] = y
	r.frustum[2] = width
	r.frustum[3] = height
}
func (r *TiledMapRenderer) SetFrustumC(camera *Camera) {
	r.frustum[0] = camera.position[0] - camera.viewportWidth/(2*camera.zoom)
	r.frustum[1] = camera.position[1] - camera.viewportHeight/(2*camera.zoom)
	r.frustum[2] = camera.position[0] + camera.viewportWidth/(2*camera.zoom)
	r.frustum[3] = camera.position[1] + camera.viewportHeight/(2*camera.zoom)
}

//TODO render right order -> save index when reading
func (r *TiledMapRenderer) Render(batch *SpriteBatch) {
	for i, l := range r.tiledMap.Layers {
		if !l.Visible || l.Empty {
			continue // don't render hidden layers
		}
		r.RenderTiledLayer(batch, i)
	}
	for i, l := range r.tiledMap.ImageLayers {
		if !l.Visible {
			continue // don't render hidden layers
		}
		r.RenderImageLayer(batch, i)
	}
}

func (r *TiledMapRenderer) RenderImageLayer(batch *SpriteBatch, index int) {
	layer := r.tiledMap.ImageLayers[index]
	batch.SetColor(Color{1, 1, 1, layer.Opacity})
	x := float32(layer.OffsetX)
	y := float32(layer.OffsetY)
	width := float32(layer.Image.Width)
	height := float32(layer.Image.Height)
	switch RenderOrder(r.tiledMap.RenderOrder) {
	case RightDown:
		y = float32(r.tiledMap.Height*r.tiledMap.TileHeight) - float32(layer.OffsetY)
		height *= -1 // flip diagonal
		break
	case RightUp:
		break
	default:
	}
	texture := r.tiledMap.images[layer.Image.Source]
	batch.DrawTexture(texture, x*r.scale, y*r.scale, width*r.scale, height*r.scale)
}

func clamp(value *int, min, max int) {
	v := *value
	if v < min {
		*value = min
	} else if v > max {
		*value = max
	}
}

func (r *TiledMapRenderer) RenderTiledLayer(batch *SpriteBatch, index int) {
	layer := r.tiledMap.Layers[index]
	batch.SetColor(Color{1, 1, 1, layer.Opacity})

	minx := int(r.frustum.X() / (float32(r.tiledMap.TileWidth) * r.scale))
	maxy := r.tiledMap.Height - int(r.frustum.Y()/(float32(r.tiledMap.TileHeight)*r.scale))
	miny := r.tiledMap.Height - int(r.frustum.W()/(float32(r.tiledMap.TileHeight)*r.scale)) - 1
	maxx := int(r.frustum.Z()/(float32(r.tiledMap.TileHeight)*r.scale)) + 1

	if minx >= r.tiledMap.Width || miny >= r.tiledMap.Height || maxx < 0 || maxy < 0 {
		//return // nothing to draw
	}

	clamp(&minx, 0, r.tiledMap.Width)
	clamp(&miny, 0, r.tiledMap.Height)
	clamp(&maxx, 0, r.tiledMap.Width)
	clamp(&maxy, 0, r.tiledMap.Height)
	count := 0

	for ty := miny; ty < maxy; ty++ {
		for tx := minx; tx < maxx; tx++ {
			t := layer.DecodedTiles[ty*r.tiledMap.Width+tx]

			if t.IsNil() {
				continue
			}

			var x, y float32
			width := float32(r.tiledMap.TileWidth)
			height := float32(r.tiledMap.TileHeight)

			x = float32(tx) * width
			y = float32(r.tiledMap.Height-ty) * height
			height *= -1 // flip diagonal
			if t.DiagonalFlip {
				height *= -1                        // flip diagonal
				y += float32(r.tiledMap.TileHeight) // translate left
			}
			if t.HorizontalFlip {
				width *= -1                        // flip horizontal
				x += float32(r.tiledMap.TileWidth) // translate right
			}
			batch.DrawRegion(r.tiledMap.tileSets[t.Tileset.Name][int(t.ID)], x*r.scale, y*r.scale, width*r.scale, height*r.scale)

			count++
		}
	}
}
