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
		return nil, errors.New("gl: Empty textures")
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

func Begin(constant GLConstant) {
	C.glBegin(C.GLenum(constant))
}

func BindTexture(target GLConstant, texture uint32) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

func Bitmap(width, height int, xorig, yorig, xmove, ymove float32, bitmap []byte) error {
	size := len(bitmap)
	if size == 0 {
		return errors.New("gl: Empty bitmap")
	}

	C.glBitmap(C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(unsafe.Pointer(&bitmap[0])))

	return nil
}

func BlendColorEXT(red, green, blue, alpha float32) {
	C.glBlendColorEXT(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

func BlendFunc(sfactor, dfactor GLConstant) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
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

func End() {
	C.glEnd()
}

func Disable(constant GLConstant) {
	C.glDisable(C.GLenum(constant))
}

func Viewport(x, y, width, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
