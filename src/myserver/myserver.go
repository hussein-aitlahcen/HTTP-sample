package myserver

import (
	"fmt"
	"net/http"
	"sync"
)

type MyHandler struct {
	sync.Mutex
	count int
}

type WelcomeHandler struct {
	Message string
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var count int

	h.Lock()
	count = h.count + 1
	h.Unlock()

	fmt.Fprintf(w, "Visitor count: %d.", count)
}

func (wh *WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to %s", wh.Message)
}
