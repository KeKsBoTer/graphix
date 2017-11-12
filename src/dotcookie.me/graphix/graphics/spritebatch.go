package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type SpriteBatch struct {
	defaultShader    ShaderProgram
	customShader     *ShaderProgram
	drawing          bool
	transformMatrix  mgl32.Mat4
	projectionMatrix mgl32.Mat4
	color            float32
	idx              int
	vao              uint32
}

func NewSpriteBatch() *SpriteBatch {
	batch := new(SpriteBatch)
	vertexShader := LoadFile("base.vert")
	fragmentShader := LoadFile("base.frag")

	// Configure the vertex and fragment shaders
	p, err := NewShaderProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}
	batch.defaultShader = *p
	color := Color{0,0,0,1}
	batch.color = color.DecodeFloatRGBA()

	gl.GenVertexArrays(1, &batch.vao)
	gl.BindVertexArray(batch.vao)

	var vbo uint32
	var x, y, fx2, fy2 float32 = -1, -1, 1, 1
	vertices := []float32{
		x, y, fy2, 1.0, 0.0,     // bottom left
		fx2, y, 1.0, 0.0, 0.0,   // bottom right
		x, fy2, 1.0, 1.0, 1.0,   // top right
		fx2, y, 1.0, 0.0, 0.0,   // bottom right
		fx2, fy2, 1.0, 0.0, 1.0, // top right
		x, fy2, 1.0, 1.0, 1.0,   // top left
	}
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	program := batch.getActiveShaderProgram()
	vertAttrib := uint32(gl.GetAttribLocation(program.GetId(), gl.Str(GlString("vert"))))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program.GetId(), gl.Str(GlString("vertTexCoord"))))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	return batch
}

func (s *SpriteBatch) IsDrawing() bool {
	return s.drawing
}

func (s *SpriteBatch) Dispose() {
	s.defaultShader.dispose()
}

func (s *SpriteBatch) Begin() {
	gl.DepthMask(false)
	s.getActiveShaderProgram().Begin()
	s.setupMatrices()
	s.drawing = true
}

func (s *SpriteBatch) Flush() {
	if s.idx == 0 {
		return // Nothing to flush
	}
	//TODO implement
}

func (s *SpriteBatch) End() {
	if !s.IsDrawing() {
		panic("SpriteBatch.begin must be called before end.")
	}
	if s.idx > 0 {
		s.Flush()
	}
	s.drawing = false
	gl.DepthMask(true)
	s.getActiveShaderProgram().End()

}

func (s *SpriteBatch) getActiveShaderProgram() *ShaderProgram {
	if s.customShader != nil {
		return s.customShader
	} else {
		return &s.defaultShader
	}
}

func (s *SpriteBatch) setupMatrices() {
	program := s.getActiveShaderProgram()
	program.SetUniformMatrix4fv("projection", &s.projectionMatrix, false)
	program.SetUniformMatrix4fv("camera", &s.transformMatrix, false)
	program.SetUniform1i("texture", 0)
}

func (s *SpriteBatch) SetShader(program *ShaderProgram) {
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

func (s *SpriteBatch) GetShader() *ShaderProgram {
	return s.getActiveShaderProgram()
}

func (s *SpriteBatch) SetTransformationMatrix(mat4 mgl32.Mat4) {
	if s.drawing {
		s.Flush()
	}
	s.transformMatrix = mat4
	if s.drawing {
		s.setupMatrices()
	}
}

func (s *SpriteBatch) GetTransformationMatrix() mgl32.Mat4 {
	return s.transformMatrix
}

func (s *SpriteBatch) SetProjectionMatrix(mat4 mgl32.Mat4) {
	if s.drawing {
		s.Flush()
	}
	s.projectionMatrix = mat4
	if s.drawing {
		s.setupMatrices()
	}
}

func (s *SpriteBatch) GetProjectionMatrix() mgl32.Mat4 {
	return s.projectionMatrix
}

func (s *SpriteBatch) SetColor(color Color) {
	s.color = color.DecodeFloatRGBA()
}

func (s *SpriteBatch) GetColor() Color {
	panic("Not implemented yet")
}

func (s *SpriteBatch) DrawTexture(texture Texture, x, y, width, height float32) {
	if !s.IsDrawing() {
		panic("SpriteBatch.begin must be called before draw")
	}
	model := mgl32.Translate3D(x,y,0).Mul4(mgl32.Scale3D(width,height,1))
	program := s.getActiveShaderProgram()
	program.Begin()
	program.SetUniformMatrix4fv("model", &model, false)

	gl.BindVertexArray(s.vao)
	program.BindFragDataLocation("outputColor", 0)

	gl.ActiveTexture(gl.TEXTURE0)
	texture.Bind()

	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)

}

func (s *SpriteBatch) DrawRegion(texture TextureRegion, x, y, width, height float32) {
	panic("implement me")
}
