package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"unsafe"
	"fmt"
)

type Buffer interface {
	create()
	Bind()
	Unbind()
	Update()
	Dispose()
}

const shortSize = int(unsafe.Sizeof(uint16(0)))

type ShortBuffer struct {
	id     uint32
	data   []uint16
	glType uint32
}

func NewShortBuffer(length int, glType uint32) *ShortBuffer {
	buffer := new(ShortBuffer)
	buffer.glType = glType
	buffer.data = make([]uint16, length)
	buffer.create()
	return buffer
}

func NewShortBufferFromData(data []uint16, glType uint32) *ShortBuffer {
	buffer := new(ShortBuffer)
	buffer.glType = glType
	buffer.data = data
	buffer.create()
	return buffer
}


func (b *ShortBuffer) create() {
	gl.GenBuffers(1, &b.id)
	b.Bind()
	gl.BufferData(b.glType, len(b.data)*shortSize, gl.Ptr(b.data), gl.STATIC_DRAW)
}

func (b *ShortBuffer) Update() {
	b.Bind()
	gl.BufferSubData(b.glType, 0, len(b.data)*shortSize, gl.Ptr(b.data))
}

func (b *ShortBuffer) Bind() {
	gl.BindBuffer(b.glType, b.id)
}

func (b *ShortBuffer) Unbind() {
	gl.BindBuffer(b.glType, 0)
}
func (b *ShortBuffer) Dispose() {
	gl.DeleteBuffers(0,&b.id)
}

func (b *ShortBuffer) GetData() *[]uint16 {
	return &b.data
}

const floatSize = int(unsafe.Sizeof(float32(0)))

type FloatBuffer struct {
	id     uint32
	data   []float32
	glType uint32

}

func (b *FloatBuffer) create() {
	gl.GenBuffers(1, &b.id)
	b.Bind()
	fmt.Println(b.id, b.data)
	gl.BufferData(b.glType, len(b.data)*floatSize, gl.Ptr(b.data), gl.STATIC_DRAW)
}

func (b *FloatBuffer) Bind() {
	gl.BindBuffer(b.glType,b.id)
}

func (b *FloatBuffer) Unbind() {
	gl.BindBuffer(b.glType,0)
}

func (b *FloatBuffer) Update() {
	b.Bind()
	gl.BufferSubData(b.glType, 0, len(b.data)*floatSize, gl.Ptr(b.data))
}

func (b *FloatBuffer) Dispose() {
	gl.DeleteBuffers(0,&b.id)
}


func (b *FloatBuffer) GetData() *[]float32{
	return &b.data
}

func NewFloatBuffer(length int,glType uint32) *FloatBuffer{
	buffer := new(FloatBuffer)
	buffer.glType = glType
	buffer.data = make([]float32,length)
	buffer.create()
	return buffer
}