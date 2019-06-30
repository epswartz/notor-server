package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type response struct {
	Note  []byte `json: "note"`
	Index string `json: "index"`
}

func main() {
	user := os.Getenv("NOTOR_USER")
	pass := os.Getenv("NOTOR_PASS")

	if user != "" && pass != "" {
		http.HandleFunc("/", basicAuth(handleIndex, user, pass, "Please enter user/pass"))
	} else {
		http.HandleFunc("/", handleIndex)

	}

	log("Server Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

// Handler for /
func handleIndex(w http.ResponseWriter, r *http.Request) {
	noteBytes, idx, err := getNote()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errlog("GETNOTE ERROR: " + err.Error())
		w.Write([]byte("500 - check server log."))
		return
	}
	resp := response{
		Note:  noteBytes,
		Index: idx,
	}
	log("SUCCESS: Sent note")
	json.NewEncoder(w).Encode(resp)

}

// BasicAuth wraps a handler requiring HTTP basic auth for it using the given
// username and password and the specified realm, which shouldn't contain quotes.
func basicAuth(handler http.HandlerFunc, username, password, message string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+message+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized.\n"))
			return
		}

		handler(w, r)
	}
}

// Gets a random note from the note directory and returns its contents and name.
func getNote() ([]byte, string, error) {
	if _, err := os.Stat("./notes"); os.IsNotExist(err) { // If no notes dir
		return []byte{}, "", err
	}

	files, err := ioutil.ReadDir("./notes")
	if err != nil {
		return []byte{}, "", err
	}

	fileName := files[rand.Intn(len(files))].Name() // random file
	b, err := ioutil.ReadFile("./notes/" + fileName)

	if err != nil {
		return []byte{}, "", err
	}
	return b, fileName, nil
}

// Easy function for logging, might replace with something more robust later
func log(s string) {
	now := time.Now()
	t := fmt.Sprintf("[%d-%02d-%02d %02d:%02d:%02d-00:00]",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Printf("%s %s\n", t, s)
}

// Easy function for logging errors, might replace with something more robust later
func errlog(s string) {
	now := time.Now()
	t := fmt.Sprintf("ERR [%d-%02d-%02d %02d:%02d:%02d-00:00]",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Printf("%s %s\n", t, s)
}
