package garbage

import "strings"

// UpDown returns str as a sequence of upper/lower characters.
func UpDown(str string) (f string) {
	what := false
	for _, c := range str {
		if what {
			f = f + strings.ToUpper(string(c))
			what = false
		} else {
			f = f + strings.ToLower(string(c))
			what = true
		}
	}

	return
}
