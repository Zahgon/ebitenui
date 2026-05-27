package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

type game struct {
	ui *ebitenui.UI
}

type pageContainer struct {
	widget    widget.PreferredSizeLocateableWidget
	titleText *widget.Text
	flipBook  *widget.FlipBook
}

func main() {
	ebiten.SetWindowSize(900, 800)
	ebiten.SetWindowTitle("Ebiten UI Demo")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetVsyncEnabled(true)

	ui, err := createUI()
	if err != nil {
		log.Fatal(err)
	}

	game := game{
		ui: ui,
	}

	err = ebiten.RunGame(&game)
	if err != nil {
		log.Print(err)
	}
}

func createUI() (*ebitenui.UI, error) { _ = "STUB: not implemented"; return nil, nil }

//This creates the root container for this UI.

// It is using a GridLayout with a single column

// It uses the Stretch parameter to define how the rows will be layed out.
// - a fixed sized header
// - a content row that stretches to fill all remaining space
// - a fixed sized footer

// Padding defines how much space to put around the outside of the grid.

// Spacing defines how much space to put between each column and row

func headerContainer(res *uiResources) widget.PreferredSizeLocateableWidget {
	_ = "STUB: not implemented"
	return *new(widget.PreferredSizeLocateableWidget)
}

func header(label string, res *uiResources, opts ...widget.ContainerOpt) widget.PreferredSizeLocateableWidget {
	_ = "STUB: not implemented"
	return *new(widget.PreferredSizeLocateableWidget)
}

func demoContainer(res *uiResources, ui func() *ebitenui.UI) widget.PreferredSizeLocateableWidget {
	_ = "STUB: not implemented"
	return *new(widget.PreferredSizeLocateableWidget)
}

func newPageContainer(res *uiResources) *pageContainer { _ = "STUB: not implemented"; return nil }

func (p *pageContainer) setPage(page *page) { _ = "STUB: not implemented"; return }

func newCheckbox(label string, changedHandler widget.CheckboxChangedHandlerFunc, res *uiResources) *widget.Checkbox {
	_ = "STUB: not implemented"
	return nil
}

func newPageContentContainer() *widget.Container { _ = "STUB: not implemented"; return nil }

func newListComboButton(entries []interface{}, buttonLabel widget.SelectComboButtonEntryLabelFunc, entryLabel widget.ListEntryLabelFunc,
	entrySelectedHandler widget.ListComboButtonEntrySelectedHandlerFunc, res *uiResources) *widget.ListComboButton {
	_ = "STUB: not implemented"
	return nil
}

func newList(entries []interface{}, res *uiResources, widgetOpts ...widget.WidgetOpt) *widget.List {
	_ = "STUB: not implemented"
	return nil
}

func newTextArea(text string, res *uiResources, widgetOpts ...widget.WidgetOpt) *widget.TextArea {
	_ = "STUB: not implemented"
	return nil
}

func newSeparator(res *uiResources, ld interface{}) widget.PreferredSizeLocateableWidget {
	_ = "STUB: not implemented"
	return *new(widget.PreferredSizeLocateableWidget)
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (g *game) Update() error { _ = "STUB: not implemented"; return nil }

func (g *game) Draw(screen *ebiten.Image) { _ = "STUB: not implemented"; return }
