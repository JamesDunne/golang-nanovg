// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 11 Jan 2018 11:32:52 CST.
// By https://git.io/c-for-go. DO NOT EDIT.

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

// BeginFrame function as declared in nvg/nanovg.h:155
func BeginFrame(ctx *Context, windowWidth int32, windowHeight int32, devicePixelRatio float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cwindowWidth, _ := (C.int)(windowWidth), cgoAllocsUnknown
	cwindowHeight, _ := (C.int)(windowHeight), cgoAllocsUnknown
	cdevicePixelRatio, _ := (C.float)(devicePixelRatio), cgoAllocsUnknown
	C.nvgBeginFrame(cctx, cwindowWidth, cwindowHeight, cdevicePixelRatio)
}

// CancelFrame function as declared in nvg/nanovg.h:158
func CancelFrame(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgCancelFrame(cctx)
}

// EndFrame function as declared in nvg/nanovg.h:161
func EndFrame(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgEndFrame(cctx)
}

// GlobalCompositeOperation function as declared in nvg/nanovg.h:171
func GlobalCompositeOperation(ctx *Context, op int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cop, _ := (C.int)(op), cgoAllocsUnknown
	C.nvgGlobalCompositeOperation(cctx, cop)
}

// GlobalCompositeBlendFunc function as declared in nvg/nanovg.h:174
func GlobalCompositeBlendFunc(ctx *Context, sfactor int32, dfactor int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	csfactor, _ := (C.int)(sfactor), cgoAllocsUnknown
	cdfactor, _ := (C.int)(dfactor), cgoAllocsUnknown
	C.nvgGlobalCompositeBlendFunc(cctx, csfactor, cdfactor)
}

// GlobalCompositeBlendFuncSeparate function as declared in nvg/nanovg.h:177
func GlobalCompositeBlendFuncSeparate(ctx *Context, srcRGB int32, dstRGB int32, srcAlpha int32, dstAlpha int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	csrcRGB, _ := (C.int)(srcRGB), cgoAllocsUnknown
	cdstRGB, _ := (C.int)(dstRGB), cgoAllocsUnknown
	csrcAlpha, _ := (C.int)(srcAlpha), cgoAllocsUnknown
	cdstAlpha, _ := (C.int)(dstAlpha), cgoAllocsUnknown
	C.nvgGlobalCompositeBlendFuncSeparate(cctx, csrcRGB, cdstRGB, csrcAlpha, cdstAlpha)
}

