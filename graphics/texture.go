package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type TextureFilter int32
type TextureWrap int32

const (
	Linear  TextureFilter = gl.LINEAR  //The key or button was released.
	MipMap  TextureFilter = gl.MIPMAP  //The key or button was pressed.
	Nearest TextureFilter = gl.NEAREST //The key was held down until it repeated.

	ClampToEdge    TextureWrap = gl.CLAMP_TO_EDGE
	MirroredRepeat TextureWrap = gl.MIRRORED_REPEAT
	Repeat         TextureWrap = gl.REPEAT
)

type Texture struct {
	id                   uint32
	width, height        int32
	filterMin, filterMax TextureFilter
	wrapS, wrapT         TextureWrap
}

func (t *Texture) GetId() uint32 {
	return t.id
}

func NewTexture(width, height int32, filterMin, filterMax TextureFilter, wrapS, warpT TextureWrap) (*Texture, error) {
	tex := Texture{
		width:     width,
		height:    height,
		filterMin: filterMin,
		filterMax: filterMax,
		wrapS:     wrapS,
		wrapT:     warpT}

	gl.GenTextures(1, &tex.id)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex.id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(tex.filterMin))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(tex.filterMax))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, int32(tex.wrapS))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, int32(tex.wrapT))
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	return &tex, nil
}

func (t *Texture) GetWidth() int32 {
	return t.width
}

func (t *Texture) GetHeight() int32 {
	return t.height
}

func (t *Texture) GetSize() (int32, int32) {
	return t.width, t.height
}

func (t *Texture) Dispose() {
	gl.DeleteTextures(1, &t.id)
}

func (t *Texture) ToRegion() *TextureRegion {
	return NewTextureRegion(t, 0, 0, t.width, t.height)
}

func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.id)
}

func (t *Texture) Split(width, height int32) [][]TextureRegion {
	var tWidth, tHeight int32 = t.GetSize()
	regions := make([][]TextureRegion, tHeight/height)
	for i := 0; i < len(regions); i++ {
		regions[i] = make([]TextureRegion, tWidth/width)
		for j := 0; j < len(regions[i]); j++ {
			regions[i][j] = *NewTextureRegion(t, int32(j)*width, int32(i)*height, width, height)
		}
	}
	return regions
}

func (t *Texture) SplitLine(width, height int32) []TextureRegion {
	var tWidth, tHeight int32 = t.GetSize()
	regions := make([]TextureRegion, tWidth/width*tHeight/height)
	xSize, ySize := int(tWidth/width), int(tHeight/height)
	for i := 0; i < ySize; i++ {
		for j := 0; j < xSize; j++ {
			regions[i*xSize+j] = *NewTextureRegion(t, int32(j)*width, int32(i)*height, width, height)
		}
	}
	return regions
}
