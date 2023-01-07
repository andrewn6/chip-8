package main

import (
	"runtime"
	"github.com/go-gl/glfw/v3.3/glfw"
  "github.com/anddddrew/chip-8/display"
  "github.com/anddddrew/chip-8/keyboard"
)

func init() {	
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "CHIP-8", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
