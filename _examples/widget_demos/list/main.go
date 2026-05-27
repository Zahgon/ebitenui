package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/utilities/constantutil"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Struct used as data for combobox. You could just use strings instead of a struct.
type ListEntry struct {
	id   int
	Name string
}

// Game object used by ebiten.
type game struct {
	ui   *ebitenui.UI
	list *widget.List
}

func main() {
	game := game{}
	// load images for button states: idle, hover, and pressed
	buttonImage, _ := loadButtonImage()

	// load button text font
	face, _ := loadFont(20)

	// Create array of list entries
	numEntries := 20
	id := 1
	entries := make([]any, 0, numEntries)
	for id = 1; id <= numEntries; id++ {
		entries = append(entries, ListEntry{id, fmt.Sprintf("Entry %d", id)})
	}

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	// Construct a list. This is one of the more complicated widgets to use since
	// it is composed of multiple widget types
	game.list = widget.NewList(
		// Set how wide the list should be
		widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(150, 0),
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
				StretchVertical:    true,
				Padding:            widget.NewInsetsSimple(50),
			}),
		)),
		// Set the entries in the list
		// widget.ListOpts.Entries(entries),
		widget.ListOpts.ScrollContainerImage(&widget.ScrollContainerImage{
			Idle:     image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			Disabled: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			Mask:     image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
		}),

		widget.ListOpts.SliderParams(&widget.SliderParams{
			// Set the background images/color for the background of the slider track
			TrackImage: &widget.SliderTrackImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
				Hover: image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			},
			HandleImage:   buttonImage,
			MinHandleSize: constantutil.ConstantToPointer(5),
			// Set how wide the track should be
			TrackPadding: widget.NewInsetsSimple(2),
		}),
		// Hide the horizontal slider
		widget.ListOpts.HideHorizontalSlider(),
		// Set the font for the list options
		widget.ListOpts.EntryFontFace(&face),
		// Set the colors for the list
		widget.ListOpts.EntryColor(&widget.ListEntryColor{
			Selected:                   color.NRGBA{R: 0, G: 255, B: 0, A: 255},     // Foreground color for the unfocused selected entry
			Unselected:                 color.NRGBA{R: 254, G: 255, B: 255, A: 255}, // Foreground color for the unfocused unselected entry
			SelectedBackground:         color.NRGBA{R: 130, G: 130, B: 200, A: 255}, // Background color for the unfocused selected entry
			SelectingBackground:        color.NRGBA{R: 130, G: 130, B: 130, A: 255}, // Background color for the unfocused being selected entry
			SelectingFocusedBackground: color.NRGBA{R: 130, G: 140, B: 170, A: 255}, // Background color for the focused being selected entry
			SelectedFocusedBackground:  color.NRGBA{R: 130, G: 130, B: 170, A: 255}, // Background color for the focused selected entry
			FocusedBackground:          color.NRGBA{R: 170, G: 170, B: 180, A: 255}, // Background color for the focused unselected entry
			DisabledUnselected:         color.NRGBA{R: 100, G: 100, B: 100, A: 255}, // Foreground color for the disabled unselected entry
			DisabledSelected:           color.NRGBA{R: 100, G: 100, B: 100, A: 255}, // Foreground color for the disabled selected entry
			DisabledSelectedBackground: color.NRGBA{R: 100, G: 100, B: 100, A: 255}, // Background color for the disabled selected entry
		}),
		// This required function returns the string displayed in the list.
		widget.ListOpts.EntryLabelFunc(func(e interface{}) string {
			return e.(ListEntry).Name
		}),
		// Padding for each entry
		widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
		// Text position for each entry
		widget.ListOpts.EntryTextPosition(widget.TextPositionStart, widget.TextPositionCenter),
		// This handler defines what function to run when a list item is selected.
		widget.ListOpts.EntrySelectedHandler(func(args *widget.ListEntrySelectedEventArgs) {
			entry := args.Entry.(ListEntry)
			fmt.Println("Entry Selected: ", entry)
		}),
		// This option will select the entry as it is focused
		// widget.ListOpts.SelectFocus(),
		// This option will select the entry when pressing the mouse button instead of releasing it
		// widget.ListOpts.SelectPressed(),

		// This option will disable default keys (up and down)
		// widget.ListOpts.DisableDefaultKeys(true),
	)

	// Add list to the root container
	rootContainer.AddChild(game.list)
	game.list.SetEntries(entries)
	game.list.AddEntry(ListEntry{999, "Entry 999"})

	buttonsContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
				Padding:            widget.NewInsetsSimple(5),
			}),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(3),
			widget.GridLayoutOpts.Spacing(10, 3),
		)),
	)

	// construct a button
	buttonAdd := widget.NewButton(

		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),

		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text("Add Entry", &face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(&widget.Insets{
			Left:   10,
			Right:  10,
			Top:    5,
			Bottom: 5,
		}),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			entryToAdd := ListEntry{id, fmt.Sprintf("Entry %d", id)}
			id++

			game.list.AddEntry(entryToAdd)
			game.list.SetSelectedEntry(entryToAdd)
		}),
	)
	buttonsContainer.AddChild(buttonAdd)

	// construct a button
	buttonRemove := widget.NewButton(

		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),

		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text("Remove Entry", &face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(&widget.Insets{
			Left:   10,
			Right:  10,
			Top:    5,
			Bottom: 5,
		}),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			game.list.RemoveEntry(game.list.SelectedEntry())
		}),
	)
	buttonsContainer.AddChild(buttonRemove)

	// construct a button
	buttonClean := widget.NewButton(

		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),

		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text("Clear", &face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(&widget.Insets{
			Left:   10,
			Right:  10,
			Top:    5,
			Bottom: 5,
		}),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			entries = nil
			game.list.SetEntries(entries)
		}),
	)
	buttonsContainer.AddChild(buttonClean)

	rootContainer.AddChild(buttonsContainer)

	// construct the UI
	ui := ebitenui.UI{
		Container: rootContainer,
	}

	// Ebiten setup
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Ebiten UI - List")

	game.ui = &ui

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
