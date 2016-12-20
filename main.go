package envy

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Parse takes a string p that is used
// as the environment variable prefix
// for each flag configured.
func Parse(p string) {
	flag.CommandLine.VisitAll(func(f *flag.Flag) {
		// Create an env var name
		// based on the supplied prefix.
		envVar := fmt.Sprintf("%s_%s", p, strings.ToUpper(f.Name))
		envVar = strings.Replace(envVar, "-", "_", -1)

		// Update the Flag.Value if the
		// env var is non "".
		if val := os.Getenv(envVar); val != "" {
			// Update the value.
			flag.Set(f.Name, val)
		}

		// Append the env var to the
		// Flag.Usage field.
		f.Usage = fmt.Sprintf("%s [%s]", f.Usage, envVar)
	})
}