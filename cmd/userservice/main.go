package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Typeform/users/transport/v1"
	"github.com/Typeform/users/user"
)

var (
	addr = flag.String("addr", ":8080", "default -addr :8080")
)

func main() {
	flag.Parse()

	userService := user.New()
	us := v1.NewUsersHTTPServer(userService)

	log.Printf("listening at %s\n", *addr)
	log.Fatal("listening at "+*addr+" failed: ", http.ListenAndServe(*addr, us))
}
