package graphics

import (
	
	"github.com/go-gl/mathgl/mgl32"
	"dotcookie.me/graphix/utils"
)

type Batch interface {
	utils.Disposable
	Begin()
	IsDrawing() bool
	Flush()
	End()

	SetShader(program ShaderProgram)
	GetShader() *ShaderProgram

	SetTransformationMatrix(mat4 mgl32.Mat4)
	GetTransformationMatrix() mgl32.Mat4

	SetProjectionMatrix(mat4 mgl32.Mat4)
	GetProjectionMatrix() mgl32.Mat4

	SetColor(color Color)
	GetColor() Color

	//Drawing
	DrawTexture(texture Texture,x,y,width,height float32)
	DrawRegion(texture TextureRegion,x,y,width,height float32)
}
