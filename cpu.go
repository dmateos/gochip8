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
	case 0x0:
		switch sub_instruction := get_low_byte(opcode); sub_instruction {
		case 0xE0:
			log.Printf("\tCLS UNIMPLEMENTED\n")
		case 0xEE:
			cpu.PC = cpu.memory.stack[cpu.memory.sp]
			cpu.memory.sp--
			log.Printf("RET to 0x%x\n", cpu.PC)
		}
	case 0x1:
		addr := get_nnn(opcode)
		cpu.PC = addr
		log.Printf("\tJMP to 0x%x\n", addr)
	case 0x2:
		addr := get_nnn(opcode)
		cpu.memory.sp += 1
		cpu.memory.stack[cpu.memory.sp] = cpu.PC
		cpu.PC = addr
		log.Printf("\tCALL 0x%x\n", addr)
	case 0x3:
		log.Printf("\tSE UNIMPLEMENTED\n")
	case 0x4:
		log.Printf("\tSNE UNIMPLEMENTED\n")
	case 0x5:
		log.Printf("\tSE UNIMPLEMENTED\n")
	case 0x6:
		x := get_x(opcode)
		lb := get_low_byte(opcode)
		cpu.V[x] = lb
		log.Printf("\tLD %d %d\n", x, lb)
	case 0x7:
		log.Printf("\tADD UNIMPLEMENTED\n")
	case 0x8:
		switch sub_instruction := get_nibble(opcode); sub_instruction {
		case 0x1:
			break
		case 0x2:
			break
		case 0x3:
			break
		case 0x4:
			break
		case 0x5:
			break
		case 0x6:
			break
		case 0x7:
			break
		case 0xE:
			break
		}
	case 0x9:
		break
	case 0xA:
		cpu.I = get_nnn(opcode)
		log.Printf("\tLD I %d UNIMPLEMENTED\n", cpu.I)
	case 0xB:
		log.Fatal("\tJP UNIMPLEMENTED\n")
		break
	case 0xC:
		log.Fatal("\tRND UNIMPLEMENTED\n")
		break
	case 0xD:
		x := get_x(opcode)
		y := get_y(opcode)
		nibble := get_nibble(opcode)
		log.Printf("\tDRW %d %d %d UNIMPLEMENTED\n", x, y, nibble)
	case 0xE:
		log.Fatal("\tUNIMP\n")
		sub_instruction := get_low_byte(opcode)
		switch sub_instruction {
		case 0x9E:
			break
		case 0xA1:
			break
		}
	case 0xF:
		sub_instruction := get_low_byte(opcode)
		log.Printf("\tLD (0x%x)\n", sub_instruction)
		switch sub_instruction {
		case 0x07:
			break
		case 0x0A:
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
