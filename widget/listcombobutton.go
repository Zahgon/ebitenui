package widget

import (
	"image"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ListComboButtonParams struct {
	List               *ListParams
	Button             *ButtonParams
	DisableDefaultKeys *bool
	MaxContentHeight   *int
}

type ListComboButton struct {
	definedParams  ListComboButtonParams
	computedParams ListComboButtonParams

	EntrySelectedEvent *event.Event

	init       *MultiOnce
	widget     *Widget
	widgetOpts []WidgetOpt
	entries    []any

	button          *SelectComboButton
	list            *List
	buttonLabelFunc SelectComboButtonEntryLabelFunc
	listLabelFunc   ListEntryLabelFunc

	selectedEntry any

	tabOrder int
	focusMap map[FocusDirection]Focuser
}

type ListComboButtonOpt func(l *ListComboButton)

type ListComboButtonEntrySelectedEventArgs struct {
	Button        *ListComboButton
	Entry         interface{}
	PreviousEntry interface{}
}

type ListComboButtonEntrySelectedHandlerFunc func(args *ListComboButtonEntrySelectedEventArgs)

type ListComboButtonOptions struct {
}

var ListComboButtonOpts ListComboButtonOptions

func NewListComboButton(opts ...ListComboButtonOpt) *ListComboButton {
	_ = "STUB: not implemented"
	return nil
}

func (l *ListComboButton) Validate() { _ = "STUB: not implemented"; return }

func (t *ListComboButton) populateComputedParams() { _ = "STUB: not implemented"; return }

// Set theme values

// Set definedParam values

// Set defaults

func (o ListComboButtonOptions) WidgetOpts(opts ...WidgetOpt) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

func (o ListComboButtonOptions) ButtonParams(buttonParams *ButtonParams) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

func (o ListComboButtonOptions) ListParams(listParams *ListParams) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

func (o ListComboButtonOptions) Text(face *text.Face, image *GraphicImage, color *ButtonTextColor) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

func (o ListComboButtonOptions) Entries(e []any) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

func (o ListComboButtonOptions) InitialEntry(e any) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

func (o ListComboButtonOptions) EntryLabelFunc(button SelectComboButtonEntryLabelFunc, list ListEntryLabelFunc) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

func (o ListComboButtonOptions) EntrySelectedHandler(f ListComboButtonEntrySelectedHandlerFunc) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

func (o ListComboButtonOptions) MaxContentHeight(h int) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

func (o ListComboButtonOptions) TabOrder(tabOrder int) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

func (o ListComboButtonOptions) DisableDefaultKeys(val bool) ListComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ListComboButtonOpt)
}

/** Focuser Interface - Start **/

func (l *ListComboButton) Focus(focused bool) { _ = "STUB: not implemented"; return }

func (l *ListComboButton) IsFocused() bool { _ = "STUB: not implemented"; return false }

func (l *ListComboButton) TabOrder() int { _ = "STUB: not implemented"; return 0 }

func (l *ListComboButton) GetFocus(direction FocusDirection) Focuser {
	_ = "STUB: not implemented"
	return *new(Focuser)
}

func (l *ListComboButton) AddFocus(direction FocusDirection, focus Focuser) {
	_ = "STUB: not implemented"
	return
}

/** Focuser Interface - End **/

func (l *ListComboButton) FocusNext() { _ = "STUB: not implemented"; return }

func (l *ListComboButton) FocusPrevious() { _ = "STUB: not implemented"; return }

func (l *ListComboButton) SelectFocused() { _ = "STUB: not implemented"; return }

func (l *ListComboButton) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (l *ListComboButton) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (l *ListComboButton) SetLocation(rect image.Rectangle) { _ = "STUB: not implemented"; return }

func (l *ListComboButton) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (l *ListComboButton) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (l *ListComboButton) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (l *ListComboButton) createWidget() { _ = "STUB: not implemented"; return }

func (l *ListComboButton) initWidget() { _ = "STUB: not implemented"; return }

func (l *ListComboButton) SetSelectedEntry(e interface{}) { _ = "STUB: not implemented"; return }

func (l *ListComboButton) SelectedEntry() interface{} { _ = "STUB: not implemented"; return nil }

func (l *ListComboButton) SetContentVisible(v bool) { _ = "STUB: not implemented"; return }

func (l *ListComboButton) ContentVisible() bool { _ = "STUB: not implemented"; return false }

func (l *ListComboButton) Label() string { _ = "STUB: not implemented"; return "" }
