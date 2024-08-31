package main

import "fyne.io/fyne/v2"

type InputLayout struct{}

func (d *InputLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h*0.5)
}

func (d *InputLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	objects[0].Resize(fyne.NewSize(containerSize.Width*0.75, containerSize.Height))
	objects[0].Move(fyne.NewPos(0, 0))

	objects[1].Resize(fyne.NewSize(containerSize.Width*0.25, containerSize.Height))
	objects[1].Move(fyne.NewPos(containerSize.Width*0.75, 0))
}
