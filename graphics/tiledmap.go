package graphics

import (
	"github.com/KeKsBoTer/go-tmx/tmx"
	"os"
	"path/filepath"
	"strings"
	"log"
)

type Orientation string

const (
	orthogonal Orientation = "orthogonal"
	isometric  Orientation = "isometric"
	staggered  Orientation = "staggered"
)

type RenderOrder string

const (
	RightUp   RenderOrder = "right-up"
	RightDown RenderOrder = "right-down"
	LeftUp    RenderOrder = "left-up"
	LeftDown  RenderOrder = "left-down"
)

type Polygon = tmx.Polygon
type PolyLine = tmx.PolyLine
type Property = tmx.Property

type TiledMap struct {
	tmx.Map
	tileSets map[string][]TextureRegion
	filePath string
}

func LoadMap(path string) (*TiledMap, error) {
	mapFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	tmxMap, err := tmx.Read(mapFile)
	if err != nil {
		return nil, err
	}
	if strings.Compare(tmxMap.Orientation, string(orthogonal)) != 0 {
		log.Fatalln("Only orthogonal maps are supported currently")
	}
	tiledMap := TiledMap{
		Map:      *tmxMap,
		tileSets: make(map[string][]TextureRegion, len(tmxMap.Tilesets)),
		filePath: filepath.Dir(path),
	}
	for _, t := range tiledMap.Tilesets {
		texture, err := LoadTexture(tiledMap.filePath + "/" + t.Image.Source)
		if err != nil {
			return nil, err
		}
		tiledMap.tileSets[t.Name] = texture.SplitLine(int32(t.TileWidth), int32(t.TileHeight))
	}
	return &tiledMap, nil
}