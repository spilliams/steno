package dictionary

import "fmt"

type NumberOption int

const (
	// NumberOptionDisabled turns off generation for this numbers option
	NumberOptionDisabled NumberOption = NumberOption(iota)
	// NumberOptionNumbers generates 1-5 and 0 for this numbers option
	// (S=1, T=2, P=3, H=4, A=5, O=0)
	NumberOptionNumbers NumberOption = NumberOption(iota)
	// NumberOptionNumbersHigh generates 5-9 and 0 for this numbers option
	// (S=6, T=7, P=8, H=9, A=5, O=0)
	NumberOptionNumbersHigh NumberOption = NumberOption(iota)
	// NumberOptionFunctions generates F1-F5 and F12 for this numbers option
	// (S=F1, T=F2, P=F3, H=F4, A=F5, O=F12)
	NumberOptionFunctions NumberOption = NumberOption(iota)
	// NumberOptionFunctionsHigh generate F6-F11 for this numbers option
	// (S=F6, T=F7, P=F8, H=F9, A=F10, O=F11)
	NumberOptionFunctionsHigh NumberOption = NumberOption(iota)
)

// FactoryOpts represents a set of options for the factory to use during
// generation.
type FactoryOpts struct {
	// NonstandardModCombinations tells the factory to generate strokes for the
	// modifier combinations ctrl-gui, shift-ctrl-gui, ctrl-alt-gui, and
	// shift-ctrl-alt-gui ("hyper")
	NonstandardModCombinations bool
	// Fingerspellings tells the factory to generate strokes for all
	// fingerspellings (with the alteration that * is not included in the
	// fingerspelling, to leave it open for modifier masks)
	Fingerspellings bool
	// NumbersLeft tells the factory to generate strokes using # and S-, T-,
	// P-, H-, A and O
	NumbersLeft NumberOption
	// NumberStarsLeft tells the factory to generate strokes using #* and S-,
	// T-, P-, H-, A and O
	NumberStarsLeft NumberOption
	// TODO: NumbersRight and NumberStarsRight?
}

// Factory allows a caller to generate a dictionary, using certain options.
type Factory struct {
	opts FactoryOpts
}

// NewFactory builds a new dictionary factory. If `opts` is nil, it will use
// DefaultFactoryOpts.
func NewFactory(opts FactoryOpts) *Factory {
	return &Factory{opts}
}

const definitionFmt = "{#%s}{^}{>}"

