// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 11 Jan 2018 09:58:29 CST.
// By https://git.io/c-for-go. DO NOT EDIT.

package nvg

/*
#cgo LDFLAGS: -lm -lGLESv2 -L/opt/vc/lib
#cgo CFLAGS: -I/opt/vc/include
#include "nanovg.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// Pi as defined in nvg/nanovg.h:26
	Pi = 3.14159265358979323846264338327
)

// Winding as declared in nvg/nanovg.h:56
type Winding int32

// Winding enumeration from nvg/nanovg.h:56
const (
	Ccw = 1
	Cw  = 2
)

// Solidity as declared in nvg/nanovg.h:61
type Solidity int32

// Solidity enumeration from nvg/nanovg.h:61
const (
	Solid = 1
	Hole  = 2
)

// LineCapKind as declared in nvg/nanovg.h:66
type LineCapKind int32

// LineCapKind enumeration from nvg/nanovg.h:66
const (
	Butt   = iota
	Round  = 1
	Square = 2
	Bevel  = 3
	Miter  = 4
)

// Align as declared in nvg/nanovg.h:74
type Align int32

// Align enumeration from nvg/nanovg.h:74
const (
	AlignLeft     = 1 << 0
	AlignCenter   = 1 << 1
	AlignRight    = 1 << 2
	AlignTop      = 1 << 3
	AlignMiddle   = 1 << 4
	AlignBottom   = 1 << 5
	AlignBaseline = 1 << 6
)

// BlendFactor as declared in nvg/nanovg.h:86
type BlendFactor int32

// BlendFactor enumeration from nvg/nanovg.h:86
const (
	Zero             = 1 << 0
	One              = 1 << 1
	SrcColor         = 1 << 2
	OneMinusSrcColor = 1 << 3
	DstColor         = 1 << 4
	OneMinusDstColor = 1 << 5
	SrcAlpha         = 1 << 6
	OneMinusSrcAlpha = 1 << 7
	DstAlpha         = 1 << 8
	OneMinusDstAlpha = 1 << 9
	SrcAlphaSaturate = 1 << 10
)

// CompositeOperation as declared in nvg/nanovg.h:100
type CompositeOperation int32

// CompositeOperation enumeration from nvg/nanovg.h:100
const (
	SourceOver      = iota
	SourceIn        = 1
	SourceOut       = 2
	Atop            = 3
	DestinationOver = 4
	DestinationIn   = 5
	DestinationOut  = 6
	DestinationAtop = 7
	Lighter         = 8
	Copy            = 9
	Xor             = 10
)

// ImageFlags as declared in nvg/nanovg.h:138
type ImageFlags int32

// ImageFlags enumeration from nvg/nanovg.h:138
const (
	ImageGenerateMipmaps = 1 << 0
	ImageRepeatx         = 1 << 1
	ImageRepeaty         = 1 << 2
	ImageFlipy           = 1 << 3
	ImagePremultiplied   = 1 << 4
	ImageNearest         = 1 << 5
)

// Texture as declared in nvg/nanovg.h:618
type Texture int32

// Texture enumeration from nvg/nanovg.h:618
const (
	TextureAlpha = 0x01
	TextureRgba  = 0x02
)

// CreateFlags as declared in nvg/nanovg_gl.h:27
type CreateFlags int32

// CreateFlags enumeration from nvg/nanovg_gl.h:27
const (
	Antialias      = 1 << 0
	StencilStrokes = 1 << 1
	Debug          = 1 << 2
)

// ImageFlagsGL as declared in nvg/nanovg_gl.h:98
type ImageFlagsGL int32

// ImageFlagsGL enumeration from nvg/nanovg_gl.h:98
const (
	ImageNodelete = 1 << 16
)
