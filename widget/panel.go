package widget

import (
	"github.com/ebitenui/ebitenui/image"
)

type PanelParams struct {
	BackgroundImage *image.NineSlice
}

type Panel struct {
	Container
}

func NewPanel(opts ...ContainerOpt) *Panel { _ = "STUB: not implemented"; return nil }

func (p *Panel) Validate() { _ = "STUB: not implemented"; return }