// Generate tells the receiver to build a dictionary.
// The base set of definitions will include the following:
// - Navigation strokes (layer-esc, layer-space, layer-tab, layer-return, layer-home, layer-pageUp, layer-pageDown, layer-end, layer-backspace, layer-delete, layer-up, layer-down, layer-left, layer-right)
// - Modifier+Navigation strokes (e.g. shift-esc, ctrl-esc, alt-esc, gui-esc, shift-ctrl-esc, shift-alt-esc, shift-gui-esc, ctrl-alt-esc, alt-gui-esc, shift-ctrl-alt-esc, shift-alt-gui-esc)
// Based on the receivers options, the factory may generate more definitions:
// - More combinations of modifiers, applied to all other definitions (ctrl-gui-, shift-ctrl-gui-, ctrl-alt-gui-, shift-ctrl-alt-gui-)
// - Fingerspelling strokes (e.g. shift-S, ctrl-S, alt-S, gui-S, shift-ctrl-S, shift-alt-S, shift-gui-S, ctrl-alt-S, alt-gui-S, shift-ctrl-alt-S, shift-alt-gui-S)
// - Left-hand Number strokes (e.g. shift-1, ctrl-1, alt-1, gui-1, shift-ctrl-1, shift-alt-1, shift-gui-1, ctrl-alt-1, alt-gui-1, shift-ctrl-alt-1, shift-alt-gui-1)
// - Left-hand Function strokes (replaces Left-hand Number strokes 1-5 and 0 with F1-F5 and F12, respectively)
func (f *Factory) Generate(r *Rules) *Dictionary {
	// mods is a map from rules entry (steno keymask) to definition format
	mods := map[Keymask]QwertyMod{
		r.Layer:                  "%s",
		r.Shift:                  Shift,
		r.Ctrl:                   Ctrl,
		r.Alt:                    Alt,
		r.Gui:                    Gui,
		r.Shift | r.Ctrl:         Shift.apply(string(Ctrl)),
		r.Shift | r.Alt:          Shift.apply(string(Alt)),
		r.Shift | r.Gui:          Shift.apply(string(Gui)),
		r.Ctrl | r.Alt:           Ctrl.apply(string(Alt)),
		r.Alt | r.Gui:            Alt.apply(string(Gui)),
		r.Shift | r.Ctrl | r.Alt: Shift.apply(string(Ctrl.apply(string(Alt)))),
		r.Shift | r.Alt | r.Gui:  Shift.apply(string(Alt.apply(string(Gui)))),
	}
	if f.opts.NonstandardModCombinations {
		mods[r.Ctrl|r.Gui] = Ctrl.apply(string(Gui))
		mods[r.Shift|r.Ctrl|r.Gui] = Shift.apply(string(Ctrl.apply(string(Gui))))
		mods[r.Ctrl|r.Alt|r.Gui] = Ctrl.apply(string(Alt.apply(string(Gui))))
		mods[r.Shift|r.Ctrl|r.Alt|r.Gui] = Shift.apply(string(Ctrl.apply(string(Alt.apply(string(Gui))))))
	}
	// keys is a map from rules entry (steno keymask) to definition string
	keys := map[Keymask]QwertyKey{
		r.Escape:    Escape,
		r.Space:     Space,
		r.Tab:       Tab,
		r.Return:    Return,
		r.Home:      Home,
		r.PageUp:    PageUp,
		r.PageDown:  PageDown,
		r.End:       End,
		r.Backspace: Backspace,
		r.Delete:    Delete,
		r.Left:      Left,
		r.Up:        Up,
		r.Down:      Down,
		r.Right:     Right,
	}
	if f.opts.Fingerspellings {
		for k, q := range allFingerspellings() {
			keys[k] = q
		}
	}
	switch f.opts.NumbersLeft {
	case NumberOptionNumbers:
		keys[Steno1] = N1
		keys[Steno2] = N2
		keys[Steno3] = N3
		keys[Steno4] = N4
		keys[Steno5] = N5
		keys[Steno0] = N0
	case NumberOptionNumbersHigh:
		keys[Steno1] = N6
		keys[Steno2] = N7
		keys[Steno3] = N8
		keys[Steno4] = N9
		keys[Steno5] = N5
		keys[Steno0] = N0
	case NumberOptionFunctions:
		keys[Steno1] = F1
		keys[Steno2] = F2
		keys[Steno3] = F3
		keys[Steno4] = F4
		keys[Steno5] = F5
		keys[Steno0] = F12
	case NumberOptionFunctionsHigh:
		keys[Steno1] = F6
		keys[Steno2] = F7
		keys[Steno3] = F8
		keys[Steno4] = F9
		keys[Steno5] = F10
		keys[Steno0] = F11
	}
	switch f.opts.NumberStarsLeft {
	case NumberOptionNumbers:
		keys[Steno1|Star] = N1
		keys[Steno2|Star] = N2
		keys[Steno3|Star] = N3
		keys[Steno4|Star] = N4
		keys[Steno5|Star] = N5
		keys[Steno0|Star] = N0
	case NumberOptionNumbersHigh:
		keys[Steno1|Star] = N6
		keys[Steno2|Star] = N7
		keys[Steno3|Star] = N8
		keys[Steno4|Star] = N9
		keys[Steno5|Star] = N5
		keys[Steno0|Star] = N0
	case NumberOptionFunctions:
		keys[Steno1|Star] = F1
		keys[Steno2|Star] = F2
		keys[Steno3|Star] = F3
		keys[Steno4|Star] = F4
		keys[Steno5|Star] = F5
		keys[Steno0|Star] = F12
	case NumberOptionFunctionsHigh:
		keys[Steno1|Star] = F6
		keys[Steno2|Star] = F7
		keys[Steno3|Star] = F8
		keys[Steno4|Star] = F9
		keys[Steno5|Star] = F10
		keys[Steno0|Star] = F11
	}

	d := make(map[*Brief]string)
	for stenoMod, qwertyMod := range mods {
		for stenoKey, qwertyKey := range keys {
			d[SingleStrokeBrief(stenoMod|stenoKey)] = fmt.Sprintf(definitionFmt, qwertyMod.apply(string(qwertyKey)))
		}
	}

	dict := Dictionary(d)
	return &dict
}
