#include "_cgo_export.h"

void errorCallback(int code, const char *description) {
	goCallErrorCB(code, (char*)description);
}

void goSetErrorCB() {
	glfwSetErrorCallback(errorCallback);
}
