package dictionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/apex/log"
)

type Rules struct {
	Escape    Keymask
	Space     Keymask
	Tab       Keymask
	Return    Keymask
	Home      Keymask
	PageUp    Keymask
	PageDown  Keymask
	End       Keymask
	Backspace Keymask
	Delete    Keymask
	Up        Keymask
	Down      Keymask
	Left      Keymask
	Right     Keymask
	Layer     Keymask
	Shift     Keymask
	Ctrl      Keymask
	Alt       Keymask
	Gui       Keymask
}

func (r *Rules) MustBeValid() []error {
	errs := make([]error, 0)
	keymaskNames := []string{
		"escape",
		"space",
		"tab",
		"return",
		"home",
		"pageUp",
		"pageDown",
		"end",
		"backspace",
		"delete",
		"up",
		"down",
		"left",
		"right",
	}
	modmaskNames := []string{
		"layer",
		"shift",
		"ctrl",
		"alt",
		"gui",
		"shift-ctrl",
		"shift-alt",
		"shift-gui",
		"ctrl-alt",
		"alt-gui",
		"shift-ctrl-alt",
		"shift-alt-gui",
		// the below aren't part of Single Stroke Commands
		"ctrl-gui",
		"shift-ctrl-gui",
		"ctrl-alt-gui",
		"shift-ctrl-alt-gui",
	}
	keymasks := []Keymask{
		r.Escape,
		r.Space,
		r.Tab,
		r.Return,
		r.Home,
		r.PageUp,
		r.PageDown,
		r.End,
		r.Backspace,
		r.Delete,
		r.Up,
		r.Down,
		r.Left,
		r.Right,
	}
	modmasks := []Keymask{
		r.Layer,
		r.Shift,
		r.Ctrl,
		r.Alt,
		r.Gui,
		r.Shift | r.Ctrl,
		r.Shift | r.Alt,
		r.Shift | r.Gui,
		r.Ctrl | r.Alt,
		r.Alt | r.Gui,
		r.Shift | r.Ctrl | r.Alt,
		r.Shift | r.Alt | r.Gui,
		// the below aren't part of Single Stroke Commands
		r.Ctrl | r.Gui,
		r.Shift | r.Ctrl | r.Gui,
		r.Ctrl | r.Alt | r.Gui,
		r.Shift | r.Ctrl | r.Alt | r.Gui,
	}
	// nothing can be empty
	for i, m := range keymasks {
		if m == 0 {
			errs = append(errs, fmt.Errorf("Mask for %s must not be blank", keymaskNames[i]))
		}
	}
	for i, m := range modmasks {
		if m == 0 {
			errs = append(errs, fmt.Errorf("Mask for %s must not be blank", modmaskNames[i]))
		}
	}
	// no two key masks can be the same
	for i, m := range keymasks {
		for j := i + 1; j < len(keymasks); j++ {
			if m == keymasks[j] {
				errs = append(errs, fmt.Errorf("Masks for %s and %s must not be the same (%s)", keymaskNames[i], keymaskNames[j], m))
			}
		}
	}
	// no two mod masks can be the same
	for i, m := range modmasks {
		for j := i + 1; j < len(modmasks); j++ {
			if m == modmasks[j] {
				errs = append(errs, fmt.Errorf("Masks for %s and %s must not be the same (%s)", modmaskNames[i], modmaskNames[j], m))
			}
		}
	}
	// no mod+key combination can be the same as another mod
	for i, m := range modmasks {
		for j, n := range keymasks {
			for k, o := range modmasks {
				if m|n == o {
					errs = append(errs, fmt.Errorf("Masks for %s+%s and %s must not be the same (%s)", modmaskNames[i], keymaskNames[j], modmaskNames[k], o))
				}
			}
		}
	}
	// no mod+key combination can be the same as another key
	for i, m := range modmasks {
		for j, n := range keymasks {
			for k, o := range keymasks {
				if m|n == o {
					errs = append(errs, fmt.Errorf("Masks for %s+%s and %s must not be the same (%s)", modmaskNames[i], keymaskNames[j], keymaskNames[k], o))
				}
			}
		}
	}
	// no two mod+key combinations can be the same
	for i, m := range modmasks {
		for j, n := range keymasks {
			for k, o := range modmasks {
				for l, p := range keymasks {
					if i == k && j == l {
						continue
					}
					if m|n == o|p {
						errs = append(errs, fmt.Errorf("Masks for %s+%s and %s+%s must not be the same (%s)", modmaskNames[i], keymaskNames[j], modmaskNames[k], keymaskNames[l], Keymask(m|n)))
					}
				}
			}
		}
	}
	// nothing can match a fingerspelling
	checkedKeymasks := false
	for i, m := range modmasks {
		if m.IsFingerspelling() {
			errs = append(errs, fmt.Errorf("Mask for %s matches a fingerspelling (%s)", modmaskNames[i], m))
		}
		for j, n := range keymasks {
			if !checkedKeymasks && n.IsFingerspelling() {
				errs = append(errs, fmt.Errorf("Mask for %s matches a fingerspelling (%s)", keymaskNames[j], n))
			}
			if (m | n).IsFingerspelling() {
				errs = append(errs, fmt.Errorf("Mask for %s+%s matches a fingerspelling (%s)", modmaskNames[i], keymaskNames[j], m|n))
			}
		}
		checkedKeymasks = true
	}

	return errs
}

func (r *Rules) UnmarshalJSON(b []byte) error {
	var stringMap map[string]string
	if err := json.Unmarshal(b, &stringMap); err != nil {
		return err
	}
	newRules := Rules{}
	for k, v := range stringMap {
		stroke, err := ParseStroke(v)
		if err != nil {
			return err
		}
		log.WithFields(log.Fields{
			"key":    k,
			"stroke": stroke,
		}).Debug("parsed")
		switch strings.ToLower(k) {
		case "escape":
			newRules.Escape = stroke
		case "space":
			newRules.Space = stroke
		case "tab":
			newRules.Tab = stroke
		case "return":
			newRules.Return = stroke
		case "home":
			newRules.Home = stroke
		case "pageup":
			newRules.PageUp = stroke
		case "pagedown":
			newRules.PageDown = stroke
		case "end":
			newRules.End = stroke
		case "backspace":
			newRules.Backspace = stroke
		case "delete":
			newRules.Delete = stroke
		case "up":
			newRules.Up = stroke
		case "down":
			newRules.Down = stroke
		case "left":
			newRules.Left = stroke
		case "right":
			newRules.Right = stroke
		case "layer":
			newRules.Layer = stroke
		case "shift":
			newRules.Shift = stroke
		case "ctrl":
			newRules.Ctrl = stroke
		case "alt":
			newRules.Alt = stroke
		case "gui":
			newRules.Gui = stroke
		}
	}
	*r = newRules
	return nil
}

func ReadRulesFile(filename string) (*Rules, error) {
	inBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	r := Rules{}
	if err = json.Unmarshal(inBytes, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
