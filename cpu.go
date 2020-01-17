package main

import (
	"log"
)

type Memory struct {
	stack  [16]uint16
	memory [4096]byte
}

type Cpu struct {
	V      [16]byte
	I      uint16
	PC     uint16
	memory *Memory
}

func NewCpu(rom []byte) *Cpu {
	cpu := Cpu{}
	memory := Memory{}

	/* Execution starts at 512 on chip8. */
	copy(memory.memory[512:], rom)
	cpu.PC = 512
	cpu.memory = &memory

	return &cpu
}

func (cpu *Cpu) Step() {
	opcode := uint16(cpu.memory.memory[cpu.PC])<<8 | uint16(cpu.memory.memory[cpu.PC+1])
	instruction := opcode & 0xF000

	log.Printf("CPU opcode is 0x%x (%b) (%d)\n", opcode, opcode, opcode)
	log.Printf("\tDecoded instruction: 0x%x (%b) (%d)\n", instruction, instruction, instruction)

	switch instruction {
	case 0x00:
		break
	case 0x1000:
		log.Printf("\tJMP instruction\n")
	case 0x2000:
		log.Printf("\tCALL instruction\n")
	case 0x3000:
		log.Printf("\tSE instruction\n")
	case 0x4000:
		log.Printf("\tSNE instruction\n")
	case 0x5000:
		log.Printf("\tSE instruction\n")
	case 0x6000:
		arg2 := get_low_byte(opcode)
		log.Printf("\tLD instruction %d\n", arg2)
	case 0x7000:
		log.Printf("\tADD instruction\n")
	case 0x8000:
		break
	case 0xa000:
		log.Printf("\tLD instruction\n")
	case 0xb000:
		break
	case 0xc000:
		break
	case 0xd000:
		log.Printf("\tDRW instruction\n")
	case 0xe000:
		break
	case 0xf000:
		break
	default:
		log.Fatal("\tUNIMP\n")
		return
	}
	cpu.PC += 2
}
