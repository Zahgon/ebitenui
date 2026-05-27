package widget

import "image"

// AnchorLayout layouts widgets anchored to either a corner or edge of a rectangle,
// optionally stretching it in one or both directions.
//
// AnchorLayout will layout all widgets  in the container to the specified locations regardless of overlap.
// The widgets in the container will be drawn in the order they were added to the container.
//
// Widget.LayoutData of widgets being layouted by AnchorLayout need to be of type AnchorLayoutData.
type AnchorLayout struct {
	padding *Insets
}

// AnchorLayoutOpt is a function that configures a.
type AnchorLayoutOpt func(a *AnchorLayout)

type AnchorLayoutOptions struct {
}

// AnchorLayoutPosition is the type used to specify an anchoring position.
type AnchorLayoutPosition int

// AnchorLayoutData specifies layout settings for a widget.
type AnchorLayoutData struct {
	// HorizontalPosition specifies the horizontal anchoring position.
	HorizontalPosition AnchorLayoutPosition

	// VerticalPosition specifies the vertical anchoring position.
	VerticalPosition AnchorLayoutPosition

	// StretchHorizontal specifies whether to stretch in the horizontal direction.
	StretchHorizontal bool

	// StretchVertical specifies whether to stretch in the vertical direction.
	StretchVertical bool

	// Sets the padding for the child.
	Padding *Insets
}

const (
	// AnchorLayoutPositionStart is the anchoring position for "left" (in the horizontal direction) or "top" (in the vertical direction.)
	AnchorLayoutPositionStart = AnchorLayoutPosition(iota)

	// AnchorLayoutPositionCenter is the center anchoring position.
	AnchorLayoutPositionCenter

	// AnchorLayoutPositionEnd is the anchoring position for "right" (in the horizontal direction) or "bottom" (in the vertical direction.)
	AnchorLayoutPositionEnd
)

// AnchorLayoutOpts contains functions that configure an AnchorLayout.
var AnchorLayoutOpts AnchorLayoutOptions

// NewAnchorLayout constructs a new AnchorLayout, configured by opts.
func NewAnchorLayout(opts ...AnchorLayoutOpt) *AnchorLayout { _ = "STUB: not implemented"; return nil }

// Padding configures an anchor layout to use padding i. This affects all children.
func (o AnchorLayoutOptions) Padding(i *Insets) AnchorLayoutOpt {
	_ = "STUB: not implemented"
	return *new(AnchorLayoutOpt)
}

// PreferredSize implements Layouter.
func (a *AnchorLayout) PreferredSize(widgets []PreferredSizeLocateableWidget) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Layout implements Layouter.
func (a *AnchorLayout) Layout(widgets []PreferredSizeLocateableWidget, rect image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (a *AnchorLayout) applyLayoutData(ld AnchorLayoutData, wx int, wy int, ww int, wh int, rect image.Rectangle) (int, int, int, int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// Do nothing

// Do nothing
