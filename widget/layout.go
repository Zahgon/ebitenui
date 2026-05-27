package widget

import (
	"image"
)

type Layouter interface {
	PreferredSize(widgets []PreferredSizeLocateableWidget) (int, int)
	Layout(widgets []PreferredSizeLocateableWidget, rect image.Rectangle)
}

type Relayoutable interface {
	RequestRelayout()
}

type Locateable interface {
	SetLocation(rect image.Rectangle)
}

type Locater interface {
	WidgetAt(x int, y int) HasWidget
}

type Insets struct {
	Top    int
	Left   int
	Right  int
	Bottom int
}

type Direction int

const (
	DirectionHorizontal = Direction(iota)
	DirectionVertical
)

func NewInsetsSimple(widthHeight int) *Insets { _ = "STUB: not implemented"; return nil }

func (i Insets) Apply(rect image.Rectangle) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (i Insets) Dx() int { _ = "STUB: not implemented"; return 0 }

func (i Insets) Dy() int { _ = "STUB: not implemented"; return 0 }
