package glfw

import (
	"fmt"
	"github.com/doxxan/glitch/glfw"
)

func ExampleCreateWindow() {
	if err := glfw.Init(); err != nil {
		fmt.Println(err)
		return
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Example", nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	window.MakeContextCurrent()
	for !window.ShouldClose() {
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			window.SetShouldClose(true)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
