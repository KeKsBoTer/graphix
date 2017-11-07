package gx

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"strings"
)

func glfwBool(value bool) int {
	if value {
		return glfw.True
	} else {
		return glfw.False
	}
}

func GlString(text string) string {
	if !strings.HasSuffix(text, "\x00") {
		return text + "\x00"
	}
	return text
}
