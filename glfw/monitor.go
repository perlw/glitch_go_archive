package glfw

//#include <GLFW/glfw3.h>
import "C"

type Monitor struct {
	internalPtr *C.GLFWmonitor
}
