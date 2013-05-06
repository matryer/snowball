/*
	This package defines a SnowballWord struct that is used
	to encapsulate most of the "state" variables we must track
	when stemming a word.  The SnowballWord struct also has
	a few methods common to stemming in a variety of languages.
*/
package snowballword

// SnowballWord represents a word that is going to be stemmed.
// 
type SnowballWord struct {

	// A slice of runes
	RS []rune

	// The index in RS where the R1 region begins
	R1start int

	// The index in RS where the R2 region begins
	R2start int
}

// Create a new SnowballWord struct
func New(in string) (word *SnowballWord) {
	word = &SnowballWord{RS: []rune(in)}
	word.R1start = len(word.RS)
	word.R2start = len(word.RS)
	return
}

// Replace a suffix and adjust R1start and R2start as needed.
// If `force` is false, check to make sure the suffix exists first.
//
func (w *SnowballWord) ReplaceSuffix(suffix, replacement string, force bool) bool {

	if force || suffix == w.FirstSuffix(suffix) {
		lenWithoutSuffix := len(w.RS) - len(suffix)
		w.RS = append(w.RS[:lenWithoutSuffix], []rune(replacement)...)

		// If R1 and R2 are now beyond the length
		// of the word, they are set to the length
		// of the word.  Otherwise, they are left
		// as they were.
		w.resetR1R2()
		return true
	}
	return false
}

// Resets R1start and R2start to ensure they 
// are within bounds of the current rune slice.
func (w *SnowballWord) resetR1R2() {
	rsLen := len(w.RS)
	if w.R1start > rsLen {
		w.R1start = rsLen
	}
	if w.R2start > rsLen {
		w.R2start = rsLen
	}
}

// Return a slice of w.RS, allowing the start
// and stop to be out of bounds.
//
func (w *SnowballWord) slice(start, stop int) []rune {
	startMin := 0
	if start < startMin {
		start = startMin
	}
	max := len(w.RS) - 1
	if start > max {
		start = max
	}
	if stop > max {
		stop = max
	}
	return w.RS[start:stop]
}

// Return the R1 region as a slice of runes
func (w *SnowballWord) R1() []rune {
	return w.RS[w.R1start:]
}

// Return the R1 region as a string
func (w *SnowballWord) R1String() string {
	return string(w.R1())
}

// Return the R2 region as a slice of runes
func (w *SnowballWord) R2() []rune {
	return w.RS[w.R2start:]
}

// Return the R2 region as a string
func (w *SnowballWord) R2String() string {
	return string(w.R2())
}

// Return the SnowballWord as a string
func (w *SnowballWord) String() string {
	return string(w.RS)
}

// Return the first prefix found or the empty string.
func (w *SnowballWord) FirstPrefix(prefixes ...string) string {
	found := false
	rsLen := len(w.RS)

	for _, prefix := range prefixes {
		found = true
		for i, r := range prefix {
			if i > rsLen-1 || (w.RS)[i] != r {
				found = false
				break
			}
		}
		if found {
			return prefix
		}
	}
	return ""
}

// Return the first suffix found or the empty string.
func (w *SnowballWord) FirstSuffix(sufficies ...string) (suffix string) {
	rsLen := len(w.RS)
	for _, suffix := range sufficies {
		numMatching := 0
		suffixRunes := []rune(suffix)
		suffixLen := len(suffixRunes)
		for i := 0; i < rsLen && i < suffixLen; i++ {
			if w.RS[rsLen-i-1] != suffixRunes[suffixLen-i-1] {
				break
			} else {
				numMatching += 1
			}
		}
		if numMatching == suffixLen {
			return suffix
		}
	}
	return ""
}
