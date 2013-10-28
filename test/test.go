package main

import (
	"fmt"
	"github.com/doxxan/glitch/glfw"
)

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Println(err)
		return
	}
	defer glfw.Terminate()

	fmt.Println(glfw.GetVersion())
	fmt.Println(glfw.GetVersionString())

	glfw.SetErrorCallback(func(code glfw.ErrorCode, description string) {
		fmt.Println(code, description)
	})

	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	window.SetIconifyCallback(func(iconified bool) {
		fmt.Println(iconified)
	})

	window.SetPositionCallback(func(x, y int) {
		fmt.Println(x, y)
	})

	fmt.Println(glfw.GetMonitors())

	if monitor, err := glfw.GetPrimaryMonitor(); err == nil {
		fmt.Println(monitor.GetVideoMode())
		fmt.Println(monitor.GetVideoModes())
		fmt.Println(monitor.GetPosition())
		fmt.Println(monitor.GetPhysicalSize())
		fmt.Println(monitor.GetName())
	}

	window.SetCharCallback(func(char rune) {
		fmt.Println(char)
	})

	window.SetCursorEnterCallback(func(flag bool) {
		fmt.Println(flag)
	})

	window.SetScrollCallback(func(xoff, yoff float64) {
		fmt.Println(xoff, yoff)
	})

	window.MakeContextCurrent()
	for !window.ShouldClose() {
		if buttons, err := glfw.GetJoystickButtons(glfw.Joystick1); err == nil {
			if buttons[1] == glfw.Press {
				fmt.Println("Hit!")
			}
			if buttons[2] == glfw.Press {
				fmt.Println(glfw.GetJoystickAxes(glfw.Joystick1))
			}
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
