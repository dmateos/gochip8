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
	DT, ST byte
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
		log.Printf("CPU opcode is 0x%x (b%b)\n", opcode, opcode)
		log.Printf("\tDecoded instruction: 0x%x (b%b)\n", instruction, instruction)
		log.Printf("\tPC: 0x%x (%d), SP: %d\n", cpu.PC, cpu.PC, cpu.memory.sp)
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
			log.Printf("RET to 0x%x (%d)\n", cpu.PC, cpu.PC)
		}
	case 0x1:
		skip_pc = true
		addr := get_nnn(opcode)
		cpu.PC = addr
		log.Printf("\tJMP to 0x%x (%d)\n", addr, addr)
	case 0x2:
		skip_pc = true
		addr := get_nnn(opcode)
		cpu.memory.sp++
		cpu.memory.stack[cpu.memory.sp] = cpu.PC + 2
		cpu.PC = addr
		log.Printf("\tCALL 0x%x (%d)\n", addr, addr)
	case 0x3:
		x := get_x(opcode)
		b := get_low_byte(opcode)
		if cpu.V[x] == b {
			skip_pc = true
			cpu.PC += 2
		}
		log.Printf("\tSEQ reg:0x%x (%d) 0x%x (%d) 0x%x (%d)\n", x, x, cpu.V[x], cpu.V[x], b, b)
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
		x := get_x(opcode)
		lb := get_low_byte(opcode)
		cpu.V[x] += lb
		log.Printf("\tADD 0x%x (%d) 0x%x (%d)\n", x, x, lb, lb)
	case 0x8:
		sub_instruction := get_nibble(opcode)
		switch sub_instruction {
		case 0x0:
			x := get_x(opcode)
			y := get_y(opcode)
			cpu.V[x] = cpu.V[y]
			log.Printf("\tLD\n")
		case 0x1:
			x := get_x(opcode)
			y := get_y(opcode)
			cpu.V[x] = cpu.V[x] | cpu.V[y]
			log.Printf("\tOR\n")
		case 0x2:
			x := get_x(opcode)
			y := get_y(opcode)
			cpu.V[x] = cpu.V[x] & cpu.V[y]
			log.Printf("\tAND\n")
		case 0x3:
			x := get_x(opcode)
			y := get_y(opcode)
			cpu.V[x] = cpu.V[x] ^ cpu.V[y]
			log.Printf("\tXOR\n")
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
		x := get_x(opcode)
		y := get_y(opcode)
		if x != y {
			skip_pc = true
			cpu.PC += 2
		}
		log.Printf("\tSNE 0x%x (%d) 0x%x (%d)\n", x, x, y, y)
	case 0xA:
		cpu.I = get_nnn(opcode)
		log.Printf("\tLD I %d\n", cpu.I)
	case 0xB:
		addr := get_nnn(opcode)
		skip_pc = true
		cpu.PC += addr + uint16(cpu.V[0])
		log.Printf("JMP 0x%x (%d) + (%d)\n", addr, addr, cpu.V[0])
	case 0xC:
		//kk := get_low_byte(opcode)
		log.Printf("\tRND UNIMPLEMENTED\n")
	case 0xD:
		x := get_x(opcode)
		y := get_y(opcode)
		nibble := get_nibble(opcode)
		log.Printf("\tDRW %d %d %d UNIMPLEMENTED\n", x, y, nibble)
	case 0xE:
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
			x := get_x(opcode)
			cpu.V[x] = cpu.DT
			log.Printf("\tLD DT %0x%x (%d) %d\n", x, x, cpu.DT)
		case 0x0A:
			log.Printf("\tLD UNIMPLEMENTED\n")
		case 0x15:
			x := get_x(opcode)
			cpu.DT = cpu.V[x]
			log.Printf("\tLD %0x%x (%d) %d DT\n", x, x, cpu.DT)
		case 0x18:
			x := get_x(opcode)
			cpu.ST = cpu.V[x]
			log.Printf("\tLD %0x%x (%d) %d ST\n", x, x, cpu.DT)
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

	if cpu.DT != 0 {
		cpu.DT--
	}

	if cpu.ST != 0 {
		cpu.ST--
	}
}
