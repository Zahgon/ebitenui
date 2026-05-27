package widget

import (
	img "image"
	"image/color"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ButtonParams struct {
	Image        *ButtonImage
	GraphicImage *GraphicImage
	TextColor    *ButtonTextColor

	TextPosition   *TextPositioning
	TextPadding    *Insets
	TextFace       *text.Face
	GraphicPadding *Insets
	MinSize        *img.Point
}

type Button struct {
	definedParams           ButtonParams
	computedParams          ButtonParams
	IgnoreTransparentPixels bool
	KeepPressedOnExit       bool
	ToggleMode              bool
	// Allows the user to disable space bar and enter automatically triggering a focused button.
	DisableDefaultKeys bool

	PressedEvent       *event.Event
	ReleasedEvent      *event.Event
	ClickedEvent       *event.Event
	CursorEnteredEvent *event.Event
	CursorMovedEvent   *event.Event
	CursorExitedEvent  *event.Event
	StateChangedEvent  *event.Event

	widgetOpts               []WidgetOpt
	autoUpdateTextAndGraphic bool

	init              *MultiOnce
	widget            *Widget
	container         *Container
	graphic           *Graphic
	text              *Text
	textLabel         string
	textProcessBBCode bool
	hovering          bool
	pressing          bool
	state             WidgetState

	tabOrder      int
	focused       bool
	justSubmitted bool

	focusMap map[FocusDirection]Focuser
}

type ButtonOpt func(b *Button)

type ButtonImage struct {
	Idle            *image.NineSlice
	Hover           *image.NineSlice
	Pressed         *image.NineSlice
	PressedHover    *image.NineSlice
	Disabled        *image.NineSlice
	PressedDisabled *image.NineSlice
}

type ButtonTextColor struct {
	Idle     color.Color
	Disabled color.Color
	Hover    color.Color
	Pressed  color.Color
}

type ButtonPressedEventArgs struct {
	Button  *Button
	OffsetX int
	OffsetY int
}

type ButtonReleasedEventArgs struct {
	Button  *Button
	Inside  bool
	OffsetX int
	OffsetY int
}

type ButtonClickedEventArgs struct {
	Button  *Button
	OffsetX int
	OffsetY int
}

type ButtonHoverEventArgs struct {
	Button  *Button
	Entered bool
	OffsetX int
	OffsetY int
	DiffX   int
	DiffY   int
}

type ButtonChangedEventArgs struct {
	Button  *Button
	State   WidgetState
	OffsetX int
	OffsetY int
}

type ButtonPressedHandlerFunc func(args *ButtonPressedEventArgs)

type ButtonReleasedHandlerFunc func(args *ButtonReleasedEventArgs)

type ButtonClickedHandlerFunc func(args *ButtonClickedEventArgs)

type ButtonCursorHoverHandlerFunc func(args *ButtonHoverEventArgs)

type ButtonChangedHandlerFunc func(args *ButtonChangedEventArgs)

type ButtonOptions struct {
}

var ButtonOpts ButtonOptions

func NewButton(opts ...ButtonOpt) *Button { _ = "STUB: not implemented"; return nil }

func (b *Button) Validate() { _ = "STUB: not implemented"; return }

func (b *Button) populateComputedParams() { _ = "STUB: not implemented"; return }

// clone the theme

func (o ButtonOptions) WidgetOpts(opts ...WidgetOpt) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) Image(i *ButtonImage) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

// IgnoreTransparentPixels disables mouse events like cursor entered,
// moved and exited if the mouse pointer is over a pixel that is transparent
// (alpha = 0). The source of pixels is Image.Idle. This options is
// especially useful, if your button does not have a rectangular shape.
func (o ButtonOptions) IgnoreTransparentPixels(ignoreTransparentPixels bool) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

// Text combines three options: TextLabel, TextFace and TextColor.
// It can be used for the inline configurations of Text object while
// separate functions are useful for a multi-step configuration.
func (o ButtonOptions) Text(label string, face *text.Face, color *ButtonTextColor) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) TextLabel(label string) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) TextFace(face *text.Face) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) TextColor(color *ButtonTextColor) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) TextProcessBBCode(enabled bool) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

// TODO: add parameter for image position (start/end).
func (o ButtonOptions) TextAndImage(label string, face *text.Face, image *GraphicImage, color *ButtonTextColor) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

