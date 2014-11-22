package glfw

//#include <stdlib.h>
//#include <GLFW/glfw3.h>
import "C"

import "unsafe"

/*
This function returns the contents of the system clipboard, if it contains or is convertible to a UTF-8 encoded string.

Returns
	The contents of the clipboard as a UTF-8 encoded string.
Note
	This function may only be called from the main thread.
*/
func (win *Window) GetClipboardString() string {
	return C.GoString(C.glfwGetClipboardString(win.internalPtr))
}

/*
This function sets the system clipboard to the specified, UTF-8 encoded string. The string is copied before returning, so you don't have to retain it afterwards.

Parameters
	str	A UTF-8 encoded string.
Note
	This function may only be called from the main thread.
See Also
	GetClipboardString
*/
func (win *Window) SetClipboardString(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.glfwSetClipboardString(win.internalPtr, cStr)
}
