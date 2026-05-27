package widget

import (
	img "image"

	"github.com/ebitenui/ebitenui/image"

	"github.com/hajimehoshi/ebiten/v2"
)

type ProgressBarParams struct {
	TrackImage   *ProgressBarImage
	FillImage    *ProgressBarImage
	TrackPadding *Insets
}

type ProgressBar struct {
	definedParams  ProgressBarParams
	computedParams ProgressBarParams

	Min     int
	Max     int
	current int

	widgetOpts []WidgetOpt
	direction  Direction
	inverted   bool

	init   *MultiOnce
	widget *Widget
}

type ProgressBarImage struct {
	Idle     *image.NineSlice
	Hover    *image.NineSlice
	Disabled *image.NineSlice
}

type ProgressBarOpt func(s *ProgressBar)

type ProgressBarOptions struct {
}

var ProgressBarOpts ProgressBarOptions

func NewProgressBar(opts ...ProgressBarOpt) *ProgressBar { _ = "STUB: not implemented"; return nil }

func (pb *ProgressBar) Validate() { _ = "STUB: not implemented"; return }

func (pb *ProgressBar) populateComputedParams() { _ = "STUB: not implemented"; return }

func (o ProgressBarOptions) WidgetOpts(opts ...WidgetOpt) ProgressBarOpt {
	_ = "STUB: not implemented"
	return *new(ProgressBarOpt)
}

// Direction sets the direction of the progress bar.
// The default is horizontal.
func (o ProgressBarOptions) Direction(d Direction) ProgressBarOpt {
	_ = "STUB: not implemented"
	return *new(ProgressBarOpt)
}

// Inverted sets whether the progress bar is inverted.
// The default is false, which means from left to right or top to bottom.
func (o ProgressBarOptions) Inverted(inverted bool) ProgressBarOpt {
	_ = "STUB: not implemented"
	return *new(ProgressBarOpt)
}

func (o ProgressBarOptions) Images(track *ProgressBarImage, fill *ProgressBarImage) ProgressBarOpt {
	_ = "STUB: not implemented"
	return *new(ProgressBarOpt)
}

func (o ProgressBarOptions) TrackImage(track *ProgressBarImage) ProgressBarOpt {
	_ = "STUB: not implemented"
	return *new(ProgressBarOpt)
}

func (o ProgressBarOptions) FillImage(fill *ProgressBarImage) ProgressBarOpt {
	_ = "STUB: not implemented"
	return *new(ProgressBarOpt)
}

func (o ProgressBarOptions) TrackPadding(i *Insets) ProgressBarOpt {
	_ = "STUB: not implemented"
	return *new(ProgressBarOpt)
}

func (o ProgressBarOptions) Values(min int, max int, current int) ProgressBarOpt {
	_ = "STUB: not implemented"
	return *new(ProgressBarOpt)
}

func (s *ProgressBar) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (s *ProgressBar) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *ProgressBar) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (s *ProgressBar) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (s *ProgressBar) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (s *ProgressBar) draw(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (s *ProgressBar) currentPercentage() float64 { _ = "STUB: not implemented"; return 0 }

func (s *ProgressBar) SetCurrent(value int) bool { _ = "STUB: not implemented"; return false }

func (s *ProgressBar) GetCurrent() int { _ = "STUB: not implemented"; return 0 }

func (s *ProgressBar) createWidget() { _ = "STUB: not implemented"; return }
