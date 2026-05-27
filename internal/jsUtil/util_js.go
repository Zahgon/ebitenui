//go:build js
// +build js

package jsUtil

import (
	"regexp"
	"syscall/js"

	"github.com/ebitenui/ebitenui/utilities/mobile"
)

var MOBILE_BROWSER_REGEX = regexp.MustCompile("(?i)Android|webOS|iPhone|iPad|iPod|BlackBerry|Windows Phone")

var document js.Value

var insertCB InsertCallBack

var selectAllCB SelectAllCallback

var started bool

var offsetTop int

func init() {
	document = js.Global().Get("document")

	//Create a hidden html input element that will capture keystrokes
	p := document.Call("createElement", "input")
	p.Set("id", "tempInput")
	p.Set("style", "height:0px; width:1px; margin:0px; position: fixed; overflow:hidden; top:-10px; border:0px; padding:0px")
	document.Get("body").Call("appendChild", p)

	//Add a listener on the hidden html input for keystrokes
	p.Call("addEventListener", "input", js.FuncOf(handleInput), false)
	p.Call("addEventListener", "select", js.FuncOf(handleSelection), false)

	//Get the canvas and attach an event listener for screen touches
	requestAnimationFrame := js.Global().Get("requestAnimationFrame")
	requestAnimationFrame.Invoke(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		canvas := document.Get("body").Call("getElementsByTagName", "canvas").Index(0)
		offsetTop = canvas.Get("offsetTop").Int()
		canvas.Call("addEventListener", "touchstart", js.FuncOf(handleClick), false)
		canvas.Call("addEventListener", "touchend", js.FuncOf(handleClick), false)
		return nil
	}))
}

func IsMobileBrowser() bool { _ = "STUB: not implemented"; return false }

func Prompt(mode mobile.InputMode, title string, value string, cursorPos int, yPos int, cb InsertCallBack, sa SelectAllCallback) {
	_ = "STUB: not implemented"
	return
}

//Configure the hidden html input element based on what our library has for the input

//Indicate we've started capturing input

// TODO: fix cursor position from being called every loop
func SetCursorPosition(cursorPos int, cursorPos2 int) { _ = "STUB: not implemented"; return }

func GetCursorPosition() int { _ = "STUB: not implemented"; return 0 }

func handleClick(this js.Value, args []js.Value) any {
	_ = "STUB: not implemented"
	// If we have clicked on one of the inputs, shift focus to the input to open the keyboard
	return *new(any)
}

var previousValue = ""
var previousPosition = 0

// Process changes on the hidden html text input
func handleInput(this js.Value, args []js.Value) any { _ = "STUB: not implemented"; return *new(any) }

func handleSelection(this js.Value, args []js.Value) any {
	_ = "STUB: not implemented"
	return *new(any)
}
