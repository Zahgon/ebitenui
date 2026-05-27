package main

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	backgroundColor = "131a22"

	textIdleColor     = "dff4ff"
	textDisabledColor = "5a7a91"

	labelIdleColor     = textIdleColor
	labelDisabledColor = textDisabledColor

	buttonIdleColor     = textIdleColor
	buttonDisabledColor = labelDisabledColor

	listSelectedBackground         = "4b687a"
	listDisabledSelectedBackground = "2a3944"

	listFocusedBackground = "2a3944"

	headerColor = textIdleColor

	textInputCaretColor         = "e7c34b"
	textInputDisabledCaretColor = "766326"

	toolTipColor = backgroundColor

	separatorColor = listDisabledSelectedBackground
)

type uiResources struct {
	fonts *fonts

	background *image.NineSlice

	separatorColor color.Color

	text        *textResources
	button      *buttonResources
	label       *labelResources
	checkbox    *checkboxResources
	comboButton *comboButtonResources
	list        *listResources
	slider      *sliderResources
	progressBar *progressBarResources
	panel       *panelResources
	tabBook     *tabBookResources
	header      *headerResources
	textInput   *textInputResources
	textArea    *textAreaResources
	toolTip     *toolTipResources
}

type textResources struct {
	idleColor     color.Color
	disabledColor color.Color
	face          *text.Face
	titleFace     *text.Face
	bigTitleFace  *text.Face
	smallFace     *text.Face
}

type buttonResources struct {
	image   *widget.ButtonImage
	text    *widget.ButtonTextColor
	face    *text.Face
	padding *widget.Insets
}

type checkboxResources struct {
	image   *widget.CheckboxImage
	spacing int
}

type labelResources struct {
	text *widget.LabelColor
	face *text.Face
}

type comboButtonResources struct {
	image   *widget.ButtonImage
	text    *widget.ButtonTextColor
	face    *text.Face
	graphic *widget.GraphicImage
	padding *widget.Insets
}

type listResources struct {
	image        *widget.ScrollContainerImage
	track        *widget.SliderTrackImage
	trackPadding *widget.Insets
	handle       *widget.ButtonImage
	handleSize   *int
	face         *text.Face
	entry        *widget.ListEntryColor
	entryPadding *widget.Insets
}

type sliderResources struct {
	trackImage *widget.SliderTrackImage
	handle     *widget.ButtonImage
	handleSize *int
}

type progressBarResources struct {
	trackImage *widget.ProgressBarImage
	fillImage  *widget.ProgressBarImage
}

type panelResources struct {
	image    *image.NineSlice
	titleBar *image.NineSlice
	padding  *widget.Insets
}

type tabBookResources struct {
	buttonFace    *text.Face
	buttonText    *widget.ButtonTextColor
	buttonPadding *widget.Insets
}

type headerResources struct {
	background *image.NineSlice
	padding    *widget.Insets
	face       *text.Face
	color      color.Color
}

type textInputResources struct {
	image   *widget.TextInputImage
	padding *widget.Insets
	face    *text.Face
	color   *widget.TextInputColor
}

type textAreaResources struct {
	image        *widget.ScrollContainerImage
	track        *widget.SliderTrackImage
	trackPadding *widget.Insets
	handle       *widget.ButtonImage
	handleSize   *int
	face         *text.Face
	entryPadding *widget.Insets
}

type toolTipResources struct {
	background *image.NineSlice
	padding    *widget.Insets
	face       *text.Face
	color      color.Color
}

func newUIResources() (*uiResources, error) { _ = "STUB: not implemented"; return nil, nil }

func newButtonResources(fonts *fonts) (*buttonResources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCheckboxResources() (*checkboxResources, error) { _ = "STUB: not implemented"; return nil, nil }

func newLabelResources(fonts *fonts) *labelResources { _ = "STUB: not implemented"; return nil }

func newComboButtonResources(fonts *fonts) (*comboButtonResources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newListResources(fonts *fonts) (*listResources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newSliderResources() (*sliderResources, error) { _ = "STUB: not implemented"; return nil, nil }

func newProgressBarResources() (*progressBarResources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newPanelResources() (*panelResources, error) { _ = "STUB: not implemented"; return nil, nil }

func newTabBookResources(fonts *fonts) (*tabBookResources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newHeaderResources(fonts *fonts) (*headerResources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTextInputResources(fonts *fonts) (*textInputResources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTextAreaResources(fonts *fonts) (*textAreaResources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newToolTipResources(fonts *fonts) (*toolTipResources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hexToColor(h string) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }
