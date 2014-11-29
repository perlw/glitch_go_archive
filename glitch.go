package glitch

import (
	"github.com/doxxan/glitch/core/gl"
	"github.com/doxxan/glitch/core/glfw"
	"runtime"
)

var window *glfw.Window
var screenWidth int = 640
var screenHeight int = 480

type MainFunc func()

func Init() error {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		return err
	}

	if runtime.GOOS == "darwin" {
		glfw.WindowHint(glfw.ContextVersionMajorConst, 3)
		glfw.WindowHint(glfw.ContextVersionMinorConst, 2)
		glfw.WindowHint(glfw.OpenGLProfileConst, glfw.OpenGLCoreProfileConst)
		glfw.WindowHint(glfw.OpenGLForwardCompatConst, glfw.TrueConst)
	}

	glfw.SwapInterval(0)
	var err error
	if window, err = glfw.CreateWindow(screenWidth, screenHeight, "glItch", nil, nil); err != nil {
		return err
	}
	window.MakeContextCurrent()

	gl.Init()
	gl.Enable(gl.CullFaceConst)
	gl.Enable(gl.DepthTestConst)
	gl.ClearColor(0.5, 0.5, 0.5, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEqualConst)

	gl.Viewport(0, 0, screenWidth, screenHeight)

	return nil
}

func Terminate() {
	glfw.Terminate()
}

func Run(fn MainFunc) {
	for !window.ShouldClose() {
		fn()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
