package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
)

const VertexSize = 2 + 1 + 2
const SpriteSize = 4 * VertexSize

type SpriteBatch struct {
	defaultShader    ShaderProgram
	customShader     *ShaderProgram
	drawing          bool
	transformMatrix  mgl32.Mat4
	projectionMatrix mgl32.Mat4
	color            float32
	idx              int
	vao              uint32
	vertices         FloatBuffer
	indices          ShortBuffer
	lastTexture      *Texture
}

func NewSpriteBatch() *SpriteBatch {
	batch := new(SpriteBatch)
	vertexShader := LoadFile("base.vert")
	fragmentShader := LoadFile("base.frag")

	// Configure the vertex and fragment shader
	p, err := NewShaderProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}
	batch.defaultShader = *p
	color := Color{0, 0, 0, 1}
	batch.color = color.DecodeFloatRGBA()
	batch.lastTexture = nil

	size := 1000

	batch.idx = 0

	gl.GenVertexArrays(1, &batch.vao)
	gl.BindVertexArray(batch.vao)

	batch.vertices = *NewFloatBuffer(size*SpriteSize, gl.ARRAY_BUFFER)

	program := batch.getActiveShaderProgram()
	vertAttrib := uint32(gl.GetAttribLocation(program.GetId(), gl.Str(GlString("vert"))))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program.GetId(), gl.Str(GlString("vertTexCoord"))))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

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

func (s *SpriteBatch) IsDrawing() bool {
	return s.drawing
}

func (s *SpriteBatch) Dispose() {
	s.defaultShader.dispose()
	s.vertices.Dispose()
	s.indices.Dispose()
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

	s.vertices.Update()
	program := s.getActiveShaderProgram()
	program.Begin()
	gl.BindVertexArray(s.vao)
	program.BindFragDataLocation("outputColor", 0)

	gl.ActiveTexture(gl.TEXTURE0)
	s.lastTexture.Bind()
	s.indices.Bind()
	gl.DrawElements(gl.TRIANGLES, int32(len(s.indices.data)), gl.UNSIGNED_SHORT, gl.PtrOffset(0))
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

	(*vertices)[idx] = vx
	(*vertices)[idx+1] = vy

	(*vertices)[idx+5] = vx
	(*vertices)[idx+6] = fy2

	(*vertices)[idx+10] = fx2
	(*vertices)[idx+11] = fy2

	(*vertices)[idx+15] = fx2
	(*vertices)[idx+16] = vy
	return idx + 17
}

func updateTextureCoords(vertices *[]float32, u, v, u2, v2 float32, idx int) int {
	color := float32(1)
	(*vertices)[idx+2] = color
	(*vertices)[idx+3] = u
	(*vertices)[idx+4] = v
	(*vertices)[idx+7] = color
	(*vertices)[idx+8] = u
	(*vertices)[idx+9] = v2
	(*vertices)[idx+12] = color
	(*vertices)[idx+13] = u2
	(*vertices)[idx+14] = v2
	(*vertices)[idx+17] = color
	(*vertices)[idx+18] = u2
	(*vertices)[idx+19] = v
	return idx + 20
}

func (s *SpriteBatch) DrawTexture(texture *Texture, x, y, width, height float32) {
	if !s.IsDrawing() {
		panic("SpriteBatch.begin must be called before draw")
	}
	if texture != s.lastTexture {
		s.switchTexture(texture)
	} else if s.idx >= len(s.vertices.data) {
		s.Flush()
	}
	updateVertices(&s.vertices.data, x, y, width, height, s.idx)
	s.idx = updateTextureCoords(&s.vertices.data, 0, 0, 1, 1, s.idx)

}

func (s *SpriteBatch) DrawRegion(region TextureRegion, x, y, width, height float32) {
	if !s.IsDrawing() {
		panic("SpriteBatch.begin must be called before draw")
	}
	texture := region.texture
	if &texture != &s.lastTexture {
		s.switchTexture(texture)
	} else if s.idx >= len(s.vertices.data) {
		s.Flush()
	}
	updateVertices(&s.vertices.data, x, y, width, height, s.idx)
	u, v, u2, v2 := region.GetBounds()
	s.idx = updateTextureCoords(&s.vertices.data, u, v, u2, v2, s.idx)
}
