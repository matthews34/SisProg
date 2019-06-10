package registers

import "fmt"

type register uint16
type word uint16

var PC register
var SP register
var SR byte // 4 bits mais significativos em branco
var A, D [8]register

func (r *register) ReadWord() word {
	return word(*r)
}

func (r *register) WriteWord(wordToWrite word) {
	*r = register(wordToWrite)
}

func (r *register) ReadByte() byte {
	return byte(*r & 0x00FF)
}

func (r *register) WriteByte(byteToWrite byte) {
	*r = register((word(*r) & 0xFF00) | word(byteToWrite))
}

func (r *register) Print() {
	fmt.Printf("%.2X\n", r.ReadWord())
}
