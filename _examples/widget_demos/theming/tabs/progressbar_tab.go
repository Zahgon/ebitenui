package tabs

import (
	"github.com/ebitenui/ebitenui/widget"
)

func NewProgressBarTab() *widget.TabBookTab { _ = "STUB: not implemented"; return nil }

// Construct a container to hold the progress bars.

// The container will use a vertical row layout to lay out the progress
// bars in a vertical row.

// Set the required anchor layout data to determine where in the root
// container to place the progress bars.

// Construct a horizontal progress bar.

// Set the minimum size for the progress bar.
// This is necessary if you wish to have the progress bar be larger than
// the provided track image. In this exampe since we are using NineSliceColor
// which is 1px x 1px we must set a minimum size.

// Set the progress images (Idle, Disabled).

// Set the min, max, and current values.

// Construct a vertical inverted progress bar.

// Set the direction of the progress bar to vertical.

// Invert the progress bar, meaning here it will fill from the bottom to the top
// since it’s vertical.

// Set the minimum size for the progress bar.
// This is necessary if you wish to have the progress bar be larger than
// the provided track image. In this example since we are using NineSliceColor
// which is 1px x 1px we must set a minimum size.

// Set the progress bar in the middle of the row container cell

// Set the min, max, and current values.
