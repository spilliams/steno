package dictionary

import (
	"encoding/json"
	"testing"
)

func TestRulesMustUnmarshal(t *testing.T) {
	jsonFile := `{
		"escape": "SKP",
		"space": "SP",
		"tab": "TPW",
		"return": "TRE",
		"home": "PWH",
		"pageUp": "TKPWU",
		"pageDown": "TKPWH",
		"end": "TKW",
		"backspace": "KPW",
		"delete": "PWR",
		"up": "PU",
		"down": "TKPH",
		"left": "TPHRE",
		"right": "TREU",
		"layer": "-FRLG",
		"shift": "-FRPLG",
		"ctrl": "-FRLGTS",
		"alt": "-FRBLG",
		"gui": "-FRLGDZ"
	}`
	expected := Rules{
		Escape:    LeftS | LeftK | LeftP,
		Space:     LeftS | LeftP,
		Tab:       LeftT | LeftP | LeftW,
		Return:    LeftT | LeftR | RightE,
		Home:      LeftP | LeftW | LeftH,
		PageUp:    LeftT | LeftK | LeftP | LeftW | RightU,
		PageDown:  LeftT | LeftK | LeftP | LeftW | LeftH,
		End:       LeftT | LeftK | LeftW,
		Backspace: LeftK | LeftP | LeftW,
		Delete:    LeftP | LeftW | LeftR,
		Up:        LeftP | RightU,
		Down:      LeftT | LeftK | LeftP | LeftH,
		Left:      LeftT | LeftP | LeftH | LeftR | RightE,
		Right:     LeftT | LeftR | RightE | RightU,
		Layer:     RightF | RightR | RightL | RightG,
		Shift:     RightF | RightR | RightP | RightL | RightG,
		Ctrl:      RightF | RightR | RightL | RightG | RightT | RightS,
		Alt:       RightF | RightR | RightB | RightL | RightG,
		Gui:       RightF | RightR | RightL | RightG | RightD | RightZ,
	}

	actual := Rules{}
	if err := json.Unmarshal([]byte(jsonFile), &actual); err != nil {
		t.Fatal(err)
	}
	t.Log("expected rules:")
	t.Log(expected)
	t.Log("actual rules:")
	t.Log(actual)
	if expected != actual {
		t.Fatal("Actual rules do not match expected")
	}
}

