package main

import (
	tk9 "modernc.org/tk9.0"
)

func main() {
	createUI()
	tk9.App.Wait()
}

func createUI() *tk9.CanvasWidget {
	photo := tk9.NewPhoto(tk9.File("sam.png"))

	tk9.App.Configure(
		tk9.Borderwidth(0),
		tk9.Highlightthickness(0),
		tk9.Padx(0),
		tk9.Pady(0),
	)

	// Size the dialog window to match the image dimensions.
	tk9.WmGeometry(tk9.App, photo.Width()+"x"+photo.Height())

	canvas := tk9.Canvas(
		tk9.Borderwidth(0),
		tk9.Highlightthickness(0),
	)

	canvas.CreateImage(0, 0, tk9.Anchor("nw"), tk9.Image(photo))
	tk9.Place(canvas, tk9.X(0), tk9.Y(0), tk9.Width(photo.Width()), tk9.Height(photo.Height()))

	close := tk9.Button(tk9.Txt("Close"), tk9.Command(func() {
		tk9.Destroy(tk9.App)
	}))
	tk9.Place(close, tk9.Relx(0.5), tk9.Rely(1.0), tk9.Anchor("s"), tk9.Y(-10))

	return canvas
}
