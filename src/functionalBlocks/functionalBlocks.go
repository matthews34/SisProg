package functionalBlocks

type mnemonic struct {
	Mnemonic string
	Code     byte
	Operand  bool // size in bytes
}

var Mnemonics = map[string]mnemonic{
	"JP": {"JP", 0x00, true},
	"JZ": {"JZ", 0x01, true},
	"JN": {"JN", 0x02, true},
	"LV": {"LV", 0x03, true},
	"+":  {"+", 0x04, true},
	"-":  {"-", 0x05, true},
	"*":  {"*", 0x06, true},
	"/":  {"/", 0x07, true},
	"LD": {"LD", 0x08, true},
	"MM": {"MM", 0x09, true},
	"SC": {"SC", 0x0A, true},
	"OS": {"OS", 0x0B, true},
	"HM": {"HM", 0x0C, true},
	"RS": {"RS", 0x0D, true},
	"GD": {"GD", 0x0E, true},
	"PD": {"PD", 0x0F, true},
}

var PseudoMnemonics = [4]string{"@", "#", "$", "K"}

func IsPseudoMnemonic(s string) bool {
	for _, p := range PseudoMnemonics {
		if s == p {
			return true
		}
	}
	return false
}
