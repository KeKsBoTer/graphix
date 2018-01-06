package graphics

import (
	gofnt "github.com/KeKsBoTer/gofnt/json"
	"io/ioutil"
	"path/filepath"
	"math"
	"github.com/go-gl/mathgl/mgl32"
)

type BitmapFont struct {
	gofnt.Font
	textures []Texture
	mapping  map[int32]int
}

func (f *BitmapFont) GetLineHeight() float32 {
	return f.Common.Base
}

func (f *BitmapFont) GetWidth(text string) (x float32) {
	for _, v := range text {
		charData := f.getCharData(v)
		x += float32(charData.XAdvanced)
	}
	return
}

func LoadBitmapFont(path string) (*BitmapFont, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	font, err := gofnt.Parse(data)
	if err != nil {
		return nil, err
	}
	btmFont := BitmapFont{
		Font:    *font,
		mapping: map[int32]int{},
	}
	for i, c := range btmFont.Font.Chars {
		btmFont.mapping[c.Id] = i
	}
	parentPath := filepath.Dir(path)
	btmFont.textures = make([]Texture, len(font.Pages))
	for i, v := range font.Pages {
		var texPath string
		if filepath.IsAbs(v) {
			texPath = v
		} else {
			texPath = parentPath + "/" + v
		}
		if tex, err := LoadTexture(texPath); err != nil {
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

func (f *BitmapFont) getCharData(c int32) gofnt.Char {
	if mp, ok := f.mapping[c]; ok {
		return f.Chars[mp]
	} else {
		return f.Chars[f.mapping[0]]
	}
}

// Returns the size, the font rendering took
func (f *BitmapFont) Draw(text string, x, y float32, batch *SpriteBatch) (float32, float32) {

	combined := batch.projectionMatrix.Mul4(batch.transformMatrix)

	startX := x
	for _, v := range text {
		if v == '\n' {
			x = startX
			y -= float32(f.Info.Size*f.Info.StretchH) / 100.0
			continue
		}
		charData := f.getCharData(v)
		region := NewTextureRegion(&f.textures[charData.Page],
			int32(charData.X),
			int32(charData.Y),
			int32(charData.Width),
			int32(charData.Height),
		)
		width := float32(charData.Height) * float32(charData.Width) / float32(charData.Height)
		xOffset := float32(charData.XOffset)
		if !math.IsNaN(float64(width)) {
			xPos := x + xOffset
			yPos := y + float32(f.Common.Base-charData.YOffset)

			lowerLeft := combined.Mul4x1(mgl32.Vec4{xPos, yPos - float32(charData.Height), 0, 1})
			upperRight := combined.Mul4x1(mgl32.Vec4{xPos + width, yPos, 0, 1})

			// Check if character is visible
			if lowerLeft.X() <= 1 && lowerLeft.Y() <= 1 && upperRight.X() >= -1 && upperRight.Y() >= -1 {
				batch.DrawRegion(*region, xPos, yPos, width, -float32(charData.Height))
			}
		}
		x += float32(charData.XAdvanced)
	}
	return x, y
}
