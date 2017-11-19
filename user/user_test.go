package user

import (
	"context"
	"testing"

	"github.com/Typeform/users/datastore"
	"github.com/Typeform/users/transport/v1"

	"github.com/Typeform/users/datastore/mockdb"
	"github.com/golang/mock/gomock"
)

func Int(i int) *int {
	return &i
}

func String(s string) *string {
	return &s
}

func TestGetUser(t *testing.T) {
	tests := []struct { // Create a testtable
		name string
		id   int
		user *v1.User
		err  error
	}{
		{
			name: "get1", // add a test that returns a user
			id:   1,
			user: &v1.User{
				ID:   Int(1),
				Name: String("sevki"),
			},
		},
		{
			name: "get2", // add a test for a user we won't find
			id:   2,
			user: nil,
			err:  datastore.ErrUserNotFound, // this is the datastore err we defined earlier
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// read the mockgen documentation here https://godoc.org/github.com/golang/mock/gomock
			mockCtrl := gomock.NewController(t) // create a new recorder
			defer mockCtrl.Finish()             // cleanup after

			mockUsersDB := mockdb.NewMockUsers(mockCtrl) // create the mockdb object
			// use EXPECT to tell the controller what we expect to see in the first call made to this controller
			// And returns to tell the controller what we expect to see in return
			mockUsersDB.EXPECT().GetUser(test.id).Return(test.user, test.err)

			// finally create the userservice and make the calls
			users := New(mockUsersDB)
			user, err := users.GetUser(context.Background(), test.id)
			// check they return correct values
			if err != nil {
				if err != test.err {
					t.Logf("was expecting err %v got %v instead", test.err, err)
					t.Fail()
				} else {
					return
				}
			}
			if user.ID != test.user.ID {
				t.Logf("was expecting userID %d got %d instead", test.user.ID, user.ID)
				t.Fail()
			}
			if user.Name != test.user.Name {
				t.Logf("was expecting user name %q got %q instead", test.user.Name, user.Name)
				t.Fail()
			}
		})
	}
}
