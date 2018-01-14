package nvgui

import (
	"github.com/JamesDunne/golang-nanovg/nvg"
)

type UIPalette [5]nvg.Color
type PaletteIndex int

var (
	// Palettes are ordered from darkest to lightest shades:
	// This palette from https://flatuicolors.com/
	//palette0 = UIPalette{
	//	nvg.RGB(44, 62, 80),    // midnight blue
	//	nvg.RGB(52, 73, 94),    // wet asphalt
	//	nvg.RGB(127, 140, 141), // asbestos
	//	nvg.RGB(149, 165, 166), // concrete
	//}
	// http://flatcolors.net/palette/724-flat-design-blue
	//palette = UIPalette{
	//	nvg.RGB(0x00, 0x00, 0x40),
	//	nvg.RGB(0x00, 0x4C, 0x79),
	//	nvg.RGB(0x00, 0x5F, 0x97),
	//	nvg.RGB(0x3C, 0x8A, 0xB8),
	//	nvg.RGB(0xE8, 0xF4, 0xFA),
	//}
	// http://flatcolors.net/palette/462-flat-existence#
	//palette = UIPalette{
	//	nvg.RGB(0x13, 0x1A, 0x1E),
	//	nvg.RGB(0x11, 0x36, 0xC7),
	//	nvg.RGB(0x1C, 0x57, 0xE1),
	//	nvg.RGB(0x59, 0x7D, 0xF7),
	//	nvg.RGB(0x77, 0x9B, 0xF0),
	//}
	palette = UIPalette{
		nvg.RGB(0x00, 0x00, 0x00),
		nvg.RGB(0x02, 0x05, 0x24),
		nvg.RGB(0x0D, 0x28, 0x4F),
		nvg.RGB(0x71, 0x78, 0x9F),
		nvg.RGB(0xb4, 0xb7, 0xF7),
	}
)

type Touch struct {
	Point
	ID int32
}

type UI struct {
	vg *nvg.Context
	p  UIPalette

	Touches []Touch
}

func NewUI(vg *nvg.Context) *UI {
	return &UI{
		vg:      vg,
		p:       palette,
		Touches: make([]Touch, 10),
	}
}

func (u *UI) Palette(p PaletteIndex) nvg.Color {
	return u.p[p]
}

func (u *UI) Save() {
	nvg.Save(u.vg)
}

func (u *UI) Restore() {
	nvg.Restore(u.vg)
}

func (u *UI) MiterLimit(limit float32) {
	nvg.MiterLimit(u.vg, limit)
}

func (u *UI) LineCap(cap int32) {
	nvg.LineCap(u.vg, cap)
}

func (u *UI) FillColor(c nvg.Color) {
	nvg.FillColor(u.vg, c)
}

func (u *UI) StrokeColor(c nvg.Color) {
	nvg.StrokeColor(u.vg, c)
}

func (u *UI) StrokeWidth(size float32) {
	nvg.StrokeWidth(u.vg, size)
}

func (u *UI) BeginPath() {
	nvg.BeginPath(u.vg)
}

func (u *UI) Fill() {
	nvg.Fill(u.vg)
}

func (u *UI) Stroke() {
	nvg.Stroke(u.vg)
}

func (u *UI) Rect(w Window) {
	nvg.Rect(u.vg, w.X, w.Y, w.W-1, w.H-1)
}

func (u *UI) RoundedRect(w Window, radius float32) {
	nvg.RoundedRect(u.vg, w.X, w.Y, w.W-1, w.H-1, radius)
}

func (u *UI) Circle(p Point, r float32) {
	nvg.Circle(u.vg, p.X, p.Y, r)
}

func (u *UI) TextPoint(p Point, size float32, align int32, string string) {
	nvg.FontSize(u.vg, size)
	nvg.TextAlign(u.vg, align)
	nvg.Text(u.vg, p.X, p.Y, string)
}

func (u *UI) Text(w Window, size float32, align int32, string string) {
	u.TextPoint(w.AlignedPoint(align), size, align, string)
}

// Angles in radians, 0 is horizontal extending right.
func (u *UI) Arc(p Point, r, a0, a1 float32, dir int32) {
	nvg.Arc(u.vg, p.X, p.Y, r, a0, a1, dir)
}

// TODO: add more nvg wrappers
