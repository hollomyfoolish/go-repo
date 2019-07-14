package p1

import(
	"github.com/satori/go.uuid"
)

func Hello() string {
	uid := uuid.NewV4()
	return "service in p1 of module go-repo["+ uid.String() +"]"
}