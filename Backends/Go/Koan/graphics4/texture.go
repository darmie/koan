package graphics4

import (
	kinc "github.com/darmie/go-kinc/Sources/kinc"
	g4 "github.com/darmie/go-kinc/Sources/kinc/graphics4"
)

type TextureUnit g4.TextureUnit

const (
	RGBA32  TextureFormat = 0
	GREY8   TextureFormat = 1
	RGB24   TextureFormat = 2
	RGBA128 TextureFormat = 3
	RGBA64  TextureFormat = 4
	A32     TextureFormat = 5
	BGRA32  TextureFormat = 6
	A16     TextureFormat = 7
)

type TextureFormat kinc.ImageFormat

type Texture struct {
	g4.Texture
}
