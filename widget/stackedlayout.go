package widget

import (
	"image"
)

// StackedLayout lays out multiple widgets stacked on top of each other in the order they are added as children
// Each child will have the dimensions equal to the max size of all the StackedLayout's children
//
// Note: Events will propogate to each layer. e.g. if you have overlapping buttons, a click event over both will trigger both events.
//
// Widget.LayoutData of widgets being layed out by StackedLayout should be left empty.
type StackedLayout struct {
	padding *Insets
}

// StackedLayoutOpt is a function that configures a.
type StackedLayoutOpt func(a *StackedLayout)

type StackedLayoutOptions struct {
}

// StackedLayoutData specifies layout settings for a widget.
type StackedLayoutData struct {
}

// StackedLayoutOpts contains functions that configure an StackedLayout.
var StackedLayoutOpts StackedLayoutOptions

// NewStackedLayout constructs a new StackedLayout, configured by opts.
func NewStackedLayout(opts ...StackedLayoutOpt) *StackedLayout {
	_ = "STUB: not implemented"
	return nil
}

// Padding configures an Stacked layout to use padding i.
func (o StackedLayoutOptions) Padding(i *Insets) StackedLayoutOpt {
	_ = "STUB: not implemented"
	return *new(StackedLayoutOpt)
}

// PreferredSize implements Layouter.
func (a *StackedLayout) PreferredSize(widgets []PreferredSizeLocateableWidget) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Layout implements Layouter.
func (a *StackedLayout) Layout(widgets []PreferredSizeLocateableWidget, rect image.Rectangle) {
	_ = "STUB: not implemented"
	return
}
