package widget

import (
	"image"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/input"
	"github.com/hajimehoshi/ebiten/v2"
)

// A Widget is an abstraction of a user interface widget, such as a button. Actual widget implementations
// "have" a Widget in their internal structure.
type Widget struct {
	// Rect specifies the widget's position on screen. It is usually not set directly, but a Layouter is
	// used to set the position in relation to other widgets or the space available.
	Rect image.Rectangle

	// If the Widget is not a perfect Rect, the specific implementation (like Button with IgnoreTransparentPixels)
	// will fill this mask that then will be used to precisely check if actions on this widget
	// are actually happening in it
	mask []byte

	// LayoutData specifies additional optional data for a Layouter that is used to layout this widget's
	// parent container. The exact type depends on the layout being used, for example, GridLayout requires
	// GridLayoutData to be used.
	LayoutData interface{}

	// The minimum width for this Widget
	MinWidth int

	// The minimum height for this Widget
	MinHeight int

	// Disabled specifies whether the widget is disabled, whatever that means. Disabled widgets should
	// usually render in some sort of "greyed out" visual state, and not react to user input.
	//
	// Not reacting to user input depends on the actual implementation. For example, List will not allow
	// entry selection via clicking, but the scrollbars will still be usable. The reasoning is that from
	// the user's perspective, scrolling does not change state, but only the display of that state.
	Disabled bool

	// Hidden specifies whether the widget is visible. Hidden widgets should
	// not render anything or react to user input.
	visibility Visibility

	// CursorEnterEvent fires an event with *WidgetCursorEnterEventArgs when the cursor enters the widget's Rect.
	CursorEnterEvent *event.Event

	// CursorMoveEvent fires an event with *WidgetCursorMoveEventArgs when the cursor moves within the widget's Rect.
	CursorMoveEvent *event.Event

	// CursorExitEvent fires an event with *WidgetCursorExitEventArgs when the cursor exits the widget's Rect.
	CursorExitEvent *event.Event

	// MouseButtonPressedEvent fires an event with *WidgetMouseButtonPressedEventArgs when a mouse button is pressed
	// while the cursor is inside the widget's Rect.
	MouseButtonPressedEvent *event.Event

	// MouseButtonPressedEvent fires an event with *WidgetMouseButtonPressedEventArgs when a mouse button is pressed
	// while the cursor is inside the widget's Rect.
	MouseButtonLongPressedEvent *event.Event

	// MouseButtonReleasedEvent fires an event with *WidgetMouseButtonReleasedEventArgs when a mouse button is released
	// while the cursor is inside the widget's Rect.
	MouseButtonReleasedEvent *event.Event

	// MouseButtonClickedEvent fires an event with *WidgetMouseButtonClickedEventArgs when a mouse button is pressed and released
	// while the cursor is inside the widget's Rect.
	MouseButtonClickedEvent *event.Event

	// ScrolledEvent fires an event with *WidgetScrolledEventArgs when the mouse wheel is scrolled while
	// the cursor is inside the widget's Rect.
	ScrolledEvent *event.Event

	FocusEvent *event.Event

	ContextMenuEvent *event.Event

	ToolTipEvent *event.Event

	DragAndDropEvent *event.Event

	OnUpdate UpdateFunc

	// Custom Data is a field to allow users to attach data to any widget
	CustomData any
	// This allows for non-focusable widgets (Containers) to report hover.
	TrackHover bool

	// This determines if the widget should use it's own layer.
	// The new layer will be added in the order that the widget is added to the render tree.
	// This means the last widiget added where this value is true will have the highest input layer.
	ElevateLayer bool

	canDrop CanDropFunc
	drop    DropFunc

	parent                      *Widget
	self                        HasWidget
	lastUpdateCursorEntered     bool
	lastUpdateCursorPosition    image.Point
	lastUpdateMouseLeftPressed  bool
	lastUpdateMouseRightPressed bool
	mouseLeftPressedInside      bool
	mouseRightPressedInside     bool
	inputLayer                  *input.Layer
	focusable                   Focuser
	theme                       *Theme
	longPressButton             ebiten.MouseButton
	longPressDuration           int
	longPressCurrent            int

	ContextMenu          *Container
	ContextMenuWindow    *Window
	ContextMenuCloseMode WindowCloseMode

	ToolTips []*ToolTip

	DragAndDrop *DragAndDrop

	CursorHovered string
	CursorPressed string

	relayoutParent bool
	debugMode      bool
}

// WidgetOpt is a function that configures w.
type WidgetOpt func(w *Widget) //nolint:golint

// HasWidget must be implemented by concrete widget types to get their Widget.
type HasWidget interface {
	GetWidget() *Widget
}

