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

func (cpu *Cpu) Step(debug bool) {
	opcode := combine_two_bytes(cpu.memory.memory[cpu.PC], cpu.memory.memory[cpu.PC+1])
	instruction := get_high_nibble(opcode)

	if debug {
		log.Printf("CPU opcode is 0x%x (%b) (%d)\n", opcode, opcode, opcode)
		log.Printf("\tDecoded instruction: 0x%x (%b) (%d)\n", instruction, instruction, instruction)
	}

	switch instruction {
	case 0x00:
		log.Fatal("\tUNIMP\n")
		break
	case 0x1:
		log.Printf("\tJMP\n")
	case 0x2:
		addr := get_nnn(opcode)
		cpu.memory.sp += 1
		cpu.memory.stack[cpu.memory.sp] = cpu.PC
		cpu.PC = addr
		log.Printf("\tCAL %d\n", addr)
	case 0x3:
		log.Printf("\tSE\n")
	case 0x4:
		log.Printf("\tSNE\n")
	case 0x5:
		log.Printf("\tSE\n")
	case 0x6:
		x := get_x(opcode)
		lb := get_low_byte(opcode)
		cpu.V[x] = lb
		log.Printf("\tLD %d %d\n", x, lb)
	case 0x7:
		log.Printf("\tADD\n")
	case 0x8:
		break
	case 0xa:
		cpu.I = get_nnn(opcode)
		log.Printf("\tLD I %d\n", cpu.I)
	case 0xb:
		log.Fatal("\tUNIMP\n")
		break
	case 0xc:
		log.Fatal("\tUNIMP\n")
		break
	case 0xd:
		x := get_x(opcode)
		y := get_y(opcode)
		nibble := get_nibble(opcode)
		/* TODO DRAW */
		log.Printf("\tDRW %d %d %d\n", x, y, nibble)
	case 0xe:
		log.Fatal("\tUNIMP\n")
		break
	case 0xf:
		sub_instruction := get_low_byte(opcode)
		log.Printf("\tLD (0x%x)\n", sub_instruction)
		switch sub_instruction {
		case 0x07:
			break
		case 0x0a:
			break
		case 0x15:
			break
		case 0x18:
			break
		case 0x1e:
			break
		case 0x29:
			break
		case 0x33:
			break
		case 0x55:
			break
		case 0x65:
			break
		}
	default:
		log.Fatal("\tUNIMP\n")
		return
	}
	cpu.PC += 2
}
