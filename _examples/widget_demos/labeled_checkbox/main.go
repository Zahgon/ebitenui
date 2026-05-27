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
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed assets
var embeddedAssets embed.FS

// Game object used by ebiten.
type game struct {
	ui               *ebitenui.UI
	labeledCheckBox1 *widget.Checkbox
}

func main() {
	g := game{}
	// load images for button states: idle, hover, and pressed
	checkboxImage, _ := loadCheckboxImage()

	// load button text font
	face, _ := loadFont(20)

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(35),
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(30)),
			),
		),
	)

	g.labeledCheckBox1 = widget.NewCheckbox(
		// Set the labeled checkbox's position
		widget.CheckboxOpts.WidgetOpts(
			// Set the location of the checkbox
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  false,
			}),
			widget.WidgetOpts.MinSize(30, 30),
		),

		// Set the images
		widget.CheckboxOpts.Image(checkboxImage),

		// Set the label
		widget.CheckboxOpts.Text("Labeled Checkbox1", &face, &widget.LabelColor{
			Idle:     color.White,
			Disabled: color.White,
		}),

		// Set the spacing between the label and the checkbox
		widget.CheckboxOpts.Spacing(15),

		// Set the label to be before the checkbox.
		// widget.CheckboxOpts.LabelFirst(),

		// Set the state change handler
		widget.CheckboxOpts.StateChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
			if args.State == widget.WidgetChecked {
				fmt.Println("Checkbox1 is Checked")
			} else {
				fmt.Println("Checkbox1 is Unchecked")
			}
		}),
	)
	rootContainer.AddChild(g.labeledCheckBox1)

	labeledCheckBox2 := widget.NewCheckbox(
		// Set the labeled checkbox's position
		widget.CheckboxOpts.WidgetOpts(
			// Set the location of the checkbox
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  false,
			}),
			// Set the minimum size of the checkbox
			widget.WidgetOpts.MinSize(30, 30),
		),

		// Set the images
		widget.CheckboxOpts.Image(checkboxImage),

		// Set the label
		widget.CheckboxOpts.Text("Labeled Checkbox2", &face, &widget.LabelColor{
			Idle:     color.White,
			Disabled: color.White,
		}),

		// Set the spacing between the label and the checkbox
		widget.CheckboxOpts.Spacing(15),

		// Set the label to be before the checkbox.
		widget.CheckboxOpts.LabelFirst(),

		// Set the state change handler
		widget.CheckboxOpts.StateChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
			if args.State == widget.WidgetChecked {
				fmt.Println("Checkbox2 is Checked")
			} else {
				fmt.Println("Checkbox2 is Unchecked")
			}
		}),
	)
	// Set this checkbox as Checked by default
	labeledCheckBox2.SetState(widget.WidgetChecked)

	rootContainer.AddChild(labeledCheckBox2)
	// construct the UI
	ui := ebitenui.UI{
		Container: rootContainer,
	}

	// Ebiten setup
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Ebiten UI - Labeled  Checkbox")

	g.ui = &ui

	// run Ebiten main loop
	err := ebiten.RunGame(&g)
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

func loadFont(size float64) (text.Face, error) {
	_ = "STUB: not implemented"
	return *new(text.Face), nil
}
