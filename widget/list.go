package widget

import (
	img "image"
	"image/color"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ListParams struct {
	EntryFace                   *text.Face
	EntryColor                  *ListEntryColor
	EntryTextPadding            *Insets
	EntryTextHorizontalPosition *TextPosition
	EntryTextVerticalPosition   *TextPosition
	ControlWidgetSpacing        *int
	MinSize                     *img.Point

	AllowReselect      *bool
	SelectFocus        *bool
	SelectPressed      *bool
	DisableDefaultKeys *bool

	Slider                 *SliderParams
	ScrollContainerImage   *ScrollContainerImage
	ScrollContainerPadding *Insets

	entryUnselectedColor     *ButtonImage
	entrySelectedColor       *ButtonImage
	entryUnselectedTextColor *ButtonTextColor
	entryTextColor           *ButtonTextColor
}

type List struct {
	definedParams  ListParams
	computedParams ListParams

	EntrySelectedEvent *event.Event

	containerOpts        []ContainerOpt
	hideHorizontalSlider bool
	hideVerticalSlider   bool

	entries        []any
	entryLabelFunc ListEntryLabelFunc
	entrySortFunc  ListEntrySortFunc

	init            *MultiOnce
	container       *Container
	listContent     *Container
	scrollContainer *ScrollContainer
	layout          *GridLayout
	vSlider         *Slider
	hSlider         *Slider
	buttons         []*Button
	selectedEntry   any
	validated       bool

	focused        bool
	tabOrder       int
	justMoved      bool
	focusIndex     int
	prevFocusIndex int

	focusMap map[FocusDirection]Focuser
}

type ListOpt func(l *List)

type ListEntryLabelFunc func(e any) string
type ListEntrySortFunc func(a, b any) int

type ListEntryColor struct {
	Unselected                 color.Color
	Selected                   color.Color
	DisabledUnselected         color.Color
	DisabledSelected           color.Color
	SelectingBackground        color.Color
	SelectedBackground         color.Color
	FocusedBackground          color.Color
	SelectingFocusedBackground color.Color
	SelectedFocusedBackground  color.Color
	DisabledSelectedBackground color.Color
}

type ListEntrySelectedEventArgs struct {
	List          *List
	Entry         any
	PreviousEntry any
}

type ListEntrySelectedHandlerFunc func(args *ListEntrySelectedEventArgs)

type ListOptions struct {
}

var ListOpts ListOptions

func NewList(opts ...ListOpt) *List { _ = "STUB: not implemented"; return nil }

func (l *List) Validate() { _ = "STUB: not implemented"; return }

func (t *List) populateComputedParams() { _ = "STUB: not implemented"; return }

// Set theme values

// Set definedParam values

// Set defaults

func (o ListOptions) ContainerOpts(opts ...ContainerOpt) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

// Specify the images for the scroll container.
func (o ListOptions) ScrollContainerImage(image *ScrollContainerImage) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

// Specify the padding for the scroll container.
func (o ListOptions) ScrollContainerPadding(padding *Insets) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

// Specify the options for the scroll bars.
func (o ListOptions) SliderParams(sliderParams *SliderParams) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (o ListOptions) ControlWidgetSpacing(s int) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (o ListOptions) HideHorizontalSlider() ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (o ListOptions) HideVerticalSlider() ListOpt { _ = "STUB: not implemented"; return *new(ListOpt) }

func (o ListOptions) Entries(e []any) ListOpt { _ = "STUB: not implemented"; return *new(ListOpt) }

func (o ListOptions) EntryLabelFunc(f ListEntryLabelFunc) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (o ListOptions) EntrySortFunc(f ListEntrySortFunc) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (o ListOptions) EntryFontFace(f *text.Face) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (o ListOptions) DisableDefaultKeys(val bool) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (o ListOptions) EntryColor(c *ListEntryColor) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (o ListOptions) EntryTextPadding(i *Insets) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

// EntryTextPosition sets the position of the text for entries.
// Defaults to both TextPositionCenter.
func (o ListOptions) EntryTextPosition(h TextPosition, v TextPosition) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (o ListOptions) EntrySelectedHandler(f ListEntrySelectedHandlerFunc) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (o ListOptions) AllowReselect() ListOpt { _ = "STUB: not implemented"; return *new(ListOpt) }

// SelectFocus automatically selects each focused entry.
func (o ListOptions) SelectFocus() ListOpt { _ = "STUB: not implemented"; return *new(ListOpt) }

// SelectPressed selects entries when pressing instead of releasing (the default).
func (o ListOptions) SelectPressed() ListOpt { _ = "STUB: not implemented"; return *new(ListOpt) }

func (o ListOptions) TabOrder(tabOrder int) ListOpt {
	_ = "STUB: not implemented"
	return *new(ListOpt)
}

func (l *List) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (l *List) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (l *List) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (l *List) RequestRelayout() { _ = "STUB: not implemented"; return }

func (l *List) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (l *List) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (l *List) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

/** Focuser Interface - Start **/

func (l *List) Focus(focused bool) { _ = "STUB: not implemented"; return }

func (l *List) IsFocused() bool { _ = "STUB: not implemented"; return false }

func (l *List) TabOrder() int { _ = "STUB: not implemented"; return 0 }

func (l *List) GetFocus(direction FocusDirection) Focuser {
	_ = "STUB: not implemented"
	return *new(Focuser)
}

func (l *List) AddFocus(direction FocusDirection, focus Focuser) { _ = "STUB: not implemented"; return }

/** Focuser Interface - End **/

func (l *List) handleInput() { _ = "STUB: not implemented"; return }

func (l *List) FocusNext() { _ = "STUB: not implemented"; return }

func (l *List) FocusPrevious() { _ = "STUB: not implemented"; return }

func (l *List) SelectFocused() { _ = "STUB: not implemented"; return }

func (l *List) resetFocusIndex() { _ = "STUB: not implemented"; return }

func (l *List) createWidget() { _ = "STUB: not implemented"; return }

func (l *List) initWidget() { _ = "STUB: not implemented"; return }

// Updates the entries in the list.
// Note: Duplicates will be removed.
func (l *List) SetEntries(newEntries []any) { _ = "STUB: not implemented"; return }

// UpdateEntry recreates the button for the passed in entry if it exists.
func (l *List) UpdateEntry(entry any) { _ = "STUB: not implemented"; return }

// Remove the passed in entry from the list if it exists.
func (l *List) RemoveEntry(entry any) { _ = "STUB: not implemented"; return }

// Add a new entry to the end of the list
// Note: Duplicates will not be added.
func (l *List) AddEntry(entry any) { _ = "STUB: not implemented"; return }

// Return the current entries in the list.
func (l *List) Entries() []any { _ = "STUB: not implemented"; return nil }

// Return the currently selected entry in the list.
func (l *List) SelectedEntry() any { _ = "STUB: not implemented"; return *new(any) }

// Set the Selected Entry to e if it is found.
func (l *List) SetSelectedEntry(entry any) { _ = "STUB: not implemented"; return }

func (l *List) setSelectedEntry(e any, user bool) { _ = "STUB: not implemented"; return }

func (l *List) setEntryButtonStyle(button *Button, selected bool) {
	_ = "STUB: not implemented"
	return
}

func (l *List) checkForDuplicates(entries []any, entry any) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *List) normalizeEntries(entries []any) []any { _ = "STUB: not implemented"; return nil }

func (l *List) insertListContentChild(index int, child PreferredSizeLocateableWidget) {
	_ = "STUB: not implemented"
	return
}

func (l *List) createEntry(entry any) *Button { _ = "STUB: not implemented"; return nil }

func (l *List) setScrollTop(t float64) { _ = "STUB: not implemented"; return }

func (l *List) setScrollLeft(left float64) { _ = "STUB: not implemented"; return }

func (l *List) scrollVisible(w HasWidget) { _ = "STUB: not implemented"; return }

func scrollClamp(targetScroll, currentScroll float64) float64 { _ = "STUB: not implemented"; return 0 }
