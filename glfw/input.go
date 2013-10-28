package glfw

//#include <GLFW/glfw3.h>
//float getAxesAtIndex(float* axes, int index);
//unsigned char getButtonAtIndex(unsigned char* buttons, int index);
import "C"

import "errors"

type Joystick int

const (
	Joystick1    Joystick = C.GLFW_JOYSTICK_1
	Joystick2             = C.GLFW_JOYSTICK_2
	Joystick3             = C.GLFW_JOYSTICK_3
	Joystick4             = C.GLFW_JOYSTICK_4
	Joystick5             = C.GLFW_JOYSTICK_5
	Joystick6             = C.GLFW_JOYSTICK_6
	Joystick7             = C.GLFW_JOYSTICK_7
	Joystick8             = C.GLFW_JOYSTICK_8
	Joystick9             = C.GLFW_JOYSTICK_9
	Joystick10            = C.GLFW_JOYSTICK_10
	Joystick11            = C.GLFW_JOYSTICK_11
	Joystick12            = C.GLFW_JOYSTICK_12
	Joystick13            = C.GLFW_JOYSTICK_13
	Joystick14            = C.GLFW_JOYSTICK_14
	Joystick15            = C.GLFW_JOYSTICK_15
	Joystick16            = C.GLFW_JOYSTICK_16
	JoystickLast          = C.GLFW_JOYSTICK_LAST
)

type InputAction int

const (
	Press   InputAction = C.GLFW_PRESS
	Release             = C.GLFW_RELEASE
	Repeat              = C.GLFW_REPEAT
)

type Key int

