package widget

import (
	img "image"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type ScrollContainer struct {
	ScrollLeft float64
	ScrollTop  float64

	widgetOpts          []WidgetOpt
	image               *ScrollContainerImage
	content             PreferredSizeLocateableWidget
	padding             *Insets
	stretchContentWidth bool

	init      *MultiOnce
	widget    *Widget
	validated bool
	renderBuf *image.MaskedRenderBuffer
}

type ScrollContainerOpt func(s *ScrollContainer)

type ScrollContainerImage struct {
	Idle     *image.NineSlice
	Disabled *image.NineSlice
	Mask     *image.NineSlice
}

type ScrollContainerOptions struct {
}

var ScrollContainerOpts ScrollContainerOptions

func NewScrollContainer(opts ...ScrollContainerOpt) *ScrollContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *ScrollContainer) Validate() { _ = "STUB: not implemented"; return }

func (o ScrollContainerOptions) WidgetOpts(opts ...WidgetOpt) ScrollContainerOpt {
	_ = "STUB: not implemented"
	return *new(ScrollContainerOpt)
}

func (o ScrollContainerOptions) Image(i *ScrollContainerImage) ScrollContainerOpt {
	_ = "STUB: not implemented"
	return *new(ScrollContainerOpt)
}

func (o ScrollContainerOptions) Content(c PreferredSizeLocateableWidget) ScrollContainerOpt {
	_ = "STUB: not implemented"
	return *new(ScrollContainerOpt)
}

func (o ScrollContainerOptions) Padding(p *Insets) ScrollContainerOpt {
	_ = "STUB: not implemented"
	return *new(ScrollContainerOpt)
}

func (o ScrollContainerOptions) StretchContentWidth() ScrollContainerOpt {
	_ = "STUB: not implemented"
	return *new(ScrollContainerOpt)
}

func (s *ScrollContainer) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (s *ScrollContainer) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (s *ScrollContainer) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *ScrollContainer) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (s *ScrollContainer) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (s *ScrollContainer) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (s *ScrollContainer) GetFocusers() []Focuser { _ = "STUB: not implemented"; return nil }

func (s *ScrollContainer) GetDropTargets() []HasWidget { _ = "STUB: not implemented"; return nil }

func (s *ScrollContainer) draw(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (s *ScrollContainer) drawImageOptions(opts *ebiten.DrawImageOptions) {
	_ = "STUB: not implemented"
	return
}

func (s *ScrollContainer) renderContent(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (s *ScrollContainer) ViewRect() img.Rectangle {
	_ = "STUB: not implemented"
	return *new(img.Rectangle)
}

func (s *ScrollContainer) ContentRect() img.Rectangle {
	_ = "STUB: not implemented"
	return *new(img.Rectangle)
}

func (s *ScrollContainer) clampScroll() { _ = "STUB: not implemented"; return }

func (s *ScrollContainer) createWidget() { _ = "STUB: not implemented"; return }
