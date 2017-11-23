package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type ShapeRenderer struct {
	defaultShader    ShaderProgram
	customShader     *ShaderProgram
	drawing          bool
	transformMatrix  mgl32.Mat4
	projectionMatrix mgl32.Mat4
	color            float32
	idx              int32
	vao              uint32
	vertices         FloatBuffer
	indices          ShortBuffer
}

func NewShapeRenderer() *ShapeRenderer {
	return NewShapeRendererS(1395)
}

func NewShapeRendererS(size int) *ShapeRenderer {
	batch := new(ShapeRenderer)
	vertexShader := LoadFile("shader/shapeRenderer.vert")
	fragmentShader := LoadFile("shader/shapeRenderer.frag")

	// Configure the vertex and fragment shader
	p, err := NewShaderProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}
	batch.defaultShader = *p
	black := Color{0, 0, 0, 1}
	batch.color = black.Pack() //white

	batch.idx = 0

	gl.GenVertexArrays(1, &batch.vao)
	gl.BindVertexArray(batch.vao)

	batch.vertices = *NewFloatBuffer(size*SpriteSize, gl.ARRAY_BUFFER)

	program := batch.getActiveShaderProgram()
	vertAttrib := uint32(gl.GetAttribLocation(program.GetId(), gl.Str(GlString("vert"))))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 2, gl.FLOAT, false, 3*4, gl.PtrOffset(0))

	colorAttrib := uint32(gl.GetAttribLocation(program.GetId(), gl.Str(GlString("a_color"))))
	gl.EnableVertexAttribArray(colorAttrib)
	gl.VertexAttribPointer(colorAttrib, 4, gl.UNSIGNED_BYTE, true, 3*4, gl.PtrOffset(2*4))

	length := size * 6
	indices := make([]uint16, length)

	var j uint16 = 0
	for i := 0; i < length; i, j = i+6, j+4 {
		indices[i] = j
		indices[i+1] = j + 1
		indices[i+2] = j + 2
		indices[i+3] = j + 2
		indices[i+4] = j + 3
		indices[i+5] = j
	}
	batch.indices = *NewShortBufferFromData(indices, gl.ELEMENT_ARRAY_BUFFER)
	return batch
}

func (s *ShapeRenderer) IsDrawing() bool {
	return s.drawing
}

func (s *ShapeRenderer) Dispose() {
	s.defaultShader.dispose()
	s.vertices.Dispose()
	s.indices.Dispose()
	gl.DeleteVertexArrays(0, &s.vao)
}

func (s *ShapeRenderer) Begin() {
	gl.DepthMask(false)
	s.getActiveShaderProgram().Begin()
	s.setupMatrices()
	gl.DepthMask(false)
	s.drawing = true
}

func (s *ShapeRenderer) Flush() {
	if s.idx == 0 {
		return // Nothing to flush
	}

	s.vertices.Update()
	program := s.getActiveShaderProgram()
	program.Begin()
	gl.BindVertexArray(s.vao)
	program.BindFragDataLocation("outputColor", 0)

	s.indices.Bind()
	gl.DrawElements(gl.TRIANGLES, s.idx/12*6, gl.UNSIGNED_SHORT, gl.PtrOffset(0))
	s.idx = 0
}

func (s *ShapeRenderer) End() {
	if !s.IsDrawing() {
		panic("ShapeRenderer.begin must be called before end.")
	}
	if s.idx > 0 {
		s.Flush()
	}
	s.drawing = false
	gl.DepthMask(true)
	s.getActiveShaderProgram().End()

}

func (s *ShapeRenderer) getActiveShaderProgram() *ShaderProgram {
	if s.customShader != nil {
		return s.customShader
	} else {
		return &s.defaultShader
	}
}

func (s *ShapeRenderer) setupMatrices() {
	program := s.getActiveShaderProgram()
	program.SetUniformMatrix4fv("projection", &s.projectionMatrix, false)
	program.SetUniformMatrix4fv("camera", &s.transformMatrix, false)
}

func (s *ShapeRenderer) SetShader(program *ShaderProgram) {
	if s.drawing {
		s.Flush()
		s.getActiveShaderProgram().End()
	}
	s.customShader = program
	if s.drawing {
		s.getActiveShaderProgram().Begin()
		s.setupMatrices()
	}
}

func (s *ShapeRenderer) GetShader() *ShaderProgram {
	return s.getActiveShaderProgram()
}

func (s *ShapeRenderer) SetTransformationMatrix(mat4 mgl32.Mat4) {
	if s.drawing {
		s.Flush()
	}
	s.transformMatrix = mat4
	if s.drawing {
		s.setupMatrices()
	}
}

func (s *ShapeRenderer) GetTransformationMatrix() mgl32.Mat4 {
	return s.transformMatrix
}

func (s *ShapeRenderer) SetProjectionMatrix(mat4 mgl32.Mat4) {
	if s.drawing {
		s.Flush()
	}
	s.projectionMatrix = mat4
	if s.drawing {
		s.setupMatrices()
	}
}

func (s *ShapeRenderer) GetProjectionMatrix() mgl32.Mat4 {
	return s.projectionMatrix
}

func (s *ShapeRenderer) SetColor(c Color) {
	s.color = c.Pack()
}

func (s *ShapeRenderer) GetColor() Color {
	panic("Not implemented yet")
}

func (s *ShapeRenderer) DrawRectangle(x, y, width, height float32) {
	if !s.IsDrawing() {
		panic("ShapeRenderer.begin must be called before draw")
	}
	if s.idx >= int32(len(s.vertices.data)) {
		s.Flush()
	}
	var vx, vy, fx2, fy2 = x, y, x+width, y+height
	vertices := s.vertices.data

	idx := s.idx
	vertices[idx] = vx
	vertices[idx+1] = vy
	vertices[idx+2] = s.color

	vertices[idx+3] = vx
	vertices[idx+4] = fy2
	vertices[idx+5] = s.color

	vertices[idx+6] = fx2
	vertices[idx+7] = fy2
	vertices[idx+8] = s.color

	vertices[idx+9] = fx2
	vertices[idx+10] = vy
	vertices[idx+11] = s.color

	s.idx += 12
}
