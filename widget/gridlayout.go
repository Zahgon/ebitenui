package widget

import (
	"image"
)

// GridLayout layouts widgets in a grid fashion, with columns or rows optionally being stretched.
//
// Widget.LayoutData of widgets being layouted by GridLayout need to be of type GridLayoutData.
type GridLayout struct {
	columns       int
	padding       *Insets
	columnSpacing int
	rowSpacing    int
	columnStretch []bool
	rowStretch    []bool

	defaultColumnStretch bool
	defaultRowStretch    bool
}

// GridLayoutOpt is a function that configures g.
type GridLayoutOpt func(g *GridLayout)

type GridLayoutOptions struct {
}

// GridLayoutData specifies layout settings for a widget.
type GridLayoutData struct {
	// MaxWidth specifies the maximum width.
	MaxWidth int

	// MaxHeight specifies the maximum height..
	MaxHeight int

	// HorizontalPosition specifies the horizontal anchoring position inside the grid cell.
	HorizontalPosition GridLayoutPosition

	// VerticalPosition specifies the vertical anchoring position inside the grid cell.
	VerticalPosition GridLayoutPosition
}

// GridLayoutPosition is the type used to specify an anchoring position.
type GridLayoutPosition int

const (
	// GridLayoutPositionStart is the anchoring position for "left" (in the horizontal direction) or "top" (in the vertical direction.)
	GridLayoutPositionStart = GridLayoutPosition(iota)

	// GridLayoutPositionStart is the center anchoring position.
	GridLayoutPositionCenter

	// GridLayoutPositionStart is the anchoring position for "right" (in the horizontal direction) or "bottom" (in the vertical direction.)
	GridLayoutPositionEnd
)

// GridLayoutOpts contains functions that configure a GridLayout.
var GridLayoutOpts GridLayoutOptions

// NewGridLayout constructs a new GridLayout, configured by opts.
func NewGridLayout(opts ...GridLayoutOpt) *GridLayout { _ = "STUB: not implemented"; return nil }

func (gl *GridLayout) validate() { _ = "STUB: not implemented"; return }

// Columns configures a grid layout to use c columns.
func (o GridLayoutOptions) Columns(c int) GridLayoutOpt {
	_ = "STUB: not implemented"
	return *new(GridLayoutOpt)
}

// Padding configures a grid layout to use padding i.
func (o GridLayoutOptions) Padding(i *Insets) GridLayoutOpt {
	_ = "STUB: not implemented"
	return *new(GridLayoutOpt)
}

// Spacing configures a grid layout to separate columns by spacing c and rows by spacing r.
func (o GridLayoutOptions) Spacing(c int, r int) GridLayoutOpt {
	_ = "STUB: not implemented"
	return *new(GridLayoutOpt)
}

// Stretch configures a grid layout to stretch columns according to c and rows according to r.
// The number of elements of c and r must correspond with the number of columns and rows in the
// layout.
func (o GridLayoutOptions) Stretch(c []bool, r []bool) GridLayoutOpt {
	_ = "STUB: not implemented"
	return *new(GridLayoutOpt)
}

// DefaultStretch will set the stretch value to the columns/rows that are
// extra not defined on the main Stretch.
func (o GridLayoutOptions) DefaultStretch(c bool, r bool) GridLayoutOpt {
	_ = "STUB: not implemented"
	return *new(GridLayoutOpt)
}

// PreferredSize implements Layouter.
func (g *GridLayout) PreferredSize(widgets []PreferredSizeLocateableWidget) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Layout implements Layouter.
func (g *GridLayout) Layout(widgets []PreferredSizeLocateableWidget, rect image.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (g *GridLayout) stretchedCellSizes(colWidths []int, rowHeights []int, rect image.Rectangle) (int, int, int, int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

func (g *GridLayout) columnStretched(c int) bool { _ = "STUB: not implemented"; return false }

func (g *GridLayout) rowStretched(r int) bool { _ = "STUB: not implemented"; return false }

func (g *GridLayout) preferredColumnWidthsAndRowHeights(widgets []PreferredSizeLocateableWidget) ([]int, []int) {
	_ = "STUB: not implemented"

	// Check if there are less widgets than columns declared
	// and if so use the len(widgets) as columns so the sizes
	// work as epxected
	return nil, nil
}

func (g *GridLayout) applyLayoutData(ld GridLayoutData, wx int, wy int, ww int, wh int, x int, y int, cw int, ch int) (int, int, int, int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// Do Nothing

// Do Nothing

func (g *GridLayout) applyMaxSize(ld GridLayoutData, ww int, wh int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func sumInts(ints []int) int { _ = "STUB: not implemented"; return 0 }
