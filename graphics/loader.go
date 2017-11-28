package graphics

import (
	"image"
	"fmt"
	"os"
	"image/draw"
	_ "image/png"
	"github.com/go-gl/gl/v3.2-core/gl"
	"io/ioutil"
)

func LoadTexture(file string) (*Texture, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	p := img.Bounds().Size()
	width, height := int32(p.X), int32(p.Y)
	texture, err := NewTexture(width, height, Linear, Linear, ClampToEdge, ClampToEdge)
	if err != nil {
		return nil, err
	}
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		width,
		height,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}

func LoadFile(name string) string {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(buf)
}
