package glfw

//#include <GLFW/glfw3.h>
//void goSetErrorCB();
import "C"

type ErrorCode int

const (
	NotInitializedConst     ErrorCode = C.GLFW_NOT_INITIALIZED     // GLFW has not been initialized.
	NoCurrentContextConst             = C.GLFW_NO_CURRENT_CONTEXT  // No context is current for this thread.
	InvalidEnumConst                  = C.GLFW_INVALID_ENUM        // One of the enum parameters for the function was given an invalid enum.
	InvalidValueConst                 = C.GLFW_INVALID_VALUE       // One of the parameters for the function was given an invalid value.
	OutOfMemoryConst                  = C.GLFW_OUT_OF_MEMORY       // A memory allocation failed.
	ApiUnavailableConst               = C.GLFW_API_UNAVAILABLE     // GLFW could not find support for the requested client API on the system.
	VersionUnavailableConst           = C.GLFW_VERSION_UNAVAILABLE // The requested client API version is not available.
	PlatformErrorConst                = C.GLFW_PLATFORM_ERROR      // A platform-specific error occurred that does not match any of the more specific categories.
	FormatUnavailableConst            = C.GLFW_FORMAT_UNAVAILABLE  // The clipboard did not contain data in the requested format.
)

var errorCB func(code ErrorCode, description string)

//export goCallErrorCB
func goCallErrorCB(code C.int, description *C.char) {
	errorCB(ErrorCode(code), C.GoString(description))
}

/*
This function sets the error callback, which is called with an error code and a human-readable description each time a GLFW error occurs.

Parameters
	callback	The new callback, or nil to remove the currently set callback.
Returns
The previously set callback, or NULL if no callback was set or an error occurred.
Remarks
	This function may be called before glfwInit.
*/
func SetErrorCallback(callback func(code ErrorCode, description string)) {
	if callback == nil {
		C.glfwSetErrorCallback(nil)
	} else {
		errorCB = callback
		C.goSetErrorCB()
	}
}
