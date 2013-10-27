package glfw

//#include <stdlib.h>
//#include <GLFW/glfw3.h>
//void goSetWindowPosCB(GLFWwindow* win);
//void goSetWindowSizeCB(GLFWwindow* win);
//void goSetFramebufferSizeCB(GLFWwindow* win);
//void goSetWindowCloseCB(GLFWwindow* win);
//void goSetWindowRefreshCB(GLFWwindow* win);
//void goSetWindowFocusCB(GLFWwindow* win);
//void goSetWindowIconifyCB(GLFWwindow* win);
import "C"

import (
	"errors"
	"unsafe"
)

type Hint int

const (
	Resizable           Hint = C.GLFW_RESIZABLE
	Visible                  = C.GLFW_VISIBLE
	Decorated                = C.GLFW_DECORATED
	RedBits                  = C.GLFW_RED_BITS
	GreenBits                = C.GLFW_GREEN_BITS
	BlueBits                 = C.GLFW_BLUE_BITS
	AlphaBits                = C.GLFW_ALPHA_BITS
	DepthBits                = C.GLFW_DEPTH_BITS
	StencilBits              = C.GLFW_STENCIL_BITS
	AccumRedBits             = C.GLFW_ACCUM_RED_BITS
	AccumGreenBits           = C.GLFW_ACCUM_GREEN_BITS
	AccumBlueBits            = C.GLFW_ACCUM_BLUE_BITS
	AccumAlphaBits           = C.GLFW_ACCUM_ALPHA_BITS
	AuxBuffers               = C.GLFW_AUX_BUFFERS
	Samples                  = C.GLFW_SAMPLES
	RefreshRate              = C.GLFW_REFRESH_RATE
	Stereo                   = C.GLFW_STEREO
	SRGBCapable              = C.GLFW_SRGB_CAPABLE
	ClientApi                = C.GLFW_CLIENT_API
	ContextVersionMajor      = C.GLFW_CONTEXT_VERSION_MAJOR
	ContextVersionMinor      = C.GLFW_CONTEXT_VERSION_MINOR
	ContextRobustness        = C.GLFW_CONTEXT_ROBUSTNESS
	OpenGLForwardCompat      = C.GLFW_OPENGL_FORWARD_COMPAT
	OpenGLDebugContext       = C.GLFW_OPENGL_DEBUG_CONTEXT
	OpenGLProfile            = C.GLFW_OPENGL_PROFILE
	True                     = C.GL_TRUE
	False                    = C.GL_FALSE
	OpenGLApi                = C.GLFW_OPENGL_API
	OpenGLEsApi              = C.GLFW_OPENGL_ES_API
	NoRobustness             = C.GLFW_NO_ROBUSTNESS
	NoResetNotification      = C.GLFW_NO_RESET_NOTIFICATION
	LoseContextOnReset       = C.GLFW_LOSE_CONTEXT_ON_RESET
	OpenGLAnyProfile         = C.GLFW_OPENGL_ANY_PROFILE
	OpenGLCompatProfile      = C.GLFW_OPENGL_COMPAT_PROFILE
	OpenGLCoreProfile        = C.GLFW_OPENGL_CORE_PROFILE
)

// Opaque window object.
type Window struct {
	internalPtr *C.GLFWwindow

	positionCB        func(x, y int)
	sizeCB            func(width, height int)
	framebufferSizeCB func(width, height int)
	closeCB           func()
	refreshCB         func()
	focusCB           func(focused bool)
	iconifyCB         func(iconified bool)
}

var windows = map[*C.GLFWwindow]*Window{}

