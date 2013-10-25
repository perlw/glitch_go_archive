#include "_cgo_export.h"

GLFWmonitor* monitorAtIndex(GLFWmonitor** monitors, int index) {
    return monitors[index];
}

GLFWvidmode vidModeAtIndex(GLFWvidmode* vidModes, int index) {
    return vidModes[index];
}

void monitorCallback(GLFWmonitor* monitor, int event) {
    goCallMonitorCB(monitor, event);
}

void goSetMonitorCB() {
    glfwSetMonitorCallback(monitorCallback);
}
