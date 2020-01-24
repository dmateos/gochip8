package main

import (
	"bufio"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func init_video() *sdl.Window {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 64*10, 32*10, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	return window
}

func get_byte() byte {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadByte()
	return text
}

func display_byte(x, y int32, window *sdl.Window) {
	rect := sdl.Rect{x * 10, y * 10, 10, 10}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	surface.FillRect(&rect, 0xFFFF0000)
	window.UpdateSurface()
}

func main() {
	program_data, err := ioutil.ReadFile("roms/pong.rom")
	if err != nil {
		log.Fatal(err)
	}

	cpu := NewCpu(program_data)
	window := init_video()
	defer window.Destroy()

	fmt.Println(program_data)
	fmt.Println(cpu)
	fmt.Println(cpu.memory)

	for {
		cpu.Step(true, get_byte)

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println(cpu)
}