const (
	KeyUnknown      Key = C.GLFW_KEY_UNKNOWN
	KeySpace            = C.GLFW_KEY_SPACE
	KeyApostrophe       = C.GLFW_KEY_APOSTROPHE
	KeyComma            = C.GLFW_KEY_COMMA
	KeyMinus            = C.GLFW_KEY_MINUS
	KeyPeriod           = C.GLFW_KEY_PERIOD
	KeySlash            = C.GLFW_KEY_SLASH
	Key0                = C.GLFW_KEY_0
	Key1                = C.GLFW_KEY_1
	Key2                = C.GLFW_KEY_2
	Key3                = C.GLFW_KEY_3
	Key4                = C.GLFW_KEY_4
	Key5                = C.GLFW_KEY_5
	Key6                = C.GLFW_KEY_6
	Key7                = C.GLFW_KEY_7
	Key8                = C.GLFW_KEY_8
	Key9                = C.GLFW_KEY_9
	KeySemicolon        = C.GLFW_KEY_SEMICOLON
	KeyEqual            = C.GLFW_KEY_EQUAL
	KeyA                = C.GLFW_KEY_A
	KeyB                = C.GLFW_KEY_B
	KeyC                = C.GLFW_KEY_C
	KeyD                = C.GLFW_KEY_D
	KeyE                = C.GLFW_KEY_E
	KeyF                = C.GLFW_KEY_F
	KeyG                = C.GLFW_KEY_G
	KeyH                = C.GLFW_KEY_H
	KeyI                = C.GLFW_KEY_I
	KeyJ                = C.GLFW_KEY_J
	KeyK                = C.GLFW_KEY_K
	KeyL                = C.GLFW_KEY_L
	KeyM                = C.GLFW_KEY_M
	KeyN                = C.GLFW_KEY_N
	KeyO                = C.GLFW_KEY_O
	KeyP                = C.GLFW_KEY_P
	KeyQ                = C.GLFW_KEY_Q
	KeyR                = C.GLFW_KEY_R
	KeyS                = C.GLFW_KEY_S
	KeyT                = C.GLFW_KEY_T
	KeyU                = C.GLFW_KEY_U
	KeyV                = C.GLFW_KEY_V
	KeyW                = C.GLFW_KEY_W
	KeyX                = C.GLFW_KEY_X
	KeyY                = C.GLFW_KEY_Y
	KeyZ                = C.GLFW_KEY_Z
	KeyLeftBracket      = C.GLFW_KEY_LEFT_BRACKET
	KeyBackslash        = C.GLFW_KEY_BACKSLASH
	KeyRightBracket     = C.GLFW_KEY_RIGHT_BRACKET
	KeyGraveAccent      = C.GLFW_KEY_GRAVE_ACCENT
	KeyWorld1           = C.GLFW_KEY_WORLD_1
	KeyWorld2           = C.GLFW_KEY_WORLD_2
	KeyEscape           = C.GLFW_KEY_ESCAPE
	KeyEnter            = C.GLFW_KEY_ENTER
	KeyTab              = C.GLFW_KEY_TAB
	KeyBackspace        = C.GLFW_KEY_BACKSPACE
	KeyInsert           = C.GLFW_KEY_INSERT
	KeyDelete           = C.GLFW_KEY_DELETE
	KeyRight            = C.GLFW_KEY_RIGHT
	KeyLeft             = C.GLFW_KEY_LEFT
	KeyDown             = C.GLFW_KEY_DOWN
	KeyUp               = C.GLFW_KEY_UP
	KeyPageUp           = C.GLFW_KEY_PAGE_UP
	KeyPageDown         = C.GLFW_KEY_PAGE_DOWN
	KeyHome             = C.GLFW_KEY_HOME
	KeyEnd              = C.GLFW_KEY_END
	KeyCapsLock         = C.GLFW_KEY_CAPS_LOCK
	KeyScrollLock       = C.GLFW_KEY_SCROLL_LOCK
	KeyNumLock          = C.GLFW_KEY_NUM_LOCK
	KeyPrintScreen      = C.GLFW_KEY_PRINT_SCREEN
	KeyPause            = C.GLFW_KEY_PAUSE
	KeyF1               = C.GLFW_KEY_F1
	KeyF2               = C.GLFW_KEY_F2
	KeyF3               = C.GLFW_KEY_F3
	KeyF4               = C.GLFW_KEY_F4
	KeyF5               = C.GLFW_KEY_F5
	KeyF6               = C.GLFW_KEY_F6
	KeyF7               = C.GLFW_KEY_F7
	KeyF8               = C.GLFW_KEY_F8
	KeyF9               = C.GLFW_KEY_F9
	KeyF10              = C.GLFW_KEY_F10
	KeyF11              = C.GLFW_KEY_F11
	KeyF12              = C.GLFW_KEY_F12
	KeyF13              = C.GLFW_KEY_F13
	KeyF14              = C.GLFW_KEY_F14
	KeyF15              = C.GLFW_KEY_F15
	KeyF16              = C.GLFW_KEY_F16
	KeyF17              = C.GLFW_KEY_F17
	KeyF18              = C.GLFW_KEY_F18
	KeyF19              = C.GLFW_KEY_F19
	KeyF20              = C.GLFW_KEY_F20
	KeyF21              = C.GLFW_KEY_F21
	KeyF22              = C.GLFW_KEY_F22
	KeyF23              = C.GLFW_KEY_F23
	KeyF24              = C.GLFW_KEY_F24
	KeyF25              = C.GLFW_KEY_F25
	KeyKp0              = C.GLFW_KEY_KP_0
	KeyKp1              = C.GLFW_KEY_KP_1
	KeyKp2              = C.GLFW_KEY_KP_2
	KeyKp3              = C.GLFW_KEY_KP_3
	KeyKp4              = C.GLFW_KEY_KP_4
	KeyKp5              = C.GLFW_KEY_KP_5
	KeyKp6              = C.GLFW_KEY_KP_6
	KeyKp7              = C.GLFW_KEY_KP_7
	KeyKp8              = C.GLFW_KEY_KP_8
	KeyKp9              = C.GLFW_KEY_KP_9
	KeyKpDecimal        = C.GLFW_KEY_KP_DECIMAL
	KeyKpDivide         = C.GLFW_KEY_KP_DIVIDE
	KeyKpMultiply       = C.GLFW_KEY_KP_MULTIPLY
	KeyKpSubtract       = C.GLFW_KEY_KP_SUBTRACT
	KeyKpAdd            = C.GLFW_KEY_KP_ADD
	KeyKpEnter          = C.GLFW_KEY_KP_ENTER
	KeyKpEqual          = C.GLFW_KEY_KP_EQUAL
	KeyLeftShift        = C.GLFW_KEY_LEFT_SHIFT
	KeyLeftControl      = C.GLFW_KEY_LEFT_CONTROL
	KeyLeftAlt          = C.GLFW_KEY_LEFT_ALT
	KeyLeftSuper        = C.GLFW_KEY_LEFT_SUPER
	KeyRightShift       = C.GLFW_KEY_RIGHT_SHIFT
	KeyRightControl     = C.GLFW_KEY_RIGHT_CONTROL
	KeyRightAlt         = C.GLFW_KEY_RIGHT_ALT
	KeyRightSuper       = C.GLFW_KEY_RIGHT_SUPER
	KeyMenu             = C.GLFW_KEY_MENU
	KeyLast             = C.GLFW_KEY_LAST
)

