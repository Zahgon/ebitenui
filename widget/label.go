package widget

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type LabelParams struct {
	Face    *text.Face
	Color   *LabelColor
	Padding *Insets
}

type Label struct {
	Label string

	definedParams  LabelParams
	computedParams LabelParams

	textOpts []TextOpt
	init     *MultiOnce
	text     *Text
}

type LabelOpt func(l *Label)

type LabelColor struct {
	Idle     color.Color
	Disabled color.Color
}

type LabelOptions struct {
}

var LabelOpts LabelOptions

func NewLabel(opts ...LabelOpt) *Label { _ = "STUB: not implemented"; return nil }

func (l *Label) Validate() { _ = "STUB: not implemented"; return }

func (l *Label) populateComputedParams() { _ = "STUB: not implemented"; return }

func (o LabelOptions) TextOpts(opts ...TextOpt) LabelOpt {
	_ = "STUB: not implemented"
	return *new(LabelOpt)
}

// Set the label text, font, and font colors.
func (o LabelOptions) Text(label string, face *text.Face, color *LabelColor) LabelOpt {
	_ = "STUB: not implemented"
	return *new(LabelOpt)
}

// Set the label text.
func (o LabelOptions) LabelText(label string) LabelOpt {
	_ = "STUB: not implemented"
	return *new(LabelOpt)
}

// Set the label font.
func (o LabelOptions) LabelFace(face *text.Face) LabelOpt {
	_ = "STUB: not implemented"
	return *new(LabelOpt)
}

// Set the label font colors.
func (o LabelOptions) LabelColor(color *LabelColor) LabelOpt {
	_ = "STUB: not implemented"
	return *new(LabelOpt)
}

// Set the label padding.
func (o LabelOptions) LabelPadding(padding *Insets) LabelOpt {
	_ = "STUB: not implemented"
	return *new(LabelOpt)
}

func (l *Label) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (l *Label) SetLocation(rect image.Rectangle) { _ = "STUB: not implemented"; return }

func (l *Label) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (l *Label) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (l *Label) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (l *Label) createWidget() { _ = "STUB: not implemented"; return }

func (l *Label) setChildComputedParams() { _ = "STUB: not implemented"; return }
