package widget

import (
	"github.com/ebitenui/ebitenui/image"
)

type TabBookTab struct {
	Container
	Disabled bool
	label    string
	image    *GraphicImage
}

type TabBookTabSelectedEventArgs struct {
	TabBook     *TabBook
	Tab         *TabBookTab
	PreviousTab *TabBookTab
}

type TabBookTabSelectedHandlerFunc func(args *TabBookTabSelectedEventArgs)

type TabParams struct {
	BackgroundImage *image.NineSlice
}

type TabBookTabOptions struct {
}

type TabBookTabOpt func(o *TabBookTab)

var TabBookTabOpts TabBookTabOptions

func (o *TabBookTabOptions) ContainerOpts(opts ...ContainerOpt) TabBookTabOpt {
	_ = "STUB: not implemented"
	return *new(TabBookTabOpt)
}

func (o *TabBookTabOptions) Image(img *GraphicImage) TabBookTabOpt {
	_ = "STUB: not implemented"
	return *new(TabBookTabOpt)
}

func (o *TabBookTabOptions) Label(label string) TabBookTabOpt {
	_ = "STUB: not implemented"
	return *new(TabBookTabOpt)
}

func NewTabBookTab(opts ...TabBookTabOpt) *TabBookTab { _ = "STUB: not implemented"; return nil }

// Set a default layout so that tabs use the full container

func (t *TabBookTab) Validate() { _ = "STUB: not implemented"; return }
