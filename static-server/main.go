package main

import(
	"net/http"
	"fmt"

	"github.com/hollomyfoolish/go-repo/utils"
)

func main(){
	args := utils.ParseArgs()
	fs := http.FileServer(http.Dir(""))

	fmt.Printf("%v\n", args)
	fmt.Printf("%v\n", fs)
}