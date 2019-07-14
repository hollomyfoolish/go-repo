package p2

import(
	"fmt"
	"github.com/hollomyfoolish/go-repo/v2/p1"
)

func Echo() string {
	return "service in p2 of module go-repo called: " + fmt.Sprintf("%s\n", p1.Hello())
}