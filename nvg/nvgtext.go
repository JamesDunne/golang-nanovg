package nvg

/*
#cgo LDFLAGS: -lm -lGLESv2 -L/opt/vc/lib
#cgo CFLAGS: -I/opt/vc/include
#include "nanovg.h"
#include "nanovg_gl_cgo.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// Text function as declared in nvg/nanovg.h:584
func Text(ctx *Context, x float32, y float32, string string) float32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	string = safeString(string)
	cstring, _ := unpackPCharString(string)
	__ret := C.nvgText(cctx, cx, cy, cstring, nil)
	runtime.KeepAlive(string)
	__v := (float32)(__ret)
	return __v
}

// TextBox function as declared in nvg/nanovg.h:589
func TextBox(ctx *Context, x float32, y float32, breakRowWidth float32, string string) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	cbreakRowWidth, _ := (C.float)(breakRowWidth), cgoAllocsUnknown
	string = safeString(string)
	cstring, _ := unpackPCharString(string)
	C.nvgTextBox(cctx, cx, cy, cbreakRowWidth, cstring, nil)
	runtime.KeepAlive(string)
}

// TextBounds function as declared in nvg/nanovg.h:595
func TextBounds(ctx *Context, x float32, y float32, string string, bounds *float32) float32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	string = safeString(string)
	cstring, _ := unpackPCharString(string)
	cbounds, _ := (*C.float)(unsafe.Pointer(bounds)), cgoAllocsUnknown
	__ret := C.nvgTextBounds(cctx, cx, cy, cstring, nil, cbounds)
	runtime.KeepAlive(string)
	__v := (float32)(__ret)
	return __v
}

// TextBoxBounds function as declared in nvg/nanovg.h:600
func TextBoxBounds(ctx *Context, x float32, y float32, breakRowWidth float32, string string, bounds *float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	cbreakRowWidth, _ := (C.float)(breakRowWidth), cgoAllocsUnknown
	string = safeString(string)
	cstring, _ := unpackPCharString(string)
	cbounds, _ := (*C.float)(unsafe.Pointer(bounds)), cgoAllocsUnknown
	C.nvgTextBoxBounds(cctx, cx, cy, cbreakRowWidth, cstring, nil, cbounds)
	runtime.KeepAlive(string)
}

// TextGlyphPositions function as declared in nvg/nanovg.h:604
func TextGlyphPositions(ctx *Context, x float32, y float32, string string, positions *GlyphPosition, maxPositions int32) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	string = safeString(string)
	cstring, _ := unpackPCharString(string)
	cpositions, _ := (*C.NVGglyphPosition)(unsafe.Pointer(positions)), cgoAllocsUnknown
	cmaxPositions, _ := (C.int)(maxPositions), cgoAllocsUnknown
	__ret := C.nvgTextGlyphPositions(cctx, cx, cy, cstring, nil, cpositions, cmaxPositions)
	runtime.KeepAlive(string)
	__v := (int32)(__ret)
	return __v
}

// TextBreakLines function as declared in nvg/nanovg.h:613
func TextBreakLines(ctx *Context, string string, breakRowWidth float32, rows *TextRow, maxRows int32) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	string = safeString(string)
	cstring, _ := unpackPCharString(string)
	cbreakRowWidth, _ := (C.float)(breakRowWidth), cgoAllocsUnknown
	crows, _ := (*C.NVGtextRow)(unsafe.Pointer(rows)), cgoAllocsUnknown
	cmaxRows, _ := (C.int)(maxRows), cgoAllocsUnknown
	__ret := C.nvgTextBreakLines(cctx, cstring, nil, cbreakRowWidth, crows, cmaxRows)
	runtime.KeepAlive(string)
	__v := (int32)(__ret)
	return __v
}