// Renderer may be implemented by concrete widget types that can render onto the screen.
type Renderer interface {
	// Render renders the widget onto screen.
	Render(screen *ebiten.Image)
}

type UpdateObject struct {
	RelayoutRequested     bool
	CloseEphemeralWindows bool
	DebugMode             bool
}

// Updater may be implemented by concrete widget types that should be updated.
type Updater interface {
	// Update updates the widget state based on input.
	Update(updObj *UpdateObject)
}

type FocusDirection int

const (
	FOCUS_NEXT FocusDirection = iota
	FOCUS_PREVIOUS
	FOCUS_NORTH
	FOCUS_NORTHEAST
	FOCUS_EAST
	FOCUS_SOUTHEAST
	FOCUS_SOUTH
	FOCUS_SOUTHWEST
	FOCUS_WEST
	FOCUS_NORTHWEST
)

type Focuser interface {
	HasWidget
	Focus(focused bool)
	IsFocused() bool
	TabOrder() int
	GetFocus(direction FocusDirection) Focuser
	AddFocus(direction FocusDirection, focus Focuser)
}

type Dropper interface {
	GetDropTargets() []HasWidget
}

type Visibility int

const (
	Visibility_Show          Visibility = iota
	Visibility_Hide_Blocking            // Hide widget, but take up space
	Visibility_Hide                     // Hide widget, but don't take up space
)

type RenderFunc func(screen *ebiten.Image)

type UpdateFunc func(w HasWidget)

// PreferredSizer may be implemented by concrete widget types that can report a preferred size.
type PreferredSizer interface {
	PreferredSize() (int, int)
}

type Containerer interface {
	Updater
	Renderer
	Dropper
	Relayoutable
	input.Layerer
	PreferredSizeLocateableWidget
	GetFocusers() []Focuser
	AddChild(children ...PreferredSizeLocateableWidget) RemoveChildFunc
	RemoveChild(child PreferredSizeLocateableWidget)
	RemoveChildren()
	Children() []PreferredSizeLocateableWidget
	IsValidated() bool
}

// WidgetCursorEnterEventArgs are the arguments for cursor enter events.
type WidgetCursorEnterEventArgs struct { //nolint:golint
	Widget *Widget

	// OffsetX is the x offset relative to the widget's Rect.
	OffsetX int

	// OffsetY is the y offset relative to the widget's Rect.
	OffsetY int
}

// WidgetCursorMoveEventArgs are the arguments for cursor move events.
type WidgetCursorMoveEventArgs struct { //nolint:golint
	Widget *Widget

	// OffsetX is the x offset relative to the widget's Rect.
	OffsetX int

	// OffsetY is the y offset relative to the widget's Rect.
	OffsetY int

	// DiffX is the x change to the old mouse cursor position.
	DiffX int

	// DiffY is the y change to the old mouse cursor position.
	DiffY int
}

// WidgetCursorExitEventArgs are the arguments for cursor exit events.
type WidgetCursorExitEventArgs struct { //nolint:golint
	Widget *Widget

	// OffsetX is the x offset relative to the widget's Rect.
	OffsetX int

	// OffsetY is the y offset relative to the widget's Rect.
	OffsetY int
}

// WidgetMouseButtonPressedEventArgs are the arguments for mouse button press events.
type WidgetMouseButtonPressedEventArgs struct { //nolint:golint
	Widget *Widget
	Button ebiten.MouseButton

	// OffsetX is the x offset relative to the widget's Rect.
	OffsetX int

	// OffsetY is the y offset relative to the widget's Rect.
	OffsetY int
}

// WidgetMouseButtonPressedEventArgs are the arguments for mouse button press events.
type WidgetMouseButtonLongPressedEventArgs struct { //nolint:golint
	Widget *Widget
	Button ebiten.MouseButton

	// OffsetX is the x offset relative to the widget's Rect.
	OffsetX int

	// OffsetY is the y offset relative to the widget's Rect.
	OffsetY int
}

// WidgetMouseButtonReleasedEventArgs are the arguments for mouse button release events.
type WidgetMouseButtonReleasedEventArgs struct { //nolint:golint
	Widget *Widget
	Button ebiten.MouseButton

	// Inside specifies whether the button has been released inside the widget's Rect.
	Inside bool

	// OffsetX is the x offset relative to the widget's Rect.
	OffsetX int

	// OffsetY is the y offset relative to the widget's Rect.
	OffsetY int
}

// WidgetMouseButtonClickedEventArgs are the arguments for mouse button press events.
type WidgetMouseButtonClickedEventArgs struct { //nolint:golint
	Widget *Widget
	Button ebiten.MouseButton

	// OffsetX is the x offset relative to the widget's Rect.
	OffsetX int

	// OffsetY is the y offset relative to the widget's Rect.
	OffsetY int
}

