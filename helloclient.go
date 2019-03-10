package main

import(
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://helloworld/")
	if err != nil {
	  fmt.Fprintf(w,"%s", err)
	} else {
	  contents, err := ioutil.ReadAll(response.Body)
	  if err != nil {
	    fmt.Fprintf(w,"%s", err)
	  }
	  fmt.Fprintf(w, "Message found: %s", contents)
	}
  })

  http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	  fmt.Fprintf(w, "I am running and new !")
  })

  http.ListenAndServe(":3001", nil)

}