type MouseButton int

const (
	MouseButton1      MouseButton = C.GLFW_MOUSE_BUTTON_1
	MouseButton2                  = C.GLFW_MOUSE_BUTTON_2
	MouseButton3                  = C.GLFW_MOUSE_BUTTON_3
	MouseButton4                  = C.GLFW_MOUSE_BUTTON_4
	MouseButton5                  = C.GLFW_MOUSE_BUTTON_5
	MouseButton6                  = C.GLFW_MOUSE_BUTTON_6
	MouseButton7                  = C.GLFW_MOUSE_BUTTON_7
	MouseButton8                  = C.GLFW_MOUSE_BUTTON_8
	MouseButtonLast               = C.GLFW_MOUSE_BUTTON_LAST
	MouseButtonLeft               = C.GLFW_MOUSE_BUTTON_LEFT
	MouseButtonRight              = C.GLFW_MOUSE_BUTTON_RIGHT
	MouseButtonMiddle             = C.GLFW_MOUSE_BUTTON_MIDDLE
)

type InputMode int

const (
	Cursor             InputMode = C.GLFW_CURSOR
	StickyKeys                   = C.GLFW_STICKY_KEYS
	StickyMouseButtons           = C.GLFW_STICKY_MOUSE_BUTTONS
)

const (
	CursorNormal   int = C.GLFW_CURSOR_NORMAL
	CursorHidden       = C.GLFW_CURSOR_HIDDEN
	CursorDisabled     = C.GLFW_CURSOR_DISABLED
	True               = C.GL_TRUE
	False              = C.GL_FALSE
)

/*
This function returns the last reported position of the cursor, in screen coordinates, relative to the upper-left corner of the client area of the specified window.

If the cursor is disabled (with CursorDisabled) then the cursor position is unbounded and limited only by the minimum and maximum values of a double.

The coordinate can be converted to their integer equivalents with the floor function. Casting directly to an integer type works for positive coordinates, but fails for negative ones.

Returns
	xpos, ypos
See Also
	SetCursorPos
*/
func (win *Window) GetCursorPosition() (float64, float64) {
	var xpos, ypos C.double
	C.glfwGetCursorPos(win.internalPtr, &xpos, &ypos)
	return float64(xpos), float64(ypos)
}

/*
Parameters
	mode	One of Cursor, StickyKeys or StickyMouseButtons.
See Also
	SetInputMode
*/
func (win *Window) GetInputMode(mode InputMode) int {
	return int(C.glfwGetInputMode(win.internalPtr, C.int(mode)))
}

/*
This function returns the values of all axes of the specified joystick.

Parameters
	joy	The joystick to query.
Returns
	An array of axis values.
*/
func GetJoystickAxes(joy Joystick) ([]float32, error) {
	var count C.int

	cAxes := C.glfwGetJoystickAxes(C.int(joy), &count)
	if cAxes == nil {
		return nil, errors.New("glfw: No such joystick.")
	}

	axes := make([]float32, int(count))
	for t := 0; t < int(count); t++ {
		axes[t] = float32(C.getAxesAtIndex(cAxes, C.int(t)))
	}

	return axes, nil
}