/*
This function creates a window and its associated context. Most of the options controlling how the window and its context should be created are specified through WindowHint.

Successful creation does not change which context is current. Before you can use the newly created context, you need to make it current using MakeContextCurrent.

Note that the created window and context may differ from what you requested, as not all parameters and hints are hard constraints. This includes the size of the window, especially for full screen windows. To retrieve the actual attributes of the created window and context, use queries like GetWindowAttrib and GetWindowSize.

To create a full screen window, you need to specify the monitor to use. If no monitor is specified, windowed mode will be used. Unless you have a way for the user to choose a specific monitor, it is recommended that you pick the primary monitor.

To create the window at a specific position, make it initially invisible using the Visible window hint, set its position and then show it.

If a full screen window is active, the screensaver is prohibited from starting.

Parameters
	width	The desired width, in screen coordinates, of the window. This must be greater than zero.
	height	The desired height, in screen coordinates, of the window. This must be greater than zero.
	title	The initial, UTF-8 encoded window title.
	monitor	The monitor to use for full screen mode, or nil to use windowed mode.
	share	The window whose context to share resources with, or nil to not share resources.
Returns
	A new window object.
Remarks
	Windows: Window creation will fail if the Microsoft GDI software OpenGL implementation is the only one available.
	Windows: If the executable has an icon resource named GLFW_ICON, it will be set as the icon for the window. If no such icon is present, the IDI_WINLOGO icon will be used instead.
	OS X: The GLFW window has no icon, as it is not a document window, but the dock icon will be the same as the application bundle's icon. Also, the first time a window is opened the menu bar is populated with common commands like Hide, Quit and About. The (minimal) about dialog uses information from the application's bundle. For more information on bundles, see the Bundle Programming Guide provided by Apple.
	X11: There is no mechanism for setting the window icon yet.
	The swap interval is not set during window creation, but is left at the default value for that platform. For more information, see SwapInterval.
Note
	This function may only be called from the main thread.
See Also
	Destroy
*/
func CreateWindow(width, height int, title string, monitor *Monitor, share *Window) (*Window, error) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	cMonPtr := (*C.GLFWmonitor)(nil)
	if monitor != nil {
		cMonPtr = monitor.internalPtr
	}

	cSharePtr := (*C.GLFWwindow)(nil)
	if share != nil {
		cSharePtr = share.internalPtr
	}

	cWinPtr := C.glfwCreateWindow(C.int(width), C.int(height), cTitle, cMonPtr, cSharePtr)
	if cWinPtr == nil {
		return nil, errors.New("glfw: Couldn't create window.")
	}

	win := &Window{internalPtr: cWinPtr}
	windows[cWinPtr] = win

	return win, nil
}

/*
This function returns the window whose context is current on the calling thread.

Returns
	The window whose context is current.
Remarks
	This function may be called from secondary threads.
See Also
	MakeContextCurrent
*/
func GetCurrentContext() (*Window, error) {
	cWinPtr := C.glfwGetCurrentContext()
	if win, ok := windows[cWinPtr]; ok {
		return win, nil
	}

	return nil, errors.New("glfw: No current context.")
}

/*
This function resets all window hints to their default values.

Note
	This function may only be called from the main thread.
See Also
	WindowHint
*/
func DefaultWindowHints() {
	C.glfwDefaultWindowHints()
}

/*
This function sets hints for the next call to CreateWindow. The hints, once set, retain their values until changed by a call to WindowHint or DefaultWindowHints, or until the library is terminated with Terminate.

Parameters
	hint	The window hint to set.
	flag	The new value of the window hint.
Note
	This function may only be called from the main thread.
See Also
	DefaultWindowHints
*/
func WindowHint(hint Hint, flag int) {
	C.glfwWindowHint(C.int(hint), C.int(flag))
}

/*
This function destroys the specified window and its context. On calling this function, no further callbacks will be called for that window.

Note
	This function may only be called from the main thread.
	This function may not be called from a callback.
	If the window's context is current on the main thread, it is detached before being destroyed.
Warning
	The window's context must not be current on any other thread.
See Also
	CreateWindow
*/
func (win *Window) Destroy() {
	delete(windows, win.internalPtr)
	C.glfwDestroyWindow(win.internalPtr)
}

/*
This function returns the value of the close flag of the specified window.

Returns
	The value of the close flag.
Remarks
	This function may be called from secondary threads.
*/
func (win *Window) ShouldClose() bool {
	if ret := C.glfwWindowShouldClose(win.internalPtr); ret == C.GL_TRUE {
		return true
	} else {
		return false
	}
}

