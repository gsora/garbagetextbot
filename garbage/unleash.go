package garbage

// List of the string -> emoji associations.
//
// Just add another tuple and OrderedSubstitution will make it available.

const (
	clap = "\xF0\x9F\x91\x8F"
)

var emojis = map[string]string{
	"\xE2\x83\xA3": "",

	"a": "\xF0\x9F\x85\xB0",
	"b": "\xF0\x9F\x85\xB1",
	"c": "\xC2\xA9",
	"e": "\xF0\x9F\x93\xA7",
	"g": "\xF0\x9F\x85\xB1",
	"h": "\xE2\x99\x93",
	"i": "\xE2\x84\xB9",
	"l": "\xF0\x9F\x95\x92",
	"m": "\xE2\x93\x82",
	"o": "\xF0\x9F\x85\xBE",
	"p": "\xF0\x9F\x85\xBF",
	"r": "\xC2\xAE",
	"s": "\xF0\x9F\x92\xB2",
	"t": "\xE2\x9C\x9D\xEF\xB8\x8F",
	"u": "\xE2\x9B\x8E",
	"v": "\xE2\x99\x88",
	"w": "\xF0\x9F\x94\xB1",
	"x": "\xE2\x9D\x8C",
	"y": "\xF0\x9F\x8C\xB1",

	"?":  "\xE2\x9D\x93",
	"!":  "\xE2\x9D\x97",
	"!?": "\xE2\x81\x89",
	"!!": "\xE2\x80\xBC",

	"1":   "\x31\xE2\x83\xA3",
	"2":   "\x32\xE2\x83\xA3",
	"3":   "\x33\xE2\x83\xA3",
	"4":   "\x34\xE2\x83\xA3",
	"5":   "\x35\xE2\x83\xA3",
	"6":   "\x36\xE2\x83\xA3",
	"7":   "\x37\xE2\x83\xA3",
	"8":   "\x38\xE2\x83\xA3",
	"9":   "\x39\xE2\x83\xA3",
	"0":   "\x30\xE2\x83\xA3",
	"100": "\xF0\x9F\x92\xAF",

	"ab":      "\xF0\x9F\x86\x8E",
	"allah":   "\xF0\x9F\x91\xB3 \xF0\x9F\x92\xA3",
	"cl":      "\xF0\x9F\x86\x91",
	"dio":     "\xF0\x9F\x90\x96",
	"id":      "\xF0\x9F\x86\x94",
	"docker":  "\xF0\x9F\x90\xB3",
	"ngul":    "\xF0\x9F\x86\x96 \xF0\x9F\x86\x92",
	"ok":      "\xF0\x9F\x86\x97",
	"peroni":  "\xF0\x9F\x8D\xBA",
	"2peroni": "\xF0\x9F\x8D\xBB",
	"sos":     "\xF0\x9F\x86\x98",
	"nsfw":    "\xF0\x9F\x94\x9E",
	"snek":    "\xF0\x9F\x90\x8D",
	"vs":      "\xF0\x9F\x86\x9A",
	"leone":   "\xE2\x99\x8C",
	"shit":    "\xF0\x9F\x92\xA9",
}