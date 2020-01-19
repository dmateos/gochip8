package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"io/ioutil"
	"log"
)

func init_video() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Chip8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

}

func main() {
	program_data, err := ioutil.ReadFile("roms/pong.rom")
	if err != nil {
		log.Fatal(err)
	}

	cpu := NewCpu(program_data)

	fmt.Println(program_data)
	fmt.Println(cpu)
	fmt.Println(cpu.memory)

	for {
		cpu.Step(true)
		fmt.Scanln()
	}

	fmt.Println(cpu)
}
