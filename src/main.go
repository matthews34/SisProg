package main

import (
	mem "memory"
)

func main() {
	mem.WriteWord(0, 0xABCD)
	mem.WriteByte(3, 0xCF)
	mem.PrintSlice(0, 8)
	mem.WriteByte(0, 0xCD)
	mem.PrintSlice(0, 8)
	mem.PrintWord(1)
	mem.PrintByte(3)
}
