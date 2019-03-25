// Package envy automatically exposes environment
// variables for all of your flags.
package envy

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Parse takes a prefix string and exposes environment variables
// for all flags in the default FlagSet (flag.CommandLine) in the
// form of PREFIX_FLAGNAME.
func Parse(prefix string) {
	ParseFlagSet(prefix, flag.CommandLine)
}

// ParseFlagSet takes a prefix string and exposes environment variables
// for all flags in the FlagSet in the form of PREFIX_FLAGNAME.
//
// Each flag in the FlagSet is exposed as an SCREAMING_SNAKE_CASE environment
// variable prefixed with prefix. Any flag that was not explicitly set
// by a user is updated to the environment variable, if set.
func ParseFlagSet(prefix string, fs *flag.FlagSet) {
	// Build a map of explicitly set flags.
	set := map[string]interface{}{}
	fs.Visit(func(f *flag.Flag) {
		set[f.Name] = nil
	})

	fs.VisitAll(func(f *flag.Flag) {
		envVar := formatName(prefix, f.Name)

		if val := os.Getenv(envVar); val != "" {
			if _, defined := set[f.Name]; !defined {
				fs.Set(f.Name, val)
			}
		}

		f.Usage = fmt.Sprintf("%s [%s]", f.Usage, envVar)
	})
}

func formatName(prefix, s string) string {
	runes := []rune(s)
	length := len(runes)

	out := &strings.Builder{}
	out.Grow(len(prefix) + length + 1)
	out.WriteString(prefix)
	out.WriteRune('_')

	for i := 0; i < length; i++ {
		if !unicode.IsLetter(runes[i]) {
			out.WriteRune('_')
			continue
		}
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out.WriteRune('_')
		}
		out.WriteRune(unicode.ToUpper(runes[i]))
	}

	return out.String()
}
