package input

import (
	"image"
	"runtime"

	"github.com/ebitenui/ebitenui/internal/jsUtil"
	"github.com/hajimehoshi/ebiten/v2"
)

type DefaultInternalHandler struct {
	LeftMouseButtonPressed   bool
	MiddleMouseButtonPressed bool
	RightMouseButtonPressed  bool
	CursorX                  int
	CursorY                  int
	WheelX                   float64
	WheelY                   float64

	LeftMouseButtonJustPressed   bool
	MiddleMouseButtonJustPressed bool
	RightMouseButtonJustPressed  bool

	LeftMouseButtonJustReleased   bool
	MiddleMouseButtonJustReleased bool
	RightMouseButtonJustReleased  bool

	LastLeftMouseButtonPressed   bool
	LastMiddleMouseButtonPressed bool
	LastRightMouseButtonPressed  bool

	InputChars    []rune
	KeyPressed    map[ebiten.Key]bool
	AnyKeyPressed bool
	isTouched     bool
	cursorImages  map[string]*ebiten.Image
	cursorOffset  map[string]image.Point

	touchscreenPlatform bool
	touchIDs            []ebiten.TouchID
}

var InternalUIHovered = false

var InputHandler *DefaultInternalHandler = &DefaultInternalHandler{
	// A touchscreenPlatform is defined as a device that doesn't have a mouse pointer,
	// but has a touchscreen input instead.
	// For native builds, there are Android and IOS; Ebitengine defines a mobile platform
	// as these two build tags (they will always return {0,0} from ebiten.CursorPosition).
	// Then we add web builds that are running on a mobile browser.
	//
	// TODO: maybe move this platform-detection code to somewhere else?
	// There should be a context-like object that would infer the preferred platform
	// input options.
	touchscreenPlatform: jsUtil.IsMobileBrowser() || runtime.GOOS == "android" || runtime.GOOS == "ios",

	KeyPressed:   make(map[ebiten.Key]bool),
	cursorImages: make(map[string]*ebiten.Image),
	cursorOffset: make(map[string]image.Point),
}

// Update updates the input system. This is called by the UI.
func (handler *DefaultInternalHandler) Update() { _ = "STUB: not implemented"; return }

// Only execute this branch on non-mobile platforms.
// This is a workaround to keep the touch position intact,
// as ebiten.CursorPosition() would set it to (0, 0).
//
// TODO: maybe get rid of this special condition when fireEvents are
// moved to the Update() tree.
// See issue #100.

func (handler *DefaultInternalHandler) AfterUpdate() { _ = "STUB: not implemented"; return }

func (handler *DefaultInternalHandler) Draw(screen *ebiten.Image) {
	_ = "STUB: not implemented"
	return
}

func (handler *DefaultInternalHandler) AfterDraw(screen *ebiten.Image) {
	_ = "STUB: not implemented"
	return
}

func (handler *DefaultInternalHandler) MouseButtonPressed(b ebiten.MouseButton) bool {
	_ = "STUB: not implemented"
	return false
}

func (handler *DefaultInternalHandler) MouseButtonJustPressed(b ebiten.MouseButton) bool {
	_ = "STUB: not implemented"
	return false
}

func (handler *DefaultInternalHandler) MouseButtonJustReleased(b ebiten.MouseButton) bool {
	_ = "STUB: not implemented"
	return false
}

func (handler *DefaultInternalHandler) CursorPosition() (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (handler *DefaultInternalHandler) GetCursorImage(name string) *ebiten.Image {
	_ = "STUB: not implemented"
	return nil
}

func (handler *DefaultInternalHandler) GetCursorOffset(name string) image.Point {
	_ = "STUB: not implemented"
	return *new(image.Point)
}

func (handler *DefaultInternalHandler) SetCursorImage(name string, cursorImage *ebiten.Image, offset image.Point) {
	_ = "STUB: not implemented"
	return
}
