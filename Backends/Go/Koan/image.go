package koan

import (
	kinc "github.com/darmie/go-kinc/Sources/kinc"
	g4 "github.com/darmie/go-kinc/Sources/kinc/graphics4"
	"github.com/darmie/koan/graphics1"
	"github.com/darmie/koan/graphics2"
	graphics4 "github.com/darmie/koan/graphics4"
)

const (
	None  ImageCompression = 0
	DXT5  ImageCompression = 1
	ASTC  ImageCompression = 2
	PVRTC ImageCompression = 3
)

type ImageCompression kinc.ImageCompression

type Image struct {
	kinc.Image
	Readable             bool
	RenderTarget         *g4.RenderTarget
	Texture              *graphics4.Texture
	TextureArray         *g4.TextureArray
	TextureArrayTextures []*Texture
	bytes                []byte
}

// CreateImage create new image object
func CreateImage(width int, height int, format *graphics4.TextureFormat, readable bool) *Image {
	if format == nil {
		_f := graphics4.RGBA32
		format = &_f
	}
	data := make([]byte, width*height*int(*format))
	img := kinc.InitImage(data, width, height, format)
	img.Data = data
	img.Texture = g4.InitTexture(width, height, format)
	img.RenderTarget = nil
	img.Readable = readable
}

// CreateImage3D create new image object
func CreateImage3D(width int, height int, depth int, format *graphics4.TextureFormat, readable bool) *Image {
	if format == nil {
		_f := graphics4.RGBA32
		format = &_f
	}

	data := make([]byte, width*height*int(*format))
	img := kinc.InitImage(data, width, height, depth, format)
	img.Data = data
	img.Texture = g4.Init3DTexture(width, height, depth, format)
	img.RenderTarget = nil
	img.Readable = readable
	return img
}

// CreateImageFromBytes create new image object
func CreateImageFromBytes(data []byte, width int, height int, format *graphics4.TextureFormat, readable bool) *Image {
	if format == nil {
		_f := graphics4.RGBA32
		format = &_f
	}
	img := kinc.InitImage(data, width, height, format)
	img.Data = data
	img.Texture = g4.InitTexture(width, height, format)
	img.RenderTarget = nil
	img.Readable = readable
}

// CreateImage3DFromBytes create new image object
func CreateImage3DFromBytes(data []byte, width int, height int, depth int, format *graphics4.TextureFormat, readable bool) *Image {
	if format == nil {
		_f := graphics4.RGBA32
		format = &_f
	}

	img := kinc.InitImage(data, width, height, depth, format)
	img.Data = data
	img.Texture = g4.Init3DTexture(width, height, depth, format)
	img.RenderTarget = nil
	img.Readable = readable
	return img
}

// CreateImageFromFile create new image object
func CreateImageFromFile(filepath string, width int, height int, format *graphics4.TextureFormat) *Image {
	if format == nil {
		_f := graphics4.RGBA32
		format = &_f
	}
	img := kinc.InitImageFromFile(filepath, width, height, *format)
	img.Texture = g4.InitTexture(width, height, format)
	img.RenderTarget = nil
	return img
}

// CreateImage3DFromFile create new image object
func CreateImage3DFromFile(filepath string, width int, height int, depth, format *graphics4.TextureFormat) *Image {
	if format == nil {
		_f := graphics4.RGBA32
		format = &_f
	}
	img := kinc.InitImageFromFile(filepath, width, height, depth, format)
	img.Texture = g4.Init3DTexture(width, height, depth, format)
	img.RenderTarget = nil
	return img
}

// CreateRenderTarget Create a new image instance and set things up so you can render to the image.
func CreateRenderTarget(width int, height int, format graphics4.TextureFormat, depthStencil graphics4.DepthStencilFormat, antiAliasingSamples int, contextID int) *Image {
	renderTarget := kinc.InitRenderTarget(width, height, getDepthBufferBits(depthStencil), antiAliasingSamples, getRenderTargetFormat(format), getStencilBufferBits(depthStencil), contextID)
	img := kinc.InitImage(data, width, height, format)
	img.RenderTarget = renderTarget
	img.Texture = g4.InitTexture(width, height, format)
	return img
}

