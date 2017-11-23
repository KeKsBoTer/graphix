package graphics

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"errors"
	"fmt"
)

type graphics struct {
	width, height int
	window        *glfw.Window
	fullScreen    bool
}

func (g *graphics) GetWidth() int {
	return g.width
}
func (g *graphics) GetHeight() int {
	return g.height
}

func (g *graphics) IsFullScreen() bool {
	return g.fullScreen
}


//TODO resolution and resizing is not working correctly. Its pretty buggy
func (g *graphics) ToggleFullScreen(isFullScreen bool) error {
	fmt.Errorf("do not use ToggleFullScreen yet")
	if isFullScreen == g.fullScreen {
		return nil
	}

	if isFullScreen {
		monitor := getMonitor(g.window)
		if monitor == nil {
			return errors.New("could not detect the current monitor")
		}
		monitorWidth, monitorHeight := monitor.GetPhysicalSize()
		monitor.GetVideoMode().Width=monitorWidth*2
		monitor.GetVideoMode().Height=monitorHeight*2
		g.window.SetMonitor(monitor,0,0,monitorWidth*2,monitorHeight*2,60)
	} else {
		monitorWidth, monitorHeight := g.window.GetMonitor().GetPhysicalSize()
		g.window.SetSize(g.width,g.height)
		g.window.SetMonitor(
			nil,
			monitorWidth/2+g.width/2,
			monitorHeight/2+g.height/2,
			g.width,
			g.height,
			-1,
		)
	}
	g.fullScreen = isFullScreen
	return nil
}

func getMonitor(w *glfw.Window) *glfw.Monitor {
	return glfw.GetPrimaryMonitor() //TODO return monitor the window is in
}
