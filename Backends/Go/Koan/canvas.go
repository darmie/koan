package koan

import (
	"github.com/darmie/koan/graphics1"
	"github.com/darmie/koan/graphics2"
	"github.com/darmie/koan/graphics4"
)

type Resource interface {
	// Unload the resource from memory. Normally called by the Loader.
	Unload()
}

type Canvas interface {
	Resource
	// GetWidth of the canvas in pixels.
	GetWidth() int
	// GetHeight of the canvas in pixels.
	GetHeight() int
	// G1 Graphics1 interface object.<br>
	// Basic setPixel operation.
	G1() *graphics1.Graphics

	// G2 Graphics2 interface object for drawing 2D graphics
	G2() *graphics2.Graphics

	// G4 graphics4 interface object for 3D drawing
	G4() *graphics4.Graphics
}
