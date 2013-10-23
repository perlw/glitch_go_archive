#include "_cgo_export.h"

void setWindowPosCallback(GLFWwindow *win, int x, int y) {
	goCallSetWindowPosCB(win, x, y);
}

void setWindowSizeCallback(GLFWwindow *win, int width, int height) {
	goCallSetWindowSizeCB(win, width, height);
}

void setFramebufferSizeCallback(GLFWwindow *win, int width, int height) {
	goCallSetFramebufferSizeCB(win, width, height);
}

void setWindowCloseCallback(GLFWwindow *win) {
	goCallSetWindowCloseCB(win);
}

void setWindowRefreshCallback(GLFWwindow *win) {
	goCallSetWindowRefreshCB(win);
}

void setWindowFocusCallback(GLFWwindow *win, int focused) {
	goCallSetWindowFocusCB(win, focused);
}

void setWindowIconifyCallback(GLFWwindow *win, int iconified) {
	goCallSetWindowIconifyCB(win, iconified);
}

void goSetWindowPosCB(GLFWwindow *win) {
	glfwSetWindowPosCallback(win, setWindowPosCallback);
}

void goSetWindowSizeCB(GLFWwindow *win) {
	glfwSetWindowSizeCallback(win, setWindowSizeCallback);
}

void goSetFramebufferSizeCB(GLFWwindow *win) {
	glfwSetFramebufferSizeCallback(win, setFramebufferSizeCallback);
}

void goSetWindowCloseCB(GLFWwindow *win) {
	glfwSetWindowCloseCallback(win, setWindowCloseCallback);
}

void goSetWindowRefreshCB(GLFWwindow *win) {
	glfwSetWindowRefreshCallback(win, setWindowRefreshCallback);
}

void goSetWindowFocusCB(GLFWwindow *win) {
	glfwSetWindowFocusCallback(win, setWindowFocusCallback);
}

void goSetWindowIconifyCB(GLFWwindow *win) {
	glfwSetWindowIconifyCallback(win, setWindowIconifyCallback);
}
