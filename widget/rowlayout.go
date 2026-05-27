package widget

import (
	"image"
)

// RowLayout layouts widgets in either a single row or a single column,
// optionally stretching them in the other direction.
//
// Widget.LayoutData of widgets being layouted by RowLayout need to be of type RowLayoutData.
type RowLayout struct {
	direction Direction
	padding   *Insets
	spacing   int
}

type RowLayoutOptions struct {
}

// RowLayoutOpt is a function that configures r.
type RowLayoutOpt func(r *RowLayout)

// RowLayoutData specifies layout settings for a widget.
type RowLayoutData struct {
	// Position specifies the anchoring position for the direction that is not the primary direction of the layout.
	Position RowLayoutPosition

	// Stretch specifies whether to stretch in the direction that is not the primary direction of the layout.
	Stretch bool

	// MaxWidth specifies the maximum width.
	MaxWidth int

	// MaxHeight specifies the maximum height.
	MaxHeight int
}

// RowLayoutPosition is the type used to specify an anchoring position.
type RowLayoutPosition int

const (
	// RowLayoutPositionStart is the anchoring position for "left" (in the horizontal direction) or "top" (in the vertical direction.)
	RowLayoutPositionStart = RowLayoutPosition(iota)

	// RowLayoutPositionCenter is the center anchoring position.
	RowLayoutPositionCenter

	// RowLayoutPositionEnd is the anchoring position for "right" (in the horizontal direction) or "bottom" (in the vertical direction.)
	RowLayoutPositionEnd
)

// RowLayoutOpts contains functions that configure a RowLayout.
var RowLayoutOpts RowLayoutOptions

// NewRowLayout constructs a new RowLayout, configured by opts.
func NewRowLayout(opts ...RowLayoutOpt) *RowLayout { _ = "STUB: not implemented"; return nil }

// Direction configures a row layout to layout widgets in the primary direction d. This will also switch the meaning
// of any widget's RowLayoutData.Position and RowLayoutData.Stretch to the other direction.
func (o RowLayoutOptions) Direction(d Direction) RowLayoutOpt {
	_ = "STUB: not implemented"
	return *new(RowLayoutOpt)
}

// Padding configures a row layout to use padding i.
func (o RowLayoutOptions) Padding(i *Insets) RowLayoutOpt {
	_ = "STUB: not implemented"
	return *new(RowLayoutOpt)
}

// Spacing configures a row layout to separate widgets by spacing s.
func (o RowLayoutOptions) Spacing(s int) RowLayoutOpt {
	_ = "STUB: not implemented"
	return *new(RowLayoutOpt)
}

// PreferredSize implements Layouter.
func (r *RowLayout) PreferredSize(widgets []PreferredSizeLocateableWidget) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Layout implements Layouter.
func (r *RowLayout) Layout(widgets []PreferredSizeLocateableWidget, rect image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (r *RowLayout) layout(widgets []PreferredSizeLocateableWidget, rect image.Rectangle, usePosition bool, locationFunc func(w PreferredSizeLocateableWidget, wr image.Rectangle)) {
	_ = "STUB: not implemented"
	return
}

func (r *RowLayout) applyLayoutData(ld RowLayoutData, wx int, wy int, ww int, wh int, usePosition bool, rect image.Rectangle, x int, y int) (int, int, int, int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

func (r *RowLayout) applyStretch(ld RowLayoutData, ww int, wh int, rect image.Rectangle) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (r *RowLayout) applyMaxSize(ld RowLayoutData, ww int, wh int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (r *RowLayout) applyPosition(ld RowLayoutData, wx int, wy int, ww int, wh int, rect image.Rectangle, x int, y int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Do Nothing
