package main

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	fontFaceRegular = "assets/fonts/notosans-regular.ttf"
	fontFaceBold    = "assets/fonts/notosans-bold.ttf"
)

type fonts struct {
	face         *text.Face
	titleFace    *text.Face
	bigTitleFace *text.Face
	toolTipFace  *text.Face
}

func loadFonts() (*fonts, error) { _ = "STUB: not implemented"; return nil, nil }

func loadFont(path string, size float64) (text.Face, error) {
	_ = "STUB: not implemented"
	return *new(text.Face), nil
}
