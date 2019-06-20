package main

import (
	asm "assembler"
	"os"
	reg "registers"
)

var PC = reg.PC

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func reverseMap(m map[fb.Operation]string) map[string]fb.Operation {
// 	n := make(map[string]fb.Operation)
// 	for k, v := range m {
// 		n[v] = k
// 	}
// 	return n
// }

func main() {
	// cli.Run()

	file, err := os.Open("/home/matheus/Documents/SisProg/src/programs/teste.asm")
	check(err)
	asm.Assemble(file)
	file.Close()
}
