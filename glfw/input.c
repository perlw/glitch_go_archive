#include "_cgo_export.h"

void setCharCallback(GLFWwindow* win, unsigned int rune) {
	goCallSetCharCB(win, rune);
}

void setCursorEnterCallback(GLFWwindow* win, int flag) {
	goCallSetCursorEnterCB(win, flag);
}

void setCursorPosCallback(GLFWwindow* win, double xpos, double ypos) {
	goCallSetCursorPosCB(win, xpos, ypos);
}

void setKeyCallback(GLFWwindow* win, int key, int scancode, int action, int mods) {
	goCallSetKeyCB(win, key, scancode, action, mods);
}

void setMouseButtonCallback(GLFWwindow* win, int button, int action, int mods) {
	goCallSetMouseButtonCB(win, button, action, mods);
}

void setScrollCallback(GLFWwindow* win, double xoff, double yoff) {
	goCallSetScrollCB(win, xoff, yoff);
}

void goSetCharCB(GLFWwindow* win) {
	glfwSetCharCallback(win, setCharCallback);
}

void goSetCursorEnterCB(GLFWwindow* win) {
	glfwSetCursorEnterCallback(win, setCursorEnterCallback);
}

void goSetCursorPosCB(GLFWwindow* win) {
	glfwSetCursorPosCallback(win, setCursorPosCallback);
}

void goSetKeyCB(GLFWwindow* win) {
	glfwSetKeyCallback(win, setKeyCallback);
}

void goSetMouseButtonCB(GLFWwindow* win) {
	glfwSetMouseButtonCallback(win, setMouseButtonCallback);
}

void goSetScrollCB(GLFWwindow* win) {
	glfwSetScrollCallback(win, setScrollCallback);
}
