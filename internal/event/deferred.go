package event

// A DeferredAction is an action that is executed at a later time.
type DeferredAction interface {
	// Do executes the action.
	Do()
}

var deferredActions []DeferredAction

// AddDeferred adds d to the queue of deferred actions.
func AddDeferred(d DeferredAction) { _ = "STUB: not implemented"; return }

// ExecuteDeferred processes the queue of deferred actions and executes them.
func ExecuteDeferred() { _ = "STUB: not implemented"; return }
