package ebitenui

import (
	"github.com/ebitenui/ebitenui/event"
	"github.com/ebitenui/ebitenui/input"
	"github.com/ebitenui/ebitenui/widget"

	"github.com/hajimehoshi/ebiten/v2"
)

// UI encapsulates a complete user interface that can be rendered onto the screen.
// There should only be exactly one UI per application.
type UI struct {
	// Container is the root container of the UI hierarchy.
	Container widget.Containerer
	// If true the default tab/shift-tab to focus will be disabled
	DisableDefaultFocus bool

	// If true the default relayering of Windows will be disabled
	DisableWindowRelayering bool

	// This exposes a Render call before the Container is drawn,
	// but after the Windows with DrawLayer < 0 are drawn.
	PreRenderHook widget.RenderFunc

	// This exposes a Render call after the Container is drawn,
	// but before the Windows with DrawLayer >= 0 (all by default) are drawn.
	PostRenderHook widget.RenderFunc

	// Theme settings
	PrimaryTheme  *widget.Theme
	previousTheme *widget.Theme

	focusedWidget      widget.Focuser
	focusedWindow      *widget.Window
	focusedWindowIndex int
	inputLayerers      []input.Layerer
	windows            []*widget.Window

	previousContainer          widget.Containerer
	previousRemoveHandlerFuncs []event.RemoveHandlerFunc
	tabWasPressed              bool
	updObj                     *widget.UpdateObject

	debugMode bool
}

// Update updates u. This method should be called in the Ebiten Update function.
func (u *UI) Update() { _ = "STUB: not implemented"; return }

// Close all Ephemeral Windows (tooltip/dnd/etc).

// If widget is not visible or disabled, change focus to next widget.

func (u *UI) resetUpdateObject() {
	_ = "STUB: not implemented"
	// Reset update object
	return
}

// Draw renders u onto screen. This function should be called in the Ebiten Draw function.
func (u *UI) Draw(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

// Render elements that pop up (like combobox) on top of everything else

func (u *UI) setupInputLayers() { _ = "STUB: not implemented"; return }

func (u *UI) render(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func (u *UI) setTheme() {
	_ = "STUB: not implemented"
	// Handle the user setting a new theme.
	return
}

// Validate the main container with the new theme.

func (u *UI) handleContextMenu(args interface{}) { _ = "STUB: not implemented"; return }

func (u *UI) handleFocusEvent(args interface{}) { _ = "STUB: not implemented"; return }

// New widget focused

// Current widget focus removed

// Clicked out of focusable widgets
// If we didnt just click on the same widget

func (u *UI) handleToolTipEvent(args interface{}) { _ = "STUB: not implemented"; return }

func (u *UI) handleDragAndDropEvent(args interface{}) { _ = "STUB: not implemented"; return }

func (u *UI) getDropTargets() []widget.HasWidget { _ = "STUB: not implemented"; return nil }

// Loop through the windows array in reverse. If we find a modal window, only loop through its droppable widgets

func (u *UI) handleFocusChangeRequest() { _ = "STUB: not implemented"; return }

func (u *UI) ChangeFocus(direction widget.FocusDirection) { _ = "STUB: not implemented"; return }

// Loop through the windows array in reverse. If we find a modal window, only loop through its focusable widgets

// AddWindow adds window w to ui for rendering. It returns a function to remove w from ui.
func (u *UI) AddWindow(w *widget.Window) widget.RemoveWindowFunc {
	_ = "STUB: not implemented"
	return *new(widget.RemoveWindowFunc)
}

// AddWindowQuietly adds window w to ui for rendering. It returns a function to remove w from ui.
// This function allows you to specify if you would like it to close any open ephemeralWindows (tooltip/dnd/etc)
func (u *UI) AddWindowQuietly(w *widget.Window, closeEphemeralWindows bool) widget.RemoveWindowFunc {
	_ = "STUB: not implemented"
	return *new(widget.RemoveWindowFunc)
}

// Close all Ephemeral Windows (tooltip/dnd/etc)

func (u *UI) addWindow(w *widget.Window) bool { _ = "STUB: not implemented"; return false }

func (u *UI) removeWindow(w *widget.Window) { _ = "STUB: not implemented"; return }

// Used to close tooltips/dnd etc
func (u *UI) closeEphemeralWindows(windowIdx int) { _ = "STUB: not implemented"; return }

// This function returns true if the provided window object is currently active in this UI.
func (u *UI) IsWindowOpen(w *widget.Window) bool { _ = "STUB: not implemented"; return false }

// This function will re-sort the current windows attached to this UI based on its DrawLayer value.
func (u *UI) SortWindows() { _ = "STUB: not implemented"; return }

// This function will return true if any widget is currently focused or a Modal window is open.
func (u *UI) HasFocus() bool { _ = "STUB: not implemented"; return false }

// This function will unfocus the currently focused widget
func (u *UI) ClearFocus() { _ = "STUB: not implemented"; return }

// This function will return the currently focused widget if available otherwise it returns nil
func (u *UI) GetFocusedWidget() widget.Focuser {
	_ = "STUB: not implemented"
	return *

	// This function will set a specific focusable widget as the currently focused widget.
	new(widget.Focuser)
}

func (u *UI) SetFocusedWidget(focused widget.Focuser) { _ = "STUB: not implemented"; return }

// SetDebugMode enables or disables DebugMode which shows the
// margins of the Widgets when hovering the cursor on them
func (u *UI) SetDebugMode(dm bool) {
	_ = "STUB: not implemented"

	// GetDebugMode get's the current value of DebugMode
	return
}

func (u *UI) GetDebugMode() bool { _ = "STUB: not implemented"; return false }
