package graphics

import "github.com/go-gl/glfw/v3.2/glfw"

type graphics struct {
	width,height int
	Window *glfw.Window
}

func (g graphics) GetWidth() int{
	return g.width
}
func (g graphics) GetHeight() int{
	return g.height
}
