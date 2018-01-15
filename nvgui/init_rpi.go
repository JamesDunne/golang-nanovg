//+build rpi

package nvgui

import (
	"github.com/JamesDunne/rpi-egl/bcm"

	gl "github.com/JamesDunne/rpi-egl/gles2"

	"github.com/JamesDunne/golang-nanovg/nvg"
)

func (ui *UI) InitDisplay() error {
	// Set up BCM display directly with an EGL context:
	//display, err := bcm.OpenDisplay(5, 6, 5)
	display, err := bcm.OpenDisplay(8, 8, 8)
	if err != nil {
		return err
	}
	ui.display = display
	display.SwapInterval(0)

	// Initialize NVG:
	ui.vg = nvg.CreateGLES2(nvg.Antialias)

	// Set up GL viewport:
	ui.w = Window{0, 0, float32(display.Width()), float32(display.Height())}
	gl.Viewport(0, 0, int32(ui.w.W), int32(ui.w.H))
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	return nil
}

func (ui *UI) Display() *bcm.Display {
	return ui.display.(*bcm.Display)
}

func (ui *UI) Close() {
	nvg.DeleteGLES2(ui.vg)
	ui.Display().Close()
}

func (ui *UI) BeginFrame() {
	// Clear background:
	gl.Clear(gl.COLOR_BUFFER_BIT)

	nvg.BeginFrame(ui.vg, int32(ui.w.W), int32(ui.w.H), 1.0)
}

func (ui *UI) EndFrame() {
	nvg.EndFrame(ui.vg)

	// Swap current surface to display:
	ui.Display().SwapBuffers()
}
