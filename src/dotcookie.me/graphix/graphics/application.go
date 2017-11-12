package graphics

import (
	"log"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"
	"runtime"
	"math"
	_ "fmt"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

type WindowConfig struct {
	Width, Height int
	Samples       int
	Title         string
	Resizeable,
	Fullscreen,
	Vsync bool
}

type Screen interface {
	Show()
	Dispose()
	Render(delta float64)
	Resize(width, height int)
}

type Application struct {
	screen     Screen
	lastRender float64
	Graphics   graphics
	Input      input
}

func (a *Application) Create() {
	if a.screen != nil {
		a.screen.Show()
	}
}

func (a *Application) Dispose() {
	if a.screen != nil {
		a.screen.Dispose()
	}
}

func (a *Application) Render() {
	if a.screen != nil {
		time := glfw.GetTime()
		a.screen.Render(time - a.lastRender) //calculate elapsed time
		a.lastRender = time
	}
}

func (a *Application) Resize(width, height int) {
	if a.screen != nil {
		a.screen.Resize(width, height)
	}
}

func (a *Application) SetScreen(screen Screen) {
	if a.screen != nil {
		a.screen.Dispose() // Dispose old screen
	}
	a.screen = screen
	if a.screen != nil {
		a.screen.Show() // Show new screen
	}
}

func (a *Application) GetScreen() Screen {
	return a.screen
}

func DesktopApplication(config WindowConfig, screen Screen) {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfwBool(config.Resizeable))
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Samples, config.Samples)
	glfw.WindowHint(glfw.Maximized, glfwBool(config.Fullscreen))

	var monitor *glfw.Monitor
	var window *glfw.Window
	var err error
	if config.Fullscreen {
		monitor = glfw.GetPrimaryMonitor()
		vidMode := monitor.GetVideoMode()
		config.Width = vidMode.Width
		config.Height = vidMode.Height
	}
	window, err = glfw.CreateWindow(config.Width, config.Height, config.Title, monitor, nil)
	if err != nil {
		panic(err)
	}

	App = Application{screen: screen}
	App.Graphics.width = config.Width
	App.Graphics.height = config.Height

	window.SetSizeCallback(func(w *glfw.Window, width int, height int) {
		App.Graphics.width = width
		App.Graphics.height = height
		App.Resize(width, height)
	})

	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		x, y := window.GetCursorPos()
		mx := int(math.Floor(x))
		my := int(math.Floor(y))
		App.Input.fireMouseEvent(mx, my, button, action)
	})

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		App.Input.fireKeyEvent(key, action)
	})

	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		mx := int(math.Floor(xpos))
		my := int(math.Floor(ypos))
		App.Input.fireMouseMoveEvent(mx, my)
	})
	window.MakeContextCurrent()
	glfw.SwapInterval(glfwBool(config.Vsync))

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	App.Create()

	for !window.ShouldClose() {
		glfw.PollEvents()
		App.Render()
		// Maintenance
		window.SwapBuffers()
	}
	App.Dispose()
	window.Destroy()
	glfw.Terminate()
}
