package unidecode

import (
	"github.com/hscells/go-unidecode/table"
	"unicode"
)

// Unidecode implements transliterate Unicode text into plain 7-bit ASCII.
// e.g. Unidecode("kožušček") => "kozuscek"
func Unidecode(s string) []uint8 {
	ret := make([]uint8, 0, len(s))
	for i, r := range s {
		if r < unicode.MaxASCII {
			ret = append(ret, s[i])
			continue
		}

		if int32(r) > 0xeffff {
			continue
		}

		section := r >> 8   // Chop off the last two hex digits
		position := r % 256 // Last two hex digits
		if tb, ok := table.Tables[section]; ok {
			if len(tb) > int(position) {
				ret = append(ret, tb[position]...)
			}
		}
	}

	return ret
}
