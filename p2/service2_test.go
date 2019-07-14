package p2

import(
	"testing"
)

func TestHello(t *testing.T) {
    want := "service in p2 of module go-repo"
    if got := Echo(); got != want {
        t.Errorf("Echo() = %q, want %q", got, want)
    }
}