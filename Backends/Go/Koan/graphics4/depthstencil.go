package graphics4

const (
	NoDepthAndStencil    DepthStencilFormat = 0
	DepthOnly            DepthStencilFormat = 1
	DepthAutoStencilAuto DepthStencilFormat = 2

	// This is platform specific, use with care!
	Depth24Stencil8 DepthStencilFormat = 3
	Depth32Stencil8 DepthStencilFormat = 4
	Depth16         DepthStencilFormat = 5
)

type DepthStencilFormat int
