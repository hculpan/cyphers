package layout

import "fyne.io/fyne"

type LabelLayout struct {
}

func NewLabelLayout() *LabelLayout {
	return &LabelLayout{}
}

func (l *LabelLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}

func (l *LabelLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	objects[0].Resize(fyne.NewSize(385, 430))
	objects[0].Move(fyne.NewPos(0, 0))
}
