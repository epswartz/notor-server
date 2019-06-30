# notor-server
Notor is a rotating notes system that shows you random notes from your past.

I often find that I learn things, and while my life is affected during the initial period where those things have a high salience in my mind, I lose the benefits of knowing them over time as I forget them, or forget to pay attention to them.

The notor server has 2 basic functions: storing notes, and regurgitating a random note. At a later date I will attach a desktop notification or mobile app to this server, so that I can randomly remind myself of things I used to know and consider important.

## Install/Usage
0. Install `go` and `curl`
1. Clone this repo into gopath (often `~/go/src`)
2. Start the server with `go run main.go`
3. Get the test note with `curl localhost:3000` (the note itself will be base64 encoded)
