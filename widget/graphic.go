package widget

import (
	"image/color"
	"image/gif"
	"time"

	"github.com/ebitenui/ebitenui/image"

	img "image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Graphic struct {
	Image          *ebiten.Image
	ImageNineSlice *image.NineSlice
	images         *GraphicImage
	gif            *gif.GIF

	gifImages         []*ebiten.Image
	gifCurrentImage   int
	gifCurrentImageAt time.Time

	widgetOpts []WidgetOpt

	init   *MultiOnce
	widget *Widget
}

type GraphicImage struct {
	Idle     *ebiten.Image
	Disabled *ebiten.Image

	// Graphic does not have a handler for this 2 states
	// and they are used through the Button
	Pressed *ebiten.Image
	Hover   *ebiten.Image
}

type GraphicOpt func(g *Graphic)

type GraphicOptions struct {
}

var GraphicOpts GraphicOptions

func NewGraphic(opts ...GraphicOpt) *Graphic { _ = "STUB: not implemented"; return nil }

func (g *Graphic) validate() { _ = "STUB: not implemented"; return }

func (o GraphicOptions) WidgetOpts(opts ...WidgetOpt) GraphicOpt {
	_ = "STUB: not implemented"
	return *new(GraphicOpt)
}

func (o GraphicOptions) Image(i *ebiten.Image) GraphicOpt {
	_ = "STUB: not implemented"
	return *new(GraphicOpt)
}

func (o GraphicOptions) Images(i *GraphicImage) GraphicOpt {
	_ = "STUB: not implemented"
	return *new(GraphicOpt)
}

func (o GraphicOptions) GIF(gif *gif.GIF) GraphicOpt {
	_ = "STUB: not implemented"
	return *new(GraphicOpt)
}

func restoreFrame(current *img.Paletted, prev img.Image, rect img.Rectangle) img.Image {
	_ = "STUB: not implemented"
	return *new(img.Image)
}

func isOpaque(c color.Color) bool { _ = "STUB: not implemented"; return false }

func isInRect(x, y int, r img.Rectangle) bool { _ = "STUB: not implemented"; return false }

func (o GraphicOptions) ImageNineSlice(i *image.NineSlice) GraphicOpt {
	_ = "STUB: not implemented"
	return *new(GraphicOpt)
}

func (g *Graphic) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (g *Graphic) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (g *Graphic) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (g *Graphic) Validate() { _ = "STUB: not implemented"; return }

func (g *Graphic) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (g *Graphic) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (g *Graphic) draw(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (g *Graphic) createWidget() { _ = "STUB: not implemented"; return }
