package gl

// #cgo linux   LDFLAGS: -lGLEW -lGL
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #include "common.h"
import "C"

import (
	"errors"
	"unsafe"
)

func Init() {
	C.glewExperimental = 1
	C.glewInit()
}

func Accum(constant GLConstant, value float32) {
	C.glAccum(C.GLenum(constant), C.GLfloat(value))
}

func AlphaFunc(constant GLConstant, value float32) {
	C.glAlphaFunc(C.GLenum(constant), C.GLclampf(value))
}

func AreTexturesResident(textures []uint32) ([]bool, error) {
	size := len(textures)
	if size == 0 {
		return nil, errors.New("gl: Empty list")
	}
	residences := make([]bool, size)

	if ret := C.glAreTexturesResident(C.GLsizei(size), (*C.GLuint)(unsafe.Pointer(&textures[0])), (*C.GLboolean)(unsafe.Pointer(&residences[0]))); ret == C.GL_TRUE {
		for t := range residences {
			residences[t] = true
		}
	}
	return residences, nil
}

func ArrayElement(index int) {
	C.glArrayElement(C.GLint(index))
}

func Clear(clearBits GLConstant) {
	C.glClear(C.GLbitfield(clearBits))
}

func ClearColor(r, g, b, a float32) {
	C.glClearColor(C.GLclampf(r), C.GLclampf(g), C.GLclampf(b), C.GLclampf(a))
}

func ClearDepth(depth float32) {
	C.glClearDepth(C.GLclampd(depth))
}

func DepthFunc(constant GLConstant) {
	C.glDepthFunc(C.GLenum(constant))
}

func Enable(constant GLConstant) {
	C.glEnable(C.GLenum(constant))
}

func Disable(constant GLConstant) {
	C.glDisable(C.GLenum(constant))
}

func Viewport(x, y, width, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
