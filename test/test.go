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

	window.MakeContextCurrent()
	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