// RGB function as declared in nvg/nanovg.h:185
func RGB(r byte, g byte, b byte) Color {
	cr, _ := (C.uchar)(r), cgoAllocsUnknown
	cg, _ := (C.uchar)(g), cgoAllocsUnknown
	cb, _ := (C.uchar)(b), cgoAllocsUnknown
	__ret := C.nvgRGB(cr, cg, cb)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// RGBf function as declared in nvg/nanovg.h:188
func RGBf(r float32, g float32, b float32) Color {
	cr, _ := (C.float)(r), cgoAllocsUnknown
	cg, _ := (C.float)(g), cgoAllocsUnknown
	cb, _ := (C.float)(b), cgoAllocsUnknown
	__ret := C.nvgRGBf(cr, cg, cb)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// RGBA function as declared in nvg/nanovg.h:192
func RGBA(r byte, g byte, b byte, a byte) Color {
	cr, _ := (C.uchar)(r), cgoAllocsUnknown
	cg, _ := (C.uchar)(g), cgoAllocsUnknown
	cb, _ := (C.uchar)(b), cgoAllocsUnknown
	ca, _ := (C.uchar)(a), cgoAllocsUnknown
	__ret := C.nvgRGBA(cr, cg, cb, ca)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// RGBAf function as declared in nvg/nanovg.h:195
func RGBAf(r float32, g float32, b float32, a float32) Color {
	cr, _ := (C.float)(r), cgoAllocsUnknown
	cg, _ := (C.float)(g), cgoAllocsUnknown
	cb, _ := (C.float)(b), cgoAllocsUnknown
	ca, _ := (C.float)(a), cgoAllocsUnknown
	__ret := C.nvgRGBAf(cr, cg, cb, ca)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// LerpRGBA function as declared in nvg/nanovg.h:199
func LerpRGBA(c0 Color, c1 Color, u float32) Color {
	cc0, _ := *(*C.NVGcolor)(unsafe.Pointer(&c0)), cgoAllocsUnknown
	cc1, _ := *(*C.NVGcolor)(unsafe.Pointer(&c1)), cgoAllocsUnknown
	cu, _ := (C.float)(u), cgoAllocsUnknown
	__ret := C.nvgLerpRGBA(cc0, cc1, cu)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// TransRGBA function as declared in nvg/nanovg.h:202
func TransRGBA(c0 Color, a byte) Color {
	cc0, _ := *(*C.NVGcolor)(unsafe.Pointer(&c0)), cgoAllocsUnknown
	ca, _ := (C.uchar)(a), cgoAllocsUnknown
	__ret := C.nvgTransRGBA(cc0, ca)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// TransRGBAf function as declared in nvg/nanovg.h:205
func TransRGBAf(c0 Color, a float32) Color {
	cc0, _ := *(*C.NVGcolor)(unsafe.Pointer(&c0)), cgoAllocsUnknown
	ca, _ := (C.float)(a), cgoAllocsUnknown
	__ret := C.nvgTransRGBAf(cc0, ca)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// HSL function as declared in nvg/nanovg.h:209
func HSL(h float32, s float32, l float32) Color {
	ch, _ := (C.float)(h), cgoAllocsUnknown
	cs, _ := (C.float)(s), cgoAllocsUnknown
	cl, _ := (C.float)(l), cgoAllocsUnknown
	__ret := C.nvgHSL(ch, cs, cl)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// HSLA function as declared in nvg/nanovg.h:213
func HSLA(h float32, s float32, l float32, a byte) Color {
	ch, _ := (C.float)(h), cgoAllocsUnknown
	cs, _ := (C.float)(s), cgoAllocsUnknown
	cl, _ := (C.float)(l), cgoAllocsUnknown
	ca, _ := (C.uchar)(a), cgoAllocsUnknown
	__ret := C.nvgHSLA(ch, cs, cl, ca)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// Save function as declared in nvg/nanovg.h:224
func Save(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgSave(cctx)
}

// Restore function as declared in nvg/nanovg.h:227
func Restore(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgRestore(cctx)
}

// Reset function as declared in nvg/nanovg.h:230
func Reset(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgReset(cctx)
}

// ShapeAntiAlias function as declared in nvg/nanovg.h:242
func ShapeAntiAlias(ctx *Context, enabled int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cenabled, _ := (C.int)(enabled), cgoAllocsUnknown
	C.nvgShapeAntiAlias(cctx, cenabled)
}

// StrokeColor function as declared in nvg/nanovg.h:245
func StrokeColor(ctx *Context, color Color) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccolor, _ := *(*C.NVGcolor)(unsafe.Pointer(&color)), cgoAllocsUnknown
	C.nvgStrokeColor(cctx, ccolor)
}

// StrokePaint function as declared in nvg/nanovg.h:248
func StrokePaint(ctx *Context, paint Paint) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cpaint, _ := *(*C.NVGpaint)(unsafe.Pointer(&paint)), cgoAllocsUnknown
	C.nvgStrokePaint(cctx, cpaint)
}

// FillColor function as declared in nvg/nanovg.h:251
func FillColor(ctx *Context, color Color) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccolor, _ := *(*C.NVGcolor)(unsafe.Pointer(&color)), cgoAllocsUnknown
	C.nvgFillColor(cctx, ccolor)
}

// FillPaint function as declared in nvg/nanovg.h:254
func FillPaint(ctx *Context, paint Paint) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cpaint, _ := *(*C.NVGpaint)(unsafe.Pointer(&paint)), cgoAllocsUnknown
	C.nvgFillPaint(cctx, cpaint)
}

// MiterLimit function as declared in nvg/nanovg.h:258
func MiterLimit(ctx *Context, limit float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	climit, _ := (C.float)(limit), cgoAllocsUnknown
	C.nvgMiterLimit(cctx, climit)
}

// StrokeWidth function as declared in nvg/nanovg.h:261
func StrokeWidth(ctx *Context, size float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	csize, _ := (C.float)(size), cgoAllocsUnknown
	C.nvgStrokeWidth(cctx, csize)
}

// LineCap function as declared in nvg/nanovg.h:265
func LineCap(ctx *Context, cap int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccap, _ := (C.int)(cap), cgoAllocsUnknown
	C.nvgLineCap(cctx, ccap)
}

// LineJoin function as declared in nvg/nanovg.h:269
func LineJoin(ctx *Context, join int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cjoin, _ := (C.int)(join), cgoAllocsUnknown
	C.nvgLineJoin(cctx, cjoin)
}

// GlobalAlpha function as declared in nvg/nanovg.h:273
func GlobalAlpha(ctx *Context, alpha float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	calpha, _ := (C.float)(alpha), cgoAllocsUnknown
	C.nvgGlobalAlpha(cctx, calpha)
}

// ResetTransform function as declared in nvg/nanovg.h:293
func ResetTransform(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgResetTransform(cctx)
}

// Transform function as declared in nvg/nanovg.h:300
func Transform(ctx *Context, a float32, b float32, c float32, d float32, e float32, f float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ca, _ := (C.float)(a), cgoAllocsUnknown
	cb, _ := (C.float)(b), cgoAllocsUnknown
	cc, _ := (C.float)(c), cgoAllocsUnknown
	cd, _ := (C.float)(d), cgoAllocsUnknown
	ce, _ := (C.float)(e), cgoAllocsUnknown
	cf, _ := (C.float)(f), cgoAllocsUnknown
	C.nvgTransform(cctx, ca, cb, cc, cd, ce, cf)
}

// Translate function as declared in nvg/nanovg.h:303
func Translate(ctx *Context, x float32, y float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	C.nvgTranslate(cctx, cx, cy)
}

// Rotate function as declared in nvg/nanovg.h:306
func Rotate(ctx *Context, angle float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cangle, _ := (C.float)(angle), cgoAllocsUnknown
	C.nvgRotate(cctx, cangle)
}

// SkewX function as declared in nvg/nanovg.h:309
func SkewX(ctx *Context, angle float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cangle, _ := (C.float)(angle), cgoAllocsUnknown
	C.nvgSkewX(cctx, cangle)
}

// SkewY function as declared in nvg/nanovg.h:312
func SkewY(ctx *Context, angle float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cangle, _ := (C.float)(angle), cgoAllocsUnknown
	C.nvgSkewY(cctx, cangle)
}

// Scale function as declared in nvg/nanovg.h:315
func Scale(ctx *Context, x float32, y float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	C.nvgScale(cctx, cx, cy)
}

// CurrentTransform function as declared in nvg/nanovg.h:322
func CurrentTransform(ctx *Context, xform *float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cxform, _ := (*C.float)(unsafe.Pointer(xform)), cgoAllocsUnknown
	C.nvgCurrentTransform(cctx, cxform)
}

// TransformIdentity function as declared in nvg/nanovg.h:329
func TransformIdentity(dst *float32) {
	cdst, _ := (*C.float)(unsafe.Pointer(dst)), cgoAllocsUnknown
	C.nvgTransformIdentity(cdst)
}

// TransformTranslate function as declared in nvg/nanovg.h:332
func TransformTranslate(dst *float32, tx float32, ty float32) {
	cdst, _ := (*C.float)(unsafe.Pointer(dst)), cgoAllocsUnknown
	ctx, _ := (C.float)(tx), cgoAllocsUnknown
	cty, _ := (C.float)(ty), cgoAllocsUnknown
	C.nvgTransformTranslate(cdst, ctx, cty)
}

// TransformScale function as declared in nvg/nanovg.h:335
func TransformScale(dst *float32, sx float32, sy float32) {
	cdst, _ := (*C.float)(unsafe.Pointer(dst)), cgoAllocsUnknown
	csx, _ := (C.float)(sx), cgoAllocsUnknown
	csy, _ := (C.float)(sy), cgoAllocsUnknown
	C.nvgTransformScale(cdst, csx, csy)
}

// TransformRotate function as declared in nvg/nanovg.h:338
func TransformRotate(dst *float32, a float32) {
	cdst, _ := (*C.float)(unsafe.Pointer(dst)), cgoAllocsUnknown
	ca, _ := (C.float)(a), cgoAllocsUnknown
	C.nvgTransformRotate(cdst, ca)
}

// TransformSkewX function as declared in nvg/nanovg.h:341
func TransformSkewX(dst *float32, a float32) {
	cdst, _ := (*C.float)(unsafe.Pointer(dst)), cgoAllocsUnknown
	ca, _ := (C.float)(a), cgoAllocsUnknown
	C.nvgTransformSkewX(cdst, ca)
}

// TransformSkewY function as declared in nvg/nanovg.h:344
func TransformSkewY(dst *float32, a float32) {
	cdst, _ := (*C.float)(unsafe.Pointer(dst)), cgoAllocsUnknown
	ca, _ := (C.float)(a), cgoAllocsUnknown
	C.nvgTransformSkewY(cdst, ca)
}

// TransformMultiply function as declared in nvg/nanovg.h:347
func TransformMultiply(dst *float32, src *float32) {
	cdst, _ := (*C.float)(unsafe.Pointer(dst)), cgoAllocsUnknown
	csrc, _ := (*C.float)(unsafe.Pointer(src)), cgoAllocsUnknown
	C.nvgTransformMultiply(cdst, csrc)
}

// TransformPremultiply function as declared in nvg/nanovg.h:350
func TransformPremultiply(dst *float32, src *float32) {
	cdst, _ := (*C.float)(unsafe.Pointer(dst)), cgoAllocsUnknown
	csrc, _ := (*C.float)(unsafe.Pointer(src)), cgoAllocsUnknown
	C.nvgTransformPremultiply(cdst, csrc)
}

// TransformInverse function as declared in nvg/nanovg.h:354
func TransformInverse(dst *float32, src *float32) int32 {
	cdst, _ := (*C.float)(unsafe.Pointer(dst)), cgoAllocsUnknown
	csrc, _ := (*C.float)(unsafe.Pointer(src)), cgoAllocsUnknown
	__ret := C.nvgTransformInverse(cdst, csrc)
	__v := (int32)(__ret)
	return __v
}

// TransformPoint function as declared in nvg/nanovg.h:357
func TransformPoint(dstx *float32, dsty *float32, xform *float32, srcx float32, srcy float32) {
	cdstx, _ := (*C.float)(unsafe.Pointer(dstx)), cgoAllocsUnknown
	cdsty, _ := (*C.float)(unsafe.Pointer(dsty)), cgoAllocsUnknown
	cxform, _ := (*C.float)(unsafe.Pointer(xform)), cgoAllocsUnknown
	csrcx, _ := (C.float)(srcx), cgoAllocsUnknown
	csrcy, _ := (C.float)(srcy), cgoAllocsUnknown
	C.nvgTransformPoint(cdstx, cdsty, cxform, csrcx, csrcy)
}

// DegToRad function as declared in nvg/nanovg.h:360
func DegToRad(deg float32) float32 {
	cdeg, _ := (C.float)(deg), cgoAllocsUnknown
	__ret := C.nvgDegToRad(cdeg)
	__v := (float32)(__ret)
	return __v
}

// RadToDeg function as declared in nvg/nanovg.h:361
func RadToDeg(rad float32) float32 {
	crad, _ := (C.float)(rad), cgoAllocsUnknown
	__ret := C.nvgRadToDeg(crad)
	__v := (float32)(__ret)
	return __v
}

// CreateImage function as declared in nvg/nanovg.h:372
func CreateImage(ctx *Context, filename string, imageFlags int32) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	filename = safeString(filename)
	cfilename, _ := unpackPCharString(filename)
	cimageFlags, _ := (C.int)(imageFlags), cgoAllocsUnknown
	__ret := C.nvgCreateImage(cctx, cfilename, cimageFlags)
	runtime.KeepAlive(filename)
	__v := (int32)(__ret)
	return __v
}

// CreateImageMem function as declared in nvg/nanovg.h:376
func CreateImageMem(ctx *Context, imageFlags int32, data *byte, ndata int32) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cimageFlags, _ := (C.int)(imageFlags), cgoAllocsUnknown
	cdata, _ := (*C.uchar)(unsafe.Pointer(data)), cgoAllocsUnknown
	cndata, _ := (C.int)(ndata), cgoAllocsUnknown
	__ret := C.nvgCreateImageMem(cctx, cimageFlags, cdata, cndata)
	__v := (int32)(__ret)
	return __v
}

// CreateImageRGBA function as declared in nvg/nanovg.h:380
func CreateImageRGBA(ctx *Context, w int32, h int32, imageFlags int32, data string) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cw, _ := (C.int)(w), cgoAllocsUnknown
	ch, _ := (C.int)(h), cgoAllocsUnknown
	cimageFlags, _ := (C.int)(imageFlags), cgoAllocsUnknown
	data = safeString(data)
	cdata, _ := unpackPUcharString(data)
	__ret := C.nvgCreateImageRGBA(cctx, cw, ch, cimageFlags, cdata)
	runtime.KeepAlive(data)
	__v := (int32)(__ret)
	return __v
}

// UpdateImage function as declared in nvg/nanovg.h:383
func UpdateImage(ctx *Context, image int32, data string) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cimage, _ := (C.int)(image), cgoAllocsUnknown
	data = safeString(data)
	cdata, _ := unpackPUcharString(data)
	C.nvgUpdateImage(cctx, cimage, cdata)
	runtime.KeepAlive(data)
}

// ImageSize function as declared in nvg/nanovg.h:386
func ImageSize(ctx *Context, image int32, w *int32, h *int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cimage, _ := (C.int)(image), cgoAllocsUnknown
	cw, _ := (*C.int)(unsafe.Pointer(w)), cgoAllocsUnknown
	ch, _ := (*C.int)(unsafe.Pointer(h)), cgoAllocsUnknown
	C.nvgImageSize(cctx, cimage, cw, ch)
}

// DeleteImage function as declared in nvg/nanovg.h:389
func DeleteImage(ctx *Context, image int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cimage, _ := (C.int)(image), cgoAllocsUnknown
	C.nvgDeleteImage(cctx, cimage)
}

// LinearGradient function as declared in nvg/nanovg.h:400
func LinearGradient(ctx *Context, sx float32, sy float32, ex float32, ey float32, icol Color, ocol Color) Paint {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	csx, _ := (C.float)(sx), cgoAllocsUnknown
	csy, _ := (C.float)(sy), cgoAllocsUnknown
	cex, _ := (C.float)(ex), cgoAllocsUnknown
	cey, _ := (C.float)(ey), cgoAllocsUnknown
	cicol, _ := *(*C.NVGcolor)(unsafe.Pointer(&icol)), cgoAllocsUnknown
	cocol, _ := *(*C.NVGcolor)(unsafe.Pointer(&ocol)), cgoAllocsUnknown
	__ret := C.nvgLinearGradient(cctx, csx, csy, cex, cey, cicol, cocol)
	__v := *NewPaintRef(unsafe.Pointer(&__ret))
	return __v
}

// BoxGradient function as declared in nvg/nanovg.h:408
func BoxGradient(ctx *Context, x float32, y float32, w float32, h float32, r float32, f float32, icol Color, ocol Color) Paint {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	cw, _ := (C.float)(w), cgoAllocsUnknown
	ch, _ := (C.float)(h), cgoAllocsUnknown
	cr, _ := (C.float)(r), cgoAllocsUnknown
	cf, _ := (C.float)(f), cgoAllocsUnknown
	cicol, _ := *(*C.NVGcolor)(unsafe.Pointer(&icol)), cgoAllocsUnknown
	cocol, _ := *(*C.NVGcolor)(unsafe.Pointer(&ocol)), cgoAllocsUnknown
	__ret := C.nvgBoxGradient(cctx, cx, cy, cw, ch, cr, cf, cicol, cocol)
	__v := *NewPaintRef(unsafe.Pointer(&__ret))
	return __v
}

// RadialGradient function as declared in nvg/nanovg.h:414
func RadialGradient(ctx *Context, cx float32, cy float32, inr float32, outr float32, icol Color, ocol Color) Paint {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccx, _ := (C.float)(cx), cgoAllocsUnknown
	ccy, _ := (C.float)(cy), cgoAllocsUnknown
	cinr, _ := (C.float)(inr), cgoAllocsUnknown
	coutr, _ := (C.float)(outr), cgoAllocsUnknown
	cicol, _ := *(*C.NVGcolor)(unsafe.Pointer(&icol)), cgoAllocsUnknown
	cocol, _ := *(*C.NVGcolor)(unsafe.Pointer(&ocol)), cgoAllocsUnknown
	__ret := C.nvgRadialGradient(cctx, ccx, ccy, cinr, coutr, cicol, cocol)
	__v := *NewPaintRef(unsafe.Pointer(&__ret))
	return __v
}

// ImagePattern function as declared in nvg/nanovg.h:420
func ImagePattern(ctx *Context, ox float32, oy float32, ex float32, ey float32, angle float32, image int32, alpha float32) Paint {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cox, _ := (C.float)(ox), cgoAllocsUnknown
	coy, _ := (C.float)(oy), cgoAllocsUnknown
	cex, _ := (C.float)(ex), cgoAllocsUnknown
	cey, _ := (C.float)(ey), cgoAllocsUnknown
	cangle, _ := (C.float)(angle), cgoAllocsUnknown
	cimage, _ := (C.int)(image), cgoAllocsUnknown
	calpha, _ := (C.float)(alpha), cgoAllocsUnknown
	__ret := C.nvgImagePattern(cctx, cox, coy, cex, cey, cangle, cimage, calpha)
	__v := *NewPaintRef(unsafe.Pointer(&__ret))
	return __v
}

// Scissor function as declared in nvg/nanovg.h:431
func Scissor(ctx *Context, x float32, y float32, w float32, h float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	cw, _ := (C.float)(w), cgoAllocsUnknown
	ch, _ := (C.float)(h), cgoAllocsUnknown
	C.nvgScissor(cctx, cx, cy, cw, ch)
}

// IntersectScissor function as declared in nvg/nanovg.h:439
func IntersectScissor(ctx *Context, x float32, y float32, w float32, h float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	cw, _ := (C.float)(w), cgoAllocsUnknown
	ch, _ := (C.float)(h), cgoAllocsUnknown
	C.nvgIntersectScissor(cctx, cx, cy, cw, ch)
}

// ResetScissor function as declared in nvg/nanovg.h:442
func ResetScissor(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgResetScissor(cctx)
}

// BeginPath function as declared in nvg/nanovg.h:462
func BeginPath(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgBeginPath(cctx)
}

// MoveTo function as declared in nvg/nanovg.h:465
func MoveTo(ctx *Context, x float32, y float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	C.nvgMoveTo(cctx, cx, cy)
}

// LineTo function as declared in nvg/nanovg.h:468
func LineTo(ctx *Context, x float32, y float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	C.nvgLineTo(cctx, cx, cy)
}

// BezierTo function as declared in nvg/nanovg.h:471
func BezierTo(ctx *Context, c1x float32, c1y float32, c2x float32, c2y float32, x float32, y float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cc1x, _ := (C.float)(c1x), cgoAllocsUnknown
	cc1y, _ := (C.float)(c1y), cgoAllocsUnknown
	cc2x, _ := (C.float)(c2x), cgoAllocsUnknown
	cc2y, _ := (C.float)(c2y), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	C.nvgBezierTo(cctx, cc1x, cc1y, cc2x, cc2y, cx, cy)
}

// QuadTo function as declared in nvg/nanovg.h:474
func QuadTo(ctx *Context, cx float32, cy float32, x float32, y float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccx, _ := (C.float)(cx), cgoAllocsUnknown
	ccy, _ := (C.float)(cy), cgoAllocsUnknown
	cccx, _ := (C.float)(x), cgoAllocsUnknown
	cccy, _ := (C.float)(y), cgoAllocsUnknown
	C.nvgQuadTo(cctx, ccx, ccy, cccx, cccy)
}

// ArcTo function as declared in nvg/nanovg.h:477
func ArcTo(ctx *Context, x1 float32, y1 float32, x2 float32, y2 float32, radius float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx1, _ := (C.float)(x1), cgoAllocsUnknown
	cy1, _ := (C.float)(y1), cgoAllocsUnknown
	cx2, _ := (C.float)(x2), cgoAllocsUnknown
	cy2, _ := (C.float)(y2), cgoAllocsUnknown
	cradius, _ := (C.float)(radius), cgoAllocsUnknown
	C.nvgArcTo(cctx, cx1, cy1, cx2, cy2, cradius)
}

// ClosePath function as declared in nvg/nanovg.h:480
func ClosePath(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgClosePath(cctx)
}

// PathWinding function as declared in nvg/nanovg.h:483
func PathWinding(ctx *Context, dir int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cdir, _ := (C.int)(dir), cgoAllocsUnknown
	C.nvgPathWinding(cctx, cdir)
}

// Arc function as declared in nvg/nanovg.h:488
func Arc(ctx *Context, cx float32, cy float32, r float32, a0 float32, a1 float32, dir int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccx, _ := (C.float)(cx), cgoAllocsUnknown
	ccy, _ := (C.float)(cy), cgoAllocsUnknown
	cr, _ := (C.float)(r), cgoAllocsUnknown
	ca0, _ := (C.float)(a0), cgoAllocsUnknown
	ca1, _ := (C.float)(a1), cgoAllocsUnknown
	cdir, _ := (C.int)(dir), cgoAllocsUnknown
	C.nvgArc(cctx, ccx, ccy, cr, ca0, ca1, cdir)
}

// Rect function as declared in nvg/nanovg.h:491
func Rect(ctx *Context, x float32, y float32, w float32, h float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	cw, _ := (C.float)(w), cgoAllocsUnknown
	ch, _ := (C.float)(h), cgoAllocsUnknown
	C.nvgRect(cctx, cx, cy, cw, ch)
}

// RoundedRect function as declared in nvg/nanovg.h:494
func RoundedRect(ctx *Context, x float32, y float32, w float32, h float32, r float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	cw, _ := (C.float)(w), cgoAllocsUnknown
	ch, _ := (C.float)(h), cgoAllocsUnknown
	cr, _ := (C.float)(r), cgoAllocsUnknown
	C.nvgRoundedRect(cctx, cx, cy, cw, ch, cr)
}

// RoundedRectVarying function as declared in nvg/nanovg.h:497
func RoundedRectVarying(ctx *Context, x float32, y float32, w float32, h float32, radTopLeft float32, radTopRight float32, radBottomRight float32, radBottomLeft float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cx, _ := (C.float)(x), cgoAllocsUnknown
	cy, _ := (C.float)(y), cgoAllocsUnknown
	cw, _ := (C.float)(w), cgoAllocsUnknown
	ch, _ := (C.float)(h), cgoAllocsUnknown
	cradTopLeft, _ := (C.float)(radTopLeft), cgoAllocsUnknown
	cradTopRight, _ := (C.float)(radTopRight), cgoAllocsUnknown
	cradBottomRight, _ := (C.float)(radBottomRight), cgoAllocsUnknown
	cradBottomLeft, _ := (C.float)(radBottomLeft), cgoAllocsUnknown
	C.nvgRoundedRectVarying(cctx, cx, cy, cw, ch, cradTopLeft, cradTopRight, cradBottomRight, cradBottomLeft)
}

// Ellipse function as declared in nvg/nanovg.h:500
func Ellipse(ctx *Context, cx float32, cy float32, rx float32, ry float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccx, _ := (C.float)(cx), cgoAllocsUnknown
	ccy, _ := (C.float)(cy), cgoAllocsUnknown
	crx, _ := (C.float)(rx), cgoAllocsUnknown
	cry, _ := (C.float)(ry), cgoAllocsUnknown
	C.nvgEllipse(cctx, ccx, ccy, crx, cry)
}

// Circle function as declared in nvg/nanovg.h:503
func Circle(ctx *Context, cx float32, cy float32, r float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ccx, _ := (C.float)(cx), cgoAllocsUnknown
	ccy, _ := (C.float)(cy), cgoAllocsUnknown
	cr, _ := (C.float)(r), cgoAllocsUnknown
	C.nvgCircle(cctx, ccx, ccy, cr)
}

// Fill function as declared in nvg/nanovg.h:506
func Fill(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgFill(cctx)
}

// Stroke function as declared in nvg/nanovg.h:509
func Stroke(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgStroke(cctx)
}

// CreateFont function as declared in nvg/nanovg.h:547
func CreateFont(ctx *Context, name string, filename string) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	name = safeString(name)
	cname, _ := unpackPCharString(name)
	filename = safeString(filename)
	cfilename, _ := unpackPCharString(filename)
	__ret := C.nvgCreateFont(cctx, cname, cfilename)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(name)
	__v := (int32)(__ret)
	return __v
}

