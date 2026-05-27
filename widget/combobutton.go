package widget

import (
	"image"

	"github.com/ebitenui/ebitenui/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type ComboButton struct {
	ContentVisible bool

	buttonOpts       []ButtonOpt
	maxContentHeight int

	init    *MultiOnce
	button  *Button
	content HasWidget
}

type ComboButtonOpt func(c *ComboButton)

type ComboButtonOptions struct {
}

var ComboButtonOpts ComboButtonOptions

func NewComboButton(opts ...ComboButtonOpt) *ComboButton { _ = "STUB: not implemented"; return nil }

func (c *ComboButton) Validate() { _ = "STUB: not implemented"; return }

func (o ComboButtonOptions) ButtonOpts(opts ...ButtonOpt) ComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ComboButtonOpt)
}

func (o ComboButtonOptions) Content(c HasWidget) ComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ComboButtonOpt)
}

func (o ComboButtonOptions) MaxContentHeight(h int) ComboButtonOpt {
	_ = "STUB: not implemented"
	return *new(ComboButtonOpt)
}

func (c *ComboButton) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (c *ComboButton) SetLocation(rect image.Rectangle) { _ = "STUB: not implemented"; return }

func (c *ComboButton) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (c *ComboButton) SetLabel(l string) { _ = "STUB: not implemented"; return }

func (c *ComboButton) Label() string { _ = "STUB: not implemented"; return "" }

func (c *ComboButton) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (c *ComboButton) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (c *ComboButton) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (c *ComboButton) handleClick() { _ = "STUB: not implemented"; return }

func (c *ComboButton) relayoutContent() { _ = "STUB: not implemented"; return }

func (c *ComboButton) createWidget() { _ = "STUB: not implemented"; return }