// WidgetScrolledEventArgs are the arguments for mouse wheel scroll events.
type WidgetScrolledEventArgs struct { //nolint:golint
	Widget *Widget
	X      float64
	Y      float64
}

type WidgetFocusEventArgs struct { //nolint:golint
	Widget   Focuser
	Focused  bool
	Location image.Point
}

type WidgetContextMenuEventArgs struct { //nolint:golint
	Widget   *Widget
	Location image.Point
}

type WidgetToolTipEventArgs struct { //nolint:golint
	Window *Window
	Show   bool
}

type WidgetDragAndDropEventArgs struct { //nolint:golint
	Window *Window
	Show   bool
	DnD    *DragAndDrop
}

type DragAndDropDroppedEventArgs struct { //nolint:golint
	Source  HasWidget
	SourceX int
	SourceY int
	Target  HasWidget
	TargetX int
	TargetY int
	Data    interface{}
}

type CanDropFunc func(args *DragAndDropDroppedEventArgs) bool
type DropFunc func(args *DragAndDropDroppedEventArgs)

type DragAndDropDroppedHandlerFunc func(args *DragAndDropDroppedEventArgs)

// WidgetCursorEnterHandlerFunc is a function that handles cursor enter events.
type WidgetCursorEnterHandlerFunc func(args *WidgetCursorEnterEventArgs) //nolint:golint

// WidgetCursorMoveHandlerFunc is a function that handles cursor move events.
type WidgetCursorMoveHandlerFunc func(args *WidgetCursorMoveEventArgs) //nolint:golint

// WidgetCursorExitHandlerFunc is a function that handles cursor exit events.
type WidgetCursorExitHandlerFunc func(args *WidgetCursorExitEventArgs) //nolint:golint

// WidgetMouseButtonPressedHandlerFunc is a function that handles mouse button press events.
type WidgetMouseButtonPressedHandlerFunc func(args *WidgetMouseButtonPressedEventArgs) //nolint:golint

// WidgetMouseButtonLongPressedHandlerFunc is a function that handles mouse button long press events (500ms).
type WidgetMouseButtonLongPressedHandlerFunc func(args *WidgetMouseButtonLongPressedEventArgs) //nolint:golint

// WidgetMouseButtonReleasedHandlerFunc is a function that handles mouse button release events.
type WidgetMouseButtonReleasedHandlerFunc func(args *WidgetMouseButtonReleasedEventArgs) //nolint:golint

// WidgetMouseButtonClickedHandlerFunc is a function that handles mouse button click events.
type WidgetMouseButtonClickedHandlerFunc func(args *WidgetMouseButtonClickedEventArgs) //nolint:golint

// WidgetScrolledHandlerFunc is a function that handles mouse wheel scroll events.
type WidgetScrolledHandlerFunc func(args *WidgetScrolledEventArgs) //nolint:golint

type WidgetOptions struct { //nolint:golint
}

// WidgetOpts contains functions that configure a Widget.
var WidgetOpts WidgetOptions

var deferredRenders []RenderFunc

// NewWidget constructs a new Widget configured with opts.
func NewWidget(opts ...WidgetOpt) *Widget { _ = "STUB: not implemented"; return nil }

