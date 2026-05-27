package widget

import (
	"image"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TabBookParams struct {
	TabButton      *ButtonParams
	TabSpacing     *int
	ContentSpacing *int
	ContentPadding *Insets
}

type TabBook struct {
	definedParams  TabBookParams
	computedParams TabBookParams

	TabSelectedEvent *event.Event

	tabs          []*TabBookTab
	containerOpts []ContainerOpt

	init         *MultiOnce
	container    *Container
	btnContainer *Container
	gridLayout   *GridLayout
	tabToButton  map[*TabBookTab]*Button
	flipBook     *FlipBook
	tab          *TabBookTab
	initialTab   *TabBookTab
}

type TabBookOpt func(t *TabBook)

type TabBookOptions struct {
}

var TabBookOpts TabBookOptions

func NewTabBook(opts ...TabBookOpt) *TabBook { _ = "STUB: not implemented"; return nil }

func (t *TabBook) Validate() { _ = "STUB: not implemented"; return }

func (t *TabBook) populateComputedParams() { _ = "STUB: not implemented"; return }

func (o TabBookOptions) ContainerOpts(opts ...ContainerOpt) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (o TabBookOptions) ContentPadding(contentPadding *Insets) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (o TabBookOptions) TabButtonImage(buttonImages *ButtonImage) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (o TabBookOptions) TabButtonSpacing(s int) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (o TabBookOptions) TabButtonText(face *text.Face, color *ButtonTextColor) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (o TabBookOptions) TabButtonTextPadding(textPadding *Insets) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (o TabBookOptions) TabButtonMinSize(minSize *image.Point) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (o TabBookOptions) ContentSpacing(s int) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (o TabBookOptions) Tabs(tabs ...*TabBookTab) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (o TabBookOptions) InitialTab(tab *TabBookTab) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (o TabBookOptions) TabSelectedHandler(f TabBookTabSelectedHandlerFunc) TabBookOpt {
	_ = "STUB: not implemented"
	return *new(TabBookOpt)
}

func (t *TabBook) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (t *TabBook) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (t *TabBook) SetLocation(rect image.Rectangle) { _ = "STUB: not implemented"; return }

func (t *TabBook) RequestRelayout() { _ = "STUB: not implemented"; return }

func (t *TabBook) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (t *TabBook) GetDropTargets() []HasWidget { _ = "STUB: not implemented"; return nil }

func (t *TabBook) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (t *TabBook) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (t *TabBook) createWidget() { _ = "STUB: not implemented"; return }

func (t *TabBook) initTabBook() { _ = "STUB: not implemented"; return }

// If we cannot find an initial tab default to to the first one

// Set the current tab for the tab book.
//
//		Note: This method should only be called after the
//	 ui is running. To set the initial tab please use the
//	 TabBookOptions.InitialTab method during tabbook creation.
func (t *TabBook) SetTab(tab *TabBookTab) { _ = "STUB: not implemented"; return }

// Return the currently selected tab.
func (t *TabBook) Tab() *TabBookTab {
	_ = "STUB: not implemented"

	// Return the button associated with the provided TabBookTab if not exists else nil.
	return nil
}

func (t *TabBook) GetTabButton(tab *TabBookTab) *Button { _ = "STUB: not implemented"; return nil }
