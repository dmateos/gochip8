package main

import ()

func combine_two_bytes(x, y byte) uint16 {
	return uint16(x)<<8 | uint16(y)
}

func get_nnn(opcode uint16) uint16 {
	return opcode & 0x0FFF
}

func get_nibble(opcode uint16) uint16 {
	return opcode & 0x000F
}

func get_x(opcode uint16) byte {
	return byte((opcode >> 8) & 0x000F)
}

func get_y(opcode uint16) byte {
	return (byte(opcode) & 0xF0) >> 4
}

func get_low_byte(opcode uint16) byte {
	return byte(opcode & 0x00FF)
}

func get_high_nibble(opcode uint16) uint16 {
	return (opcode & 0xF000) >> 12
}
