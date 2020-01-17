package main

import (
	"log"
)

func get_nnn(opcode uint16) {
}

func get_nibble(opcode uint16) {

}

func get_x(opcode uint16) {

}

func get_y(opcode uint16) {

}

func get_low_byte(opcode uint16) byte {
	return byte(opcode & 0x000F)
}
