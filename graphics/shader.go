package graphics

import (
	"strings"
	"fmt"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"errors"
)

type ShaderProgram struct {
	id                           uint32
	fragmentShader, vertexShader uint32
}

func (p *ShaderProgram) Begin() {
	gl.UseProgram(p.id)
}

func (p *ShaderProgram) End() {
	gl.UseProgram(0)
}

func (p *ShaderProgram) dispose() {
	gl.DeleteProgram(p.id)
}

func (p *ShaderProgram) GetId() uint32 {
	return p.id
}

func (p *ShaderProgram) findUniform(location string, onFound func(id int32)) error {
	gl.UseProgram(p.id)
	if uniform := gl.GetUniformLocation(p.id, gl.Str(GlString(location))); uniform != gl.INVALID_VALUE && uniform != gl.INVALID_OPERATION && uniform != -1 {
		onFound(uniform)
		return nil
	} else {
		switch uniform {
		case -1:
			return errors.New("failed to find uniform location '" + location + "'")
		default:
			return errors.New("the shader-program ist not valid")
		}
	}
}

func (p *ShaderProgram) SetUniformMatrix4fv(location string, mat *mgl32.Mat4, transpose bool) error {
	return p.findUniform(location, func(id int32) {
		gl.UniformMatrix4fv(id, 1, transpose, &(mat[0]))
	})
}

func (p *ShaderProgram) BindFragDataLocation(location string, color uint32) {
	gl.UseProgram(p.id)
	gl.BindFragDataLocation(p.id, color, gl.Str(GlString(location)))
}

func (p *ShaderProgram) SetUniform1i(location string, value int32) error {
	return p.findUniform(location, func(id int32) {
		gl.Uniform1i(id, value)
	})
}

func NewShaderProgram(vertexShaderSource, fragmentShaderSource string) (*ShaderProgram, error) {
	vertexShaderSource = GlString(vertexShaderSource)
	fragmentShaderSource = GlString(fragmentShaderSource)
	shader := new(ShaderProgram)
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}
	shader.vertexShader = vertexShader

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}
	shader.fragmentShader = fragmentShader

	shader.id = gl.CreateProgram()

	gl.AttachShader(shader.id, vertexShader)
	gl.AttachShader(shader.id, fragmentShader)
	gl.LinkProgram(shader.id)

	var status int32
	gl.GetProgramiv(shader.id, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shader.id, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shader.id, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return shader, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	cSources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, cSources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
