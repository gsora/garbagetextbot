package garbage

import "strings"

type nibber struct {
	substitutions orderedSubstitution
	replacer      strings.Replacer
}

var n nibber

func init() {
	subs := mapToOrderedSubstitution(emojis)

	n = nibber{
		substitutions: subs,
		replacer:      *strings.NewReplacer(subs.toReplacerArray()...),
	}
}

// Nibber replaces each character in str with its emoji representation, defined in `nibber.emojis`.
func Nibber(str string) string {
	return n.replacer.Replace(strings.ToUpper(str))
}
