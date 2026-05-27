package image

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

// A NineSlice is an image that can be drawn with any width and height. It is basically a 3x3 grid of image tiles:
// The corner tiles are drawn as-is, while the center columns and rows of tiles will be stretched to fit the desired
// width and height.
type NineSlice struct {
	image       *ebiten.Image
	widths      [3]int
	heights     [3]int
	transparent bool

	init  sync.Once
	tiles [9]*ebiten.Image
}

type borders struct {
	borderTop    int
	borderRight  int
	borderBottom int
	borderLeft   int
	borderColor  color.Color
}

// This function returns a new borders struct instance for use with the NewAdvancedNineSlice* functions.
func NewBorder(borderTop int, borderRight int, borderBottom int, borderLeft int, borderColor color.Color) borders {
	_ = "STUB: not implemented"
	return *new(borders)
}

// A DrawImageOptionsFunc is responsible for setting DrawImageOptions when drawing an image.
// This is usually used to translate the image.
type DrawImageOptionsFunc func(opts *ebiten.DrawImageOptions)

var colorImages map[color.Color]*ebiten.Image = map[color.Color]*ebiten.Image{}

var colorNineSlices map[color.Color]*NineSlice = map[color.Color]*NineSlice{}

// NewNineSlice constructs a new NineSlice from i, having columns widths w and row heights h.
func NewNineSlice(i *ebiten.Image, w [3]int, h [3]int) *NineSlice {
	_ = "STUB: not implemented"
	return nil
}

func NewFixedNineSlice(i *ebiten.Image) *NineSlice { _ = "STUB: not implemented"; return nil }

// NewNineSliceSimple constructs a new NineSlice from image. borderWidthHeight specifies the width of the
// left and right column and the height of the top and bottom row. centerWidthHeight specifies the width
// of the center column and row.
func NewNineSliceSimple(image *ebiten.Image, borderWidthHeight int, centerWidthHeight int) *NineSlice {
	_ = "STUB: not implemented"
	return nil
}

// NewNineSliceSimple constructs a new NineSlice from image. borderWidthHeight specifies the width of the
// left and right column and the height of the top and bottom row. The center width and height is computed as
// the width of image minus 2*borderWidthHeight.
func NewNineSliceBorder(image *ebiten.Image, borderWidthHeight int) *NineSlice {
	_ = "STUB: not implemented"
	return nil
}

// NewNineSliceColor constructs a new NineSlice that when drawn fills with color c.
func NewNineSliceColor(c color.Color) *NineSlice { _ = "STUB: not implemented"; return nil }

// NewBorderedNineSliceColor constructs a new NineSlice that when drawn fills with color c and has a border with
// the specified color and width.
func NewBorderedNineSliceColor(bodyColor color.Color, borderColor color.Color, borderWidth int) *NineSlice {
	_ = "STUB: not implemented"
	return nil
}

// NewBorderedNineSliceColor constructs a new NineSlice that when drawn fills with the provided image and has a border with
// the specified color and width.
func NewBorderedNineSliceImage(img *ebiten.Image, borderColor color.Color, borderWidth int) *NineSlice {
	_ = "STUB: not implemented"
	return nil
}

// Set the border color.

// Draw the image in the middle.

// NewAdvancedNineSliceColor constructs a new NineSlice that when drawn fills with color c and
// has a border defined by the borders struct.
func NewAdvancedNineSliceColor(bodyColor color.Color, border borders) *NineSlice {
	_ = "STUB: not implemented"
	return nil
}

// NewAdvancedNineSliceImage constructs a new NineSlice that when drawn fills with the provided image and
// has a border defined by the borders struct.
func NewAdvancedNineSliceImage(img *ebiten.Image, border borders) *NineSlice {
	_ = "STUB: not implemented"
	return nil
}

// Set the border color.

// Draw the image in the middle.

// NewImageColor constructs a new Image that when drawn fills with color c.
func NewImageColor(c color.Color) *ebiten.Image { _ = "STUB: not implemented"; return nil }

// Draw draws n onto screen, with the size specified by width and height. If optsFunc is not nil, it is used to set
// DrawImageOptions for each tile drawn.
func (n *NineSlice) Draw(screen *ebiten.Image, width int, height int, optsFunc DrawImageOptionsFunc) {
	_ = "STUB: not implemented"
	return
}

func (n *NineSlice) drawTiles(screen *ebiten.Image, width int, height int, optsFunc DrawImageOptionsFunc) {
	_ = "STUB: not implemented"
	return
}

func (n *NineSlice) drawTile(screen *ebiten.Image, tile *ebiten.Image, tx int, ty int, sw int, sh int, tw int, th int, optsFunc DrawImageOptionsFunc) {
	_ = "STUB: not implemented"
	return
}

func (n *NineSlice) createTiles() { _ = "STUB: not implemented"; return }

func (n *NineSlice) centerOnly() bool { _ = "STUB: not implemented"; return false }

// MinSize returns the minimum width and height to draw n correctly. If n is drawn with a smaller size,
// the corner or edge tiles will overlap.
func (n *NineSlice) MinSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }
