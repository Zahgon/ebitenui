package widget

import (
	img "image"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type Container struct {
	definedParams       PanelParams
	computedParams      PanelParams
	AutoDisableChildren bool

	widgetOpts  []WidgetOpt
	layout      Layouter
	layoutDirty bool
	validated   bool

	relayoutParent        bool
	closeEphemeralWindows bool

	init     *MultiOnce
	widget   *Widget
	children []PreferredSizeLocateableWidget
}

type ContainerOpt func(c *Container)

type RemoveChildFunc func()

type ContainerOptions struct {
}

var ContainerOpts ContainerOptions

type PreferredSizeLocateableWidget interface {
	HasWidget
	PreferredSizer
	Locateable
	Validate()
}

func NewContainer(opts ...ContainerOpt) *Container { _ = "STUB: not implemented"; return nil }

func (o ContainerOptions) WidgetOpts(opts ...WidgetOpt) ContainerOpt {
	_ = "STUB: not implemented"
	return *new(ContainerOpt)
}

// This will set the background image to the provided NineSlice. If this is set then
// we will automatically track that the UI has been hovered over for this container
// Use widget.WidgetOpts.TrackHover(false) to turn this off if desired.
func (o ContainerOptions) BackgroundImage(i *image.NineSlice) ContainerOpt {
	_ = "STUB: not implemented"
	return *new(ContainerOpt)
}

func (o ContainerOptions) AutoDisableChildren() ContainerOpt {
	_ = "STUB: not implemented"
	return *new(ContainerOpt)
}

func (o ContainerOptions) Layout(layout Layouter) ContainerOpt {
	_ = "STUB: not implemented"
	return *new(ContainerOpt)
}

func (c *Container) addChildInit(child PreferredSizeLocateableWidget) {
	_ = "STUB: not implemented"
	return
}

func (c *Container) AddChild(children ...PreferredSizeLocateableWidget) RemoveChildFunc {
	_ = "STUB: not implemented"
	return *new(RemoveChildFunc)
}

func (c *Container) ReplaceChild(remove PreferredSizeLocateableWidget, add PreferredSizeLocateableWidget) {
	_ = "STUB: not implemented"
	return
}

func closeWidget(w *Widget) { _ = "STUB: not implemented"; return }

func (c *Container) RemoveChild(child PreferredSizeLocateableWidget) {
	_ = "STUB: not implemented"
	return
}

func (c *Container) RemoveChildren() { _ = "STUB: not implemented"; return }

func (c *Container) Children() []PreferredSizeLocateableWidget {
	_ = "STUB: not implemented"
	return nil
}

func (c *Container) RequestRelayout() { _ = "STUB: not implemented"; return }

func (c *Container) GetWidget() *Widget { _ = "STUB: not implemented"; return nil }

func (c *Container) PreferredSize() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// Start with the background image min size if one is set

// If the preferred layout for the children is greater than the background image
// min size then use that

// If the set MinHeight or MinWidth are greater than calculated, use that

func (c *Container) SetLocation(rect img.Rectangle) { _ = "STUB: not implemented"; return }

func (c *Container) IsValidated() bool { _ = "STUB: not implemented"; return false }

func (c *Container) Validate() { _ = "STUB: not implemented"; return }

func (c *Container) SetBackgroundImage(image *image.NineSlice) { _ = "STUB: not implemented"; return }

func (c *Container) Render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (c *Container) Update(updObj *UpdateObject) { _ = "STUB: not implemented"; return }

func (c *Container) doLayout() { _ = "STUB: not implemented"; return }

func (c *Container) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	_ = "STUB: not implemented"
	return
}

func (c *Container) draw(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (c *Container) createWidget() { _ = "STUB: not implemented"; return }

func (c *Container) GetFocusers() []Focuser { _ = "STUB: not implemented"; return nil }

func (c *Container) GetDropTargets() []HasWidget { _ = "STUB: not implemented"; return nil }

// If the Widget has 'drop' implemented then
// we have to push them to the 'result' as
// it means it has a handler for it

// WidgetAt implements WidgetLocator.
func (c *Container) WidgetAt(x int, y int) HasWidget {
	_ = "STUB: not implemented"
	return *new(HasWidget)
}
