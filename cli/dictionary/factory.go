package dictionary

import "fmt"

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
	// Numbers tells the factory to generate strokes using # and S-, T-, P-, H-,
	// A and O.
	NumbersLeft bool
	// NumbersAsFunctions tells the factory that number strokes should be the
	// F1-F5 keys (and with `#O` meaning F12).
	NumbersAsFunctions bool
}

// DefaultFactoryOpts represents the default options for a factory. It enables
// non-standard modifier combinations, fingerspellings, and left-hand numbers.
var DefaultFactoryOpts = &FactoryOpts{
	NonstandardModCombinations: true,
	Fingerspellings:            true,
	NumbersLeft:                true,
	NumbersAsFunctions:         false,
}

// Factory allows a caller to generate a dictionary, using certain options.
type Factory struct {
	opts *FactoryOpts
}

// NewFactory builds a new dictionary factory. If `opts` is nil, it will use
// DefaultFactoryOpts.
func NewFactory(opts *FactoryOpts) *Factory {
	if opts == nil {
		opts = DefaultFactoryOpts
	}
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
	if f.opts.NumbersLeft {
		if f.opts.NumbersAsFunctions {
			keys[Steno1] = F1
			keys[Steno2] = F2
			keys[Steno3] = F3
			keys[Steno4] = F4
			keys[Steno5] = F5
			keys[Steno0] = F12
		} else {
			keys[Steno1] = N1
			keys[Steno2] = N2
			keys[Steno3] = N3
			keys[Steno4] = N4
			keys[Steno5] = N5
			keys[Steno0] = N0
		}
	}

	d := make(map[Keymask]string)
	for stenoMod, qwertyMod := range mods {
		for stenoKey, qwertyKey := range keys {
			d[stenoMod|stenoKey] = fmt.Sprintf(definitionFmt, qwertyMod.apply(string(qwertyKey)))
		}
	}

	dict := Dictionary(d)
	return &dict
}
