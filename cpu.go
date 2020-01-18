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
	var skip_pc bool = false

	opcode := combine_two_bytes(cpu.memory.memory[cpu.PC], cpu.memory.memory[cpu.PC+1])
	instruction := get_high_nibble(opcode)

	if debug {
		log.Printf("CPU opcode is 0x%x (%b) (%d)\n", opcode, opcode, opcode)
		log.Printf("\tDecoded instruction: 0x%x (%b) (%d)\n", instruction, instruction, instruction)
		log.Printf("\tPC: 0x%x (%d)\n", cpu.PC, cpu.PC)
	}

	switch instruction {
	case 0x0:
		sub_instruction := get_low_byte(opcode)
		switch sub_instruction {
		case 0xE0:
			log.Printf("\tCLS UNIMPLEMENTED\n")
		case 0xEE:
			skip_pc = true
			cpu.PC = cpu.memory.stack[cpu.memory.sp]
			cpu.memory.sp--
			log.Printf("RET to 0x%x\n", cpu.PC)
		}
	case 0x1:
		skip_pc = true
		addr := get_nnn(opcode)
		cpu.PC = addr
		log.Printf("\tJMP to 0x%x\n", addr)
	case 0x2:
		skip_pc = true
		addr := get_nnn(opcode)
		cpu.memory.sp += 1
		cpu.memory.stack[cpu.memory.sp] = cpu.PC + 2
		cpu.PC = addr
		log.Printf("\tCALL 0x%x\n", addr)
	case 0x3:
		x := get_x(opcode)
		b := get_low_byte(opcode)
		if cpu.V[x] == b {
			skip_pc = true
			cpu.PC += 2
		}
		log.Printf("\tSE 0x%x 0x%x (%d) 0x%x (%d)\n", x, cpu.V[x], cpu.V[x], b, b)
	case 0x4:
		x := get_x(opcode)
		b := get_low_byte(opcode)
		if cpu.V[x] != b {
			skip_pc = true
			cpu.PC += 2
		}
		log.Printf("\tSNE 0x%x 0x%x (%d) 0x%x (%d)\n", x, cpu.V[x], cpu.V[x], b, b)
	case 0x5:
		x := get_x(opcode)
		y := get_y(opcode)
		if x == y {
			skip_pc = true
			cpu.PC += 2
		}
		log.Printf("\tSE 0x%x (%d) 0x%x (%d)\n", x, x, y, y)
	case 0x6:
		x := get_x(opcode)
		lb := get_low_byte(opcode)
		cpu.V[x] = lb
		log.Printf("\tLD %d %d\n", x, lb)
	case 0x7:
		log.Printf("\tADD UNIMPLEMENTED\n")
	case 0x8:
		sub_instruction := get_nibble(opcode)
		switch sub_instruction {
		case 0x0:
			log.Printf("\tLD UNIMPLEMENTED\n")
		case 0x1:
			log.Printf("\tOR UNIMPLEMENTED\n")
		case 0x2:
			log.Printf("\tAND UNIMPLEMENTED\n")
		case 0x3:
			log.Printf("\tXOR UNIMPLEMENTED\n")
		case 0x4:
			log.Printf("\tADD UNIMPLEMENTED\n")
		case 0x5:
			log.Printf("\tSUB UNIMPLEMENTED\n")
		case 0x6:
			log.Printf("\tSHR UNIMPLEMENTED\n")
		case 0x7:
			log.Printf("\tSUBN UNIMPLEMENTED\n")
		case 0xE:
			log.Printf("\tSHL UNIMPLEMENTED\n")
		}
	case 0x9:
		log.Printf("\tSNE UNIMPLEMENTED\n")
	case 0xA:
		cpu.I = get_nnn(opcode)
		log.Printf("\tLD I %d\n", cpu.I)
	case 0xB:
		log.Printf("\tJP UNIMPLEMENTED\n")
		break
	case 0xC:
		log.Printf("\tRND UNIMPLEMENTED\n")
		break
	case 0xD:
		x := get_x(opcode)
		y := get_y(opcode)
		nibble := get_nibble(opcode)
		log.Printf("\tDRW %d %d %d UNIMPLEMENTED\n", x, y, nibble)
	case 0xE:
		log.Printf("\tUNIMP\n")
		sub_instruction := get_low_byte(opcode)
		switch sub_instruction {
		case 0x9E:
			log.Printf("\tSKIP UNIMPLEMENTED\n")
		case 0xA1:
			log.Printf("\tNSKIP UNIMPLEMENTED\n")
		}
	case 0xF:
		sub_instruction := get_low_byte(opcode)
		switch sub_instruction {
		case 0x07:
			log.Printf("\tLD UNIMPLEMENTED\n")
		case 0x0A:
			log.Printf("\tLD UNIMPLEMENTED\n")
		case 0x15:
			log.Printf("\tLD UNIMPLEMENTED\n")
		case 0x18:
			log.Printf("\tLD UNIMPLEMENTED\n")
		case 0x1E:
			log.Printf("\tADD UNIMPLEMENTED\n")
		case 0x29:
			log.Printf("\tLD UNIMPLEMENTED\n")
		case 0x33:
			log.Printf("\tLD UNIMPLEMENTED\n")
		case 0x55:
			log.Printf("\tLD UNIMPLEMENTED\n")
		case 0x65:
			log.Printf("\tLD UNIMPLEMENTED\n")
		}
	default:
		log.Fatal("\tUNIMP\n")
		return
	}

	if !skip_pc {
		cpu.PC += 2
	}
}
