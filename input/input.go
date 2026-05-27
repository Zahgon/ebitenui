package input

import (
	"image"

	internalinput "github.com/ebitenui/ebitenui/internal/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type CursorUpdater interface {
	// Called every Update call from Ebiten
	// Note that before this is called the current cursor shape is reset to DEFAULT every cycle
	Update()
	// Called at the end of every Update call
	AfterUpdate()
	// Called at the beginning of every Draw call.
	Draw(screen *ebiten.Image)
	// Called at the end of every Draw call
	AfterDraw(screen *ebiten.Image)
	// MouseButtonPressed returns whether mouse button b is currently pressed.
	MouseButtonPressed(b ebiten.MouseButton) bool
	// MouseButtonJustPressed returns whether mouse button b has just been pressed.
	// It only returns true during the first frame that the button is pressed.
	MouseButtonJustPressed(b ebiten.MouseButton) bool
	// MouseButtonJustReleased returns whether mouse button b has just been released.
	// It only returns true during the first frame that the button is released.
	MouseButtonJustReleased(b ebiten.MouseButton) bool
	// CursorPosition returns the current cursor position.
	// If you define a CursorPosition that doesn't align with a system cursor you will need to
	// set the CursorDrawMode to Custom. This is because ebiten doesn't have a way to set the
	// cursor location manually
	CursorPosition() (int, int)
	// Returns the image to use as the cursor
	// EbitenUI by default will look for the following cursors:
	//  "EWResize"
	//  "NSResize"
	//  "Default"
	GetCursorImage(name string) *ebiten.Image
	// Returns how far from the CursorPosition to offset the cursor image.
	// This is best used with cursors such as resizing.
	GetCursorOffset(name string) image.Point
}

// This flag allows you to disable ebitenui's cursor management
var CursorManagementEnabled = true

// This variable indicates if the UI has currently being hovered over
var UIHovered = false

var currentCursorUpdater CursorUpdater = internalinput.InputHandler
var windowSize image.Point

// If the system cannot find a cursor image, it will revert to the system defaults.
// If cursorUpdater is nil the system will revert to the standard InputHandler system
//
// EbitenUI by default will look for the following cursors:
//
//	CURSOR_EWRESIZE  : "Cursor_EWResize"
//	CURSOR_NSRESIZE  : "Cursor_NSResize"
//	CURSOR_DEFAULT   : "Cursor_Default"
//	CURSOR_POINTER   : "Cursor_Pointer"
//	CURSOR_TEXT      : "Cursor_Text"
//	CURSOR_CROSSHAIR : "Cursor_Crosshair"
func SetCursorUpdater(cursorUpdater CursorUpdater) { _ = "STUB: not implemented"; return }

const (
	CURSOR_DEFAULT   = "Cursor_Default"
	CURSOR_EWRESIZE  = "Cursor_EWResize"
	CURSOR_NSRESIZE  = "Cursor_NSResize"
	CURSOR_POINTER   = "Cursor_Pointer"
	CURSOR_TEXT      = "Cursor_Text"
	CURSOR_CROSSHAIR = "Cursor_Crosshair"
	CURSOR_NONE      = "Cursor_None"
)

var currentCursor string = CURSOR_DEFAULT

func SetCursorShape(name string) { _ = "STUB: not implemented"; return }

func SetCursorImage(name string, cursorImage *ebiten.Image) { _ = "STUB: not implemented"; return }

func SetCursorImageWithOffset(name string, cursorImage *ebiten.Image, offset image.Point) {
	_ = "STUB: not implemented"
	return
}

// MouseButtonPressed returns whether mouse button b is currently pressed.
func MouseButtonPressed(b ebiten.MouseButton) bool { _ = "STUB: not implemented"; return false }

// MouseButtonJustPressed returns whether mouse button b has just been pressed.
// It only returns true during the first frame that the button is pressed.
func MouseButtonJustPressed(b ebiten.MouseButton) bool { _ = "STUB: not implemented"; return false }

// MouseButtonJustPressed returns whether mouse button b has just been pressed.
// It only returns true during the first frame that the button is pressed.
func MouseButtonJustReleased(b ebiten.MouseButton) bool { _ = "STUB: not implemented"; return false }

// MouseButtonPressedLayer returns whether mouse button b is currently pressed if input layer l is
// eligible to handle it.
func MouseButtonPressedLayer(b ebiten.MouseButton, l *Layer) bool {
	_ = "STUB: not implemented"
	return false
}

// MouseButtonJustPressedLayer returns whether mouse button b has just been pressed if input layer l
// is eligible to handle it. It only returns true during the first frame that the button is pressed.
func MouseButtonJustPressedLayer(b ebiten.MouseButton, l *Layer) bool {
	_ = "STUB: not implemented"
	return false
}

// MouseButtonJustPressedLayer returns whether mouse button b has just been pressed if input layer l
// is eligible to handle it. It only returns true during the first frame that the button is pressed.
func MouseButtonJustReleasedLayer(b ebiten.MouseButton, l *Layer) bool {
	_ = "STUB: not implemented"
	return false
}

// CursorPosition returns the current cursor position.
func CursorPosition() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// Wheel returns current mouse wheel movement.
func Wheel() (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

// WheelLayer returns current mouse wheel movement if input layer l is eligible to handle it.
// If l is not eligible, it returns 0, 0.
func WheelLayer(l *Layer) (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

// InputChars returns user keyboard input.
func InputChars() []rune {
	_ = "STUB: not implemented" //nolint:golint
	return nil
}

// KeyPressed returns whether key k is currently pressed.
func KeyPressed(k ebiten.Key) bool { _ = "STUB: not implemented"; return false }

// AnyKeyPressed returns whether any key is currently pressed.
func AnyKeyPressed() bool { _ = "STUB: not implemented"; return false }

// This method returns the drawable screen size whether it is fullscreen or not.
func GetWindowSize() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

func Update() { _ = "STUB: not implemented"; return }

func AfterUpdate() { _ = "STUB: not implemented"; return }

func Draw(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

func AfterDraw(screen *ebiten.Image) { _ = "STUB: not implemented"; return }

// Process Cursor

//If cursor outside the window do nothing

// If we have a cursor image hide current cursor and use it

// If we don't have an image use the system shapes.
