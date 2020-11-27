package dictionary

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/apex/log"
)

// Keymask is a bitmask for steno keys
type Keymask uint32

// The steno keyboard keys
const (
	Num       Keymask = 1 << 22
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

func ParseStroke(in string) (Keymask, error) {
	in = strings.ToUpper(in)
	letterOrder := regexp.MustCompile(`^S?T?K?P?W?H?R?A?O?\*?-?E?U?F?R?P?B?L?G?T?S?D?Z?$`)
	numberOrder := regexp.MustCompile(`^#?1?2?K?3?W?4?R?5?0?\*?-?E?U?6?R?7?B?8?G?9?S?D?Z?$`)
	numberWithLettersOrder := regexp.MustCompile(`^#[S1]?[T2]?K?[P3]?W?[H4]?R?[A5]?[O0]?\*?-?E?U?[F6]?R?[P7]?B?[L8]?G?[T9]?S?D?Z?$`)
	cmpLet := "STKPWHRAO*-EUFRPBLGTSDZ"
	cmpNum := "12K3W4R50*-EU6R7B8G9SDZ"
	validLetters := letterOrder.MatchString(in)
	validNumbers := numberOrder.MatchString(in) || numberWithLettersOrder.MatchString(in)
	log.WithFields(log.Fields{
		"validLetters": validLetters,
		"validNumbers": validNumbers,
	}).Debug("input validated")

	numberFlag := false
	if !validLetters {
		if !validNumbers {
			return 0, fmt.Errorf("Input keys %s did not seem to be in steno order (%s or %s)", in, cmpLet, cmpNum)
		}

		// remove # from input, if it's there
		log.Debug("# is on")
		if in[0] == '#' {
			in = in[1:]
		}
		numberFlag = true
	}

	cursor := 0
	for len(in) < len(cmpLet) {
		if cursor == len(in) {
			// we must have matched all of in already
			break
		}
		log.Debugf("%s", cmpNum)
		log.Debugf("%s", cmpLet)
		log.Debugf("%s", in)
		// print cursor
		cstring := ""
		for i := 0; i < cursor; i++ {
			cstring += " "
		}
		cstring += "^"
		log.Debugf("%s", cstring)

		if in[cursor] != cmpNum[cursor] && in[cursor] != cmpLet[cursor] {
			// insert a space just before the cursor
			in = in[0:cursor] + " " + in[cursor:]
		}
		cursor++
	}
	log.Debugf("%s", cmpNum)
	log.Debugf("%s", cmpLet)
	log.Debugf("%s", in)

	// remove the space for the hyphen
	if len(in) > 10 {
		in = in[0:10] + in[11:]
	}

	// we now have an input string that has spaces where there is no match with
	// the steno order.
	// we can turn this into a keymask
	var mask Keymask
	if numberFlag {
		mask = 1 << 22
	}

	shift := 21
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

	if k&Num == Num {
		return k.numberString()
	}

	return k.letterString()
}

func (k Keymask) numberString() string {
	str := ""
	if !k.hasNumbers() {
		str += "#"
	}
	str += stringIfContains(k, LeftS, "1")
	str += stringIfContains(k, LeftT, "2")
	str += stringIfContains(k, LeftK, "K")
	str += stringIfContains(k, LeftP, "3")
	str += stringIfContains(k, LeftW, "W")
	str += stringIfContains(k, LeftH, "4")
	str += stringIfContains(k, LeftR, "R")
	str += stringIfContains(k, A, "5")
	str += stringIfContains(k, O, "0")
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
	str += stringIfContains(k, RightF, "6")
	str += stringIfContains(k, RightR, "R")
	str += stringIfContains(k, RightP, "7")
	str += stringIfContains(k, RightB, "B")
	str += stringIfContains(k, RightL, "8")
	str += stringIfContains(k, RightG, "G")
	str += stringIfContains(k, RightT, "9")
	str += stringIfContains(k, RightS, "S")
	str += stringIfContains(k, RightD, "D")
	str += stringIfContains(k, RightZ, "Z")
	return str
}

func (k Keymask) letterString() string {
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

func (k Keymask) hasNumbers() bool {
	return k&(LeftS|LeftT|LeftP|LeftH|A|O|RightF|RightP|RightL|RightT) > 0
}

func (k Keymask) IsFingerspelling() bool {
	fingerspellings := []Keymask{
		A,
		LeftP | LeftW,
		LeftK | LeftR,
		LeftT | LeftK,
		E,
		LeftT | LeftP,
		LeftT | LeftK | LeftP | LeftW,
		LeftH,
		E | U,
		LeftS | LeftK | LeftW | LeftR,
		LeftK,
		LeftH | LeftR,
		LeftP | LeftH,
		LeftT | LeftP | LeftH,
		O,
		LeftP,
		LeftK | LeftW,
		LeftR,
		LeftS,
		LeftT,
		U,
		LeftS | LeftR,
		LeftW,
		LeftK | LeftP,
		LeftK | LeftW | LeftR,
		LeftS | LeftT | LeftK | LeftP | LeftW,
	}
	for _, fingerspelling := range fingerspellings {
		if k == fingerspelling {
			return true
		}
	}
	return false
}
