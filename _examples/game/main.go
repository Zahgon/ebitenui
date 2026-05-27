// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This file was modified from the example found: https://github.com/hajimehoshi/ebiten/blob/main/examples/tiles/main.go with permission from the author
package main

import (
	_ "image/png"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth  = 480
	screenHeight = 480
	tileSize     = 16
	tileMapWidth = 15
)

func main() {
	g := &Game{
		layers:     getLayers(),
		tilesImage: getTileImage(),
	}
	g.ui = g.getEbitenUI()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game Demo")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	tilesImage *ebiten.Image
	layers     [][]int

	ui        *ebitenui.UI
	headerLbl *widget.Text
}

func (g *Game) Update() error {
	_ = "STUB: not implemented"
	// Ensure that the UI is updated to receive events
	return nil
}

// Update the Label text to indicate if the ui is currently being hovered over or not

// Log out if we have clicked on the gamefield and NOT the ui

func (g *Game) Draw(screen *ebiten.Image) {
	_ = "STUB: not implemented"
	// Draw the tilemap
	return
}

// Ensure ui.Draw is called after the gameworld is drawn

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (g *Game) getEbitenUI() *ebitenui.UI {
	_ = "STUB: not implemented"
	// load label text font
	return nil
}

// construct a new container that serves as the root of the UI hierarchy

// the container will use an anchor layout to layout its single child widget

// Because this container has a backgroundImage set we track that the ui is hovered over.

// Uncomment this to not track that you are hovering over this header
// widget.WidgetOpts.TrackHover(false),

// Uncomment to force tracking hover of this element
// widget.WidgetOpts.TrackHover(true),

// Set the minimum size for the progress bar.
// This is necessary if you wish to have the progress bar be larger than
// the provided track image. In this exampe since we are using NineSliceColor
// which is 1px x 1px we must set a minimum size.

// Set this parameter to indicate we want do not want to track that this ui element is being hovered over.
// widget.WidgetOpts.TrackHover(false),

// Set the track images (Idle, Disabled).

// Set the progress images (Idle, Disabled).

// Set the min, max, and current values.

// Set how much of the track is displayed when the bar is overlayed.

// Create a label to show the percentage on top of the progress bar

//Call a render method after the rootContainer is drawn but before any ebitenui.Windows are drawn
//PostRenderHook: g.Render,

func (g *Game) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (g *Game) drawGameWorld(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func getTileImage() *ebiten.Image {
	_ = "STUB: not implemented"
	// Decode an image from the image file's byte slice.
	return nil
}

func getLayers() [][]int { _ = "STUB: not implemented"; return nil }

func loadFont(size float64) (text.Face, error) {
	_ = "STUB: not implemented"
	return *new(text.Face), nil
}
