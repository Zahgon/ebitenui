package widget

import (
	img "image"
	"image/color"
	"sync/atomic"
	"time"

	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
)

type Caret struct {
	Width         int
	Height        int
	Color         color.Color
	blinkInterval time.Duration

	init    *MultiOnce
	widget  *Widget
	image   *image.NineSlice
	state   caretBlinkState
	visible bool
}

type CaretOpt func(c *Caret)

type CaretOptions struct {
}

var CaretOpts CaretOptions

type caretBlinkState func() caretBlinkState

func NewCaret(opts ...CaretOpt) *Caret { _ = "STUB: not implemented"; return nil }

func (c *Caret) Validate() { _ = "STUB: not implemented"; return }

func (o CaretOptions) Color(c color.Color) CaretOpt {
	_ = "STUB: not implemented"
	return *new(CaretOpt)
}

func (o CaretOptions) Size(height int, width int) CaretOpt {
	_ = "STUB: not implemented"
	return *new(CaretOpt)
}

func (c *Caret) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (c *Caret) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (c *Caret) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (c *Caret) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (c *Caret) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (c *Caret) ResetBlinking() { _ = "STUB: not implemented"; return }

func (c *Caret) resetBlinking() { _ = "STUB: not implemented"; return }

func (c *Caret) blinkState(visible bool, timer *time.Timer, expired *atomic.Value) caretBlinkState {
	_ = "STUB: not implemented"
	return *new(caretBlinkState)
}

func (c *Caret) createWidget() { _ = "STUB: not implemented"; return }
