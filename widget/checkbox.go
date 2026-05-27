package widget

import (
	img "image"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type LabelOrder int

const (
	CHECKBOX_FIRST LabelOrder = iota
	LABEL_FIRST
)

type CheckboxParams struct {
	Image *CheckboxImage
	Label *LabelParams
}

type Checkbox struct {
	definedParams  CheckboxParams
	computedParams CheckboxParams

	init              *MultiOnce
	widget            *Widget
	widgetOpts        []WidgetOpt
	triState          bool
	StateChangedEvent *event.Event

	state    WidgetState
	hovering bool

	labelString string
	label       *Label
	spacing     int
	order       LabelOrder

	// Allows the user to disable space bar and enter automatically triggering a focused checkbox.
	DisableDefaultKeys bool

	tabOrder int
	focused  bool
	focusMap map[FocusDirection]Focuser
}

type CheckboxOpt func(c *Checkbox)

type CheckboxImage struct {
	Unchecked         *image.NineSlice
	UncheckedHovered  *image.NineSlice
	UncheckedDisabled *image.NineSlice
	Checked           *image.NineSlice
	CheckedHovered    *image.NineSlice
	CheckedDisabled   *image.NineSlice
	Greyed            *image.NineSlice
	GreyedHovered     *image.NineSlice
	GreyedDisabled    *image.NineSlice
}

type CheckboxChangedEventArgs struct {
	Active *Checkbox
	State  WidgetState
}

type CheckboxChangedHandlerFunc func(args *CheckboxChangedEventArgs)

type CheckboxOptions struct {
}

var CheckboxOpts CheckboxOptions

func NewCheckbox(opts ...CheckboxOpt) *Checkbox { _ = "STUB: not implemented"; return nil }

func (c *Checkbox) Validate() { _ = "STUB: not implemented"; return }

func (c *Checkbox) populateComputedParams() { _ = "STUB: not implemented"; return }

func (o CheckboxOptions) WidgetOpts(opts ...WidgetOpt) CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

// This option allows you to specify a label to be shown before or after the checkbox.
func (o CheckboxOptions) Text(labelString string, face *text.Face, color *LabelColor) CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

func (o CheckboxOptions) TextLabel(label string) CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

func (o CheckboxOptions) TextFace(face *text.Face) CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

func (o CheckboxOptions) TextColor(color *LabelColor) CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

// This option defines how far the checkbox and label should be spaced horizontally if there is a label.
func (o CheckboxOptions) Spacing(s int) CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

// This option indicates that the label should be before the checkbox.
func (o CheckboxOptions) LabelFirst() CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

// This option defines the images to show for the checkbox.
// i.Checked and i.Unchecked are required.
func (o CheckboxOptions) Image(i *CheckboxImage) CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

// This option indicates this checkbox should have 3 states.
// If this option is specified a Greyed image is required.
func (o CheckboxOptions) TriState() CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

func (o CheckboxOptions) TabOrder(tabOrder int) CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

// This option allows you to specify a callback to be called when the checkbox state is changed.
func (o CheckboxOptions) StateChangedHandler(f CheckboxChangedHandlerFunc) CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

// This option sets the initial state for the checkbox.
func (o CheckboxOptions) InitialState(state WidgetState) CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

// This option will disable enter and space from submitting a focused checkbox.
func (o CheckboxOptions) DisableDefaultKeys() CheckboxOpt {
	_ = "STUB: not implemented"
	return *new(CheckboxOpt)
}

// This function will return the internal Text object if this checkbox has a label, otherwise nil.
func (tw *Checkbox) Text() *Text { _ = "STUB: not implemented"; return nil }

// This function will return the current state the checkbox is in.
func (tw *Checkbox) State() WidgetState {
	_ = "STUB: not implemented"

	// This function will allow you to update the checkbox's current state.
	return *new(WidgetState)
}

func (tw *Checkbox) SetState(state WidgetState) { _ = "STUB: not implemented"; return }

// This method is required for this to be part of a radio button group.
func (tw *Checkbox) getStateChangedEvent() *event.Event { _ = "STUB: not implemented"; return nil }

func (c *Checkbox) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (c *Checkbox) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (c *Checkbox) checkboxPreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (c *Checkbox) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (c *Checkbox) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (c *Checkbox) RequestRelayout() { _ = "STUB: not implemented"; return }

func (c *Checkbox) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

// If the label is larger than the checkbox, center the checkbox

func (c *Checkbox) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (c *Checkbox) handleDefaultInput() { _ = "STUB: not implemented"; return }

/** Focuser Interface - Start **/

func (c *Checkbox) Focus(focused bool) { _ = "STUB: not implemented"; return }

func (c *Checkbox) IsFocused() bool { _ = "STUB: not implemented"; return false }

func (c *Checkbox) TabOrder() int { _ = "STUB: not implemented"; return 0 }

func (c *Checkbox) GetFocus(direction FocusDirection) Focuser {
	_ = "STUB: not implemented"
	return *new(Focuser)
}

func (c *Checkbox) AddFocus(direction FocusDirection, focus Focuser) {
	_ = "STUB: not implemented"
	return
}

/** Focuser Interface - End **/

func (c *Checkbox) Click() { _ = "STUB: not implemented"; return }

func (c *Checkbox) createWidget() { _ = "STUB: not implemented"; return }

func (s WidgetState) Advance(triState bool) WidgetState {
	_ = "STUB: not implemented"
	return *new(WidgetState)
}

func (c *Checkbox) currentImage() *image.NineSlice {
	_ = "STUB: not implemented"
	// Handle disabled
	return nil
}

// Handle hovered

// Fallback to default images

func (c *Checkbox) setChildComputedParams() { _ = "STUB: not implemented"; return }
