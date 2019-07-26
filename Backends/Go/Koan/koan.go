package koan

// +build cgo

import (
	"fmt"

	"github.com/darmie/go-kinc/Sources/kinc"
	"github.com/darmie/go-kinc/Sources/kinc/graphics4"
)

var win *kinc.Window
var pipe *graphics4.Pipeline
var vertices *graphics4.VertexBuffer
var indices *graphics4.IndexBuffer

func update() {
	fmt.Println("update")

	graphics4.Begin(window.Index)
	graphics4.Clear(graphics4.ClearColor, 0, graphics4.ClearDepth, graphics4.ClearStencil)

	graphics4.SetPipeline(pipe)
	graphics4.SetVertexBuffer(vertices)
	graphics4.SetIndexBuffer(indices)
	graphics4.DrawIndexedVertices()

	graphics4.End(window.Index)
	graphics4.SwapBuffers()
}

func main() {
	win = kinc.Init("Shader", 1024, 768, nil, nil)

	kinc.SetUpdateCallback(update)
	vertexShader := graphics4.InitShader("shader.vert", graphics4.Vertex)
	fragmentShader := graphics4.InitShader("shader.frag", graphics4.Fragment)

	structure := graphics4.InitVertexStructure()
	structure.Add("pos", graphics4.Float3)

	pipe = graphics4.InitPipeline()
	pipe.InputLayout[0] = structure
	pipe.InputLayout[1] = nil
	pipe.VertexShader = vertexShader
	pipe.FragmentShader = fragmentShader
	pipe.Compile()

	vertices = graphics4.InitVertexBuffer(3, structure, graphics4.Static, 0)
	defer vertices.Unlock(3)
	_v := make([]float32, 9)
	_v[0] = -1
	_v[1] = -1
	_v[2] = 0.5
	_v[3] = 1
	_v[4] = -1
	_v[5] = 0.5
	_v[6] = -1
	_v[7] = 1
	_v[8] = 0.5
	vertices.Lock(_v, 0, 3)

	indices = graphics4.InitIndexBuffer(3)
	defer indices.Unlock()
	_i := make([]float32, 3)
	_i[0] = 0
	_i[1] = 1
	_i[2] = 2
	indices.Lock(_i)

	kinc.Start()
}
