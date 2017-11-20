package graphics

import "fmt"

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

func (r *TiledMapRenderer) Render(batch *SpriteBatch) {
	for i, l := range r.tiledMap.Layers {
		fmt.Println(l.Name,l.Visible)
		if /*!l.Visible ||*/  l.Empty {
			continue
		}
		r.RenderTiledLayer(batch, i)
	}
}
func (r *TiledMapRenderer) RenderTiledLayer(batch *SpriteBatch, index int) {
	layer := r.tiledMap.Layers[index]
	//batch.SetColor(Color{1,1,1,layer.Opacity})
	for i, t := range layer.DecodedTiles {
		if t.IsNil() {
			continue
		}

		var x, y float32
		width := float32(t.Tileset.TileWidth)
		height := float32(t.Tileset.TileHeight)

		switch RenderOrder(r.tiledMap.RenderOrder) {
		case RightDown:
			x = float32((i - int(i/r.tiledMap.Width)*r.tiledMap.Width) * t.Tileset.TileWidth)
			y = float32((r.tiledMap.Height - i/r.tiledMap.Width) * t.Tileset.TileHeight)
			height *= -1 // flip diagonal
			break
		case RightUp:
			x = float32((i - int(i/r.tiledMap.Width)*r.tiledMap.Width) * t.Tileset.TileWidth)
			y = float32(r.tiledMap.Height - i/r.tiledMap.Width*t.Tileset.TileHeight)
			break
		default:
			x = float32((i - int(i/r.tiledMap.Width)*r.tiledMap.Width) * t.Tileset.TileWidth)
			y = float32(r.tiledMap.Height - i/r.tiledMap.Width*t.Tileset.TileHeight)
		}
		if t.DiagonalFlip {
			height *= -1 // flip diagonal
			y += float32(r.tiledMap.TileHeight) // translate left
		}
		if t.HorizontalFlip {
			width *= -1 // flip horizontal
			x += float32(r.tiledMap.TileWidth) // translate right
		}
		batch.DrawRegion(r.tiledMap.tileSets[t.Tileset.Name][int(t.ID)], x*r.scale, y*r.scale, width*r.scale, height*r.scale)
	}
}
