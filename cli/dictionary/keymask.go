package dictionary

import (
	"encoding/json"
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
	LeftA     Keymask = 1 << 14
	LeftO     Keymask = 1 << 13
	Star      Keymask = 1 << 12
	RightE    Keymask = 1 << 11
	RightU    Keymask = 1 << 10
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
	Steno1    Keymask = Num | LeftS
	Steno2    Keymask = Num | LeftT
	Steno3    Keymask = Num | LeftP
	Steno4    Keymask = Num | LeftH
	Steno5    Keymask = Num | LeftA
	Steno6    Keymask = Num | RightF
	Steno7    Keymask = Num | RightP
	Steno8    Keymask = Num | RightL
	Steno9    Keymask = Num | RightT
	Steno0    Keymask = Num | LeftO
)

// ParseStroke takes in a string (e.g. "STPH") and returns a Keymask or an error.
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

// numberString returns the string representation of the receiver using the
// steno number order (#12K3W4R50*EU6R7B8G9SDZ)
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
	str += stringIfContains(k, LeftA, "5")
	str += stringIfContains(k, LeftO, "0")
	if k&Star == Star {
		str += "*"
	} else {
		// hyphens never go alongside stars. also:
		if k.hasRight() && !k.hasVowel() {
			str += "-"
		}
	}
	str += stringIfContains(k, RightE, "E")
	str += stringIfContains(k, RightU, "U")
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

// letterString returns the string representation of the receiver using the
// steno letter order (STKPWHRAO*EUFRPBLGTSDZ)
func (k Keymask) letterString() string {
	str := ""
	str += stringIfContains(k, LeftS, "S")
	str += stringIfContains(k, LeftT, "T")
	str += stringIfContains(k, LeftK, "K")
	str += stringIfContains(k, LeftP, "P")
	str += stringIfContains(k, LeftW, "W")
	str += stringIfContains(k, LeftH, "H")
	str += stringIfContains(k, LeftR, "R")
	str += stringIfContains(k, LeftA, "A")
	str += stringIfContains(k, LeftO, "O")
	if k&Star == Star {
		str += "*"
	} else {
		// hyphens never go alongside stars. also:
		if k.hasRight() && !k.hasVowel() {
			str += "-"
		}
	}
	str += stringIfContains(k, RightE, "E")
	str += stringIfContains(k, RightU, "U")
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

// stringIfContains returns the given string if the first Keymask contains the
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
	return k&(LeftA|LeftO|RightE|RightU) > 0
}

func (k Keymask) hasNumbers() bool {
	return k&(LeftS|LeftT|LeftP|LeftH|LeftA|LeftO|RightF|RightP|RightL|RightT) > 0
}

// allFingerspellings returns the mapping of each of Plover's left-hand
// fingerspellings (including alternate definitions) to their corresponging
// Qwerty keys. Note: the steno chords do not include `*`.
func allFingerspellings() map[Keymask]QwertyKey {
	return map[Keymask]QwertyKey{
		LeftA:                                 QwertyA,
		LeftP | LeftW:                         QwertyB,
		LeftK | LeftR:                         QwertyC,
		LeftT | LeftK:                         QwertyD,
		RightE:                                QwertyE,
		LeftT | LeftP:                         QwertyF,
		LeftT | LeftK | LeftP | LeftW:         QwertyG,
		LeftH:                                 QwertyH,
		RightE | RightU:                       QwertyI,
		LeftS | LeftK | LeftW | LeftR:         QwertyJ,
		LeftK:                                 QwertyK,
		LeftH | LeftR:                         QwertyL,
		LeftP | LeftH:                         QwertyM,
		LeftT | LeftP | LeftH:                 QwertyN,
		LeftO:                                 QwertyO,
		LeftP:                                 QwertyP,
		LeftK | LeftW:                         QwertyQ,
		LeftR:                                 QwertyR,
		LeftS:                                 QwertyS,
		LeftT:                                 QwertyT,
		RightU:                                QwertyU,
		LeftS | LeftR:                         QwertyV,
		LeftW:                                 QwertyW,
		LeftK | LeftP:                         QwertyX,
		LeftK | LeftW | LeftR:                 QwertyY,
		LeftS | LeftT | LeftK | LeftP | LeftW: QwertyZ,
		LeftS | LeftT | LeftK:                 QwertyZ, // alternate z
	}
}

// isFingerspelling returns true if the receiver matches a Plover fingerspelling
func (k Keymask) isFingerspelling() bool {
	for stroke := range allFingerspellings() {
		if k == stroke {
			return true
		}
	}
	return false
}

func (k *Keymask) UnmarshalJSON(b []byte) error {
	var in string
	if err := json.Unmarshal(b, &in); err != nil {
		return err
	}
	newK, err := ParseStroke(in)
	if err != nil {
		return err
	}
	*k = newK
	return nil
}

func (k *Keymask) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.String())
}
