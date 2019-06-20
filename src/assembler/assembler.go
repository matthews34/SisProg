package assembler

import (
	"bufio"
	"fmt"
	fb "functionalBlocks"
	"io"
	mem "memory"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type label_t struct {
	label string
	addr  uint16
}

var labelTable []label_t

func checkLabels() {
	for i, label1 := range labelTable {
		for j, label2 := range labelTable {
			if i != j && label1 == label2 {
				panic("Rótulos repetidos")
			}
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printSlice(array [mem.MEMORY_SIZE]byte, from uint16, to uint16) {
	bytes := array[from:to]
	fmt.Printf("[")
	for _, byte := range bytes {
		fmt.Printf("%.2X ", byte)
	}
	fmt.Printf("\b]\n")
}

func findAddress(label string) uint16 {
	for _, l := range labelTable {
		if l.label == label {
			return l.addr
		}
	}
	err := fmt.Sprintf("Rótulo %s não encontrado na tabela de rótulos\n %v", label, labelTable)
	panic(err)
}

func readLine(file *bufio.Reader) (string, bool) {
	line, err := file.ReadString('\n')
	if err == io.EOF {
		return line, true
	} else if err != nil {
		panic(err)
	} else {
		return line[:len(line)-1], false
	}
}

func Assemble(f *os.File) {
	file := bufio.NewReader(f)
	var addr uint16 = 0
	var hex [mem.MEMORY_SIZE]byte
	instruction, err := regexp.Compile("(?:	+| +)(.+)(?:	+| +)(.+)")
	check(err)
	label, err := regexp.Compile("(^[a-zA-Z0-9]+)")
	check(err)

	// first step:
	for line, EOF := readLine(file); ; line, EOF = readLine(file) {
		line = strings.ToUpper(line)
		if label := label.FindString(line); len(label) != 0 {
			labelTable = append(labelTable, label_t{label, addr})
		}
		if mnemonic := instruction.FindStringSubmatch(line); len(mnemonic) != 0 {
			if instruction, exists := fb.Mnemonics[mnemonic[1]]; exists { // instruction
				addr += 1
				if instruction.Operand {
					if _, exists := fb.Mnemonics[mnemonic[2]]; exists {
						addr += 1
					} else {
						addr += 2
					}
				}
			} else if fb.IsPseudoMnemonic(mnemonic[1]) { // pseudo instruction
				switch mnemonic[1] {
				case "@":
					i, err := strconv.Atoi(mnemonic[2])
					check(err)
					addr = uint16(i)
				case "#":
					break
				case "$":
					i, err := strconv.ParseUint(mnemonic[2], 16, 16)
					check(err)
					addr += uint16(i)
				case "K":
					addr += 2
				}
			} else { // error
				err := fmt.Sprintf("%s\n	%s é um símbolo inválido", line, mnemonic[1])
				panic(err)
			}
		}
		if EOF {
			break
		}
	}

	// second step
	f.Seek(0, 0)
	addr = 0
	for line, EOF := readLine(file); ; line, EOF = readLine(file) {
		line = strings.ToUpper(line)
		if mnemonic := instruction.FindStringSubmatch(line); len(mnemonic) != 0 {
			if instruction, exists := fb.Mnemonics[mnemonic[1]]; exists { // instruction
				fmt.Printf("%v: %s ", mnemonic[1:], instruction.Mnemonic)
				hex[addr] = instruction.Code
				addr += 1
				if instruction.Operand {
					if m, exists := fb.Mnemonics[mnemonic[2]]; exists {
						fmt.Printf("(m):%s ", m.Mnemonic)
						hex[addr] = m.Code
						addr += 1
					} else {
						fmt.Printf("(l):%s", mnemonic[2])
						hex[addr] = byte(findAddress(mnemonic[2]) >> 8 & 0x00FF)
						addr++
						hex[addr] = byte(findAddress(mnemonic[2]) & 0x00FF)
						addr++
					}
				}
				fmt.Printf("\n")
			} else if fb.IsPseudoMnemonic(mnemonic[1]) { // pseudo instruction
				switch mnemonic[1] {
				case "@":
					i, err := strconv.ParseUint(mnemonic[2], 16, 16)
					check(err)
					addr = uint16(i)
				case "#":
					// descobrir o que fazer
					break
				case "$":
					i, err := strconv.ParseUint(mnemonic[2], 16, 16)
					check(err)
					addr += uint16(i)
				case "K":
					i, err := strconv.ParseUint(mnemonic[2], 16, 16)
					check(err)
					hex[addr] = byte(uint16(i) >> 8 & 0x00FF)
					addr++
					hex[addr] = byte(uint16(i) & 0x00FF)
					addr++
				}
			}
		}
		if EOF {
			break
		}
	}
	printSlice(hex, 0, 64)
	fmt.Println(labelTable)
}
