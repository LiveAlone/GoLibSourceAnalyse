package net

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestFileServe(t *testing.T) {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))

	// hijack
	http.HandleFunc("/hijack", func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Don't forget to close the connection:
		defer conn.Close()
		bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
		bufrw.Flush()
		s, err := bufrw.ReadString('\n')
		if err != nil {
			log.Printf("error reading string: %v", err)
			return
		}
		fmt.Fprintf(bufrw, "You said: %q\nBye.\n", s)
		bufrw.Flush()
	})

	http.NotFoundHandler()
}

func TestLocalServe(t *testing.T) {
	s := &http.Server{
		Addr: ":8080",
		//Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func TestServerHttpListener(t *testing.T) {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TestBasic(t *testing.T) {
	resp, _ := http.Get("http://www.baidu.com/")
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func TestChanConfig(t *testing.T) {
	c := make(chan int, 2)
	c <- 2
	fmt.Println(c)
}
