package glfw

//#include <stdlib.h>
//#include <GLFW/glfw3.h>
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
}

var windows = map[*C.GLFWwindow]*Window{}

func CreateWindow(width, height int, title string, monitor *Monitor, share *Window) (*Window, error) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	cWinPtr := C.glfwCreateWindow(C.int(width), C.int(height), cTitle, nil, nil)
	if cWinPtr == nil {
		return nil, errors.New("glfw: Couldn't create window.")
	}

	win := &Window{internalPtr: cWinPtr}
	windows[cWinPtr] = win

	return win, nil
}

func GetCurrentContext() (*Window, error) {
	cWinPtr := C.glfwGetCurrentContext()
	if item, ok := windows[cWinPtr]; ok {
		return item, nil
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

func (win *Window) MakeContextCurrent() {
	C.glfwMakeContextCurrent(win.internalPtr)
}

func (win *Window) SwapBuffers() {
	C.glfwSwapBuffers(win.internalPtr)
}
