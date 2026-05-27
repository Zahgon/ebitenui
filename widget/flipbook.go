package widget

import (
	img "image"

	"github.com/ebitenui/ebitenui/input"
	"github.com/hajimehoshi/ebiten/v2"
)

// A FlipBook is a container that always renders exactly one child widget: the current page.
// The current page will be embedded in a AnchorLayout.
type FlipBook struct {
	containerOpts    []ContainerOpt
	anchorLayoutOpts []AnchorLayoutOpt

	init          *MultiOnce
	container     *Container
	removeCurrent RemoveChildFunc
}

// FlipBookOpt is a function that configures f.
type FlipBookOpt func(f *FlipBook)

type FlipBookOptions struct {
}

// FlipBookOpts contains functions that configure a FlipBook.
var FlipBookOpts FlipBookOptions

// NewFlipBook constructs a new FlipBook configured with opts.
func NewFlipBook(opts ...FlipBookOpt) *FlipBook { _ = "STUB: not implemented"; return nil }

// WithContainerOpts configures a FlipBook with opts.
func (o FlipBookOptions) ContainerOpts(opts ...ContainerOpt) FlipBookOpt {
	_ = "STUB: not implemented"
	return *new(FlipBookOpt)
}

// WithPadding configures a FlipBook with padding i.
func (o FlipBookOptions) Padding(i *Insets) FlipBookOpt {
	_ = "STUB: not implemented"
	return *new(FlipBookOpt)
}

// GetWidget implements HasWidget.
func (f *FlipBook) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

// PreferredSize implements PreferredSizer.
func (f *FlipBook) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// SetLocation implements Locateable.
func (f *FlipBook) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (f *FlipBook) Validate() {
	_ = "STUB: not implemented"

	// RequestRelayout implements Relayoutable.
	return
}

func (f *FlipBook) RequestRelayout() { _ = "STUB: not implemented"; return }

// SetupInputLayer implements InputLayerer.
func (f *FlipBook) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

// Render implements Renderer.
func (f *FlipBook) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (f *FlipBook) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

// WidgetAt implements WidgetLocator.
func (f *FlipBook) WidgetAt(x int, y int) HasWidget {
	_ = "STUB: not implemented"
	return *new(HasWidget)
}

func (f *FlipBook) GetFocusers() []Focuser { _ = "STUB: not implemented"; return nil }

func (f *FlipBook) GetDropTargets() []HasWidget { _ = "STUB: not implemented"; return nil }

func (f *FlipBook) createWidget() { _ = "STUB: not implemented"; return }

// SetPage sets the current page to be rendered to page. The previous page will no longer be rendered.
//
// Note that when switching to a new page, it may be necessary to re-layout parent containers if the pages
// are of different sizes.
func (f *FlipBook) SetPage(page PreferredSizeLocateableWidget) { _ = "STUB: not implemented"; return }
