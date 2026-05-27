package input

import (
	"image"
)

// Layerer may be implemented by widgets that need to set up input layers by calling AddLayer.
type Layerer interface {
	// SetupInputLayer sets up input layers. def may be called to defer additional input layer setup.
	SetupInputLayer(def DeferredSetupInputLayerFunc)
}

// SetupInputLayerFunc is a function that sets up input layers by calling AddLayer.
// def may be called to defer additional input layer setup.
type SetupInputLayerFunc func(def DeferredSetupInputLayerFunc)

// DeferredSetupInputLayerFunc is a function that stores s for deferred execution.
type DeferredSetupInputLayerFunc func(s SetupInputLayerFunc)

// A Layer is an input layer that can be used to block user input from lower layers of the user interface.
// For example, if two clickable areas overlap each other, clicking on the overlapping part should only result
// in a click event sent to the upper area instead of both. Input layers can be used to achieve this.
//
// Input layers are stacked: Lower layers may be eligible to handle an event if upper layers are not, or if
// upper layers specify to pass events on to lower layers regardless.
//
// Input layers may specify a screen Rectangle as their area of interest, or they may specify to cover the
// full screen.
type Layer struct {
	// DebugLabel is a label used in debugging to distinguish input layers. It is not used in any other way.
	DebugLabel string

	// EventTypes is a bit mask that specifies the types of events the input layer is eligible for.
	EventTypes LayerEventType

	// BlockLower specifies if events will be passed on to lower input layers even if the current layer
	// is eligible to handle them.
	BlockLower bool

	// FullScreen specifies if the input layer covers the full screen.
	FullScreen bool

	// RectFunc is a function that returns the input layer's screen area of interest. This function is only
	// called if FullScreen is false.
	RectFunc LayerRectFunc

	invalid bool
}

// LayerRectFunc is a function that returns a Layer's screen area of interest.
type LayerRectFunc func() image.Rectangle

// LayerEventType is a type of input event, such as mouse button press or release, wheel click, and so on.
type LayerEventType uint16

const (
	// LayerEventTypeAny is used for ActiveFor to indicate no special event types.
	LayerEventTypeAny = LayerEventType(0)
)

const (
	// LayerEventTypeMouseButton indicates an interest in mouse button events.
	LayerEventTypeMouseButton = LayerEventType(1 << iota)

	// LayerEventTypeWheel indicates an interest in mouse wheel events.
	LayerEventTypeWheel

	// LayerEventTypeAll indicates an interest in all event types.
	LayerEventTypeAll = LayerEventType(^uint16(0))
)

// DefaultLayer is the bottom-most input layer. It is a full screen layer that is eligible for all event types.
var DefaultLayer = Layer{
	DebugLabel: "default",
	EventTypes: LayerEventTypeAll,
	BlockLower: true,
	FullScreen: true,
}

var layers []*Layer

var deferredSetupInputLayers []SetupInputLayerFunc

// AddLayer adds l at the top of the layer stack.
//
// Layers are only valid for the duration of a frame. Layers are removed automatically for the next frame.
func AddLayer(l *Layer) { _ = "STUB: not implemented"; return }

// Valid returns whether l is still valid, that is, it has not been added to the layer stack in previous frames.
func (l *Layer) Valid() bool {
	_ = "STUB: not implemented"

	// ActiveFor returns whether l is eligible for an event of type eventType, according to l.EventTypes. It returns
	// false if l is not a fullscreen layer and does not contain the position x,y.
	return false
}

func (l *Layer) ActiveFor(x int, y int, eventType LayerEventType) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *Layer) contains(x int, y int) bool { _ = "STUB: not implemented"; return false }

// SetupInputLayersWithDeferred calls ls to set up input layers. This function is called by the UI.
func SetupInputLayersWithDeferred(ls []Layerer) { _ = "STUB: not implemented"; return }

func setupDeferredInputLayers() { _ = "STUB: not implemented"; return }

func appendToDeferredSetupInputLayerQueue(s SetupInputLayerFunc) { _ = "STUB: not implemented"; return }
