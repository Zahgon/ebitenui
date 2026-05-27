package widget

import (
	"image"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type WindowCloseMode int

const (
	// The window will not automatically close.
	NONE WindowCloseMode = iota
	// The window will close when you click anywhere.
	CLICK
	// The window will close when you click outside the window.
	CLICK_OUT
)

type RemoveWindowFunc func()

type WindowChangedEventArgs struct {
	Window *Window
	Rect   image.Rectangle
}

type WindowChangedHandlerFunc func(args *WindowChangedEventArgs)

type WindowClosedEventArgs struct {
	Window *Window
}

type WindowClosedHandlerFunc func(args *WindowClosedEventArgs)

type Window struct {
	ResizeEvent *event.Event
	MoveEvent   *event.Event
	ClosedEvent *event.Event

	Modal      bool
	Contents   Containerer
	TitleBar   Containerer
	Draggable  bool
	Resizeable bool
	MinSize    *image.Point
	MaxSize    *image.Point
	Dynamic    bool
	DrawLayer  int
	// Used to indicate this window should close if other windows close.
	Ephemeral bool

	// INTERNAL USE ONLY: Used to indicate that this window is the currently focused window.
	FocusedWindow bool
	// Sets whether the window should move to the top of the draw order when clicked
	// Default: false.
	DisableRelayering bool

	closeMode WindowCloseMode
	closeFunc RemoveWindowFunc
	container Containerer

	titleBarHeight int

	startingPoint  image.Point
	dragging       bool
	resizing       bool
	resizingWidth  bool
	resizingHeight bool
	blockLower     bool
	originalSize   image.Point
	init           *MultiOnce
}

type WindowOpt func(w *Window)

type WindowOptions struct {
}

var WindowOpts WindowOptions

func NewWindow(opts ...WindowOpt) *Window { _ = "STUB: not implemented"; return nil }

func (w *Window) Validate() { _ = "STUB: not implemented"; return }

// This is the container with the body of this window.
func (o WindowOptions) Contents(c Containerer) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// Sets the container for the TitleBar and its fixed height.
func (o WindowOptions) TitleBar(tb Containerer, height int) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// Sets the window to be modal. Blocking UI interactions on anything else.
func (o WindowOptions) Modal() WindowOpt { _ = "STUB: not implemented"; return *new(WindowOpt) }

// Sets the window to be draggable. The handle for this is the titleBar.
// If you haven't provided a titleBar this option is ignored.
func (o WindowOptions) Draggable() WindowOpt { _ = "STUB: not implemented"; return *new(WindowOpt) }

// Sets the window to be resizeable.
func (o WindowOptions) Resizeable() WindowOpt { _ = "STUB: not implemented"; return *new(WindowOpt) }

// Sets whether the window should block input beneath the window or not.
// Default: true.
func (o WindowOptions) BlockLower(blockLower bool) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// Sets whether the window should move to the top of the draw order when clicked
// Default: false.
func (o WindowOptions) DisableRelayering(disableRelayering bool) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// Sets the minimum size that the window can be reszied to.
func (o WindowOptions) MinSize(width int, height int) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// Set the maximum size that the window can be resized to.
func (o WindowOptions) MaxSize(width int, height int) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// Dynamic will check to change the default size of the Window
// if the content sizes goes over or below the current one
// Will not work with Resizeable
func (o WindowOptions) Dynamic() WindowOpt { _ = "STUB: not implemented"; return *new(WindowOpt) }

// Set the way this window should close.
func (o WindowOptions) CloseMode(mode WindowCloseMode) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// This handler is triggered when a move event is completed.
func (o WindowOptions) MoveHandler(f WindowChangedHandlerFunc) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// This handler is triggered when a resize event is completed.
func (o WindowOptions) ResizeHandler(f WindowChangedHandlerFunc) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// This handler is triggered when window is closed.
// The window can be closed by either calling [Window.Close] or
// by invoking an [UI.AddWindow] returned close function.
//
// This handler is called after the window is closed.
// The provided Window object is still accessible, but the window
// is already removed from UI.
func (o WindowOptions) ClosedHandler(f WindowClosedHandlerFunc) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// This option sets the size and location of the window.
// This method will account for specified MinSize and MaxSize values.
func (o WindowOptions) Location(rect image.Rectangle) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// This option sets order the window will be drawn
//   - &lt; 0 will have the window drawn before the container
//   - &gt;= 0 will have the window drawn after the container
func (o WindowOptions) DrawLayer(layer int) WindowOpt {
	_ = "STUB: not implemented"
	return *new(WindowOpt)
}

// This method is used to be able to close the window.
func (w *Window) Close() { _ = "STUB: not implemented"; return }

// This method will set the size and location of this window.
// This method will account for specified MinSize and MaxSize values.
func (w *Window) SetLocation(rect image.Rectangle) { _ = "STUB: not implemented"; return }

// Typically used internally.
//
//	Returns the root container that holds the provided titlebar and contents.
func (w *Window) GetContainer() Containerer {
	_ = "STUB: not implemented"

	// Typically used internally.
	return *new(Containerer)
}

func (w *Window) SetCloseFunction(removeWindowFunc RemoveWindowFunc) {
	_ = "STUB: not implemented"
	return
}

// Typically used internally.
func (w *Window) GetCloseFunction() RemoveWindowFunc {
	_ = "STUB: not implemented"
	return *

	// Typically used internally.
	new(RemoveWindowFunc)
}

func (w *Window) RequestRelayout() { _ = "STUB: not implemented"; return }

// Typically used internally.
func (w *Window) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

// Typically used internally.
func (w *Window) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (w *Window) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (w *Window) createWidget() { _ = "STUB: not implemented"; return }