/*
This function sets the value of the close flag of the specified window. This can be used to override the user's attempt to close the window, or to signal that it should be closed.

Parameters
	shouldClose	The new value.
Remarks
	This function may be called from secondary threads.
*/
func (win *Window) SetShouldClose(shouldClose bool) {
	val := C.GL_FALSE
	if shouldClose {
		val = C.GL_TRUE
	}

	C.glfwSetWindowShouldClose(win.internalPtr, C.int(val))
}

/*
This function sets the window title, encoded as UTF-8, of the specified window.

Parameters
	title	The UTF-8 encoded window title.
Note
	This function may only be called from the main thread.
*/
func (win *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	C.glfwSetWindowTitle(win.internalPtr, cTitle)
}

/*
This function retrieves the position, in screen coordinates, of the upper-left corner of the client area of the specified window.

Returns
	xpos, ypos
See Also
	SetPosition
*/
func (win *Window) GetPosition() (int, int) {
	var x, y C.int
	C.glfwGetWindowPos(win.internalPtr, &x, &y)
	return int(x), int(y)
}

/*
This function sets the position, in screen coordinates, of the upper-left corner of the client area of the window.

If the specified window is a full screen window, this function does nothing.

If you wish to set an initial window position you should create a hidden window (using WindowHint and Visible), set its position and then show it.

Parameters
	xpos	The x-coordinate of the upper-left corner of the client area.
	ypos	The y-coordinate of the upper-left corner of the client area.
Note
	It is very rarely a good idea to move an already visible window, as it will confuse and annoy the user.
	This function may only be called from the main thread.
	The window manager may put limits on what positions are allowed.
Bug
	X11: Some window managers ignore the set position of hidden (i.e. unmapped) windows, instead placing them where it thinks is appropriate once they are shown.
See Also
	GetPosition
*/
func (win *Window) SetPosition(x, y int) {
	C.glfwSetWindowPos(win.internalPtr, C.int(x), C.int(y))
}

/*
This function retrieves the size, in screen coordinates, of the client area of the specified window. If you wish to retrieve the size of the framebuffer in pixels, see GetFramebufferSize.

Returns
	width, height
See Also
	SetSize
*/
func (win *Window) GetSize() (int, int) {
	var width, height C.int
	C.glfwGetWindowSize(win.internalPtr, &width, &height)
	return int(width), int(height)
}

/*
This function sets the size, in screen coordinates, of the client area of the specified window.

For full screen windows, this function selects and switches to the resolution closest to the specified size, without affecting the window's context. As the context is unaffected, the bit depths of the framebuffer remain unchanged.

Parameters
	width	The desired width of the specified window.
	height	The desired height of the specified window.
Note
	This function may only be called from the main thread.
	The window manager may put limits on what window sizes are allowed.
See Also
	GetSize
*/
func (win *Window) SetSize(width, height int) {
	C.glfwSetWindowSize(win.internalPtr, C.int(width), C.int(height))
}

/*
This function retrieves the size, in pixels, of the framebuffer of the specified window. If you wish to retrieve the size of the window in screen coordinates, see GetWindowSize.

Returns
	width, height
See Also
	SetFramebufferSizeCallback
*/
func (win *Window) GetFramebufferSize() (int, int) {
	var width, height C.int
	C.glfwGetFramebufferSize(win.internalPtr, &width, &height)
	return int(width), int(height)
}

/*
This function iconifies/minimizes the specified window, if it was previously restored. If it is a full screen window, the original monitor resolution is restored until the window is restored. If the window is already iconified, this function does nothing.

Note
	This function may only be called from the main thread.
See Also
	Restore
*/
func (win *Window) Iconify() {
	C.glfwIconifyWindow(win.internalPtr)
}

/*
This function restores the specified window, if it was previously iconified/minimized. If it is a full screen window, the resolution chosen for the window is restored on the selected monitor. If the window is already restored, this function does nothing.

Note
	This function may only be called from the main thread.
See Also
	Iconify
*/
func (win *Window) Restore() {
	C.glfwRestoreWindow(win.internalPtr)
}

