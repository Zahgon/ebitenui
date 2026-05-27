// toolbar.go
//
// Toolbar struct and related functions.
//

package main

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

// NOTE: It's not strictly necessary to store references to all the buttons in the toolbar struct, but this example does
// so for completeness' sake. When you keep a reference to buttons in the struct, you can later configure them to respond
// to certain events in your application, and keep your program's logic outside the toolbar.
type toolbar struct {
	container   *widget.Container
	fileMenu    *widget.Button
	editMenu    *widget.Button
	helpButton  *widget.Button
	saveButton  *widget.Button
	quitButton  *widget.Button
	loadButton  *widget.Button
	undoButton  *widget.Button
	redoButton  *widget.Button
	cutButton   *widget.Button
	copyButton  *widget.Button
	pasteButton *widget.Button
}

func newToolbar(ui *ebitenui.UI, res *resources) *toolbar {
	_ = "STUB: not implemented"
	// Create a root container for the toolbar.
	return nil
}

// Use black background for the toolbar.

// Toolbar components must be aligned horizontally.

// Make the toolbar fill the whole horizontal space of the screen.

//
// "File" menu
//

// Make the toolbar entry open a menu with our "save" and "load" entries  when the user clicks it.

//
// "Edit" menu
// This is the same thing as the "File" menu, just with more entries.
//

//
// "Help" button
// Unlike the "File" and "Edit" menu, this is just a regular button on the toolbar - it does not open a menu.
// You can configure it to do something else when it's pressed, like opening a "Help" window.
//

func newToolbarButton(res *resources, label string) *widget.Button {
	_ = "STUB: not implemented"
	// Create a button for the toolbar.
	return nil
}

func newToolbarMenuEntry(res *resources, label string) *widget.Button {
	_ = "STUB: not implemented"
	// Create a button for a menu entry.
	return nil
}

func openToolbarMenu(opener *widget.Widget, ui *ebitenui.UI, entries ...*widget.Button) {
	_ = "STUB: not implemented"
	return
}

// Set the background to a translucent black.

// Menu entries should be arranged vertically.

// Set the minimum size for the menu.

// Set the menu to be a modal. This makes it block UI interactions to anything ese.

// Close the menu if the user clicks outside of it.

// Position the menu below the menu button that it belongs to.

// Immediately add the menu to the UI.
