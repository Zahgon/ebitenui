package main

import (
	"embed"
	"image"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	e_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/input"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed assets
var embeddedAssets embed.FS

// Game object used by ebiten
type game struct {
	ui *ebitenui.UI
}

func main() {
	// load images for button states: idle, hover, and pressed
	buttonImage, _ := loadButtonImage()

	// load button text font
	face, _ := loadFont(20)

	ui := ebitenui.UI{}
	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	winContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{155, 155, 0, 255})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()), // the container will use a plain color as its background
	)
	winContainer.AddChild(widget.NewText(widget.TextOpts.Text("Click outside to close.\nResizable.\nUses System cursor for E/W.\nUses Custom cursor for N/S", &face, color.White)))

	win := widget.NewWindow(widget.WindowOpts.CloseMode(widget.CLICK_OUT), widget.WindowOpts.Contents(winContainer), widget.WindowOpts.Resizeable())

	// construct a button
	button := widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			widget.WidgetOpts.CursorHovered("buttonHover"),
			widget.WidgetOpts.CursorPressed("buttonPressed"),
		),

		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),

		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text("Hello, World!", &face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(&widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println("button clicked")
			win.SetLocation(image.Rect(50, 50, 350, 150))
			ui.AddWindow(win)
		}),
	)

	// add the button as a child of the container
	rootContainer.AddChild(button)

	// construct the UI
	ui.Container = rootContainer

	// Ebiten setup
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Ebiten UI - Cursor Shape")

	game := game{
		ui: &ui,
	}

	// Set the main cursor used within the application
	input.SetCursorImage(input.CURSOR_DEFAULT, loadNormalCursorImage())

	// Set the custom hover image
	input.SetCursorImage("buttonHover", loadHoverCursorImage())

	input.SetCursorImage("buttonPressed", loadPressedCursorImage())

	// Set the NS resize cursor with an offset so that it shows up a little above the cursor point
	input.SetCursorImageWithOffset(input.CURSOR_NSRESIZE, loadNSCursorImage(), image.Point{0, -6})

	// Disable cursor management by ebitenui
	// input.CursorManagementEnabled = false

	// run Ebiten main loop
	err := ebiten.RunGame(&game)
	if err != nil {
		log.Println(err)
	}
}

// Layout implements Game.
func (g *game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Update implements Game.
func (g *game) Update() error {
	_ = "STUB: not implemented"
	// update the UI
	return nil
}

// Draw implements Ebiten's Draw method.
func (g *game) Draw(screen *ebiten.Image) {
	_ = "STUB: not implemented"
	// draw the UI onto the screen
	return
}

func loadButtonImage() (*widget.ButtonImage, error) { _ = "STUB: not implemented"; return nil, nil }

func loadFont(size float64) (text.Face, error) {
	_ = "STUB: not implemented"
	return *new(text.Face), nil
}

func loadNormalCursorImage() *ebiten.Image { _ = "STUB: not implemented"; return nil }

func loadNSCursorImage() *ebiten.Image { _ = "STUB: not implemented"; return nil }

func loadHoverCursorImage() *ebiten.Image { _ = "STUB: not implemented"; return nil }

func loadPressedCursorImage() *ebiten.Image { _ = "STUB: not implemented"; return nil }
