package garbage

import "strings"

var uwuReplacer *strings.Replacer = strings.NewReplacer(
	"r", "w",
	"R", "W",
	"l", "w",
	"L", "W",
	"ove", "uv",
	"th", "ff",
	"na", "ny",
	"ne", "ny",
	"ni", "ny",
	"no", "ny",
	"nu", "ny",
)

// Uwu returns str as a uwu-ified string.
func Uwu(str string) string {
	return uwuReplacer.Replace(str)
}
