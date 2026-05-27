package main

import (
	"github.com/ebitenui/ebitenui/widget"
)

type dragContents struct {
	res *uiResources

	text *widget.Text
}

func (d *dragContents) Create(sourceWidget widget.HasWidget) (*widget.Container, interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dragContents) Update(isDroppable bool, _ widget.HasWidget, _ interface{}) {
	_ = "STUB: not implemented"
	return
}
