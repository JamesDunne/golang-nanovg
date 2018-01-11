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

// Context as declared in nvg/nanovg.h:33
type Context C.NVGcontext

// Color as declared in nvg/nanovg.h:43
type Color C.NVGcolor

// Paint as declared in nvg/nanovg.h:54
type Paint C.NVGpaint

// CompositeOperationState as declared in nvg/nanovg.h:120
type CompositeOperationState C.NVGcompositeOperationState

// GlyphPosition as declared in nvg/nanovg.h:127
type GlyphPosition C.NVGglyphPosition

// TextRow as declared in nvg/nanovg.h:136
type TextRow C.NVGtextRow

// Scissoring as declared in nvg/nanovg.h:627
type Scissoring C.NVGscissor

// Vertex as declared in nvg/nanovg.h:632
type Vertex C.NVGvertex

// Path as declared in nvg/nanovg.h:646
type Path C.NVGpath

// Params as declared in nvg/nanovg.h:664
type Params C.NVGparams