// CreateFontMem function as declared in nvg/nanovg.h:551
func CreateFontMem(ctx *Context, name string, data *byte, ndata int32, freeData int32) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	name = safeString(name)
	cname, _ := unpackPCharString(name)
	cdata, _ := (*C.uchar)(unsafe.Pointer(data)), cgoAllocsUnknown
	cndata, _ := (C.int)(ndata), cgoAllocsUnknown
	cfreeData, _ := (C.int)(freeData), cgoAllocsUnknown
	__ret := C.nvgCreateFontMem(cctx, cname, cdata, cndata, cfreeData)
	runtime.KeepAlive(name)
	__v := (int32)(__ret)
	return __v
}

// FindFont function as declared in nvg/nanovg.h:554
func FindFont(ctx *Context, name string) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	name = safeString(name)
	cname, _ := unpackPCharString(name)
	__ret := C.nvgFindFont(cctx, cname)
	runtime.KeepAlive(name)
	__v := (int32)(__ret)
	return __v
}

// AddFallbackFontId function as declared in nvg/nanovg.h:557
func AddFallbackFontId(ctx *Context, baseFont int32, fallbackFont int32) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cbaseFont, _ := (C.int)(baseFont), cgoAllocsUnknown
	cfallbackFont, _ := (C.int)(fallbackFont), cgoAllocsUnknown
	__ret := C.nvgAddFallbackFontId(cctx, cbaseFont, cfallbackFont)
	__v := (int32)(__ret)
	return __v
}