/*
This function makes the specified window visible, if it was previously hidden. If the window is already visible or is in full screen mode, this function does nothing.

Note
	This function may only be called from the main thread.
See Also
	Hide
*/
func (win *Window) Show() {
	C.glfwShowWindow(win.internalPtr)
}

/*
This function hides the specified window, if it was previously visible. If the window is already hidden or is in full screen mode, this function does nothing.

Note
	This function may only be called from the main thread.
See Also
	Show
*/
func (win *Window) Hide() {
	C.glfwHideWindow(win.internalPtr)
}

/*
This function returns the handle of the monitor that the specified window is in full screen on.

Returns
	The monitor.
*/
func (win *Window) GetMonitor() (*Monitor, error) {
	if mon := C.glfwGetWindowMonitor(win.internalPtr); mon != nil {
		return &Monitor{internalPtr: mon}, nil
	}
	return nil, errors.New("glfw: Can't get monitor, or window is windowed.")
}

/*
This function returns an attribute of the specified window. There are many attributes, some related to the window and others to its context.

Parameters
	attrib	The window attribute whose value to return.
Returns
	The value of the attribute, or zero if an error occurred.
*/
func (win *Window) GetAttrib(hint Hint) int {
	return int(C.glfwGetWindowAttrib(win.internalPtr, C.int(hint)))
}

/*
This function sets the user-defined pointer of the specified window. The current value is retained until the window is destroyed. The initial value is nil.

Parameters
	ptr	The new value.
See Also
	GetUserPointer
*/
func (win *Window) SetUserPointer(ptr unsafe.Pointer) {
	C.glfwSetWindowUserPointer(win.internalPtr, ptr)
}

/*
This function returns the current value of the user-defined pointer of the specified window. The initial value is nil.

See Also
	SetUserPointer
*/
func (win *Window) GetUserPointer() unsafe.Pointer {
	return unsafe.Pointer(C.glfwGetWindowUserPointer(win.internalPtr))
}

func (win *Window) MakeContextCurrent() {
	C.glfwMakeContextCurrent(win.internalPtr)
}

func (win *Window) SwapBuffers() {
	C.glfwSwapBuffers(win.internalPtr)
}

/*
This function sets the position callback of the specified window, which is called when the window is moved. The callback is provided with the screen position of the upper-left corner of the client area of the window.

Parameters
	callback	The new callback, or nil to remove the currently set callback.
*/
func (win *Window) SetPositionCallback(callback func(x, y int)) {
	if callback == nil {
		C.glfwSetWindowPosCallback(win.internalPtr, nil)
	} else {
		win.positionCB = callback
		C.goSetWindowPosCB(win.internalPtr)
	}
}

/*
This function sets the size callback of the specified window, which is called when the window is resized. The callback is provided with the size, in screen coordinates, of the client area of the window.

Parameters
	callback	The new callback, or nil to remove the currently set callback.
*/
func (win *Window) SetSizeCallback(callback func(width, height int)) {
	if callback == nil {
		C.glfwSetWindowSizeCallback(win.internalPtr, nil)
	} else {
		win.sizeCB = callback
		C.goSetWindowSizeCB(win.internalPtr)
	}
}

/*
This function sets the framebuffer resize callback of the specified window, which is called when the framebuffer of the specified window is resized.

Parameters
	callback	The new callback, or nil to remove the currently set callback.
*/
func (win *Window) SetFramebufferSizeCallback(callback func(width, height int)) {
	if callback == nil {
		C.glfwSetFramebufferSizeCallback(win.internalPtr, nil)
	} else {
		win.framebufferSizeCB = callback
		C.goSetFramebufferSizeCB(win.internalPtr)
	}
}

/*
This function sets the close callback of the specified window, which is called when the user attempts to close the window, for example by clicking the close widget in the title bar.

The close flag is set before this callback is called, but you can modify it at any time with SetShouldClose.

The close callback is not triggered by Destroy.

Parameters
	callback	The new callback, or nil to remove the currently set callback.
Remarks
	OS X: Selecting Quit from the application menu will trigger the close callback for all windows.
*/
func (win *Window) SetCloseCallback(callback func()) {
	if callback == nil {
		C.glfwSetWindowCloseCallback(win.internalPtr, nil)
	} else {
		win.closeCB = callback
		C.goSetWindowCloseCB(win.internalPtr)
	}
}

