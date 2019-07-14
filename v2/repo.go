package repo

import(
	"github.com/hollomyfoolish/go-repo/v2/p2"
)

func Repo() string {
	return "github.com/hollomyfoolish/go-repo/v2: " + p2.Echo()
}