package widget

import (
	"image"

	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type SelectComboButton struct {
	EntrySelectedEvent *event.Event

	buttonOpts     []ComboButtonOpt
	entryLabelFunc SelectComboButtonEntryLabelFunc

	init          *MultiOnce
	button        *ComboButton
	selectedEntry interface{}
}

type SelectComboButtonOpt func(s *SelectComboButton)

type SelectComboButtonEntryLabelFunc func(e interface{}) string

type SelectComboButtonEntrySelectedEventArgs struct {
	Button        *SelectComboButton
	Entry         interface{}
	PreviousEntry interface{}
}

type SelectComboButtonEntrySelectedHandlerFunc func(args *SelectComboButtonEntrySelectedEventArgs)

type SelectComboButtonOptions struct {
}

var SelectComboButtonOpts SelectComboButtonOptions

func NewSelectComboButton(opts ...SelectComboButtonOpt) *SelectComboButton {
	_ = "STUB: not implemented"
	return nil
}

func (s *SelectComboButton) Validate() { _ = "STUB: not implemented"; return }

func (o SelectComboButtonOptions) ComboButtonOpts(opts ...ComboButtonOpt) SelectComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(SelectComboButtonOpt)
}

func (o SelectComboButtonOptions) EntrySelectedHandler(f SelectComboButtonEntrySelectedHandlerFunc) SelectComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(SelectComboButtonOpt)
}

func (o SelectComboButtonOptions) EntryLabelFunc(f SelectComboButtonEntryLabelFunc) SelectComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(SelectComboButtonOpt)
}

func (s *SelectComboButton) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (s *SelectComboButton) SetLocation(rect image.Rectangle) { _ = "STUB: not implemented"; return }

func (s *SelectComboButton) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (s *SelectComboButton) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *SelectComboButton) SetLabel(l string) { _ = "STUB: not implemented"; return }

func (s *SelectComboButton) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (s *SelectComboButton) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (s *SelectComboButton) createWidget() { _ = "STUB: not implemented"; return }

func (s *SelectComboButton) SetSelectedEntry(e interface{}) { _ = "STUB: not implemented"; return }

func (s *SelectComboButton) SelectedEntry() interface{} { _ = "STUB: not implemented"; return nil }

func (s *SelectComboButton) SetContentVisible(v bool) { _ = "STUB: not implemented"; return }

func (s *SelectComboButton) ContentVisible() bool { _ = "STUB: not implemented"; return false }

func (s *SelectComboButton) Label() string { _ = "STUB: not implemented"; return "" }
