package gl

// #cgo linux   LDFLAGS: -lGLEW -lGL
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #include "common.h"
import "C"

func Init() {
	C.glewExperimental = 1
	C.glewInit()
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
