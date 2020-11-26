package dictionary

import "testing"

func TestParseKeymask(t *testing.T) {
	cases := map[string]Keymask{
		"S":   0b1000000000000000000000,
		"T":   0b0100000000000000000000,
		"K":   0b0010000000000000000000,
		"P":   0b0001000000000000000000,
		"W":   0b0000100000000000000000,
		"H":   0b0000010000000000000000,
		"R":   0b0000001000000000000000,
		"A":   0b0000000100000000000000,
		"O":   0b0000000010000000000000,
		"*":   0b0000000001000000000000,
		"E":   0b0000000000100000000000,
		"U":   0b0000000000010000000000,
		"-F":  0b0000000000001000000000,
		"-R":  0b0000000000000100000000,
		"-P":  0b0000000000000010000000,
		"-B":  0b0000000000000001000000,
		"-L":  0b0000000000000000100000,
		"-G":  0b0000000000000000010000,
		"-T":  0b0000000000000000001000,
		"-S":  0b0000000000000000000100,
		"-D":  0b0000000000000000000010,
		"-Z":  0b0000000000000000000001,
		"R-R": 0b0000001000000100000000,
		"R*R": 0b0000001001000100000000,
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			actual, err := ParseKeymask(input)
			if err != nil {
				t.Errorf("ParseKeymask returned unexpected error: %v", err)
			}
			if actual != expected {
				t.Errorf("Expected %#b (%s), got %#b (%s)", expected, expected, actual, actual)
			}
		})
	}
}
