package widget

import (
	img "image"
	"image/color"
	"sync/atomic"
	"time"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/utilities/mobile"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextInputParams struct {
	Image                *TextInputImage
	Color                *TextInputColor
	Padding              *Insets
	Face                 *text.Face
	RepeatDelay          *time.Duration
	RepeatInterval       *time.Duration
	Secure               *bool
	ClearOnSubmit        *bool
	IgnoreEmptySubmit    *bool
	AllowDuplicateSubmit *bool
	SubmitOnEnter        *bool
	ScrollSensitivity    *int
	CaretWidth           *int
}

type TextInput struct {
	definedParams  TextInputParams
	computedParams TextInputParams

	ChangedEvent *event.Event
	SubmitEvent  *event.Event

	inputText       string
	validationFunc  TextInputValidationFunc
	placeholderText string
	mobileInputMode mobile.InputMode

	widgetOpts            []WidgetOpt
	init                  *MultiOnce
	commandToFunc         map[textInputControlCommand]textInputCommandFunc
	widget                *Widget
	caret                 *Caret
	text                  *Text
	renderBuf             *image.MaskedRenderBuffer
	mask                  *image.NineSlice
	cursorPosition        int
	state                 textInputState
	scrollOffset          int
	lastInputText         string
	previousSubmittedText *string
	dragStartIndex        int

	tabOrder int
	focused  bool
	focusMap map[FocusDirection]Focuser
}

type TextInputOpt func(t *TextInput)

type TextInputOptions struct {
}

type TextInputChangedEventArgs struct {
	TextInput *TextInput
	InputText string
}

type TextInputChangedHandlerFunc func(args *TextInputChangedEventArgs)

type TextInputImage struct {
	Idle     *image.NineSlice
	Disabled *image.NineSlice
	// Highlight defaults to image.NewNineSliceColor(color.NRGBA{6, 67, 161, 100}).
	Highlight *image.NineSlice
}

type TextInputColor struct {
	Idle          color.Color
	Disabled      color.Color
	Caret         color.Color
	DisabledCaret color.Color
}

type TextInputValidationFunc func(newInputText string) (bool, *string)

type textInputState func() (textInputState, bool)

type textInputControlCommand int

type textInputCommandFunc func()

var TextInputOpts TextInputOptions

const (
	textInputGoLeft = textInputControlCommand(iota + 1)
	textInputGoRight
	textInputGoStart
	textInputGoEnd
	textInputBackspace
	textInputDelete
	textInputEnter
	textInputEscape
)

var textInputKeyToCommand = map[ebiten.Key]textInputControlCommand{
	ebiten.KeyLeft:        textInputGoLeft,
	ebiten.KeyRight:       textInputGoRight,
	ebiten.KeyHome:        textInputGoStart,
	ebiten.KeyEnd:         textInputGoEnd,
	ebiten.KeyBackspace:   textInputBackspace,
	ebiten.KeyDelete:      textInputDelete,
	ebiten.KeyEnter:       textInputEnter,
	ebiten.KeyNumpadEnter: textInputEnter,
	ebiten.KeyEscape:      textInputEscape,
}

func NewTextInput(opts ...TextInputOpt) *TextInput { _ = "STUB: not implemented"; return nil }

func (t *TextInput) Validate() { _ = "STUB: not implemented"; return }

func (t *TextInput) populateComputedParams() { _ = "STUB: not implemented"; return }

// Set theme values

// Set Defined values

// Set Default values

func (o TextInputOptions) WidgetOpts(opts ...WidgetOpt) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) ChangedHandler(f TextInputChangedHandlerFunc) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) SubmitHandler(f TextInputChangedHandlerFunc) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) ClearOnSubmit(clearOnSubmit bool) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) IgnoreEmptySubmit(ignoreEmptySubmit bool) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) AllowDuplicateSubmit(allowDuplicateSubmit bool) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) Image(i *TextInputImage) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) Color(c *TextInputColor) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) Padding(i *Insets) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) Face(f *text.Face) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) RepeatInterval(i time.Duration) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) Validation(f TextInputValidationFunc) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) Placeholder(s string) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) Secure(b bool) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) CaretWidth(caretWidth int) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

func (o TextInputOptions) TabOrder(to int) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

// Sets if the input will submit when pressing enter or not.
func (o TextInputOptions) SubmitOnEnter(submitOnEnter bool) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

