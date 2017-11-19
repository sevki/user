package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Typeform/users/datastore/memstore"
	"github.com/Typeform/users/transport/v1"
	"github.com/Typeform/users/user"
)

var (
	addr = flag.String("addr", ":8080", "default -addr :8080")
)

func Int(i int) *int {
	return &i
}

func String(s string) *string {
	return &s
}

func main() {
	flag.Parse()

	x := map[int]v1.User{
		1: v1.User{
			ID:   Int(1),
			Name: String("sevki"),
		},
	}
	userService := user.New(memstore.New(x))
	us := v1.NewUsersHTTPServer(userService)

	log.Printf("listening at %s\n", *addr)
	log.Fatal("listening at "+*addr+" failed: ", http.ListenAndServe(*addr, us))
}
