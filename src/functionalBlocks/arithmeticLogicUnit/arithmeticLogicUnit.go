package arithmeticLogicUnit

import (
	"errors"
	reg "registers"
)

type word uint16

type Operation byte

func newZero() error {
	reg.SR |= 0x04
	return errors.New("Zero")
}
func newNegative() error {
	reg.SR |= 0x02
	return errors.New("Negative")
}
func newOverflow() error {
	reg.SR |= 0x01
	return errors.New("Overflow")
}

var OpCode = map[Operation]string {
	0x00: "nop",
	0x01: "jmp",
	0x02: "jmz",
	0x03: "jmn",
	0x04: "jsr",
	0x10: "add",
	0x11: "sub",
	0x12: "mul",
	// 0x13: "div",
	0x20: "ld",
	0x21: "str",
	0x30: "ctl",
	0x31: "io",
}

func (op Operation) B(a byte, b byte) (byte, error) {
	switch OpCode[op] {
	case "add": // sets only overflow
		reg.SR = 0x00
		if (int(a) + int(b) > 0xFF){ // overflow
			return a + b, newOverflow()
		} else {
			return a + b, nil
		}
	case "sub": // sets zero, negative and overflow
		reg.SR = 0x00
		if (int(a) - int(b) < 0) { // overflow
			return a - b, newOverflow()
		} else if (int(a) - int(b) >= 0x80) { // negative
			return a - b, newNegative()
		} else if (a == b) { // zero
			return a - b, newZero()
		} else {
			return a - b, nil
		}
	case "mul":
		reg.SR = 0x00
		if (int(a) * int(b) > 0xFF) { // overflow
			return a * b, newOverflow()
		} else if (int(a) * int(b) >= 0x80) { // negative
			return a * b, newNegative()
		} else if (a == 0 || b == 0) { // zero
			return a * b, newZero()
		} else {
			return a * b, nil
		}
	case "div":
		reg.SR = 0x00
		if 
	default:
		return 0, errors.New("Operation not implemented")
	}
}

// func Operation(a interface{}, b interface{}, opCode byte) interface{} {
// 	return func(a byte, b byte) byte {
// 		return a + b
// 	}
// }