// AddFallbackFont function as declared in nvg/nanovg.h:560
func AddFallbackFont(ctx *Context, baseFont string, fallbackFont string) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	baseFont = safeString(baseFont)
	cbaseFont, _ := unpackPCharString(baseFont)
	fallbackFont = safeString(fallbackFont)
	cfallbackFont, _ := unpackPCharString(fallbackFont)
	__ret := C.nvgAddFallbackFont(cctx, cbaseFont, cfallbackFont)
	runtime.KeepAlive(fallbackFont)
	runtime.KeepAlive(baseFont)
	__v := (int32)(__ret)
	return __v
}

// FontSize function as declared in nvg/nanovg.h:563
func FontSize(ctx *Context, size float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	csize, _ := (C.float)(size), cgoAllocsUnknown
	C.nvgFontSize(cctx, csize)
}

// FontBlur function as declared in nvg/nanovg.h:566
func FontBlur(ctx *Context, blur float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cblur, _ := (C.float)(blur), cgoAllocsUnknown
	C.nvgFontBlur(cctx, cblur)
}

// TextLetterSpacing function as declared in nvg/nanovg.h:569
func TextLetterSpacing(ctx *Context, spacing float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cspacing, _ := (C.float)(spacing), cgoAllocsUnknown
	C.nvgTextLetterSpacing(cctx, cspacing)
}

