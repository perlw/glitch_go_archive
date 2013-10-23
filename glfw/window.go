package glfw

//#include <stdlib.h>
//#include <GLFW/glfw3.h>
//void goSetWindowPosCB(GLFWwindow *win);
//void goSetWindowSizeCB(GLFWwindow *win);
//void goSetFramebufferSizeCB(GLFWwindow *win);
//void goSetWindowCloseCB(GLFWwindow *win);
//void goSetWindowRefreshCB(GLFWwindow *win);
//void goSetWindowFocusCB(GLFWwindow *win);
//void goSetWindowIconifyCB(GLFWwindow *win);
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

func GetCurrentContext() (*Window, error) {
	cWinPtr := C.glfwGetCurrentContext()
	if win, ok := windows[cWinPtr]; ok {
		return win, nil
	}

	return nil, errors.New("glfw: No current context.")
}

func DefaultWindowHints() {
	C.glfwDefaultWindowHints()
}

func WindowHint(hint Hint, flag int) {
	C.glfwWindowHint(C.int(hint), C.int(flag))
}

func (win *Window) Destroy() {
	delete(windows, win.internalPtr)
	C.glfwDestroyWindow(win.internalPtr)
}

func (win *Window) ShouldClose() bool {
	if ret := C.glfwWindowShouldClose(win.internalPtr); ret == C.GL_TRUE {
		return true
	} else {
		return false
	}
}

func (win *Window) SetShouldClose(shouldClose bool) {
	val := C.GL_FALSE
	if shouldClose {
		val = C.GL_TRUE
	}

	C.glfwSetWindowShouldClose(win.internalPtr, C.int(val))
}

func (win *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	C.glfwSetWindowTitle(win.internalPtr, cTitle)
}

func (win *Window) GetPosition() (int, int) {
	var x, y C.int
	C.glfwGetWindowPos(win.internalPtr, &x, &y)
	return int(x), int(y)
}

func (win *Window) SetPosition(x, y int) {
	C.glfwSetWindowPos(win.internalPtr, C.int(x), C.int(y))
}

func (win *Window) GetSize() (int, int) {
	var width, height C.int
	C.glfwGetWindowSize(win.internalPtr, &width, &height)
	return int(width), int(height)
}

func (win *Window) SetSize(width, height int) {
	C.glfwSetWindowSize(win.internalPtr, C.int(width), C.int(height))
}

func (win *Window) GetFramebufferSize() (int, int) {
	var width, height C.int
	C.glfwGetFramebufferSize(win.internalPtr, &width, &height)
	return int(width), int(height)
}

func (win *Window) Iconify() {
	C.glfwIconifyWindow(win.internalPtr)
}

func (win *Window) Restore() {
	C.glfwRestoreWindow(win.internalPtr)
}

func (win *Window) Show() {
	C.glfwShowWindow(win.internalPtr)
}

func (win *Window) Hide() {
	C.glfwHideWindow(win.internalPtr)
}

func (win *Window) GetMonitor() (*Monitor, error) {
	if mon := C.glfwGetWindowMonitor(win.internalPtr); mon != nil {
		return &Monitor{internalPtr: mon}, nil
	}
	return nil, errors.New("glfw: Can't get monitor.")
}

func (win *Window) GetAttrib(hint Hint) int {
	return int(C.glfwGetWindowAttrib(win.internalPtr, C.int(hint)))
}

func (win *Window) SetUserPointer(ptr unsafe.Pointer) {
	C.glfwSetWindowUserPointer(win.internalPtr, ptr)
}

func (win *Window) GetUserPointer() unsafe.Pointer {
	return unsafe.Pointer(C.glfwGetWindowUserPointer(win.internalPtr))
}

func (win *Window) MakeContextCurrent() {
	C.glfwMakeContextCurrent(win.internalPtr)
}

func (win *Window) SwapBuffers() {
	C.glfwSwapBuffers(win.internalPtr)
}

func (win *Window) SetPositionCallback(callback func(x, y int)) {
	if callback == nil {
		C.glfwSetWindowPosCallback(win.internalPtr, nil)
	} else {
		win.positionCB = callback
		C.goSetWindowPosCB(win.internalPtr)
	}
}

func (win *Window) SetSizeCallback(callback func(width, height int)) {
	if callback == nil {
		C.glfwSetWindowSizeCallback(win.internalPtr, nil)
	} else {
		win.sizeCB = callback
		C.goSetWindowSizeCB(win.internalPtr)
	}
}

func (win *Window) SetFramebufferSizeCallback(callback func(width, height int)) {
	if callback == nil {
		C.glfwSetFramebufferSizeCallback(win.internalPtr, nil)
	} else {
		win.framebufferSizeCB = callback
		C.goSetFramebufferSizeCB(win.internalPtr)
	}
}

func (win *Window) SetCloseCallback(callback func()) {
	if callback == nil {
		C.glfwSetWindowCloseCallback(win.internalPtr, nil)
	} else {
		win.closeCB = callback
		C.goSetWindowCloseCB(win.internalPtr)
	}
}

func (win *Window) SetRefreshCallback(callback func()) {
	if callback == nil {
		C.glfwSetWindowRefreshCallback(win.internalPtr, nil)
	} else {
		win.refreshCB = callback
		C.goSetWindowRefreshCB(win.internalPtr)
	}
}

func (win *Window) SetFocusCallback(callback func(focused bool)) {
	if callback == nil {
		C.glfwSetWindowFocusCallback(win.internalPtr, nil)
	} else {
		win.focusCB = callback
		C.goSetWindowFocusCB(win.internalPtr)
	}
}

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