/*
This function returns the state of all buttons of the specified joystick.

Parameters
	joy	The joystick to query.
Returns
	An array of button states.
*/
func GetJoystickButtons(joy Joystick) ([]InputAction, error) {
	var count C.int

	cButtons := C.glfwGetJoystickButtons(C.int(joy), &count)
	if cButtons == nil {
		return nil, errors.New("glfw: No such joystick.")
	}

	buttons := make([]InputAction, int(count))
	for t := 0; t < int(count); t++ {
		buttons[t] = InputAction(C.getButtonAtIndex(cButtons, C.int(t)))
	}

	return buttons, nil
}

/*
This function returns the name, encoded as UTF-8, of the specified joystick.

Parameters
	joy	The joystick to query.
Returns
	The UTF-8 encoded name of the joystick.
*/
func GetJoystickName(joy Joystick) (string, error) {
	if cName := C.glfwGetJoystickName(C.int(joy)); cName != nil {
		return C.GoString(cName), nil
	}
	return "", errors.New("glfw: Joystick not present.")
}

/*
This function returns the last state reported for the specified key to the specified window. The returned state is one of Press or Release. The higher-level state Repeat is only reported to the key callback.

If the StickKeys input mode is enabled, this function returns Press the first time you call this function after a key has been pressed, even if the key has already been released.

The key functions deal with physical keys, with key tokens named after their use on the standard US keyboard layout. If you want to input text, use the Unicode character callback instead.

Parameters
	key	The desired keyboard key.
Returns
	One of Press or Release.
Note
	KeyUnknown is not a valid key for this function.
*/
func (win *Window) GetKey(key Key) InputAction {
	return InputAction(C.glfwGetKey(win.internalPtr, C.int(key)))
}

/*
This function returns the last state reported for the specified mouse button to the specified window.

If the StickyMouseButtons input mode is enabled, this function returns Press the first time you call this function after a mouse button has been pressed, even if the mouse button has already been released.

Parameters
	button	The desired mouse button.
Returns
	One of Press or Release.
*/
func (win *Window) GetMouseButton(button MouseButton) InputAction {
	return InputAction(C.glfwGetMouseButton(win.internalPtr, C.int(button)))
}

/*
This function returns whether the specified joystick is present.

Parameters
	joy	The joystick to query.
Returns
	If joystick is present.
*/
func JoystickPresent(joy Joystick) bool {
	if cVal := C.glfwJoystickPresent(C.int(joy)); cVal == C.GL_TRUE {
		return true
	}
	return false
}

/*
Parameters
	mode	One of GLFW_CURSOR, GLFW_STICKY_KEYS or GLFW_STICKY_MOUSE_BUTTONS.
	value	The new value of the specified input mode.
If mode is Cursor, the value must be one of the supported input modes:

CursorNormal makes the cursor visible and behaving normally.
CursorHidden makes the cursor invisible when it is over the client area of the window.
CursorDisabled disables the cursor and removes any limitations on cursor movement.
If mode is StickyKeys, the value must be either True to enable sticky keys, or False to disable it. If sticky keys are enabled, a key press will ensure that glfwGetKey returns GLFW_PRESS the next time it is called even if the key had been released before the call. This is useful when you are only interested in whether keys have been pressed but not when or in which order.

If mode is StickyMouseButtons, the value must be either True to enable sticky mouse buttons, or False to disable it. If sticky mouse buttons are enabled, a mouse button press will ensure that glfwGetMouseButton returns GLFW_PRESS the next time it is called even if the mouse button had been released before the call. This is useful when you are only interested in whether mouse buttons have been pressed but not when or in which order.

See Also
	GetInputMode
*/
func (win *Window) SetInputMode(mode InputMode, value int) {
	C.glfwSetInputMode(win.internalPtr, C.int(mode), C.int(value))
}
