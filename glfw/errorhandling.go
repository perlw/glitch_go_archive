package glfw

//#include <GLFW/glfw3.h>
//void goSetErrorCB();
import "C"

type ErrorCode int

const (
	NotInitialized     ErrorCode = C.GLFW_NOT_INITIALIZED
	NoCurrentContext             = C.GLFW_NO_CURRENT_CONTEXT
	InvalidEnum                  = C.GLFW_INVALID_ENUM
	InvalidValue                 = C.GLFW_INVALID_VALUE
	OutOfMemory                  = C.GLFW_OUT_OF_MEMORY
	ApiUnavailable               = C.GLFW_API_UNAVAILABLE
	VersionUnavailable           = C.GLFW_VERSION_UNAVAILABLE
	PlatformError                = C.GLFW_PLATFORM_ERROR
	FormatUnavailable            = C.GLFW_FORMAT_UNAVAILABLE
)

var errorCB func(code ErrorCode, description string)

//export goCallErrorCB
func goCallErrorCB(code C.int, description *C.char) {
	errorCB(ErrorCode(code), C.GoString(description))
}

func SetErrorCallback(callback func(code ErrorCode, description string)) {
	if callback == nil {
		C.glfwSetErrorCallback(nil)
	} else {
		errorCB = callback
		C.goSetErrorCB()
	}
}
