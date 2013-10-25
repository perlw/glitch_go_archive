package glfw

//#include <GLFW/glfw3.h>
//GLFWmonitor* monitorAtIndex(GLFWmonitor** monitors, int index);
//GLFWvidmode vidModeAtIndex(GLFWvidmode* vidModes, int index);
//void goSetMonitorCB();
import "C"

import (
	"errors"
	"unsafe"
)

type MonitorEvent int

const (
	Connected    MonitorEvent = C.GLFW_CONNECTED
	Disconnected              = C.GLFW_DISCONNECTED
)

type VideoMode struct {
	Width       int
	Height      int
	RedBits     int
	GreenBits   int
	BlueBits    int
	RefreshRate int
}

type GammaRamp struct {
	Red   []uint16
	Green []uint16
	Blue  []uint16
}

type Monitor struct {
	internalPtr *C.GLFWmonitor
}

var monitorCB func(monitor *Monitor, event MonitorEvent)

func GetMonitors() ([]*Monitor, error) {
	var count C.int
	if cMonitors := C.glfwGetMonitors(&count); cMonitors != nil {
		monitors := make([]*Monitor, int(count))
		for t := 0; t < int(count); t++ {
			monitors[t] = &Monitor{internalPtr: C.monitorAtIndex(cMonitors, C.int(t))}
		}
		return monitors, nil
	}
	return nil, errors.New("glfw: Could not get monitors.")
}

func GetPrimaryMonitor() (*Monitor, error) {
	if cMonitor := C.glfwGetPrimaryMonitor(); cMonitor != nil {
		return &Monitor{internalPtr: cMonitor}, nil
	}
	return nil, errors.New("glfw: Could not get primary monitor.")
}

func (mon *Monitor) GetPosition() (int, int) {
	var x, y C.int
	C.glfwGetMonitorPos(mon.internalPtr, &x, &y)
	return int(x), int(y)
}

func (mon *Monitor) GetPhysicalSize() (int, int) {
	var width, height C.int
	C.glfwGetMonitorPhysicalSize(mon.internalPtr, &width, &height)
	return int(width), int(height)
}

func (mon *Monitor) GetName() string {
	return C.GoString(C.glfwGetMonitorName(mon.internalPtr))
}

//export goCallMonitorCB
func goCallMonitorCB(monitor unsafe.Pointer, event C.int) {
	monitorCB(&Monitor{internalPtr: (*C.GLFWmonitor)(monitor)}, MonitorEvent(event))
}

func SetMonitorCallback(callback func(monitor *Monitor, event MonitorEvent)) {
	if callback == nil {
		C.glfwSetMonitorCallback(nil)
	} else {
		monitorCB = callback
		C.goSetMonitorCB()
	}
}

func cVidModeToVideoMode(cVidMode *C.GLFWvidmode) *VideoMode {
	return &VideoMode{
		Width:       int(cVidMode.width),
		Height:      int(cVidMode.height),
		RedBits:     int(cVidMode.redBits),
		GreenBits:   int(cVidMode.greenBits),
		BlueBits:    int(cVidMode.blueBits),
		RefreshRate: int(cVidMode.refreshRate),
	}
}

func (mon *Monitor) GetVideoMode() (*VideoMode, error) {
	if cVidMode := C.glfwGetVideoMode(mon.internalPtr); cVidMode != nil {
		return cVidModeToVideoMode(cVidMode), nil
	}
	return nil, errors.New("glfw: Could not get video mode.")
}

func (mon *Monitor) GetVideoModes() ([]*VideoMode, error) {
	var count C.int
	if cVidModes := C.glfwGetVideoModes(mon.internalPtr, &count); cVidModes != nil {
		vidModes := make([]*VideoMode, int(count))
		for t := 0; t < int(count); t++ {
			cVidMode := C.vidModeAtIndex(cVidModes, C.int(t))
			vidModes[t] = cVidModeToVideoMode(&cVidMode)
		}
		return vidModes, nil

	}
	return nil, errors.New("glfw: Could not get video modes.")
}

func (mon *Monitor) SetGamma(gamma float32) {

}

func (mon *Monitor) GetGammaRamp() (*GammaRamp, error) {
	return nil, nil
}

func (mon *Monitor) SetGammaRamp(ramp GammaRamp) {

}
