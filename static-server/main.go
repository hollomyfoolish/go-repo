package main

import(
	"net/http"
	"fmt"
	// "os"
	"log"
	"io/ioutil"
	"text/template"
	"strings"

	"github.com/hollomyfoolish/go-repo/utils"
)

const MV_DIR_ARG = "dir"
var mvDir string

const templat = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <ol>
	{{range .Files}}
		<li><a href="/play?m={{.}}">{{.}}</a></li>
	{{end}}
    </ol>
</body>
</html>
`
const playTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
	<video src="/static/{{.}}" controls="controls" autoplay="autoplay">
</body>
</html>
`

var mvTemp = template.Must(template.New("movie").Parse(templat))
var playTemp = template.Must(template.New("play").Parse(playTemplate))

type Movies struct{
	Base string
	Files []string
}

func getMovieList(rsp http.ResponseWriter, r *http.Request){
	//Content-Type: text/html; charset=UTF-8
	rsp.Header().Set("Content-Type", "text/html; charset=utf-8")
	rsp.WriteHeader(200)
	mvTemp.Execute(rsp, getAllMovies())
}

func getAllMovies() Movies {
	files, err := ioutil.ReadDir(mvDir)
	if err != nil{
		fmt.Printf("%v\n", err)
		return Movies{}
	}
	
	var names []string
	for _, f := range files{
		n := f.Name()
		if strings.HasSuffix(n, ".mp4"){
			names = append(names, n)
		}
	}

	return Movies{
		Base: mvDir,
		Files: names,
	}
}

func playMovie(rsp http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	name := query.Get("m")
	rsp.Header().Set("Content-Type", "text/html; charset=utf-8")
	rsp.WriteHeader(200)
	playTemp.Execute(rsp, name)
}

func main(){
	args := utils.ParseArgs()
	dir, ok := args[MV_DIR_ARG]
	if !ok {
		log.Fatal("movie directory is reqired")
	}
	mvDir = dir

	fmt.Printf("move directory is: %s\n", mvDir)
	fs := http.FileServer(http.Dir(dir))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/movies", getMovieList)
	http.HandleFunc("/play", playMovie)

	err := http.ListenAndServe(":13000", nil)
	if err != nil{
		log.Fatal(err)
	}
}