package widget

import (
	"image"

	"github.com/ebitenui/ebitenui/input"
)

type DragAndDropAnchor int

const (
	// Anchor at the start of the element.
	DND_ANCHOR_START DragAndDropAnchor = iota
	// Anchor in the middle of the element.
	DND_ANCHOR_MIDDLE
	// Anchor at the end of the element.
	DND_ANCHOR_END
)

type DragAndDrop struct {
	ContentsOriginVertical   DragAndDropAnchor
	ContentsOriginHorizontal DragAndDropAnchor
	Offset                   image.Point

	AvailableDropTargets []HasWidget
	contentsCreater      DragContentsCreater
	minDragStartDistance int
	state                dragAndDropState
	dragWidget           *Container
	window               *Window
	dndTriggered         bool
	dndStopped           bool
	dragDisabled         bool
}

type DragAndDropOpt func(d *DragAndDrop)

type DragAndDropOptions struct {
}

var DragAndDropOpts DragAndDropOptions

type DragContentsCreater interface {
	Create(HasWidget) (*Container, interface{})
}

type DragContentsUpdater interface {
	// arg1 - isDroppable
	// arg2 - HasWidget if droppable
	// arg3 - DragData
	Update(bool, HasWidget, interface{})
}

type DragContentsEnder interface {
	// arg1 - Drop was successful
	// arg2 - Source Widget
	// arg3 - DragData
	EndDrag(bool, HasWidget, interface{})
}

type dragAndDropState func(HasWidget) (dragAndDropState, bool)

func NewDragAndDrop(opts ...DragAndDropOpt) *DragAndDrop { _ = "STUB: not implemented"; return nil }

func (d *DragAndDrop) Validate() { _ = "STUB: not implemented"; return }

func (o DragAndDropOptions) ContentsCreater(c DragContentsCreater) DragAndDropOpt {
	_ = "STUB: not implemented"
	return *new(DragAndDropOpt)
}

// The minimum distance in pixels a user must drag their cursor to display the dragged element.
//
//	Optional - Defaults to 15 pixels
func (o DragAndDropOptions) MinDragStartDistance(d int) DragAndDropOpt {
	_ = "STUB: not implemented"
	return *new(DragAndDropOpt)
}

// The vertical position of the anchor on the tooltip.
//
//	Optional - Defaults to DND_ANCHOR_MIDDLE
func (o DragAndDropOptions) ContentsOriginVertical(contentsOriginVertical DragAndDropAnchor) DragAndDropOpt {
	_ = "STUB: not implemented"
	return *new(DragAndDropOpt)
}

// The horizontal position of the anchor on the tooltip.
//
//	Optional - Defaults to DND_ANCHOR_MIDDLE
func (o DragAndDropOptions) ContentsOriginHorizontal(contentsOriginHorizontal DragAndDropAnchor) DragAndDropOpt {
	_ = "STUB: not implemented"
	return *new(DragAndDropOpt)
}

// The X/Y offsets from the Tooltip anchor point.
func (o DragAndDropOptions) Offset(off image.Point) DragAndDropOpt {
	_ = "STUB: not implemented"
	return *new(DragAndDropOpt)
}

// Disable Drag to start Drag and Drop.
// You may use the "StartDrag()" method on this object or the
// to begin the drag operation.
//
//	Expected use-case: click to pick up, click to drop.
func (o DragAndDropOptions) DisableDrag() DragAndDropOpt {
	_ = "STUB: not implemented"
	return *new(DragAndDropOpt)
}

// To avoid conflicting with dragging, if you trigger it on left click you should put the trigger in the button released event.
func (d *DragAndDrop) StartDrag() { _ = "STUB: not implemented"; return }

func (d *DragAndDrop) StopDrag() { _ = "STUB: not implemented"; return }

func (d *DragAndDrop) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (d *DragAndDrop) Update(parent HasWidget) { _ = "STUB: not implemented"; return }

func (d *DragAndDrop) idleState() dragAndDropState {
	_ = "STUB: not implemented"
	return *new(dragAndDropState)
}

func (d *DragAndDrop) dragArmedState(srcX int, srcY int) dragAndDropState {
	_ = "STUB: not implemented"
	return *new(dragAndDropState)
}

func (d *DragAndDrop) draggingState(srcX int, srcY int, dragWidget *Container, dragData interface{}, mousePressed bool) dragAndDropState {
	_ = "STUB: not implemented"
	return *new(dragAndDropState)
}

func (d *DragAndDrop) droppingState(srcX int, srcY int, x int, y int, dragData interface{}) dragAndDropState {
	_ = "STUB: not implemented"
	return *new(dragAndDropState)
}

func (d *DragAndDrop) processContentsPosition(p image.Point, sx int, sy int) image.Point {
	_ = "STUB: not implemented"
	return *new(image.Point)
}

// Do nothing
