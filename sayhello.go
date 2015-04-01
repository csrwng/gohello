package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func helloFunc(port string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("<html><body><h1>Hello World on Port %s</h1></body></html>", port)
		w.Write([]byte(str))
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
