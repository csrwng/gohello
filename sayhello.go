package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func helloFunc(port string) http.HandlerFunc {
	sprintf := func(w http.ResponseWriter, format string, args ...interface{}) {
		w.Write([]byte(fmt.Sprintf(format, args...)))
	}
	return func(w http.ResponseWriter, r *http.Request) {
		sprintf(w, `<html><body>
		<p>Version 2</p>
		<h1>Hello World on Port %s</h1>`, port)
		for k, v := range r.Header {
			sprintf(w, "Header %s ==> %s<br />", k, v)
		}
		sprintf(w, "<br />")
		for k, v := range r.URL.Query() {
			sprintf(w, "Parameter %s ==> %s<br />", k, v)
		}
		sprintf(w, "</body></html>")
	}
}

func listenOn(port string) {
	fmt.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, helloFunc(port)))
}

func main() {
	ports := []string{"8080"}
	if len(os.Args) > 1 {
		ports = os.Args[1:]
	}
	for _, p := range ports {
		go listenOn(p)
	}
	i := 0
	for {
		fmt.Printf("%s: Looping for %d times\n", time.Now().Format(time.UnixDate), i)
		time.Sleep(5 * time.Second)
		i++
	}
}
