package event

// Event encapsulates an arbitrary event that event handlers may be interested in.
type Event struct {
	idCounter uint32
	handlers  []handler
}

// A HandlerFunc is a function that receives and handles an event. When firing an event using
// Event.Fire, arbitrary event arguments may be passed that are in turn passed on to the handler function.
type HandlerFunc func(args interface{})

// RemoveHandlerFunc is a function that removes a handler from an event.
type RemoveHandlerFunc func()

type handler struct {
	id uint32
	h  HandlerFunc
}

type deferredEvent struct {
	event *Event
	args  interface{}
}

type deferredAddHandler struct {
	event   *Event
	handler handler
}

// WrapHandler accepts a function of one argument and converts it into a HandlerFunc.
// Use this function when passing adding a new handler to an event object, such as
// button.ClickedEvent.AddHandler(WrapHandler(func (args *widget.ButtonClickedEventArgs){ ... }))
func WrapHandler[T any](f func(T)) HandlerFunc { _ = "STUB: not implemented"; return *new(HandlerFunc) }

// AddHandler registers event handler h with e. It returns a function to remove h from e if desired.
func (e *Event) AddHandler(h HandlerFunc) RemoveHandlerFunc {
	_ = "STUB: not implemented"
	return *new(RemoveHandlerFunc)
}

func (e *Event) removeHandler(id uint32) { _ = "STUB: not implemented"; return }

// Fire fires an event to all registered handlers. Arbitrary event arguments may be passed
// which are in turn passed on to event handlers.
//
// Events are not fired directly, but are put into a deferred queue. This queue is then
// processed by the UI.
func (e *Event) Fire(args interface{}) { _ = "STUB: not implemented"; return }

func (e *Event) handle(args interface{}) { _ = "STUB: not implemented"; return }

// Do implements DeferredAction.
func (e *deferredEvent) Do() { _ = "STUB: not implemented"; return }

// Do implements DeferredAction.
func (a *deferredAddHandler) Do() { _ = "STUB: not implemented"; return }

// AddEventHandlerOneShot registers event handler h with e. When e fires an event, h is removed from e immediately.
func AddEventHandlerOneShot(e *Event, h HandlerFunc) { _ = "STUB: not implemented"; return }
