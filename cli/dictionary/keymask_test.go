package dictionary

import "testing"

func TestParseKeymask(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		mask   Keymask
		output string
	}{
		{
			name:   "S should be S",
			input:  "S",
			mask:   0b01000000000000000000000,
			output: "S",
		},
		{
			name:   "T should be T",
			input:  "T",
			mask:   0b00100000000000000000000,
			output: "T",
		},
		{
			name:   "K should be K",
			input:  "K",
			mask:   0b00010000000000000000000,
			output: "K",
		},
		{
			name:   "P should be P",
			input:  "P",
			mask:   0b00001000000000000000000,
			output: "P",
		},
		{
			name:   "W should be W",
			input:  "W",
			mask:   0b00000100000000000000000,
			output: "W",
		},
		{
			name:   "H should be H",
			input:  "H",
			mask:   0b00000010000000000000000,
			output: "H",
		},
		{
			name:   "R should be R",
			input:  "R",
			mask:   0b00000001000000000000000,
			output: "R",
		},
		{
			name:   "A should be A",
			input:  "A",
			mask:   0b00000000100000000000000,
			output: "A",
		},
		{
			name:   "O should be O",
			input:  "O",
			mask:   0b00000000010000000000000,
			output: "O",
		},
		{
			name:   "* should be *",
			input:  "*",
			mask:   0b00000000001000000000000,
			output: "*",
		},
		{
			name:   "E should be E",
			input:  "E",
			mask:   0b00000000000100000000000,
			output: "E",
		},
		{
			name:   "U should be U",
			input:  "U",
			mask:   0b00000000000010000000000,
			output: "U",
		},
		{
			name:   "-F should be -F",
			input:  "-F",
			mask:   0b00000000000001000000000,
			output: "-F",
		},
		{
			name:   "-R should be -R",
			input:  "-R",
			mask:   0b00000000000000100000000,
			output: "-R",
		},
		{
			name:   "-P should be -P",
			input:  "-P",
			mask:   0b00000000000000010000000,
			output: "-P",
		},
		{
			name:   "-B should be -B",
			input:  "-B",
			mask:   0b00000000000000001000000,
			output: "-B",
		},
		{
			name:   "-L should be -L",
			input:  "-L",
			mask:   0b00000000000000000100000,
			output: "-L",
		},
		{
			name:   "-G should be -G",
			input:  "-G",
			mask:   0b00000000000000000010000,
			output: "-G",
		},
		{
			name:   "-T should be -T",
			input:  "-T",
			mask:   0b00000000000000000001000,
			output: "-T",
		},
		{
			name:   "-S should be -S",
			input:  "-S",
			mask:   0b00000000000000000000100,
			output: "-S",
		},
		{
			name:   "-D should be -D",
			input:  "-D",
			mask:   0b00000000000000000000010,
			output: "-D",
		},
		{
			name:   "-Z should be -Z",
			input:  "-Z",
			mask:   0b00000000000000000000001,
			output: "-Z",
		},
		{
			name:   "a simple word",
			input:  "R-R",
			mask:   0b00000001000000100000000,
			output: "R-R",
		},
		{
			name:   "same simple word, with star",
			input:  "R*R",
			mask:   0b00000001001000100000000,
			output: "R*R",
		},
		{
			name:   "number input",
			input:  "4-6",
			mask:   0b10000010000001000000000,
			output: "4-6",
		},
		{
			name:   "incorrect number input should come out correct",
			input:  "#H-F",
			mask:   0b10000010000001000000000,
			output: "4-6",
		},
		{
			name:   "number input with no numbers present",
			input:  "#-D",
			mask:   0b10000000000000000000010,
			output: "#-D",
		},
		{
			name:   "number with only a star",
			input:  "#*",
			mask:   0b10000000001000000000000,
			output: "#*",
		},
		{
			name:   "number by itself",
			input:  "#",
			mask:   0b10000000000000000000000,
			output: "#",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			actual, err := ParseStroke(c.input)
			if err != nil {
				t.Errorf("ParseKeymask returned unexpected error: %v", err)
			}
			if actual != c.mask {
				t.Errorf("Expected %#b (%s), got %#b (%s)", c.mask, c.mask, actual, actual)
			}
			if actual.String() != c.output {
				t.Errorf("Expected printed string to be %s, got %s", c.output, actual)
			}
		})
	}
}
