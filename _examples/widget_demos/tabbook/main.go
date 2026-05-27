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

// Game object used by ebiten
type game struct {
	UI       ebitenui.UI
	TabRed   *widget.TabBookTab
	TabGreen *widget.TabBookTab
	TabBlue  *widget.TabBookTab
	TabBook  *widget.TabBook
}

func main() {
	game := game{}

	// load images for button states: idle, hover, and pressed
	buttonImage, _ := loadButtonImage()

	// load text font
	face, _ := loadFont(16)

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	// Create the first tab
	// A TabBookTab is a labelled container. The text here is what will show up in the tab button
	game.TabRed = widget.NewTabBookTab(
		widget.TabBookTabOpts.Label("Red Tab"),
		widget.TabBookTabOpts.ContainerOpts(
			widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{255, 0, 0, 255})),
			widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		),
	)

	redBtn := widget.NewText(
		widget.TextOpts.Text("Red Tab Button\nPress 'R' to select this tab.", &face, color.White),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	)
	game.TabRed.AddChild(redBtn)

	game.TabGreen = widget.NewTabBookTab(
		widget.TabBookTabOpts.Label("Green Tab"),
		widget.TabBookTabOpts.ContainerOpts(
			widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{0, 255, 0, 0xff})),
			widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		),
	)
	greenBtn := widget.NewText(
		widget.TextOpts.Text("Green Tab Button\nThis is configured as the initial tab.\nPress 'G' to select this tab.", &face, color.Black),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	)
	game.TabGreen.AddChild(greenBtn)

	blueImage := ebiten.NewImage(10, 10)
	blueImage.Fill(color.NRGBA{0, 0, 255, 255})

	game.TabBlue = widget.NewTabBookTab(
		widget.TabBookTabOpts.Label("Blue Tab"),
		widget.TabBookTabOpts.Image(&widget.GraphicImage{
			Idle:     blueImage,
			Disabled: blueImage,
			Pressed:  blueImage,
			Hover:    blueImage,
		}),
		widget.TabBookTabOpts.ContainerOpts(
			widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{0, 0, 255, 0xff})),
			widget.ContainerOpts.Layout(widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(5),
			)),
		),
	)
	blueBtn1 := widget.NewText(
		widget.TextOpts.Text("Blue Tab Button 1", &face, color.White),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		})),
	)
	game.TabBlue.AddChild(blueBtn1)
	blueBtn2 := widget.NewText(
		widget.TextOpts.Text("Press 'B' to select this tab.", &face, color.White),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		})),
	)
	game.TabBlue.AddChild(blueBtn2)

	tabDisabled := widget.NewTabBookTab(
		widget.TabBookTabOpts.Label("Disabled Tab"),
		widget.TabBookTabOpts.ContainerOpts(
			widget.ContainerOpts.BackgroundImage(e_image.NewNineSliceColor(color.NRGBA{R: 80, G: 80, B: 140, A: 255})),
		),
	)
	tabDisabled.Disabled = true
	game.TabBook = widget.NewTabBook(
		widget.TabBookOpts.TabButtonImage(buttonImage),
		widget.TabBookOpts.TabButtonText(&face, &widget.ButtonTextColor{Idle: color.White, Disabled: color.White}),
		widget.TabBookOpts.TabButtonSpacing(5),
		widget.TabBookOpts.ContentPadding(widget.NewInsetsSimple(5)),
		widget.TabBookOpts.ContentSpacing(10),
		widget.TabBookOpts.TabButtonMinSize(&image.Point{98, 40}),
		widget.TabBookOpts.ContainerOpts(
			widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				StretchHorizontal:  true,
				StretchVertical:    true,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			),
		),

		widget.TabBookOpts.Tabs(tabDisabled, game.TabRed, game.TabGreen, game.TabBlue),

		// Set the Initial Tab
		widget.TabBookOpts.InitialTab(game.TabGreen),
	)
	// add the tabBook as a child of the container
	rootContainer.AddChild(game.TabBook)

	// construct the UI
	game.UI = ebitenui.UI{
		Container: rootContainer,
	}

	// Ebiten setup
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Ebiten UI - Tabbook")

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

//Test that you can call Click on the focused widget.

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