func CreateArray(images []*Image, format *graphics4.TextureFormat) {
	if format == nil {
		_f := graphics4.RGBA32
		format = &_f
	}
	img := kinc.InitImage(data, width, height, format)
	img.Texture = g4.InitTexture(width, height, format)
	img.RenderTarget = nil

	initArrayTexture(img, images)

}

func initArrayTexture(source *Image, images []*Image) {
	source.TextureArrayTextures = make([]*kinc.Image, len(images))
	for i, img := range images {
		source.TextureArrayTextures[i] = (*kinc.Image)images[i].Texture
	}
	source.TextureArray = g4.InitTexureArray(source.TextureArrayTextures, len(images))
}

func (img *Image) GetDepth() int {
	if img.Texture != nil {
		return img.Texture.Depth
	}
	return 0
}

// GetWidth of the canvas in pixels.
func (img *Image) GetWidth() int {
	if img.Texture != nil {
		return img.Texture.Width
	}
	return img.RenderTarget.Width
}

// GetHeight of the canvas in pixels.
func (img *Image) GetHeight() int {
	if img.Texture != nil {
		return img.Texture.TexHeight
	}
	return img.RenderTarget.Height
}

// G1 Graphics1 interface object.<br>
// Basic setPixel operation.
func (img *Image) G1() *graphics1.Graphics {
	return nil
}

// G2 Graphics2 interface object for drawing 2D graphics
func (img *Image) G2() *graphics2.Graphics {
	return nil
}

// G4 graphics4 interface object for 3D drawing
func (img *Image) G4() *graphics4.Graphics {
	return &graphics4.Graphics{}
}

// Unload the resource from memory. Normally called by the Loader.
func (img *Image) Unload() {
	img.Texture = nil
	img.RenderTarget = nil
	img.TextureArray = nil
	img.TextureArrayTextures = nil
}

func (img *Image) Lock() []byte {
	size := img.SizeOf(img.Texture.Format) * img.Texture.Width * img.Texture.Height
	img.bytes = make([]byte, size)
	return img.bytes
}

func (img *Image) Unlock() {
	b := img.bytes
	tex := img.Texture.Lock()
	stride := img.Texture.Stride()
	for y := 0; y < img.Texture.Height; y++ {
		for x := 0; x < img.Texture.Width; x++ {
			for i := 0; i < size; i++ {
				// Todo: case #KORE_DIRECT3D
				tex[y * stride + x * size + i] = b[(y * texture->width + x) * size + i];
			}
		}
	}
	img.Texture.Unlock()
}


func getDepthBufferBits(depthAndStencil graphics4.DepthStencilFormat) int {
	switch depthAndStencil {
	case NoDepthAndStencil:
		return -1
	case DepthOnly:
		return 24
	case DepthAutoStencilAuto:
		return 24
	case Depth24Stencil8:
		return 24
	case Depth32Stencil8:
		return 32
	case Depth16:
		return 16
	}

	return 0
}

func getStencilBufferBits(depthAndStencil graphics4.DepthStencilFormat) int {
	switch depthAndStencil {
	case NoDepthAndStencil:
		return -1
	case DepthOnly:
		return -1
	case DepthAutoStencilAuto:
		return 8
	case Depth24Stencil8:
		return 8
	case Depth32Stencil8:
		return 8
	case Depth16:
		return 0
	}

	return 0
}

func getTextureFormat(format graphics4.TextureFormat) int {
	switch format {
	case RGBA32:
		return 0
	case RGBA128:
		return 3
	case RGBA64:
		return 4
	case A32:
		return 5
	case A16:
		return 7
	default:
		return 1 // Grey8
	}
}

func getRenderTargetFormat(format graphics4.TextureFormat) int {
	switch format {
	case RGBA32: // Target32Bit
		return 0
	case RGBA64: // Target64BitFloat
		return 1
	case A32: // Target32BitRedFloat
		return 2
	case RGBA128: // Target128BitFloat
		return 3
	case DEPTH16: // Target16BitDepth
		return 4
	case L8:
		return 5 // Target8BitRed
	case A16:
		return 6 // Target16BitRedFloat
	default:
		return 0
	}
}
