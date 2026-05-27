//go:build !js
// +build !js

package jsUtil

import "github.com/ebitenui/ebitenui/utilities/mobile"

func IsMobileBrowser() bool { _ = "STUB: not implemented"; return false }

func Prompt(mode mobile.InputMode, title string, value string, cursorPos int, yPos int, callback InsertCallBack, selectAll SelectAllCallback) {
	_ = "STUB: not implemented"
	return
}

func SetCursorPosition(cursorPos int, cursorPos2 int) { _ = "STUB: not implemented"; return }

func GetCursorPosition() int { _ = "STUB: not implemented"; return 0 }
