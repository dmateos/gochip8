package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	program_data, err := ioutil.ReadFile("roms/pong.rom")
	if err != nil {
		log.Fatal(err)
	}

	cpu := NewCpu(program_data)

	fmt.Println(program_data)
	fmt.Println(cpu)
	fmt.Println(cpu.memory)

	for i := 0; i < 16; i++ {
		cpu.Step(false)
	}

	fmt.Println(cpu)
}
