package widget

import (
	img "image"
	"image/color"

	"github.com/ebitenui/ebitenui/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextAreaParams struct {
	Face                   *text.Face
	ForegroundColor        color.Color
	TextPadding            *Insets
	TextPosition           *TextPositioning
	ControlWidgetSpacing   *int
	StripBBCode            *bool
	LinkColor              *TextLinkColor
	Slider                 *SliderParams
	ScrollContainerImage   *ScrollContainerImage
	ScrollContainerPadding *Insets
}

type TextArea struct {
	definedParams  TextAreaParams
	computedParams TextAreaParams

	containerOpts []ContainerOpt

	processBBCode        bool
	initialText          string
	verticalScrollMode   ScrollMode
	horizontalScrollMode ScrollMode
	showHorizontalSlider bool
	showVerticalSlider   bool

	linkClickedFunc       LinkHandlerFunc
	linkCursorEnteredFunc LinkHandlerFunc
	linkCursorExitedFunc  LinkHandlerFunc

	init            *MultiOnce
	container       *Container
	layout          *GridLayout
	scrollContainer *ScrollContainer
	vSlider         *Slider
	hSlider         *Slider
	text            *Text
}

type ScrollMode int

const (
	// Default. Scrolling is not automatically handled.
	None ScrollMode = iota

	// The TextArea is automatically scrolled to the beginning on change.
	ScrollBeginning

	// The TextArea is automatically scrolled to the end on change.
	ScrollEnd

	// The TextArea will initially position the text at the end of the scroll area.
	PositionAtEnd
)

type TextAreaOpt func(l *TextArea)

type TextAreaEntrySelectedEventArgs struct {
	TextArea      *TextArea
	Entry         interface{}
	PreviousEntry interface{}
}

type TextAreaEntrySelectedHandlerFunc func(args *TextAreaEntrySelectedEventArgs)

type TextAreaOptions struct {
}

var TextAreaOpts TextAreaOptions

func NewTextArea(opts ...TextAreaOpt) *TextArea { _ = "STUB: not implemented"; return nil }

func (t *TextArea) Validate() { _ = "STUB: not implemented"; return }

func (t *TextArea) populateComputedParams() { _ = "STUB: not implemented"; return }

// Set theme values

// Set definedParam values

// Set defaults

// Specify the Container options for the text area.
func (o TextAreaOptions) ContainerOpts(opts ...ContainerOpt) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Specify the images for the scroll container.
func (o TextAreaOptions) ScrollContainerImage(image *ScrollContainerImage) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Specify the padding for the scroll container.
func (o TextAreaOptions) ScrollContainerPadding(padding *Insets) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Specify the options for the scroll bars.
func (o TextAreaOptions) SliderParams(sliderParams *SliderParams) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Specify spacing between the text container and scrollbars.
func (o TextAreaOptions) ControlWidgetSpacing(s int) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Show the horizontal scrollbar.
func (o TextAreaOptions) ShowHorizontalScrollbar() TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Show the vertical scrollbar.
func (o TextAreaOptions) ShowVerticalScrollbar() TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Set how vertical scrolling should be handled.
func (o TextAreaOptions) VerticalScrollMode(scrollMode ScrollMode) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Set how horizontal scrolling should be handled.
func (o TextAreaOptions) HorizontalScrollMode(scrollMode ScrollMode) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Set the font face for this text area.
func (o TextAreaOptions) FontFace(f *text.Face) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Set the default color for the text area.
func (o TextAreaOptions) FontColor(color color.Color) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Set how far from the edges of the textarea the text should be set.
func (o TextAreaOptions) TextPadding(i Insets) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Set the positioning of the text within the text area
func (o TextAreaOptions) TextPosition(textPosition TextPositioning) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Set the initial Text for the text area.
func (o TextAreaOptions) Text(initialText string) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// This option tells the textarea object to process BBCodes.
//
// Currently the system supports the following BBCodes:
//   - color - [color=#FFFFFF] text [/color] - defines a color code for the enclosed text
//   - link - [link=id arg1:value1 ... argX:valueX] text [/link] - defines a clickable section of text,
//     that will trigger a callback.
func (o TextAreaOptions) ProcessBBCode(processBBCode bool) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Set whether or not the text area should automatically strip out BBCodes from being displayed.
func (o TextAreaOptions) StripBBCode(stripBBCode bool) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// This option sets the idle and hover color for text that is wrapped in a
// [link][/link] bbcode.
//
// Note: this is only used if ProcessBBCode is true.
func (o TextAreaOptions) LinkColor(linkColor *TextLinkColor) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Defines the handler to be called when a BBCode defined link is clicked.
//
// Note: this is only used if ProcessBBCode is true.
func (o TextAreaOptions) LinkClickedEvent(linkClickedFunc LinkHandlerFunc) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Defines the handler to be called when the cursor enters a BBCode defined link.
//
// Note: this is only used if ProcessBBCode is true.
func (o TextAreaOptions) LinkCursorEnteredEvent(linkCursorEnteredFunc LinkHandlerFunc) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

// Defines the handler to be called when the cursor enters a BBCode defined link.
//
// Note: this is only used if ProcessBBCode is true.
func (o TextAreaOptions) LinkCursorExitedEvent(linkCursorExitedFunc LinkHandlerFunc) TextAreaOpt {
	_ = "STUB: not implemented"
	return *new(TextAreaOpt)
}

func (l *TextArea) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (l *TextArea) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (l *TextArea) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (l *TextArea) RequestRelayout() { _ = "STUB: not implemented"; return }

func (l *TextArea) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (l *TextArea) GetFocusers() []Focuser { _ = "STUB: not implemented"; return nil }

func (l *TextArea) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (l *TextArea) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (l *TextArea) createWidget() { _ = "STUB: not implemented"; return }

func (l *TextArea) initWidget() { _ = "STUB: not implemented"; return }

func (l *TextArea) PrependText(value string) { _ = "STUB: not implemented"; return }

func (l *TextArea) AppendText(value string) { _ = "STUB: not implemented"; return }

func (l *TextArea) SetText(value string) { _ = "STUB: not implemented"; return }

func (l *TextArea) GetText() string { _ = "STUB: not implemented"; return "" }
