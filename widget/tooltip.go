package widget

import (
	"image"
	"image/color"
	"sync/atomic"
	"time"

	e_image "github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ToolTipPosition int

const (
	// The tooltip will follow the cursor around while visible.
	TOOLTIP_POS_CURSOR_FOLLOW ToolTipPosition = iota
	// The tooltip will stick to where the cursor was when the tooltip was made visible.
	TOOLTIP_POS_CURSOR_STICKY
	// The tooltip will display based on the Widget and Content anchor settings.
	// It defaults to opening right aligned and directly under the widget.
	TOOLTIP_POS_WIDGET
	// The tooltip will display based on x/y (offset is required)
	TOOLTIP_POS_ABSOLUTE
	// The tooltip will display based on the Widget and Content anchor settings.
	// It defaults to opening right aligned and directly under the x: 0, y: 0.
	TOOLTIP_POS_SCREEN
)

type ToolTipAnchor int

const (
	// Anchor at the start of the element.
	TOOLTIP_ANCHOR_START ToolTipAnchor = iota
	// Anchor in the middle of the element.
	TOOLTIP_ANCHOR_MIDDLE
	// Anchor at the end of the element.
	TOOLTIP_ANCHOR_END
)

type ToolTipDirection int

type ToolTip struct {
	Position ToolTipPosition
	// WidgetOriginVertical was renamed to AnchorOriginVertical to make the it more generic and reuse it for TOOLTIP_POS_SCREEN
	AnchorOriginVertical ToolTipAnchor
	// WidgetOriginHorizontal was renamed to AnchorOriginHorizontal to make the it more generic and reuse it for TOOLTIP_POS_SCREEN
	AnchorOriginHorizontal  ToolTipAnchor
	ContentOriginVertical   ToolTipAnchor
	ContentOriginHorizontal ToolTipAnchor
	Delay                   time.Duration
	Offset                  image.Point
	KeepOnHover             bool
	content                 Containerer
	window                  *Window
	visible                 bool

	state          toolTipState
	ToolTipUpdater ToolTipUpdater
}
type ToolTipOpt func(t *ToolTip)
type ToolTipOptions struct {
}

var ToolTipOpts ToolTipOptions

type toolTipState func(*Widget) toolTipState

type ToolTipUpdater func(Containerer)

// Create a new Tooltip. This method allows you to specify
// every aspect of the displayed tooltip's content.
func NewToolTip(opts ...ToolTipOpt) *ToolTip { _ = "STUB: not implemented"; return nil }

func (t *ToolTip) Validate() { _ = "STUB: not implemented"; return }

// Create a new Text Tooltip with the following defaults:
//   - ProcessBBCode = true
//   - Padding = Top/Bottom: 5px Left/Right: 10px
//   - Delay = 800ms
//   - Offset = 0, 20
//   - ContentOriginHorizontal = TOOLTIP_ANCHOR_END
//   - ContentOriginVertical = TOOLTIP_ANCHOR_START
func NewTextToolTip(label string, face *text.Face, color color.Color, background *e_image.NineSlice) *ToolTip {
	_ = "STUB: not implemented"
	return nil
}

// The container to be displayed.
func (o ToolTipOptions) Content(c Containerer) ToolTipOpt {
	_ = "STUB: not implemented"
	return *new(ToolTipOpt)
}

// The X/Y offsets from the Tooltip anchor point.
func (o ToolTipOptions) Offset(off image.Point) ToolTipOpt {
	_ = "STUB: not implemented"
	return *new(ToolTipOpt)
}

// The vertical position of the anchor on the widget. Only used when Postion = WIDGET.
func (o ToolTipOptions) AnchorOriginVertical(anchorOriginVertical ToolTipAnchor) ToolTipOpt {
	_ = "STUB: not implemented"
	return *new(ToolTipOpt)
}

// The horizontal position of the anchor on the widget. Only used when Postion = WIDGET.
func (o ToolTipOptions) AnchorOriginHorizontal(anchorOriginHorizontal ToolTipAnchor) ToolTipOpt {
	_ = "STUB: not implemented"
	return *new(ToolTipOpt)
}

// The vertical position of the anchor on the tooltip.
func (o ToolTipOptions) ContentOriginVertical(contentOriginVertical ToolTipAnchor) ToolTipOpt {
	_ = "STUB: not implemented"
	return *new(ToolTipOpt)
}

// The horizontal position of the anchor on the tooltip.
func (o ToolTipOptions) ContentOriginHorizontal(contentOriginHorizontal ToolTipAnchor) ToolTipOpt {
	_ = "STUB: not implemented"
	return *new(ToolTipOpt)
}

// Where to display the tooltip.
func (o ToolTipOptions) Position(position ToolTipPosition) ToolTipOpt {
	_ = "STUB: not implemented"
	return *new(ToolTipOpt)
}

// How long to wait before displaying the tooltip.
func (o ToolTipOptions) Delay(d time.Duration) ToolTipOpt {
	_ = "STUB: not implemented"
	return *new(ToolTipOpt)
}

// A method that is called every draw call that the tooltip is visible.
// This allows you to hook into the draw loop to update the tooltip if necessary.
func (o ToolTipOptions) ToolTipUpdater(toolTipUpdater ToolTipUpdater) ToolTipOpt {
	_ = "STUB: not implemented"
	return *new(ToolTipOpt)
}

// KeepOnHover will make it so if the user cursor is on the ToolTip it'll
// not be hidden as it does by default
func (o ToolTipOptions) KeepOnHover(b bool) ToolTipOpt {
	_ = "STUB: not implemented"
	return *new(ToolTipOpt)
}

func (t *ToolTip) Update(parent *Widget) { _ = "STUB: not implemented"; return }

func (t *ToolTip) idleState() toolTipState { _ = "STUB: not implemented"; return *new(toolTipState) }

func (t *ToolTip) armedState(p image.Point, timer *time.Timer, expired *atomic.Value) toolTipState {
	_ = "STUB: not implemented"
	return *new(toolTipState)
}

func (t *ToolTip) showingState(p image.Point) toolTipState {
	_ = "STUB: not implemented"
	return *new(toolTipState)
}

func (t *ToolTip) processWidgetPosition(widgetRect image.Rectangle) image.Point {
	_ = "STUB: not implemented"
	return *new(image.Point)
}

func (t *ToolTip) processScreenPosition() image.Point {
	_ = "STUB: not implemented"
	return *new(image.Point)
}

func (t *ToolTip) processContentPosition(p image.Point, sx, sy int, widgetRect image.Rectangle) image.Point {
	_ = "STUB: not implemented"
	return *new(image.Point)
}

func processContentPositionWorker(p image.Point, sx int, sy int, originHorizontal ToolTipAnchor, originVertical ToolTipAnchor) image.Point {
	_ = "STUB: not implemented"
	return *new(image.Point)
}

// Do nothing
