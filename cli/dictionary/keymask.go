package dictionary

import (
	"fmt"
	"regexp"
	"strings"
)

// Keymask is a bitmask for steno keys
type Keymask uint32

// The steno keyboard keys
const (
	LeftS     Keymask = 1 << 21
	LeftT     Keymask = 1 << 20
	LeftK     Keymask = 1 << 19
	LeftP     Keymask = 1 << 18
	LeftW     Keymask = 1 << 17
	LeftH     Keymask = 1 << 16
	LeftR     Keymask = 1 << 15
	A         Keymask = 1 << 14
	O         Keymask = 1 << 13
	Star      Keymask = 1 << 12
	E         Keymask = 1 << 11
	U         Keymask = 1 << 10
	RightF    Keymask = 1 << 9
	RightR    Keymask = 1 << 8
	RightP    Keymask = 1 << 7
	RightB    Keymask = 1 << 6
	RightL    Keymask = 1 << 5
	RightG    Keymask = 1 << 4
	RightT    Keymask = 1 << 3
	RightS    Keymask = 1 << 2
	RightD    Keymask = 1 << 1
	RightZ    Keymask = 1
	allLefts  Keymask = LeftS | LeftT | LeftK | LeftP | LeftW | LeftH | LeftR
	allRights Keymask = RightF | RightR | RightP | RightB | RightL | RightG | RightT | RightS | RightD | RightZ
)

func ParseKeymask(in string) (Keymask, error) {
	in = strings.ToUpper(in)
	valid := regexp.MustCompile(`S?T?K?P?W?H?R?A?O?\*?-?E?U?F?R?P?B?L?G?T?S?D?Z?`)
	cmp := "STKPWHRAO*-EUFRPBLGTSDZ"
	if !valid.MatchString(in) {
		return 0, fmt.Errorf("Input keys %s did not seem to be in steno order (%s)", in, cmp)
	}

	cursor := 0
	for len(in) < len(cmp) {
		if cursor == len(in) {
			// we must have matched all of in already
			break
		}
		// fmt.Println(cmp)
		// fmt.Println(in)
		// print cursor
		// for i := 0; i < cursor; i++ {
		// 	fmt.Printf(" ")
		// }
		// fmt.Println("^")
		if in[cursor] != cmp[cursor] {
			// fmt.Printf("no match (%c vs %c)\n", in[cursor], cmp[cursor])
			// insert a space just before the cursor
			in = in[0:cursor] + " " + in[cursor:]
		} else {
			// fmt.Printf("%c matches %c\n", in[cursor], cmp[cursor])
		}
		cursor++

		// fmt.Println()
		// fmt.Println()
	}
	// fmt.Println(cmp)
	// fmt.Println(in)

	// we now have an input string that has spaces where there is no matck with
	// the steno order.
	// we can turn this into a keymask
	shift := 21
	var mask Keymask
	for _, c := range in {
		if c != ' ' && c != '-' {
			mask = mask | 1<<shift
		}
		shift--
	}

	return mask, nil
}

func (k Keymask) String() string {
	if k == 0 {
		return ""
	}

	str := ""
	str += stringIfContains(k, LeftS, "S")
	str += stringIfContains(k, LeftT, "T")
	str += stringIfContains(k, LeftK, "K")
	str += stringIfContains(k, LeftP, "P")
	str += stringIfContains(k, LeftW, "W")
	str += stringIfContains(k, LeftH, "H")
	str += stringIfContains(k, LeftR, "R")
	str += stringIfContains(k, A, "A")
	str += stringIfContains(k, O, "O")
	if k&Star == Star {
		str += "*"
	} else {
		// hyphens never go alongside stars. also:
		if k.hasRight() && !k.hasVowel() {
			str += "-"
		}
	}
	str += stringIfContains(k, E, "E")
	str += stringIfContains(k, U, "U")
	str += stringIfContains(k, RightF, "F")
	str += stringIfContains(k, RightR, "R")
	str += stringIfContains(k, RightP, "P")
	str += stringIfContains(k, RightB, "B")
	str += stringIfContains(k, RightL, "L")
	str += stringIfContains(k, RightG, "G")
	str += stringIfContains(k, RightT, "T")
	str += stringIfContains(k, RightS, "S")
	str += stringIfContains(k, RightD, "D")
	str += stringIfContains(k, RightZ, "Z")
	return str
}

// stringIfEqual returns the given string if the first Keymask contains the
// second, otherwise it returns empty string
func stringIfContains(k1, k2 Keymask, s string) string {
	if k1&k2 == k2 {
		return s
	}
	return ""
}

func (k Keymask) hasLeft() bool {
	return k&allLefts > 0
}

func (k Keymask) hasRight() bool {
	return k&allRights > 0
}

func (k Keymask) hasStar() bool {
	return k&Star == Star
}

func (k Keymask) hasVowel() bool {
	return k&(A|O|E|U) > 0
}
