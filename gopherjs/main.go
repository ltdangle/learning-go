package main
// gopherjs build main.go -o main.js
// export GOPHERJS_GOROOT="$(go1.18.10 env GOROOT)"
// gopherjs serve
import (
	"io"
	"net/http"

	"github.com/gopherjs/gopherjs/js"
)


func api() {
	// Prepare parameters.

	// Call api.
	var apiResponse string
	url:="https://pokeapi.co/api/v2/pokemon/ditto"

	resp, err := http.Get(url)
	if err != nil {
		js.Global.Call("alert", "The HTTP request failed with error %s\n")
	} else {
		respByte, _ := io.ReadAll(resp.Body)
		apiResponse = string(respByte)
	}

	js.Global.Get("console").Call("log", apiResponse)
}
func main() {
	js.Global.Get("console").Call("log", "Hello, World!")
	api()
}
