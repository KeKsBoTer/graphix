package graphics

type TiledMapRenderer struct {
	tiledMap TiledMap
	scale    float32
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
	batch.DrawTexture(&texture, x*r.scale, y*r.scale, width*r.scale, height*r.scale)
}

func (r *TiledMapRenderer) RenderTiledLayer(batch *SpriteBatch, index int) {
	layer := r.tiledMap.Layers[index]
	batch.SetColor(Color{1, 1, 1, layer.Opacity})
	for i, t := range layer.DecodedTiles {
		if t.IsNil() {
			continue
		}

		var x, y float32
		width := float32(t.Tileset.TileWidth)
		height := float32(t.Tileset.TileHeight)

		x = float32((i - int(i/r.tiledMap.Width)*r.tiledMap.Width) * t.Tileset.TileWidth)
		y = float32((r.tiledMap.Height - i/r.tiledMap.Width) * t.Tileset.TileHeight)
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
	}
}
