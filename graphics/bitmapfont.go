package graphics

import (
	"github.com/KeKsBoTer/gofnt"
	"io/ioutil"
	"path/filepath"
	"math"
)

type BitmapFont struct {
	gofnt.Font
	textures []Texture
}

func LoadBitmapFont(path string) (*BitmapFont, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	font, err := gofnt.Parse(string(data))
	if err != nil {
		return nil, err
	}
	btmFont := BitmapFont{
		Font: *font,
	}
	parentPath := filepath.Dir(path)
	btmFont.textures = make([]Texture, len(font.Pages))
	for i, v := range font.Pages {
		if tex, err := LoadTexture(parentPath + "/" + v.File); err != nil {
			return nil, err
		} else {
			btmFont.textures[i] = *tex
		}
	}
	return &btmFont, nil
}

func (f *BitmapFont) Dispose() {
	for _, t := range f.textures {
		t.Dispose()
	}
}

func (f *BitmapFont) Draw(text string, batch *SpriteBatch) {
	x := float32(0)
	y := float32(0)
	for _, v := range text {
		charData := gofnt.Char{Id: -1}
		for _, c := range f.Chars {
			if c.Id == v {
				charData = c
				break
			}
		}
		if charData.Id != -1 {
			region := NewTextureRegion(&f.textures[charData.Page],
				int32(charData.X),
				int32(charData.Y),
				int32(charData.Width),
				int32(charData.Height),
			)
			width := float32(charData.Height) * float32(charData.Width) / float32(charData.Height)
			xOffset := float32(charData.XOffset)
			x-=float32(f.Info.Padding[0]+f.Info.Padding[2])
			if !math.IsNaN(float64(width)) {
				batch.DrawRegion(*region, x+xOffset, float32(f.Common.Base-charData.YOffset)+y, width, -float32(charData.Height))
			}
			x += float32(charData.XAdvanced) + float32(f.Info.Spacing[0])
		}
	}
}
