package main

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

type page struct {
	title   string
	content widget.PreferredSizeLocateableWidget
}

func buttonPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func checkboxPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func listPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func textAreaPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func comboButtonPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func tabBookPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func gridLayoutPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func rowLayoutPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func sliderPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func progressBarPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func toolTipPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func dragAndDropPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func textInputPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func radioGroupPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func windowPage(res *uiResources, ui func() *ebitenui.UI) *page {
	_ = "STUB: not implemented"
	return nil
}

func openWindow(res *uiResources, ui func() *ebitenui.UI) { _ = "STUB: not implemented"; return }

func openWindow2(res *uiResources, ui func() *ebitenui.UI) { _ = "STUB: not implemented"; return }

func anchorLayoutPage(res *uiResources) *page { _ = "STUB: not implemented"; return nil }

func indexCheckbox(cs []*widget.Checkbox, c widget.RadioGroupElement) int {
	_ = "STUB: not implemented"
	return 0
}