func TestRulesMustBeValid(t *testing.T) {
	cases := []struct {
		name  string
		rJSON string
		errs  []string
	}{
		{
			name: "happy path",
			rJSON: `{
				"escape": "SKP",
				"space": "SP",
				"tab": "TPW",
				"return": "TRE",
				"home": "PWH",
				"pageUp": "TKPWU",
				"pageDown": "TKPWH",
				"end": "TKW",
				"backspace": "KPW",
				"delete": "PWR",
				"up": "PU",
				"down": "TKPH",
				"left": "TPHRE",
				"right": "TREU",
				"layer": "-FRLG",
				"shift": "-FRPLG",
				"ctrl": "*FRLG",
				"alt": "-FRBLG",
				"gui": "-FRLGTS"
			}`,
			errs: []string{},
		},
		{
			name: "overlapping modifiers",
			rJSON: `{
				"escape": "SKP",
				"space": "SP",
				"tab": "TPW",
				"return": "TRE",
				"home": "PWH",
				"pageUp": "TKPWU",
				"pageDown": "TKPWH",
				"end": "TKW",
				"backspace": "KPW",
				"delete": "PWR",
				"up": "PU",
				"down": "TKPH",
				"left": "TPHRE",
				"right": "TREU",
				"layer": "-FRLG",
				"shift": "-FRPLG",
				"ctrl": "*FRLG",
				"alt": "-FRLGTS",
				"gui": "-TSDZ"
			}`,
			errs: []string{
				"Masks for shift-gui and shift-alt-gui must not be the same (-FRPLGTSDZ)",
				"Masks for ctrl-gui and ctrl-alt-gui must not be the same (*FRLGTSDZ)",
				"Masks for shift-ctrl-gui and shift-ctrl-alt-gui must not be the same (*FRPLGTSDZ)",
				"Masks for shift-gui+escape and shift-alt-gui+escape must not be the same (SKP-FRPLGTSDZ)",
				"Masks for shift-gui+space and shift-alt-gui+space must not be the same (SP-FRPLGTSDZ)",
				"Masks for shift-gui+tab and shift-alt-gui+tab must not be the same (TPW-FRPLGTSDZ)",
				"Masks for shift-gui+return and shift-alt-gui+return must not be the same (TREFRPLGTSDZ)",
				"Masks for shift-gui+home and shift-alt-gui+home must not be the same (PWH-FRPLGTSDZ)",
				"Masks for shift-gui+pageUp and shift-alt-gui+pageUp must not be the same (TKPWUFRPLGTSDZ)",
				"Masks for shift-gui+pageDown and shift-alt-gui+pageDown must not be the same (TKPWH-FRPLGTSDZ)",
				"Masks for shift-gui+end and shift-alt-gui+end must not be the same (TKW-FRPLGTSDZ)",
				"Masks for shift-gui+backspace and shift-alt-gui+backspace must not be the same (KPW-FRPLGTSDZ)",
				"Masks for shift-gui+delete and shift-alt-gui+delete must not be the same (PWR-FRPLGTSDZ)",
				"Masks for shift-gui+up and shift-alt-gui+up must not be the same (PUFRPLGTSDZ)",
				"Masks for shift-gui+down and shift-alt-gui+down must not be the same (TKPH-FRPLGTSDZ)",
				"Masks for shift-gui+left and shift-alt-gui+left must not be the same (TPHREFRPLGTSDZ)",
				"Masks for shift-gui+right and shift-alt-gui+right must not be the same (TREUFRPLGTSDZ)",
				"Masks for shift-alt-gui+escape and shift-gui+escape must not be the same (SKP-FRPLGTSDZ)",
				"Masks for shift-alt-gui+space and shift-gui+space must not be the same (SP-FRPLGTSDZ)",
				"Masks for shift-alt-gui+tab and shift-gui+tab must not be the same (TPW-FRPLGTSDZ)",
				"Masks for shift-alt-gui+return and shift-gui+return must not be the same (TREFRPLGTSDZ)",
				"Masks for shift-alt-gui+home and shift-gui+home must not be the same (PWH-FRPLGTSDZ)",
				"Masks for shift-alt-gui+pageUp and shift-gui+pageUp must not be the same (TKPWUFRPLGTSDZ)",
				"Masks for shift-alt-gui+pageDown and shift-gui+pageDown must not be the same (TKPWH-FRPLGTSDZ)",
				"Masks for shift-alt-gui+end and shift-gui+end must not be the same (TKW-FRPLGTSDZ)",
				"Masks for shift-alt-gui+backspace and shift-gui+backspace must not be the same (KPW-FRPLGTSDZ)",
				"Masks for shift-alt-gui+delete and shift-gui+delete must not be the same (PWR-FRPLGTSDZ)",
				"Masks for shift-alt-gui+up and shift-gui+up must not be the same (PUFRPLGTSDZ)",
				"Masks for shift-alt-gui+down and shift-gui+down must not be the same (TKPH-FRPLGTSDZ)",
				"Masks for shift-alt-gui+left and shift-gui+left must not be the same (TPHREFRPLGTSDZ)",
				"Masks for shift-alt-gui+right and shift-gui+right must not be the same (TREUFRPLGTSDZ)",
				"Masks for ctrl-gui+escape and ctrl-alt-gui+escape must not be the same (SKP*FRLGTSDZ)",
				"Masks for ctrl-gui+space and ctrl-alt-gui+space must not be the same (SP*FRLGTSDZ)",
				"Masks for ctrl-gui+tab and ctrl-alt-gui+tab must not be the same (TPW*FRLGTSDZ)",
				"Masks for ctrl-gui+return and ctrl-alt-gui+return must not be the same (TR*EFRLGTSDZ)",
				"Masks for ctrl-gui+home and ctrl-alt-gui+home must not be the same (PWH*FRLGTSDZ)",
				"Masks for ctrl-gui+pageUp and ctrl-alt-gui+pageUp must not be the same (TKPW*UFRLGTSDZ)",
				"Masks for ctrl-gui+pageDown and ctrl-alt-gui+pageDown must not be the same (TKPWH*FRLGTSDZ)",
				"Masks for ctrl-gui+end and ctrl-alt-gui+end must not be the same (TKW*FRLGTSDZ)",
				"Masks for ctrl-gui+backspace and ctrl-alt-gui+backspace must not be the same (KPW*FRLGTSDZ)",
				"Masks for ctrl-gui+delete and ctrl-alt-gui+delete must not be the same (PWR*FRLGTSDZ)",
				"Masks for ctrl-gui+up and ctrl-alt-gui+up must not be the same (P*UFRLGTSDZ)",
				"Masks for ctrl-gui+down and ctrl-alt-gui+down must not be the same (TKPH*FRLGTSDZ)",
				"Masks for ctrl-gui+left and ctrl-alt-gui+left must not be the same (TPHR*EFRLGTSDZ)",
				"Masks for ctrl-gui+right and ctrl-alt-gui+right must not be the same (TR*EUFRLGTSDZ)",
				"Masks for shift-ctrl-gui+escape and shift-ctrl-alt-gui+escape must not be the same (SKP*FRPLGTSDZ)",
				"Masks for shift-ctrl-gui+space and shift-ctrl-alt-gui+space must not be the same (SP*FRPLGTSDZ)",
				"Masks for shift-ctrl-gui+tab and shift-ctrl-alt-gui+tab must not be the same (TPW*FRPLGTSDZ)",
				"Masks for shift-ctrl-gui+return and shift-ctrl-alt-gui+return must not be the same (TR*EFRPLGTSDZ)",
				"Masks for shift-ctrl-gui+home and shift-ctrl-alt-gui+home must not be the same (PWH*FRPLGTSDZ)",
				"Masks for shift-ctrl-gui+pageUp and shift-ctrl-alt-gui+pageUp must not be the same (TKPW*UFRPLGTSDZ)",
				"Masks for shift-ctrl-gui+pageDown and shift-ctrl-alt-gui+pageDown must not be the same (TKPWH*FRPLGTSDZ)",
				"Masks for shift-ctrl-gui+end and shift-ctrl-alt-gui+end must not be the same (TKW*FRPLGTSDZ)",
				"Masks for shift-ctrl-gui+backspace and shift-ctrl-alt-gui+backspace must not be the same (KPW*FRPLGTSDZ)",
				"Masks for shift-ctrl-gui+delete and shift-ctrl-alt-gui+delete must not be the same (PWR*FRPLGTSDZ)",
				"Masks for shift-ctrl-gui+up and shift-ctrl-alt-gui+up must not be the same (P*UFRPLGTSDZ)",
				"Masks for shift-ctrl-gui+down and shift-ctrl-alt-gui+down must not be the same (TKPH*FRPLGTSDZ)",
				"Masks for shift-ctrl-gui+left and shift-ctrl-alt-gui+left must not be the same (TPHR*EFRPLGTSDZ)",
				"Masks for shift-ctrl-gui+right and shift-ctrl-alt-gui+right must not be the same (TR*EUFRPLGTSDZ)",
				"Masks for ctrl-alt-gui+escape and ctrl-gui+escape must not be the same (SKP*FRLGTSDZ)",
				"Masks for ctrl-alt-gui+space and ctrl-gui+space must not be the same (SP*FRLGTSDZ)",
				"Masks for ctrl-alt-gui+tab and ctrl-gui+tab must not be the same (TPW*FRLGTSDZ)",
				"Masks for ctrl-alt-gui+return and ctrl-gui+return must not be the same (TR*EFRLGTSDZ)",
				"Masks for ctrl-alt-gui+home and ctrl-gui+home must not be the same (PWH*FRLGTSDZ)",
				"Masks for ctrl-alt-gui+pageUp and ctrl-gui+pageUp must not be the same (TKPW*UFRLGTSDZ)",
				"Masks for ctrl-alt-gui+pageDown and ctrl-gui+pageDown must not be the same (TKPWH*FRLGTSDZ)",
				"Masks for ctrl-alt-gui+end and ctrl-gui+end must not be the same (TKW*FRLGTSDZ)",
				"Masks for ctrl-alt-gui+backspace and ctrl-gui+backspace must not be the same (KPW*FRLGTSDZ)",
				"Masks for ctrl-alt-gui+delete and ctrl-gui+delete must not be the same (PWR*FRLGTSDZ)",
				"Masks for ctrl-alt-gui+up and ctrl-gui+up must not be the same (P*UFRLGTSDZ)",
				"Masks for ctrl-alt-gui+down and ctrl-gui+down must not be the same (TKPH*FRLGTSDZ)",
				"Masks for ctrl-alt-gui+left and ctrl-gui+left must not be the same (TPHR*EFRLGTSDZ)",
				"Masks for ctrl-alt-gui+right and ctrl-gui+right must not be the same (TR*EUFRLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+escape and shift-ctrl-gui+escape must not be the same (SKP*FRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+space and shift-ctrl-gui+space must not be the same (SP*FRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+tab and shift-ctrl-gui+tab must not be the same (TPW*FRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+return and shift-ctrl-gui+return must not be the same (TR*EFRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+home and shift-ctrl-gui+home must not be the same (PWH*FRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+pageUp and shift-ctrl-gui+pageUp must not be the same (TKPW*UFRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+pageDown and shift-ctrl-gui+pageDown must not be the same (TKPWH*FRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+end and shift-ctrl-gui+end must not be the same (TKW*FRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+backspace and shift-ctrl-gui+backspace must not be the same (KPW*FRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+delete and shift-ctrl-gui+delete must not be the same (PWR*FRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+up and shift-ctrl-gui+up must not be the same (P*UFRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+down and shift-ctrl-gui+down must not be the same (TKPH*FRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+left and shift-ctrl-gui+left must not be the same (TPHR*EFRPLGTSDZ)",
				"Masks for shift-ctrl-alt-gui+right and shift-ctrl-gui+right must not be the same (TR*EUFRPLGTSDZ)",
			},
		},
		{
			name: "fingerspellings 1",
			rJSON: `{
				"escape": "A",
				"space": "PW",
				"tab": "KR",
				"return": "E",
				"home": "TP",
				"pageUp": "TKPW",
				"pageDown": "H",
				"end": "EU",
				"backspace": "SKWR",
				"delete": "K",
				"up": "HR",
				"down": "PH",
				"left": "TPH",
				"right": "O",
				"layer": "-FRLG",
				"shift": "-FRPLG",
				"ctrl": "*FRLG",
				"alt": "-FRBLG",
				"gui": "-FRLGTS"
			}`,
			errs: []string{
				"Mask for escape matches a fingerspelling (A)",
				"Mask for space matches a fingerspelling (PW)",
				"Mask for tab matches a fingerspelling (KR)",
				"Mask for return matches a fingerspelling (E)",
				"Mask for home matches a fingerspelling (TP)",
				"Mask for pageUp matches a fingerspelling (TKPW)",
				"Mask for pageDown matches a fingerspelling (H)",
				"Mask for end matches a fingerspelling (EU)",
				"Mask for backspace matches a fingerspelling (SKWR)",
				"Mask for delete matches a fingerspelling (K)",
				"Mask for up matches a fingerspelling (HR)",
				"Mask for down matches a fingerspelling (PH)",
				"Mask for left matches a fingerspelling (TPH)",
				"Mask for right matches a fingerspelling (O)",
			},
		},
		{
			name: "fingerspellings 2",
			rJSON: `{
				"escape": "P",
				"space": "KW",
				"tab": "R",
				"return": "S",
				"home": "T",
				"pageUp": "U",
				"pageDown": "SR",
				"end": "W",
				"backspace": "KP",
				"delete": "KWR",
				"up": "STKPW",
				"down": "STK",
				"left": "STKH",
				"right": "STKR",
				"layer": "-FRLG",
				"shift": "-FRPLG",
				"ctrl": "*FRLG",
				"alt": "-FRBLG",
				"gui": "-FRLGTS"
			}`,
			errs: []string{
				"Mask for escape matches a fingerspelling (P)",
				"Mask for space matches a fingerspelling (KW)",
				"Mask for tab matches a fingerspelling (R)",
				"Mask for return matches a fingerspelling (S)",
				"Mask for home matches a fingerspelling (T)",
				"Mask for pageUp matches a fingerspelling (U)",
				"Mask for pageDown matches a fingerspelling (SR)",
				"Mask for end matches a fingerspelling (W)",
				"Mask for backspace matches a fingerspelling (KP)",
				"Mask for delete matches a fingerspelling (KWR)",
				"Mask for up matches a fingerspelling (STKPW)",
				"Mask for down matches a fingerspelling (STK)",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := &Rules{}
			if err := json.Unmarshal([]byte(c.rJSON), r); err != nil {
				t.Fatalf("couldn't unmarshal rules: %v", err)
			}
			errs := r.MustBeValid()
			t.Log("expected errors:")
			for _, err := range c.errs {
				t.Log(err)
			}
			t.Log("actual errors:")
			for _, err := range errs {
				t.Log(err.Error())
			}
			if len(errs) != len(c.errs) {
				t.Fatalf("expected %d errors, got %d", len(c.errs), len(errs))
			}
			for i, expected := range c.errs {
				actual := errs[i].Error()
				if expected != actual {
					t.Errorf("expected %dth error to be '%s', got '%s'", i, expected, actual)
				}
			}
		})
	}
}
