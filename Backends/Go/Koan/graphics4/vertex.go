package graphics4

import g4 "github.com/darmie/go-kinc/Sources/kinc/graphics4"

const (
	None       VertexData = 0
	Float1     VertexData = 1
	Float2     VertexData = 2
	Float3     VertexData = 3
	Float4     VertexData = 4
	Float4x4   VertexData = 5
	Short2Norm VertexData = 6
	Short4Norm VertexData = 7
	Color      VertexData = 8
)

type VertexData g4.VertexData


type VertexElement struct {
	g4.VertexElement
}


type VertexStructure struct {
	g4.VertexStructure
}


type VertexBuffer  {
	g4.VertexBuffer
}


func InitVertexBuffer(count int, structure *VertexStructure, usage Usage, instanceDataStepRate int) *VertexBuffer {
	v := g4.InitVertexBuffer(count, structure, usage, instanceDataStepRate)
	return v
}


func InitVertexStructure() *VertexStructure {
	_structure := g4.InitVertexStructure()
	return _structure
}


func (vs *VertexStructure) Get(index int) *VertexElement {
	return vs.Elements[i]
}