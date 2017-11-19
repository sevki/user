package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	jennyerrors "github.com/Typeform/jenny/errors"
	"github.com/Typeform/users/datastore/mockdb"
	"github.com/Typeform/users/transport/v1"
	"github.com/Typeform/users/user"
	"github.com/golang/mock/gomock"
)

func TestGetUserHTTP(t *testing.T) {
	tests := []struct { // Create a testtable
		name   string
		id     int
		user   *v1.User
		code   int
		errors bool
	}{
		{
			name: "user1", // add a test that returns a user
			id:   1,
			user: &v1.User{
				ID:   Int(1),
				Name: String("sevki"),
			},
			code: http.StatusOK,
		},
		{
			name:   "user2",
			id:     2,
			user:   nil,
			code:   http.StatusInternalServerError,
			errors: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUsersDB := mockdb.NewMockUsers(mockCtrl)
			if test.errors {
				mockUsersDB.EXPECT().GetUser(test.id).Return(test.user, jennyerrors.NewHTTPError(errors.New("something"), http.StatusInternalServerError))
			} else {
				mockUsersDB.EXPECT().GetUser(test.id).Return(test.user, nil)
			}

			users := user.New(mockUsersDB)

			usersServer := v1.NewUsersHTTPServer(users)
			testServer := httptest.NewServer(usersServer)
			u, err := url.Parse(testServer.URL)
			u.Path = "/user"
			q := u.Query()
			q.Add("id", strconv.Itoa(test.id))
			u.RawQuery = q.Encode()
			if err != nil {
				panic(err)
			}
			req, _ := http.NewRequest("GET", u.String(), nil)
			req.Header.Add("Accept", "application/json")
			resp, _ := testServer.Client().Do(req)

			if resp.StatusCode != test.code {
				bytz, _ := ioutil.ReadAll(resp.Body)

				t.Logf("was expecting %d got %d instead", test.code, resp.StatusCode)
				t.Log("\tResponse:", string(bytz))
				t.Fail()
			}
		})
	}
}
