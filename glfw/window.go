package glfw

//#include <stdlib.h>
//#include <GLFW/glfw3.h>
import "C"

import (
	"errors"
	"unsafe"
)

type Window struct {
	internalPtr *C.GLFWwindow
}

func CreateWindow(width, height int, title string, monitor *Monitor, share *Window) (*Window, error) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	win := C.glfwCreateWindow(C.int(width), C.int(height), cTitle, nil, nil)
	if win == nil {
		return nil, errors.New("Couldn't create window.")
	}

	return &Window{internalPtr: win}, nil
}

func (win *Window) Destroy() {
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