// TextLineHeight function as declared in nvg/nanovg.h:572
func TextLineHeight(ctx *Context, lineHeight float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	clineHeight, _ := (C.float)(lineHeight), cgoAllocsUnknown
	C.nvgTextLineHeight(cctx, clineHeight)
}

// TextAlign function as declared in nvg/nanovg.h:575
func TextAlign(ctx *Context, align int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	calign, _ := (C.int)(align), cgoAllocsUnknown
	C.nvgTextAlign(cctx, calign)
}

// FontFaceId function as declared in nvg/nanovg.h:578
func FontFaceId(ctx *Context, font int32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cfont, _ := (C.int)(font), cgoAllocsUnknown
	C.nvgFontFaceId(cctx, cfont)
}

// FontFace function as declared in nvg/nanovg.h:581
func FontFace(ctx *Context, font string) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	font = safeString(font)
	cfont, _ := unpackPCharString(font)
	C.nvgFontFace(cctx, cfont)
	runtime.KeepAlive(font)
}

// TextMetrics function as declared in nvg/nanovg.h:608
func TextMetrics(ctx *Context, ascender *float32, descender *float32, lineh *float32) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cascender, _ := (*C.float)(unsafe.Pointer(ascender)), cgoAllocsUnknown
	cdescender, _ := (*C.float)(unsafe.Pointer(descender)), cgoAllocsUnknown
	clineh, _ := (*C.float)(unsafe.Pointer(lineh)), cgoAllocsUnknown
	C.nvgTextMetrics(cctx, cascender, cdescender, clineh)
}

