package nvgui

import (
	"github.com/JamesDunne/golang-nanovg/nvg"
)

const pad = 2
const size = 28.0
const round = 4.0

func (u *UI) isTouched(w Window) bool {
	for _, t := range u.Touches {
		// Skip released touch points:
		if t.ID <= 0 {
			continue
		}

		p := Point{t.X, t.Y}
		if w.IsPointInside(p) {
			return true
		}
	}
	return false
}

func (ui *UI) Button(w Window, label string) bool {
	touched := ui.isTouched(w)
	if touched {
		ui.StrokeColor(ui.Palette(2))
		ui.FillColor(ui.Palette(3))
	} else {
		ui.StrokeColor(ui.Palette(3))
		ui.FillColor(ui.Palette(2))
	}

	ui.BeginPath()
	ui.RoundedRect(w, round)
	ui.Stroke()
	ui.Fill()

	if touched {
		ui.FillColor(ui.Palette(1))
	} else {
		ui.FillColor(ui.Palette(4))
	}
	ui.Text(w, size, nvg.AlignCenter|nvg.AlignMiddle, label)

	// TODO: use state for latching behavior
	return touched
}

func (ui *UI) Pane(w Window) {
	ui.BeginPath()
	ui.RoundedRect(w, round)
	ui.Stroke()
}

func (ui *UI) Label(w Window, string string, align int32) {
	ui.BeginPath()
	ui.RoundedRect(w, round)
	ui.StrokeColor(ui.Palette(3))
	ui.FillColor(ui.Palette(1))
	ui.Fill()
	ui.Stroke()

	lblText := w.Inner(pad*2, 0, pad*2, 0)
	ui.FillColor(ui.Palette(4))
	ui.Text(lblText, size, align, string)
}

func (ui *UI) Dial(w Window, label string, value float32, valueStr string) bool {
	top, bottom := w.SplitH(size + 8)
	w, bottom = bottom.SplitH(bottom.H - (size + 8))

	c := w.AlignedPoint(nvg.AlignCenter | nvg.AlignMiddle)
	r := w.RadiusMin()

	// Labels on top and bottom:
	top = top.Inner(w.W*0.5-r, 0, w.W*0.5-r, 0)
	bottom = bottom.Inner(w.W*0.5-r, 0, w.W*0.5-r, 0)

	ui.Label(top, valueStr, nvg.AlignCenter|nvg.AlignMiddle)
	ui.Label(bottom, label, nvg.AlignCenter|nvg.AlignMiddle)

	arcWidth := r * 0.05

	// Clamp value between 0 and 1:
	clampedValue := value
	if clampedValue > 1.0 {
		clampedValue = 1.0
	} else if clampedValue < 0.0 {
		clampedValue = 0.0
	}

	ui.Save()

	// Filled center:
	ui.BeginPath()
	ui.Circle(c, r-arcWidth)
	ui.FillColor(ui.Palette(2))
	ui.Fill()

	// Highlighted arc:
	ui.BeginPath()
	ui.Arc(c, r-arcWidth*0.5, nvg.Pi*0.8, nvg.Pi*0.8+nvg.Pi*1.4*clampedValue, nvg.Cw)
	ui.StrokeWidth(arcWidth)
	ui.MiterLimit(1.0)
	ui.LineCap(nvg.Square)
	ui.StrokeColor(ui.Palette(3))
	ui.Stroke()

	ui.Restore()

	return ui.isTouched(w)
}