// LayoutData configures a Widget with layout data ld.
func (o WidgetOptions) LayoutData(ld interface{}) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// CursorEnterHandler configures a Widget with cursor enter event handler f.
func (o WidgetOptions) CursorEnterHandler(f WidgetCursorEnterHandlerFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// CursorMoveHandler configures a Widget with cursor move event handler f.
func (o WidgetOptions) CursorMoveHandler(f WidgetCursorMoveHandlerFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// CursorExitHandler configures a Widget with cursor exit event handler f.
func (o WidgetOptions) CursorExitHandler(f WidgetCursorExitHandlerFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// MouseButtonPressedHandler configures a Widget with mouse button press event handler f.
func (o WidgetOptions) MouseButtonPressedHandler(f WidgetMouseButtonPressedHandlerFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// MouseButtonLongPressedHandler configures a Widget with mouse button long press event handler f.
// Triggered after holding down left or right mouse button 500ms.
func (o WidgetOptions) MouseButtonLongPressedHandler(f WidgetMouseButtonLongPressedHandlerFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// MouseButtonReleasedHandler configures a Widget with mouse button release event handler f.
func (o WidgetOptions) MouseButtonReleasedHandler(f WidgetMouseButtonReleasedHandlerFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// MouseButtonClickedHandler configures a Widget with mouse button release event handler f.
func (o WidgetOptions) MouseButtonClickedHandler(f WidgetMouseButtonClickedHandlerFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// ScrolledHandler configures a Widget with mouse wheel scroll event handler f.
func (o WidgetOptions) ScrolledHandler(f WidgetScrolledHandlerFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// CustomData configures a Widget with custom data cd.
func (o WidgetOptions) CustomData(cd any) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

func (o WidgetOptions) MinSize(minWidth int, minHeight int) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

func (o WidgetOptions) ContextMenu(contextMenu *Container) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

func (o WidgetOptions) ContextMenuCloseMode(contextMenuCloseMode WindowCloseMode) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

func (o WidgetOptions) ToolTip(toolTips ...*ToolTip) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// This sets the source of a Drag and Drop action.
func (o WidgetOptions) EnableDragAndDrop(d *DragAndDrop) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// This sets the widget as a target of a Drag and Drop action
//
//	The Drop function must return true if it accepts this drop and false if it does not accept the drop.
func (o WidgetOptions) CanDrop(candropFunc CanDropFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// This is the function that is run if an item is dropped on this widget.
func (o WidgetOptions) Dropped(dropFunc DropFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

func (o WidgetOptions) CursorHovered(cursorHovered string) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

func (o WidgetOptions) CursorPressed(cursorPressed string) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// This allows for non-focusable widgets (Containers) to report hover.
func (o WidgetOptions) TrackHover(trackHover bool) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// This tells the system to create a new input layer for this focusable widget.
// The new layer will be added in the order that the widget is added to the render tree.
// This means the last widiget added where this value is true will have the highest input layer.
func (o WidgetOptions) ElevateLayer(elevate bool) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// This specifies a function to be called each update loop for this widget.
func (o WidgetOptions) OnUpdate(updateFunc UpdateFunc) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

// This specifies how long in seconds a Long Press is.
// Must be greater than 0.
func (o WidgetOptions) LongPressDuration(seconds float64) WidgetOpt {
	_ = "STUB: not implemented"
	return *new(WidgetOpt)
}

func (w *Widget) drawImageOptions(opts *ebiten.DrawImageOptions) { _ = "STUB: not implemented"; return }

// EffectiveInputLayer returns w's effective input layer. If w does not have an input layer,
// or if the input layer is no longer valid, it returns w's parent widget's effective input layer.
// If w does not have a parent widget, it returns input.DefaultLayer.
func (w *Widget) EffectiveInputLayer() *input.Layer { _ = "STUB: not implemented"; return nil }

// always call this method first before rendering themselves.
func (w *Widget) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (w *Widget) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

// SetVisibility changes the visibility of the Widget
func (w *Widget) SetVisibility(v Visibility) { _ = "STUB: not implemented"; return }

// GetVisibility changes the visibility of the Widget
func (w *Widget) GetVisibility() Visibility { _ = "STUB: not implemented"; return *new(Visibility) }

func (w *Widget) fireEvents() { _ = "STUB: not implemented"; return }

// SetLocation sets w's position to rect. This is usually not called directly, but by a layout.
func (w *Widget) SetLocation(rect image.Rectangle) {
	_ = "STUB: not implemented"

	// ElevateToNewInputLayer adds l to the top of the input layer stack, then sets w's input layer to l.
	return
}

func (w *Widget) ElevateToNewInputLayer(l *input.Layer) { _ = "STUB: not implemented"; return }

func (w *Widget) Parent() *Widget { _ = "STUB: not implemented"; return nil }

func (widget *Widget) FireFocusEvent(w Focuser, focused bool, location image.Point) {
	_ = "STUB: not implemented" //nolint:golint
	return
}

func (widget *Widget) FireContextMenuEvent(w *Widget, l image.Point) {
	_ = "STUB: not implemented" //nolint:golint
	return
}

func (widget *Widget) FireToolTipEvent(w *Window, show bool) {
	_ = "STUB: not implemented" //nolint:golint
	return
}

func (widget *Widget) FireDragAndDropEvent(w *Window, show bool, dnd *DragAndDrop) {
	_ = "STUB: not implemented" //nolint:golint
	return
}

// IsVisible will check if this particular widget is visible by checking Visibility of it and
// all the parents it has, as if one of the parents is not visible this widget will not be visible
// even if it has Visibility_Show.
func (widget *Widget) IsVisible() bool { _ = "STUB: not implemented"; return false }

// RenderWithDeferred renders r to screen. This function should not be called directly.
func RenderDeferred(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func AppendToDeferredRenderQueue(r RenderFunc) { _ = "STUB: not implemented"; return }

// In checks if the x and y are inside of the widget
// even if they have a mask.
func (widget *Widget) In(x, y int) bool { _ = "STUB: not implemented"; return false }

func (widget *Widget) SetTheme(theme *Theme) { _ = "STUB: not implemented"; return }

func (widget *Widget) GetTheme() *Theme { _ = "STUB: not implemented"; return nil }