// CreateInternal function as declared in nvg/nanovg.h:667
func CreateInternal(params *Params) *Context {
	cparams, _ := (*C.NVGparams)(unsafe.Pointer(params)), cgoAllocsUnknown
	__ret := C.nvgCreateInternal(cparams)
	__v := *(**Context)(unsafe.Pointer(&__ret))
	return __v
}

// DeleteInternal function as declared in nvg/nanovg.h:668
func DeleteInternal(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgDeleteInternal(cctx)
}

// InternalParams function as declared in nvg/nanovg.h:670
func InternalParams(ctx *Context) *Params {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	__ret := C.nvgInternalParams(cctx)
	__v := NewParamsRef(unsafe.Pointer(__ret))
	return __v
}

// DebugDumpPathCache function as declared in nvg/nanovg.h:673
func DebugDumpPathCache(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgDebugDumpPathCache(cctx)
}

// CreateGLES2 function as declared in nvg/nanovg_gl.h:79
func CreateGLES2(flags int32) *Context {
	cflags, _ := (C.int)(flags), cgoAllocsUnknown
	__ret := C.nvgCreateGLES2(cflags)
	__v := *(**Context)(unsafe.Pointer(&__ret))
	return __v
}

// DeleteGLES2 function as declared in nvg/nanovg_gl.h:80
func DeleteGLES2(ctx *Context) {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	C.nvgDeleteGLES2(cctx)
}

// CreateImageFromHandleGLES2 function as declared in nvg/nanovg_gl.h:82
func CreateImageFromHandleGLES2(ctx *Context, textureId uint32, w int32, h int32, flags int32) int32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	ctextureId, _ := (C.GLuint)(textureId), cgoAllocsUnknown
	cw, _ := (C.int)(w), cgoAllocsUnknown
	ch, _ := (C.int)(h), cgoAllocsUnknown
	cflags, _ := (C.int)(flags), cgoAllocsUnknown
	__ret := C.nvglCreateImageFromHandleGLES2(cctx, ctextureId, cw, ch, cflags)
	__v := (int32)(__ret)
	return __v
}

// ImageHandleGLES2 function as declared in nvg/nanovg_gl.h:83
func ImageHandleGLES2(ctx *Context, image int32) uint32 {
	cctx, _ := (*C.NVGcontext)(unsafe.Pointer(ctx)), cgoAllocsUnknown
	cimage, _ := (C.int)(image), cgoAllocsUnknown
	__ret := C.nvglImageHandleGLES2(cctx, cimage)
	__v := (uint32)(__ret)
	return __v
}
