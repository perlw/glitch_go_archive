// TODO
package glfw

//#cgo windows LDFLAGS: -lglfw3dll -lopengl32 -lgdi32
//#cgo linux LDFLAGS: -lglfw
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GLFW/glfw3.h>
//float getAxesAtIndex(float* axes, int index);
import "C"

import "errors"

type Joystick int

const (
	Joystick1    Joystick = C.GLFW_JOYSTICK_1
	Joystick2             = C.GLFW_JOYSTICK_2
	Joystick3             = C.GLFW_JOYSTICK_3
	Joystick4             = C.GLFW_JOYSTICK_4
	Joystick5             = C.GLFW_JOYSTICK_5
	Joystick6             = C.GLFW_JOYSTICK_6
	Joystick7             = C.GLFW_JOYSTICK_7
	Joystick8             = C.GLFW_JOYSTICK_8
	Joystick9             = C.GLFW_JOYSTICK_9
	Joystick10            = C.GLFW_JOYSTICK_10
	Joystick11            = C.GLFW_JOYSTICK_11
	Joystick12            = C.GLFW_JOYSTICK_12
	Joystick13            = C.GLFW_JOYSTICK_13
	Joystick14            = C.GLFW_JOYSTICK_14
	Joystick15            = C.GLFW_JOYSTICK_15
	Joystick16            = C.GLFW_JOYSTICK_16
	JoystickLast          = C.GLFW_JOYSTICK_LAST
)

/*
This function initializes the GLFW library. Before most GLFW functions can be used, GLFW must be initialized, and before a program terminates GLFW should be terminated in order to free any resources allocated during or after initialization.

If this function fails, it calls Terminate before returning. If it succeeds, you should call Terminate before the program exits.

Additional calls to this function after successful initialization but before termination will succeed but will do nothing.

Returns
	non-nil if error
Note
	This function may only be called from the main thread.
	This function may take several seconds to complete on some systems, while on other systems it may take only a fraction of a second to complete.
	OS X: This function will change the current directory of the application to the Contents/Resources subdirectory of the application's bundle, if present.
See Also
	Terminate
*/
func Init() error {
	if ret := C.glfwInit(); ret != C.GL_TRUE {
		return errors.New("glitch: Could not initialize glfw.")
	}

	return nil
}

/*
This function destroys all remaining windows, frees any allocated resources and sets the library to an uninitialized state. Once this is called, you must again call Init successfully before you will be able to use most GLFW functions.

If GLFW has been successfully initialized, this function should be called before the program exits. If initialization fails, there is no need to call this function, as it is called by Init before it returns failure.

Remarks
	This function may be called before Init.
Note
	This function may only be called from the main thread.
Warning
	No window's context may be current on another thread when this function is called.
See Also
	Init
*/
func Terminate() {
	C.glfwTerminate()
}

/*
This function retrieves the major, minor and revision numbers of the GLFW library. It is intended for when you are using GLFW as a shared library and want to ensure that you are using the minimum required version.

Returns
	major, minor, revision
Remarks
	This function may be called before glfwInit.
	This function may be called from any thread.
See Also
	GetVersionString
*/
func GetVersion() (int, int, int) {
	var major, minor, revision C.int
	C.glfwGetVersion(&major, &minor, &revision)
	return int(major), int(minor), int(revision)
}

/*
This function returns a static string generated at compile-time according to which configuration macros were defined. This is intended for use when submitting bug reports, to allow developers to see which code paths are enabled in a binary.

The format of the string is as follows:
	The version of GLFW
	The name of the window system API
	The name of the context creation API
	Any additional options or APIs

For example, when compiling GLFW 3.0 with MinGW using the Win32 and WGL back ends, the version string may look something like this:
	3.0.0 Win32 WGL MinGW

Returns
	The GLFW version string.
Remarks
	This function may be called before glfwInit.
	This function may be called from any thread.
See Also
	GetVersion
*/
func GetVersionString() string {
	return C.GoString(C.glfwGetVersionString())
}

/*
This function processes only those events that have already been received and then returns immediately. Processing events will cause the window and input callbacks associated with those events to be called.

This function is not required for joystick input to work.

Note
	This function may only be called from the main thread.
	This function may not be called from a callback.
	On some platforms, certain callbacks may be called outside of a call to one of the event processing functions.
See Also
	WaitEvents
*/
func PollEvents() {
	C.glfwPollEvents()
}

/*
This function puts the calling thread to sleep until at least one event has been received. Once one or more events have been received, it behaves as if glfwPollEvents was called, i.e. the events are processed and the function then returns immediately. Processing events will cause the window and input callbacks associated with those events to be called.

Since not all events are associated with callbacks, this function may return without a callback having been called even if you are monitoring all callbacks.

This function is not required for joystick input to work.

Note
	This function may only be called from the main thread.
	This function may not be called from a callback.
	On some platforms, certain callbacks may be called outside of a call to one of the event processing functions.
See Also
	PollEvents
*/
func WaitEvents() {
	C.glfwWaitEvents()
}

/*
This function sets the swap interval for the current context, i.e. the number of screen updates to wait before swapping the buffers of a window and returning from glfwSwapBuffers. This is sometimes called 'vertical synchronization', 'vertical retrace synchronization' or 'vsync'.

Contexts that support either of the WGL_EXT_swap_control_tear and GLX_EXT_swap_control_tear extensions also accept negative swap intervals, which allow the driver to swap even if a frame arrives a little bit late. You can check for the presence of these extensions using glfwExtensionSupported. For more information about swap tearing, see the extension specifications.

Parameters
	interval	The minimum number of screen updates to wait for until the buffers are swapped by glfwSwapBuffers.
Remarks
	This function may be called from secondary threads.
Note
	This function is not called during window creation, leaving the swap interval set to whatever is the default on that platform. This is done because some swap interval extensions used by GLFW do not allow the swap interval to be reset to zero once it has been set to a non-zero value.
	Some GPU drivers do not honor the requested swap interval, either because of user settings that override the request or due to bugs in the driver.
See Also
	SwapBuffers
*/
func SwapInterval(interval int) {
	C.glfwSwapInterval(C.int(interval))
}

/*
This function returns the value of the GLFW timer. Unless the timer has been set using SetTime, the timer measures time elapsed since GLFW was initialized.

Returns
	The current value, in seconds, or zero if an error occurred.
Remarks
	This function may be called from secondary threads.
*/
func GetTime() float64 {
	return float64(C.glfwGetTime())
}

/*
This function sets the value of the GLFW timer. It then continues to count up from that value.

Parameters
	time	The new value, in seconds.
*/
func SetTime(time float64) {
	C.glfwSetTime(C.double(time))
}

/*
This function returns the values of all axes of the specified joystick.

Parameters
	joy	The joystick to query.
Returns
	An array of axis values.
*/
func GetJoystickAxes(joy Joystick) ([]float32, error) {
	var count C.int

	cAxes := C.glfwGetJoystickAxes(C.int(joy), &count)
	if cAxes == nil {
		return nil, errors.New("glfw: No such joystick.")
	}

	axes := make([]float32, int(count))
	for t := 0; t < int(count); t++ {
		axes[t] = float32(C.getAxesAtIndex(cAxes, C.int(t)))
	}

	return axes, nil
}
