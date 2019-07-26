package graphics4

const (
	Undefined               BlendingFactor = 0
	BlendOne                BlendingFactor = 1
	BlendZero               BlendingFactor = 2
	SourceAlpha             BlendingFactor = 3
	DestinationAlpha        BlendingFactor = 4
	InverseSourceAlpha      BlendingFactor = 5
	InverseDestinationAlpha BlendingFactor = 6
	SourceColor             BlendingFactor = 7
	DestinationColor        BlendingFactor = 8
	InverseSourceColor      BlendingFactor = 9
	InverseDestinationColor BlendingFactor = 10
)

type BlendingFactor int

const (
	Add             BlendingOperation = 0
	Subtract        BlendingOperation = 1
	ReverseSubtract BlendingOperation = 2
	Min             BlendingOperation = 3
	Max             BlendingOperation = 4
)

type BlendingOperation int
