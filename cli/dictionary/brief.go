package dictionary

import "strings"

type Brief struct {
	strokes []Keymask
}

func SingleStrokeBrief(k Keymask) *Brief {
	return &Brief{
		strokes: []Keymask{k},
	}
}

const separator = "/"

func (b *Brief) String() string {
	masks := make([]string, len(b.strokes))
	for i, stroke := range b.strokes {
		masks[i] = stroke.String()
	}
	return strings.Join(masks, separator)
}

func (b *Brief) isEqual(other *Brief) bool {
	if len(b.strokes) != len(other.strokes) {
		return false
	}
	for _, stroke := range b.strokes {
		for _, otherStroke := range other.strokes {
			if stroke != otherStroke {
				return false
			}
		}
	}
	return true
}

func ParseBrief(in string) (*Brief, error) {
	strokes := strings.Split(in, separator)
	masks := make([]Keymask, len(strokes))
	for i, stroke := range strokes {
		mask, err := ParseStroke(stroke)
		if err != nil {
			return nil, err
		}
		masks[i] = mask
	}
	return &Brief{masks}, nil
}
