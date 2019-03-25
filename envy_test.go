package envy

import (
	"flag"
	"os"
	"testing"
)

func TestParseFlagSet(t *testing.T) {
	var (
		flags   = flag.NewFlagSet("myapp", flag.ExitOnError)
		tlsCert = flags.String("tlsCert", "", "TLS Certificate")
		tlsKey  = flags.String("tlsKey", "", "TLS Key")
	)

	os.Setenv("MYAPP_TLS_CERT", "foo.crt")
	os.Setenv("MYAPP_TLS_KEY", "foo.key")
	ParseFlagSet("MYAPP", flags)

	if *tlsCert != "foo.crt" {
		t.Error("Expected tlsCert to be initialized")
	}

	if *tlsKey != "foo.key" {
		t.Error("Expected tlsKey to be initialized")
	}
}
