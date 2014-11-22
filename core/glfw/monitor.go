package glfw

//#include <GLFW/glfw3.h>
//GLFWmonitor* monitorAtIndex(GLFWmonitor** monitors, int index);
//GLFWvidmode vidModeAtIndex(GLFWvidmode* vidModes, int index);
//void setGammaAtIndex(unsigned short* color, int index, unsigned short value);
//unsigned int getGammaAtIndex(unsigned short* color, int index);
//void goSetMonitorCB();
import "C"

import (
	"errors"
	"unsafe"
)

type MonitorEvent int

const (
	ConnectedConst    MonitorEvent = C.GLFW_CONNECTED
	DisconnectedConst              = C.GLFW_DISCONNECTED
)

// This describes a single video mode.
type VideoMode struct {
	Width       int
	Height      int
	RedBits     int
	GreenBits   int
	BlueBits    int
	RefreshRate int
}

/*
This describes the gamma ramp for a monitor.

See Also
	GetGammaRamp SetGammaRamp
*/
type GammaRamp struct {
	Red   []uint16
	Green []uint16
	Blue  []uint16
}

// Opaque monitor object.
type Monitor struct {
	internalPtr *C.GLFWmonitor
}

var monitorCB func(monitor *Monitor, event MonitorEvent)

/*
This function returns an array of handles for all currently connected monitors.

Returns
	A slice of monitors
Note
	The returned slice is valid only until the monitor configuration changes. See SetMonitorCallback to receive notifications of configuration changes.
See Also
	GetPrimaryMonitor
*/
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

/*
This function returns the primary monitor. This is usually the monitor where elements like the Windows task bar or the OS X menu bar is located.

Returns
	The primary monitor.
See Also
	GetMonitors
*/
func GetPrimaryMonitor() (*Monitor, error) {
	if cMonitor := C.glfwGetPrimaryMonitor(); cMonitor != nil {
		return &Monitor{internalPtr: cMonitor}, nil
	}
	return nil, errors.New("glfw: Could not get primary monitor.")
}

/*
This function returns the position, in screen coordinates, of the upper-left corner of the specified monitor.

Returns
	xpos, ypos
*/
func (mon *Monitor) GetPosition() (int, int) {
	var x, y C.int
	C.glfwGetMonitorPos(mon.internalPtr, &x, &y)
	return int(x), int(y)
}

/*
This function returns the size, in millimetres, of the display area of the specified monitor.

Returns
	width, height
Note
	Some operating systems do not provide accurate information, either because the monitor's EDID data is incorrect, or because the driver does not report it accurately.
*/
func (mon *Monitor) GetPhysicalSize() (int, int) {
	var width, height C.int
	C.glfwGetMonitorPhysicalSize(mon.internalPtr, &width, &height)
	return int(width), int(height)
}

/*
This function returns a human-readable name, encoded as UTF-8, of the specified monitor.

Returns
	The UTF-8 encoded name of the monitor.
*/
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

/*
This function returns the current video mode of the specified monitor. If you are using a full screen window, the return value will therefore depend on whether it is focused.

Returns
	The current mode of the monitor.
See Also
	GetVideoModes
*/
func (mon *Monitor) GetVideoMode() (*VideoMode, error) {
	if cVidMode := C.glfwGetVideoMode(mon.internalPtr); cVidMode != nil {
		return cVidModeToVideoMode(cVidMode), nil
	}
	return nil, errors.New("glfw: Could not get video mode.")
}

/*
This function returns an array of all video modes supported by the specified monitor. The returned array is sorted in ascending order, first by color bit depth (the sum of all channel depths) and then by resolution area (the product of width and height).

Returns
	A slice of video modes.
See Also
	GetVideoMode
*/
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

/*
This function generates a 256-element gamma ramp from the specified exponent and then calls SetGammaRamp with it.

Parameters
	gamma	The desired exponent.
*/
func (mon *Monitor) SetGamma(gamma float32) {
	C.glfwSetGamma(mon.internalPtr, C.float(gamma))
}

/*
This function retrieves the current gamma ramp of the specified monitor.

Returns
	The current gamma ramp
*/
func (mon *Monitor) GetGammaRamp() (*GammaRamp, error) {
	var ramp GammaRamp

	if cGammaRamp := C.glfwGetGammaRamp(mon.internalPtr); cGammaRamp != nil {
		length := int(cGammaRamp.size)
		ramp.Red = make([]uint16, length)
		ramp.Green = make([]uint16, length)
		ramp.Blue = make([]uint16, length)

		for t := 0; t < length; t++ {
			ramp.Red[t] = uint16(C.getGammaAtIndex(cGammaRamp.red, C.int(t)))
			ramp.Green[t] = uint16(C.getGammaAtIndex(cGammaRamp.green, C.int(t)))
			ramp.Blue[t] = uint16(C.getGammaAtIndex(cGammaRamp.blue, C.int(t)))
		}

		return &ramp, nil
	}

	return nil, errors.New("glfw: Can't get gamma ramp.")
}

/*
This function sets the current gamma ramp for the specified monitor.

Parameters
	ramp	The gamma ramp to use.
Note
	Gamma ramp sizes other than 256 are not supported by all hardware.
*/
func (mon *Monitor) SetGammaRamp(ramp *GammaRamp) {
	var cGammaRamp C.GLFWgammaramp

	cGammaRamp.size = C.uint(len(ramp.Red))

	for t := 0; t < len(ramp.Red); t++ {
		C.setGammaAtIndex(cGammaRamp.red, C.int(t), C.ushort(ramp.Red[t]))
		C.setGammaAtIndex(cGammaRamp.green, C.int(t), C.ushort(ramp.Green[t]))
		C.setGammaAtIndex(cGammaRamp.blue, C.int(t), C.ushort(ramp.Blue[t]))
	}

	C.glfwSetGammaRamp(mon.internalPtr, &cGammaRamp)
}
