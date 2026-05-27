package main

import (
	"image"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	e_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Game object used by ebiten.
type game struct {
	ui *ebitenui.UI
}

func main() {
	// construct the UI
	ui := ebitenui.UI{}

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	// construct a new row layout container that is centered on the page
	centerContainer := widget.NewContainer(
		// Configure the container to be centered in it's parent
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		// Configure the container to be a Vertical Row Layout.
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(10),
			)),
	)

	window1 := createWindow(&ui, "Window 1")
	centerContainer.AddChild(createButton(&ui, window1, "Window 1", image.Pt(100, 50)))

	window2 := createWindow(&ui, "Window 2")
	centerContainer.AddChild(createButton(&ui, window2, "Window 2", image.Pt(100, 260)))

	rootContainer.AddChild(centerContainer)

	// Set Root Container
	ui.Container = rootContainer

	// Ebiten setup
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Ebiten UI - Window")

	game := game{
		ui: &ui,
	}

	// run Ebiten main loop
	err := ebiten.RunGame(&game)
	if err != nil {
		log.Println(err)
	}
}

func createButton(ui *ebitenui.UI, win *widget.Window, label string, winPos image.Point) *widget.Button {
	_ = "STUB: not implemented"
	// load images for button states: idle, hover, and pressed
	return nil
}

// load button text font

// set general widget options

// instruct the container's anchor layout to center the button both horizontally and vertically

// specify the images to use

// specify the button's text, the font face, and the color

// specify that the button's text needs some padding for correct display

// add a handler that reacts to clicking the button

// Get the preferred size of the content

// Create a rect with the preferred size of the content

// Use the Add method to move the window to the specified point

// Set the windows location to the rect.

// Add the window to the UI.
// Note: If the window is already added, this will just move the window and not add a duplicate.

func createWindow(ui *ebitenui.UI, label string) *widget.Window {
	_ = "STUB: not implemented"
	// load the font for the window title
	return nil
}

// load button text font

// Create the contents of the window

//widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255})),

// specify the images to use

// specify the button's text, the font face, and the color

// specify that the button's text needs some padding for correct display

// specify the images to use

// specify the button's text, the font face, and the color

// specify that the button's text needs some padding for correct display

// specify the images to use

// specify the button's text, the font face, and the color

// specify that the button's text needs some padding for correct display

// add a handler that reacts to clicking the button

// Create the titlebar for the window

// Create the new window object. The window object is not tied to a container. Its location and
// size are set manually using the SetLocation method on the window and added to the UI with ui.AddWindow()
// Set the Button callback below to see how the window is added to the UI.

// Set the main contents of the window

// Set the titlebar for the window (Optional)

// Set the window above everything else and block input elsewhere
// widget.WindowOpts.Modal(),
// Set how to close the window. CLICK_OUT will close the window when clicking anywhere
// that is not a part of the window object
// widget.WindowOpts.CloseMode(widget.CLICK_OUT),
// Indicates that the window is draggable. It must have a TitleBar for this to work

// Set the window resizeable
//widget.WindowOpts.Resizeable(),
// Set the minimum size the window can be

// Set the maximum size a window can be
//widget.WindowOpts.MaxSize(300, 300),

// Set the callback that triggers when a move is complete

// Set the callback that triggers when a resize is complete

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
