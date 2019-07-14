package app
/**
import(
	"net/http"
	"log"
)
**/
func Foo() string {
	return "foo from app"
}
/**
func api(rsp http.ResponseWriter, req *http.Request){
	rsp.Header().Set("Content-Type", "application/json; charset=utf-8")
	rsp.WriteHeader(200)
	rsp.Write([]byte("{\"stataus\": 0}"))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/api/", http.HandlerFunc(api))
  
	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil{
		log.Fatal(err)
	}
  }**/
