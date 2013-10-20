package glitch

// #cgo darwin  LDFLAGS: -framework OpenGL -lGLEW
// #cgo linux   LDFLAGS: -lGLEW -lGL
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #include <stdlib.h>
// #include <GL/glew.h>
import "C"

func Init() {
	C.glewExperimental = 1
	C.glewInit()
}