// TextPosition sets the horizontal and vertical position of the text within the button.
// Default is TextPositionCenter for both.
func (o ButtonOptions) TextPosition(h TextPosition, v TextPosition) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) TextPadding(p *Insets) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) Graphic(image *GraphicImage) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) GraphicPadding(i Insets) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) KeepPressedOnExit() ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) ToggleMode() ButtonOpt { _ = "STUB: not implemented"; return *new(ButtonOpt) }

// This option will disable enter and space from submitting a focused button.
func (o ButtonOptions) DisableDefaultKeys() ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) PressedHandler(f ButtonPressedHandlerFunc) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) ReleasedHandler(f ButtonReleasedHandlerFunc) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) ClickedHandler(f ButtonClickedHandlerFunc) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) CursorEnteredHandler(f ButtonCursorHoverHandlerFunc) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) CursorMovedHandler(f ButtonCursorHoverHandlerFunc) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) CursorExitedHandler(f ButtonCursorHoverHandlerFunc) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) StateChangedHandler(f ButtonChangedHandlerFunc) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (o ButtonOptions) TabOrder(tabOrder int) ButtonOpt {
	_ = "STUB: not implemented"
	return *new(ButtonOpt)
}

func (b *Button) State() WidgetState { _ = "STUB: not implemented"; return *new(WidgetState) }

func (b *Button) SetState(state WidgetState) { _ = "STUB: not implemented"; return }

func (b *Button) getStateChangedEvent() *event.Event {
	_ = "STUB: not implemented"
	return nil

	/** Focuser Interface - Start **/
}

func (b *Button) Focus(focused bool) { _ = "STUB: not implemented"; return }

func (b *Button) IsFocused() bool { _ = "STUB: not implemented"; return false }

func (b *Button) TabOrder() int { _ = "STUB: not implemented"; return 0 }

func (b *Button) GetFocus(direction FocusDirection) Focuser {
	_ = "STUB: not implemented"
	return *new(Focuser)
}

func (b *Button) AddFocus(direction FocusDirection, focus Focuser) {
	_ = "STUB: not implemented"
	return
}

/** Focuser Interface - End **/

func (b *Button) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (b *Button) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (b *Button) SetLocation(rect img.Rectangle) {
	_ = "STUB: not implemented"

	// If we are ignoring transparent pixels and the mask isn't set to the current Image/Size, rebuild the mask.
	// If the rect is 0,0 we need to return because the 'ebiten.NewImage' will panic.
	// This is the case when we are on a window context
	return
}

func (b *Button) RequestRelayout() { _ = "STUB: not implemented"; return }

func (b *Button) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (b *Button) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

// We set the defaults first and then if needed
// they'll be overwritten by the other states

func (b *Button) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (b *Button) draw(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (b *Button) Click() { _ = "STUB: not implemented"; return }

// Press presses the button emulating a Mouse Left click.
func (b *Button) Press() { _ = "STUB: not implemented"; return }

// This means that there are some pixels that are not clickable.

// Release releases the button emulating a Mouse Left release.
func (b *Button) Release() { _ = "STUB: not implemented"; return }

// This means that there are some pixels that are not clickable.

func (b *Button) handleSubmit() { _ = "STUB: not implemented"; return }

func (b *Button) drawImageOptions(opts *ebiten.DrawImageOptions) { _ = "STUB: not implemented"; return }

func (b *Button) Text() *Text { _ = "STUB: not implemented"; return nil }

func (b *Button) SetText(text string) { _ = "STUB: not implemented"; return }

// This returns the currently defined GraphicImage object.
// This may be nil. Any changes to this reference will be reflected by the button
// but may be overwritten if the button is re-validated.
func (b *Button) GraphicImage() *GraphicImage { _ = "STUB: not implemented"; return nil }

// Set the GraphicImage for this button.
func (b *Button) SetGraphicImage(graphicImage *GraphicImage) { _ = "STUB: not implemented"; return }

// This returns the currently defined Image object.
// This may be nil. Any changes to this reference will be reflected by the button
// but may be overwritten if the button is re-validated.
func (b *Button) Image() *ButtonImage { _ = "STUB: not implemented"; return nil }

// Set the GraphicImage for this button.
func (b *Button) SetImage(image *ButtonImage) { _ = "STUB: not implemented"; return }

func (b *Button) initWidget() { _ = "STUB: not implemented"; return }

func (b *Button) createWidget() { _ = "STUB: not implemented"; return }
