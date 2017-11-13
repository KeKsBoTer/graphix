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
	vao, vbo         uint32
	vertices         []float32
	lastTexture      *Texture
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
	color := Color{0, 0, 0, 1}
	batch.color = color.DecodeFloatRGBA()
	batch.lastTexture = nil

	gl.GenVertexArrays(1, &batch.vao)
	gl.BindVertexArray(batch.vao)

	perRectangle := 30
	size := 1000 * perRectangle
	batch.vertices = make([]float32, size)

	batch.idx = 0
	gl.GenBuffers(1, &batch.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, batch.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(batch.vertices)*4, gl.Ptr(batch.vertices), gl.STATIC_DRAW)

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
	gl.DeleteBuffers(0, &s.vbo)
	gl.DeleteVertexArrays(0, &s.vao)
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

	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(s.vertices)*4, gl.Ptr(s.vertices))

	program := s.getActiveShaderProgram()
	program.Begin()
	gl.BindVertexArray(s.vao)
	program.BindFragDataLocation("outputColor", 0)

	gl.ActiveTexture(gl.TEXTURE0)
	s.lastTexture.Bind()

	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
	s.idx = 0
}

func (s *SpriteBatch) switchTexture(texture *Texture) {
	s.Flush()
	s.lastTexture = texture
}

func (s *SpriteBatch) End() {
	if !s.IsDrawing() {
		panic("SpriteBatch.begin must be called before end.")
	}
	if s.idx > 0 {
		s.Flush()
	}
	s.lastTexture = nil
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

// returns next free element in array
func updateVertices(vertices *[]float32, x, y, width, height float32, idx int) int {
	var vx, vy, fx2, fy2 = x, y, x+width, y+height
	(*vertices)[idx+0] = vx
	(*vertices)[idx+1] = vy
	(*vertices)[idx+2] = 1

	(*vertices)[idx+5] = fx2
	(*vertices)[idx+6] = vy
	(*vertices)[idx+7] = 1

	(*vertices)[idx+10] = vx
	(*vertices)[idx+11] = fy2
	(*vertices)[idx+12] = 1

	(*vertices)[idx+15] = fx2
	(*vertices)[idx+16] = vy
	(*vertices)[idx+17] = 1

	(*vertices)[idx+20] = fx2
	(*vertices)[idx+21] = fy2
	(*vertices)[idx+22] = 1

	(*vertices)[idx+25] = vx
	(*vertices)[idx+26] = fy2
	(*vertices)[idx+27] = 1
	return idx + 28
}

func updateTextureCoords(vertices *[]float32, u, v, u2, v2 float32, idx int) int {
	(*vertices)[idx+3] = u
	(*vertices)[idx+4] = v2

	(*vertices)[idx+8] = u2
	(*vertices)[idx+9] = v2

	(*vertices)[idx+13] = u
	(*vertices)[idx+14] = v

	(*vertices)[idx+18] = u2
	(*vertices)[idx+19] = v2

	(*vertices)[idx+23] = u2
	(*vertices)[idx+24] = v

	(*vertices)[idx+28] = u
	(*vertices)[idx+29] = v
	return idx + 30
}

func (s *SpriteBatch) DrawTexture(texture *Texture, x, y, width, height float32) {
	if !s.IsDrawing() {
		panic("SpriteBatch.begin must be called before draw")
	}
	if texture != s.lastTexture {
		s.switchTexture(texture)
	} else if s.idx >= len(s.vertices) {
		s.Flush()
	}
	updateVertices(&s.vertices, x, y, width, height, s.idx)
	s.idx = updateTextureCoords(&s.vertices, 0, 0, 1, 1, s.idx)

}

func (s *SpriteBatch) DrawRegion(region TextureRegion, x, y, width, height float32) {
	if !s.IsDrawing() {
		panic("SpriteBatch.begin must be called before draw")
	}
	texture := region.texture
	if texture != s.lastTexture {
		s.switchTexture(texture)
	} else if s.idx >= len(s.vertices) {
		s.Flush()
	}
	updateVertices(&s.vertices, x, y, width, height, s.idx)
	u, v, u2, v2 := region.GetBounds()
	s.idx = updateTextureCoords(&s.vertices, u, v, u2, v2, s.idx)

}
