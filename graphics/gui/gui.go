package gui

import (
	"github.com/KeKsBoTer/graphix/utils"
	"github.com/KeKsBoTer/graphix/graphics"
)

//TODO plan it before implementation

type GUI struct {
	children []Component
	batch    *graphics.SpriteBatch
	camera   *graphics.Camera
}

func (g *GUI) Dispose() {
	g.batch.Dispose()
	for _, c := range g.children {
		c.Dispose()
	}
}

func NewGUI(viewportWidth, viewportHeight int) *GUI {
	gui := GUI{
		children: []Component{},
		batch:    graphics.NewSpriteBatch(),
		camera:   graphics.NewCamera(float32(viewportWidth), float32(viewportHeight)),
	}
	gui.Resize(int32(viewportWidth), int32(viewportHeight))
	return &gui
}

func (g *GUI) AddComponent(comp Component) {
	g.children = append(g.children, comp)
}

func (g *GUI) Resize(width, height int32) {
	// Put origin to top left
	g.camera.SetViewport(float32(width), -float32(height))
	g.camera.SetPosition(float32(width)/2, float32(height)/2)

	g.camera.Update()
	g.batch.SetProjectionMatrix(g.camera.GetProjection())
	g.batch.SetTransformationMatrix(g.camera.GetView())
}

func (g *GUI) Render(delta float64) {
	for _, v := range g.children {
		v.Update(delta)
	}
	g.batch.Begin()
	for _, v := range g.children {
		v.Render(g.batch)
	}
	g.batch.End()
}

type Component interface {
	utils.Disposable
	Render(batch *graphics.SpriteBatch)
	Update(delta float64)
	GetPos() (float32, float32)
	GetSize() (float32, float32)
	GetBounds() (float32, float32, float32, float32)
	SetPos(x, y float32)
	SetSize(width, height float32)
}