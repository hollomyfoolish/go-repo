package repo

import(
	"github.com/hollomyfoolish/go-repo/p2"
)

func Repo() string {
	return "github.com/hollomyfoolish/go-repo: [update v1 but tag v2]" + p2.Echo()
}