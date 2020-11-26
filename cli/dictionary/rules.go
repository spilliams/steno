package dictionary

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
}

// TODO: read rules from json file
// TODO: validation (a valid rules cannot have two modifiers sharing the same keymask)
