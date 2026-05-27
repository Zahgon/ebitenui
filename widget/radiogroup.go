package widget

import (
	"github.com/ebitenui/ebitenui/event"
)

type WidgetState int

const (
	WidgetUnchecked = WidgetState(iota)
	WidgetChecked
	WidgetGreyed
)

type RadioGroupElement interface {
	SetState(state WidgetState)
	getStateChangedEvent() *event.Event
}

type RadioGroup struct {
	ChangedEvent *event.Event

	elements  []RadioGroupElement
	active    RadioGroupElement
	initial   RadioGroupElement
	listen    bool
	doneEvent *event.Event
}

type RadioGroupOpt func(r *RadioGroup)

type RadioGroupOptions struct {
}

type RadioGroupChangedEventArgs struct {
	Active RadioGroupElement
}

type RadioGroupChangedHandlerFunc func(args *RadioGroupChangedEventArgs)

var RadioGroupOpts RadioGroupOptions

func NewRadioGroup(opts ...RadioGroupOpt) *RadioGroup { _ = "STUB: not implemented"; return nil }

// use deferred event to initialize

func (o RadioGroupOptions) Elements(e ...RadioGroupElement) RadioGroupOpt {
	_ = "STUB: not implemented"
	return *new(RadioGroupOpt)
}

func (o RadioGroupOptions) ChangedHandler(f RadioGroupChangedHandlerFunc) RadioGroupOpt {
	_ = "STUB: not implemented"
	return *new(RadioGroupOpt)
}

// This function allows you to select which element should be selected initialially.
// Otherwise it will select the first element in the Elements array.
func (o RadioGroupOptions) InitialElement(e RadioGroupElement) RadioGroupOpt {
	_ = "STUB: not implemented"
	return *new(RadioGroupOpt)
}

func (r *RadioGroup) Active() RadioGroupElement {
	_ = "STUB: not implemented"
	return *new(RadioGroupElement)
}

func (r *RadioGroup) SetActive(a RadioGroupElement) { _ = "STUB: not implemented"; return }

// ignore unchecking and reset to checked

// SetState() fires deferred events, so we need something *after* those to tell us we should listen again

func (r *RadioGroup) create() { _ = "STUB: not implemented"; return }
