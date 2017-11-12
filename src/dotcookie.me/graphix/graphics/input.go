package graphics

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type input struct {
	mouseListeners []MouseListener
	keyListeners   []KeyListener
}

func (i *input) AddMouseListener(listener MouseListener) {
	i.mouseListeners = append(i.mouseListeners, listener)
}

func (i *input) fireMouseEvent(x, y int, button glfw.MouseButton, action glfw.Action) {
	if x < 0 || y < 0 {
		return
	}
	for _, e := range App.Input.mouseListeners {
		if action == glfw.Press {
			if e.MousePressed(x, y, button) {
				return
			}
		} else if action == glfw.Release {
			if e.MouseReleased(x, y, button) {
				return
			}
		}
	}
}

func (i *input) fireMouseMoveEvent(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for _, e := range App.Input.mouseListeners {
		e.MouseMoved(x, y)
	}
}

func (i *input) RemoveMouseListener(listener MouseListener) {
	for index, e := range i.mouseListeners {
		if e == listener {
			i.mouseListeners = append(i.mouseListeners[:index], i.mouseListeners[index+1:]...)
			return
		}
	}
}

func (i *input) fireKeyEvent(key glfw.Key, action glfw.Action) {
	if action == glfw.Press {
		for _, e := range App.Input.keyListeners {
			e.KeyPressed(key)
		}
	} else if action == glfw.Release {
		for _, e := range App.Input.keyListeners {
			e.KeyReleased(key)
		}
	}
}
func (i *input) AddKeyListener(listener KeyListener) {
	i.keyListeners = append(i.keyListeners, listener)
}

func (i *input) RemoveKeyListener(listener KeyListener) {
	for index, e := range i.keyListeners {
		if e == listener {
			i.keyListeners = append(i.keyListeners[:index], i.keyListeners[index+1:]...)
			return
		}
	}
}

type MouseListener interface {
	MouseMoved(x, y int) bool
	MousePressed(x, y int, button glfw.MouseButton) bool //TODO replace glfw with own
	MouseReleased(x, y int, button glfw.MouseButton) bool
}

type KeyListener interface {
	KeyPressed(key glfw.Key)
	KeyReleased(key glfw.Key)
}