// Sets the keyboard type to use when viewed on a mobile browser.
//
// https://css-tricks.com/everything-you-ever-wanted-to-know-about-inputmode
func (o TextInputOptions) MobileInputMode(mobileInputMode mobile.InputMode) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

// Sets how many pixels from the edge the cursor must be dragged prior to it scrolling in that direction.
//
// Default: 15.
func (o TextInputOptions) ScrollSensitivity(scrollSensitivity int) TextInputOpt {
	_ = "STUB: not implemented"
	return *new(TextInputOpt)
}

/*********** End of Configuration *****************/

func (t *TextInput) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (t *TextInput) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (t *TextInput) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (t *TextInput) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (t *TextInput) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (t *TextInput) idleState(newKeyOrCommand bool) textInputState {
	_ = "STUB: not implemented"
	return *new(textInputState)
}

func textInputCheckForCommand(t *TextInput, newKeyOrCommand bool) textInputState {
	_ = "STUB: not implemented"
	return *new(textInputState)
}

func (t *TextInput) charsInputState(c string) textInputState {
	_ = "STUB: not implemented"
	return *new(textInputState)
}

func (t *TextInput) commandState(cmd textInputControlCommand, key ebiten.Key, delay time.Duration, timer *time.Timer, expired *atomic.Value) textInputState {
	_ = "STUB: not implemented"
	return *new(textInputState)
}

func (t *TextInput) Insert(c string) { _ = "STUB: not implemented"; return }

func (t *TextInput) CursorMoveLeft() { _ = "STUB: not implemented"; return }

func (t *TextInput) CursorMoveRight() { _ = "STUB: not implemented"; return }

func (t *TextInput) CursorMoveStart() { _ = "STUB: not implemented"; return }

func (t *TextInput) CursorMoveEnd() { _ = "STUB: not implemented"; return }

func (t *TextInput) Backspace() { _ = "STUB: not implemented"; return }

func (t *TextInput) Delete() { _ = "STUB: not implemented"; return }

func (t *TextInput) submitWithEnter() { _ = "STUB: not implemented"; return }

func (t *TextInput) Submit() { _ = "STUB: not implemented"; return }

func (t *TextInput) SelectedText() string { _ = "STUB: not implemented"; return "" }

func (t *TextInput) DeselectText() { _ = "STUB: not implemented"; return }

func (t *TextInput) SelectAll() { _ = "STUB: not implemented"; return }

func (t *TextInput) DeleteSelectedText() { _ = "STUB: not implemented"; return }

func insertChars(r []rune, c []rune, pos int) []rune { _ = "STUB: not implemented"; return nil }

func removeChar(r []rune, pos int) []rune { _ = "STUB: not implemented"; return nil }

func (t *TextInput) renderImage(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (t *TextInput) renderTextAndCaret(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (t *TextInput) drawTextAndCaret(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

// Change the Dx and the tr.Min.X based on selection

func (t *TextInput) GetText() string { _ = "STUB: not implemented"; return "" }

func (t *TextInput) SetText(text string) { _ = "STUB: not implemented"; return }

func (t *TextInput) setJSText(text string) string { _ = "STUB: not implemented"; return "" }

func (t *TextInput) setText(text string, isJS bool) { _ = "STUB: not implemented"; return }

/** Focuser Interface - Start **/

func (t *TextInput) Focus(focused bool) { _ = "STUB: not implemented"; return }

func (t *TextInput) IsFocused() bool { _ = "STUB: not implemented"; return false }

func (t *TextInput) TabOrder() int { _ = "STUB: not implemented"; return 0 }

func (t *TextInput) GetFocus(direction FocusDirection) Focuser {
	_ = "STUB: not implemented"
	return *new(Focuser)
}

func (t *TextInput) AddFocus(direction FocusDirection, focus Focuser) {
	_ = "STUB: not implemented"
	return
}

/** Focuser Interface - End **/

func (t *TextInput) createWidget() { _ = "STUB: not implemented"; return }

func (t *TextInput) initWidget() { _ = "STUB: not implemented"; return }

func fontAdvance(s string, f *text.Face) int { _ = "STUB: not implemented"; return 0 }

// fontStringIndex returns an index into r that corresponds closest to pixel position x
// when string(r) is drawn using f. Pixel position x==0 corresponds to r[0].
func fontStringIndex(r []rune, f *text.Face, x int) int { _ = "STUB: not implemented"; return 0 }

// x is right of advance

// x is left of advance

// x matches advance exactly
