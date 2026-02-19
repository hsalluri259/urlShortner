package main

import (
	"fmt"
	"net/http"
	"store"
)

const addForm = `
<html><body>
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
</body></html>
`

var urlStore = store.NewURLStore()

func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8080", nil)
}

// Redirect to long URL
func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := urlStore.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

// Generate a short URL
func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	url := r.FormValue("url")
	if url == "" {
		fmt.Fprint(w, addForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "%s", key)
}
