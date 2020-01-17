package main

import (
	"log"
)

type Memory struct {
	memory [4096]byte
	stack  [16]uint16
	sp     byte
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
	opcode := combine_two_bytes(cpu.memory.memory[cpu.PC], cpu.memory.memory[cpu.PC+1])
	instruction := get_high_nibble(opcode)

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
		x := get_x(opcode)
		lb := get_low_byte(opcode)
		cpu.V[x] = lb
		log.Printf("\tLD instruction %d %d\n", x, lb)
	case 0x7000:
		log.Printf("\tADD instruction\n")
	case 0x8000:
		break
	case 0xa000:
		cpu.I = get_nnn(opcode)
		log.Printf("\tLD instruction I %d\n", cpu.I)
	case 0xb000:
		break
	case 0xc000:
		break
	case 0xd000:
		x := get_x(opcode)
		y := get_y(opcode)
		nibble := get_nibble(opcode)
		log.Printf("\tDRW instruction %d %d %d\n", x, y, nibble)
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
