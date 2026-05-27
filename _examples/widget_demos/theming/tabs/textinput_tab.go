package tabs

import (
	"github.com/ebitenui/ebitenui/widget"
)

func NewTextInputTab() *widget.TabBookTab { _ = "STUB: not implemented"; return nil }

// construct a standard textinput widget

// Set the layout information to center the textbox in the parent

// This text is displayed if the input is empty

// This is called when the user hits the "Enter" key.
// There are other options that can configure this behavior.

// This is called whenver there is a change to the text

// construct a disabled textinput widget

// Set the layout information to center the textbox in the parent

// This text is displayed if the input is empty

// This is called when the user hits the "Enter" key.
// There are other options that can configure this behavior.

// This is called whenver there is a change to the text

// construct a secure textinput widget

// This parameter indicates that the inputted text should be hidden

// This method is called whenever there is a text change.
// It allows the developer to allow or deny a change.
// In this case we are limiting the string to 5 runes.
// The first return parameter is whether or not to accept the text as is.
// The second return parameter is what to replace the text with if it is not accepted (optional)

// This will do nothing because the validation above prevents this from being set.

// This method is called whenever there is a text change.
// It allows the developer to allow or deny a change.
// In this case we are forcing the string to be all caps.
// The first return parameter is whether or not to accept the text as is.
// The second return parameter is what to replace the text with if it is not accepted (optional)

// This will show in all caps due to validation function above
