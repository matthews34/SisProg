package main

import (
	"fmt"
	fb "functionalBlocks"
	reg "registers"
)

var PC = reg.PC

func reverseMap(m map[fb.Operation]string) map[string]fb.Operation {
    n := make(map[string]fb.Operation)
    for k, v := range m {
        n[v] = k
    }
    return n
}

func main() {
	// mem.WriteWord(0, 0xABCD)
	// // mem.WriteByte(3, 0xCF)
	// // mem.WriteByte(0, 0xCD)
	// // mem.PrintSlice(0, 8)
	// // PC.WriteWord(0xABCD)
	// PC = 0xabcd
	// PC.WriteByte(0xab)
	// PC.Print()
	opCode := reverseMap(fb.OpCode)
	var sub fb.Operation = opCode["sub"]
	result, err := sub.B(10,11)
	if (err == nil) {
		fmt.Println(result)
	} else {
		fmt.Println(err, reg.SR)
	}
	
}
