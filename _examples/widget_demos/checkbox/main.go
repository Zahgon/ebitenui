package main

import (
	"embed"
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets
var embeddedAssets embed.FS

// Game object used by ebiten.
type game struct {
	ui       *ebitenui.UI
	checkBox *widget.Checkbox
}

func main() {
	game := game{}
	// load images for button states: idle, hover, and pressed
	checkboxImage, _ := loadCheckboxImage()

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	game.checkBox = widget.NewCheckbox(
		widget.CheckboxOpts.WidgetOpts(
			// Set the location of the checkbox
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		// Set the check object images
		widget.CheckboxOpts.Image(checkboxImage),
		// Set the state change handler
		widget.CheckboxOpts.StateChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
			switch args.State {
			case widget.WidgetChecked:
				fmt.Println("Checkbox is Checked")
			case widget.WidgetGreyed:
				fmt.Println("Checkbox is Greyed")
			case widget.WidgetUnchecked:
				fmt.Println("Checkbox is Unchecked")
			}
		}),
		widget.CheckboxOpts.TriState(),
		widget.CheckboxOpts.InitialState(widget.WidgetChecked),
	)

	rootContainer.AddChild(game.checkBox)

	// construct the UI
	game.ui = &ebitenui.UI{
		Container: rootContainer,
	}

	// Ebiten setup
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Ebiten UI - Checkbox")

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

func loadCheckboxImage() (*widget.CheckboxImage, error) { _ = "STUB: not implemented"; return nil, nil }
