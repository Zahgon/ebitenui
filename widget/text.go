package widget

import (
	"image"
	"image/color"
	"regexp"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/utilities/datastructures"
	"github.com/frustra/bbcode"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var bbcodeRegex = regexp.MustCompile(`\[color=#[0-9a-fA-F]{6}]|\[\/color]|\[link]|\[\/link]|\[link=[^\]]*]`)

const COLOR_TAG = "color"
const LINK_TAG = "link"

type TextParams struct {
	Face      *text.Face
	Color     color.Color
	Padding   *Insets
	LinkColor *TextLinkColor
	Position  *TextPositioning
}

type Text struct {
	definedParams  TextParams
	computedParams TextParams
	Label          string
	MaxWidth       float64
	ProcessBBCode  bool
	StripBBCode    bool

	widgetOpts []WidgetOpt

	init         *MultiOnce
	widget       *Widget
	measurements textMeasurements
	colorList    *datastructures.Stack[color.Color]
	linkStack    *datastructures.Stack[linkData]
	currentLink  *bbCodeText
	previousLink *bbCodeText

	LinkClickedEvent       *event.Event
	LinkCursorEnteredEvent *event.Event
	LinkCursorExitedEvent  *event.Event
}

type textMeasurements struct {
	label         string
	face          *text.Face
	maxWidth      float64
	ProcessBBCode bool

	processedLines      [][]*bbCodeText
	processedLineWidths []float64
	lineHeight          float64
	ascent              float64
	boundingBoxWidth    float64
	boundingBoxHeight   float64
}

type bbCodeText struct {
	text      string
	color     color.Color
	linkValue *linkData
	hovered   bool
}

type linkData struct {
	id         string
	text       string
	args       map[string]string
	textBlocks []*bbCodeText
}

type TextLinkColor struct {
	Idle  color.Color
	Hover color.Color
}

type LinkEventArgs struct {
	Text    *Text
	Id      string
	Value   string
	Args    map[string]string
	OffsetX int
	OffsetY int
}

type LinkHandlerFunc func(args *LinkEventArgs)

type TextPosition int

const (
	TextPositionStart = TextPosition(iota)
	TextPositionCenter
	TextPositionEnd
)

type TextPositioning struct {
	VTextPosition TextPosition
	HTextPosition TextPosition
}

type TextOpt func(t *Text)

type TextOptions struct {
}

var TextOpts TextOptions

func NewText(opts ...TextOpt) *Text { _ = "STUB: not implemented"; return nil }

func (t *Text) Validate() { _ = "STUB: not implemented"; return }

func (t *Text) populateComputedParams() { _ = "STUB: not implemented"; return }

func (o TextOptions) WidgetOpts(opts ...WidgetOpt) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

// Text combines three options: TextLabel, TextFace and TextColor.
// It can be used for the inline configurations of Text object while
// separate functions are useful for a multi-step configuration.
func (o TextOptions) Text(label string, face *text.Face, color color.Color) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

func (o TextOptions) TextLabel(label string) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

func (o TextOptions) TextFace(face *text.Face) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

func (o TextOptions) TextColor(color color.Color) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

func (o TextOptions) Padding(padding *Insets) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

func (o TextOptions) Position(h TextPosition, v TextPosition) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

// This option tells the text object to process BBCodes.
//
// Currently the system supports the following BBCodes:
//   - color - [color=#FFFFFF] text [/color] - defines a color code for the enclosed text
//   - link - [link=id arg1:value1 ... argX:valueX] text [/link] - defines a clickable section of text,
//     that will trigger a callback.
func (o TextOptions) ProcessBBCode(processBBCode bool) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

// Set whether or not the text area should automatically strip out BBCodes from being displayed.
func (o TextOptions) StripBBCode(stripBBCode bool) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

// This option sets the idle and hover color for text that is wrapped in a
// [link][/link] bbcode.
//
// Note: this is only used if ProcessBBCode is true.
func (o TextOptions) LinkColor(linkColor *TextLinkColor) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

// MaxWidth sets the max width the text will allow before wrapping to the next line.
func (o TextOptions) MaxWidth(maxWidth float64) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

// Defines the handler to be called when a BBCode defined link is clicked.
//
// Note: this is only used if ProcessBBCode is true.
func (o TextOptions) LinkClickedHandler(f LinkHandlerFunc) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

// Defines the handler to be called when the cursor enters a BBCode defined link.
//
// Note: this is only used if ProcessBBCode is true.
func (o TextOptions) LinkCursorEnteredHandler(f LinkHandlerFunc) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

// Defines the handler to be called when the cursor enters a BBCode defined link.
//
// Note: this is only used if ProcessBBCode is true.
func (o TextOptions) LinkCursorExitedHandler(f LinkHandlerFunc) TextOpt {
	_ = "STUB: not implemented"
	return *new(TextOpt)
}

func (t *Text) SetColor(color color.Color) { _ = "STUB: not implemented"; return }

func (t *Text) SetFace(face *text.Face) { _ = "STUB: not implemented"; return }

func (t *Text) SetPadding(padding *Insets) { _ = "STUB: not implemented"; return }

func (t *Text) SetPosition(position *TextPositioning) { _ = "STUB: not implemented"; return }

func (t *Text) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (t *Text) SetLocation(rect image.Rectangle) { _ = "STUB: not implemented"; return }

func (t *Text) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (t *Text) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (t *Text) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (t *Text) handleLinkEvents() { _ = "STUB: not implemented"; return }

// Reset Hovered

// Process Hovered

func (t *Text) draw(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

// Draw text

func (t *Text) handleBBCodeColor(line string) ([]*bbCodeText, color.Color, *linkData) {
	_ = "STUB: not implemented"
	return nil, *new(color.Color), nil
}

func (t *Text) processTree(node *bbcode.BBCodeNode, newColor color.Color, linkVal *linkData) ([]*bbCodeText, color.Color, *linkData) {
	_ = "STUB: not implemented"
	return nil, *new(color.Color), nil
}

// Handle changing color back.

func (t *Text) measure() { _ = "STUB: not implemented"; return }

// Don't add the space to the last chunk.

// If the new word doesn't push this past the max width continue adding to the current line

// If the new word would push this past the max width save off the current line and start a new one

// Save the final line

func (t *Text) getCursorOffset() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

func (t *Text) createWidget() { _ = "STUB: not implemented"; return }
