package main

import (
	"fmt"
	"github.com/doxxan/glitch/gl"
	"github.com/doxxan/glitch/glfw"
)

var screenWidth int = 640
var screenHeight int = 480

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Println(err)
		return
	}
	defer glfw.Terminate()

	glfw.SwapInterval(0)
	window, err := glfw.CreateWindow(screenWidth, screenHeight, "Testing", nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	window.MakeContextCurrent()

	gl.Init()
	gl.Enable(gl.CullFaceConst)
	gl.Enable(gl.DepthTest)
	gl.ClearColor(0.5, 0.5, 0.5, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEqual)
	gl.Viewport(0, 0, screenWidth, screenHeight)

	for !window.ShouldClose() {
		gl.Clear(gl.ColorBufferBit | gl.DepthBufferBit)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