/*
This function sets the refresh callback of the specified window, which is called when the client area of the window needs to be redrawn, for example if the window has been exposed after having been covered by another window.

On compositing window systems such as Aero, Compiz or Aqua, where the window contents are saved off-screen, this callback may be called only very infrequently or never at all.

Parameters
	callback	The new callback, or nil to remove the currently set callback.
Note
	On compositing window systems such as Aero, Compiz or Aqua, where the window contents are saved off-screen, this callback may be called only very infrequently or never at all.
*/
func (win *Window) SetRefreshCallback(callback func()) {
	if callback == nil {
		C.glfwSetWindowRefreshCallback(win.internalPtr, nil)
	} else {
		win.refreshCB = callback
		C.goSetWindowRefreshCB(win.internalPtr)
	}
}

/*
This function sets the focus callback of the specified window, which is called when the window gains or loses focus.

After the focus callback is called for a window that lost focus, synthetic key and mouse button release events will be generated for all such that had been pressed. For more information, see SetKeyCallback and SetMouseButtonCallback.

Parameters
	callback	The new callback, or nil to remove the currently set callback.
*/
func (win *Window) SetFocusCallback(callback func(focused bool)) {
	if callback == nil {
		C.glfwSetWindowFocusCallback(win.internalPtr, nil)
	} else {
		win.focusCB = callback
		C.goSetWindowFocusCB(win.internalPtr)
	}
}

/*
This function sets the iconification callback of the specified window, which is called when the window is iconified or restored.

Parameters
	callback	The new callback, or nil to remove the currently set callback.
*/
func (win *Window) SetIconifyCallback(callback func(iconified bool)) {
	if callback == nil {
		C.glfwSetWindowIconifyCallback(win.internalPtr, nil)
	} else {
		win.iconifyCB = callback
		C.goSetWindowIconifyCB(win.internalPtr)
	}
}

//export goCallSetWindowPosCB
func goCallSetWindowPosCB(window unsafe.Pointer, x, y C.int) {
	if win, ok := windows[(*C.GLFWwindow)(window)]; ok {
		win.positionCB(int(x), int(y))
	}
}

//export goCallSetWindowSizeCB
func goCallSetWindowSizeCB(window unsafe.Pointer, width, height C.int) {
	if win, ok := windows[(*C.GLFWwindow)(window)]; ok {
		win.sizeCB(int(width), int(height))
	}
}

//export goCallSetFramebufferSizeCB
func goCallSetFramebufferSizeCB(window unsafe.Pointer, width, height C.int) {
	if win, ok := windows[(*C.GLFWwindow)(window)]; ok {
		win.framebufferSizeCB(int(width), int(height))
	}
}

//export goCallSetWindowCloseCB
func goCallSetWindowCloseCB(window unsafe.Pointer) {
	if win, ok := windows[(*C.GLFWwindow)(window)]; ok {
		win.closeCB()
	}
}

//export goCallSetWindowRefreshCB
func goCallSetWindowRefreshCB(window unsafe.Pointer) {
	if win, ok := windows[(*C.GLFWwindow)(window)]; ok {
		win.refreshCB()
	}
}

//export goCallSetWindowFocusCB
func goCallSetWindowFocusCB(window unsafe.Pointer, focused C.int) {
	if win, ok := windows[(*C.GLFWwindow)(window)]; ok {
		val := false
		if focused == C.GL_TRUE {
			val = true
		}
		win.focusCB(val)
	}
}

//export goCallSetWindowIconifyCB
func goCallSetWindowIconifyCB(window unsafe.Pointer, iconified C.int) {
	if win, ok := windows[(*C.GLFWwindow)(window)]; ok {
		val := false
		if iconified == C.GL_TRUE {
			val = true
		}
		win.iconifyCB(val)
	}
}
