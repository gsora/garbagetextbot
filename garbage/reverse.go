package garbage

// Get a string and reverse it
func Reverse(text string) string {
    characters := []rune(text)
    charsLength := len(characters)
    
    // Single char cannot be reversed duh!
    if charsLength <= 1 {
        return text
    }

    i, j := 0, 0
    for i < charsLength && j < charsLength {
        j = i + 1
        for j < charsLength && isMark(characters[j]) {
            j++
        }

        if isAMarker(characters[j-1]) {
            // Reverse combined chars
            reverse(characters[i:j], j-i)
        } 

        i = j
    }

    // Reverse it
    reverse(characters, charsLength)

    return string(characters)
}

func reverse(runes []rune, length int) {
    for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
}

func isAMarker(r rune) bool {
    return unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Mc, r)
}
