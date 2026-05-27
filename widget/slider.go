package widget

import (
	img "image"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type SliderParams struct {
	Orientation     *Direction
	TrackImage      *SliderTrackImage
	TrackPadding    *Insets
	MinHandleSize   *int
	FixedHandleSize *int
	TrackOffset     *int
	HandleImage     *ButtonImage
	PageSizeFunc    SliderPageSizeFunc
}

type Slider struct {
	definedParams  SliderParams
	computedParams SliderParams

	Min                int
	Max                int
	Current            int
	DrawTrackDisabled  bool
	disableDefaultKeys bool

	widgetOpts []WidgetOpt

	ChangedEvent *event.Event

	init   *MultiOnce
	widget *Widget
	handle *Button

	lastCurrent                  int
	hovering                     bool
	dragging                     bool
	handlePressedCursorX         int
	handlePressedCursorY         int
	handlePressedOffsetX         int
	handlePressedOffsetY         int
	handlePressedInternalCurrent float64

	tabOrder  int
	justMoved bool
	focusMap  map[FocusDirection]Focuser
}

type SliderTrackImage struct {
	Idle     *image.NineSlice
	Hover    *image.NineSlice
	Disabled *image.NineSlice
}

type SliderOpt func(s *Slider)

type SliderPageSizeFunc func() int

type SliderChangedEventArgs struct {
	Slider   *Slider
	Current  int
	Dragging bool
}

type SliderChangedHandlerFunc func(args *SliderChangedEventArgs)

type SliderOptions struct {
}

var SliderOpts SliderOptions

func NewSlider(opts ...SliderOpt) *Slider { _ = "STUB: not implemented"; return nil }

func (s *Slider) Validate() { _ = "STUB: not implemented"; return }

func (s *Slider) populateComputedParams() { _ = "STUB: not implemented"; return }

// Set Theme

// Set defined params

// Set defaults

func (o SliderOptions) WidgetOpts(opts ...WidgetOpt) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (so SliderOptions) Orientation(o Direction) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

// Deprecated: Use Orientation(o *Direction) instead.
func (o SliderOptions) Direction(d Direction) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

// This sets the track images (not required) and the handle images (required).
func (o SliderOptions) Images(track *SliderTrackImage, handle *ButtonImage) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

// This sets the track images (not required).
func (o SliderOptions) TrackImage(track *SliderTrackImage) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

// This sets the handle images (required).
func (o SliderOptions) HandleImage(handle *ButtonImage) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (o SliderOptions) TrackPadding(i *Insets) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (o SliderOptions) TrackOffset(i int) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (o SliderOptions) MinHandleSize(s int) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (o SliderOptions) FixedHandleSize(s int) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (o SliderOptions) MinMax(minValue int, maxValue int) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (o SliderOptions) InitialCurrent(value int) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (o SliderOptions) PageSizeFunc(f SliderPageSizeFunc) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (o SliderOptions) ChangedHandler(f SliderChangedHandlerFunc) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (o SliderOptions) TabOrder(tabOrder int) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

func (o SliderOptions) DisableDefaultKeys(val bool) SliderOpt {
	_ = "STUB: not implemented"
	return *new(SliderOpt)
}

/** Focuser Interface - Start **/

func (s *Slider) Focus(focused bool) { _ = "STUB: not implemented"; return }

func (s *Slider) IsFocused() bool { _ = "STUB: not implemented"; return false }

func (s *Slider) TabOrder() int { _ = "STUB: not implemented"; return 0 }

func (s *Slider) GetFocus(direction FocusDirection) Focuser {
	_ = "STUB: not implemented"
	return *new(Focuser)
}

func (s *Slider) AddFocus(direction FocusDirection, focus Focuser) {
	_ = "STUB: not implemented"
	return
}

/** Focuser Interface - End **/

func (s *Slider) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (s *Slider) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *Slider) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (s *Slider) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (s *Slider) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (s *Slider) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (s *Slider) drawTrack(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (s *Slider) handleOrientation() { _ = "STUB: not implemented"; return }

func (s *Slider) fireEvents() { _ = "STUB: not implemented"; return }

func (s *Slider) updateHandleSize(handleLength float64) { _ = "STUB: not implemented"; return }

func (s *Slider) updateHandleLocation(handleLength float64, trackLength float64) {
	_ = "STUB: not implemented"
	return
}

func (s *Slider) handleLengthAndTrackLength() (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (s *Slider) currentToInternal(c int) float64 { _ = "STUB: not implemented"; return 0 }

func (s *Slider) internalToCurrent(i float64) int { _ = "STUB: not implemented"; return 0 }

func (s *Slider) clampCurrentMinMax() { _ = "STUB: not implemented"; return }

func (s *Slider) createWidget() { _ = "STUB: not implemented"; return }

// TODO: keeping the mouse button pressed should move the handle repeatedly (in PageSize steps) until it stops under the cursor

func (s *Slider) setChildComputedParams() { _ = "STUB: not implemented"; return }
