package memory

import (
	"fmt"
)

//memory size in bytes
const MEMORY_SIZE = 65536

var memory [MEMORY_SIZE]byte

func ReadWord(address uint16) uint16 {
	byte1, byte2 := memory[address], memory[address+1]
	return uint16(byte2) | uint16(byte1)<<8
}

func WriteWord(address uint16, wordToWrite uint16) {
	byte2 := byte(wordToWrite & 0x00FF)
	byte1 := byte(wordToWrite >> 8 & 0x00FF)
	memory[address], memory[address+1] = byte1, byte2
}

func ReadByte(address uint16) byte {
	return memory[address]
}

func WriteByte(address uint16, byteToWrite byte) {
	memory[address] = byteToWrite
}

func PrintSlice(from uint16, to uint16) {
	bytes := memory[from:to]
	fmt.Printf("[")
	for _, byte := range bytes {
		fmt.Printf("%.2X ", byte)
	}
	fmt.Printf("\b]\n")
}

func PrintWord(address uint16) {
	fmt.Printf("%.2X\n", ReadWord(address))
}

func PrintByte(address uint16) {
	fmt.Printf("%.2X\n", ReadByte(address))
}
