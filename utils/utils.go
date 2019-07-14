package utils

import(
	"strings"
	"os"
)

func ParseArgs() map[string]string {
	args := make(map[string]string)

	for _, arg := range os.Args {
		idx := strings.Index(arg, "-D")
		if idx == 0 { 
			idx2 := strings.Index(arg, "=")
			if idx2 >= 0 && idx2 < (len(arg) - 1){
				args[arg[idx + 2:idx2]] = arg[idx2 + 1:]
			}
		}
	}
	return args
}