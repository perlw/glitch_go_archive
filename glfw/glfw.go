package glfw

//#cgo windows LDFLAGS: -lglfw3dll -lopengl32 -lgdi32
//#cgo linux LDFLAGS: -lglfw
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GLFW/glfw3.h>
import "C"

import "errors"

func Init() error {
	if ret := C.glfwInit(); ret != C.GL_TRUE {
		return errors.New("Could not initialize glfw.")
	}

	return nil
}

func Terminate() {
	C.glfwTerminate()
}

func GetVersion() (int, int, int) {
	var major, minor, revision C.int
	C.glfwGetVersion(&major, &minor, &revision)
	return int(major), int(minor), int(revision)
}

func GetVersionString() string {
	return C.GoString(C.glfwGetVersionString())
}
