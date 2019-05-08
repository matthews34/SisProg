package memory

import (
	"fmt"
	"strings"
)

type word uint16

//memory size in bytes
const MEMORY_SIZE = 65536

var memory [MEMORY_SIZE]byte

func ReadWord(address word) word {
	byte1, byte2 := memory[address], memory[address+1]
	return word(byte2) | word(byte1)<<8
}

func WriteWord(address word, wordToWrite word) {
	byte2 := byte(wordToWrite & 0x00FF)
	byte1 := byte(wordToWrite >> 8 & 0x00FF)
	memory[address], memory[address+1] = byte1, byte2
}

func ReadByte(address word) byte {
	return memory[address]
}

func WriteByte(address word, byteToWrite byte) {
	memory[address] = byteToWrite
}

func PrintSlice(from word, to word) {
	bytes := memory[from:to]
	fmt.Printf("[")
	for _, byte := range bytes {
		fmt.Printf(strings.ToUpper("%.2x "), byte)
	}
	fmt.Printf("\b]\n")
}

func PrintWord(address word) {
	fmt.Printf(strings.ToUpper("%.2x\n"), ReadWord(address))
}

func PrintByte(address word) {
	fmt.Printf(strings.ToUpper("%.2x\n"), ReadByte(address))
}
