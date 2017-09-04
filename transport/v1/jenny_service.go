// Automatically generated by Jenny. DO NOT EDIT!

// Package v1 as generated by Jenny
// Please read about it https://localhost:8080/_spec
package v1

import (
	"context"

	"github.com/Typeform/jenny/options"
	"github.com/go-kit/kit/endpoint"
)

// Users as generated by Jenny
// Please read more at https://localhost:8080/_spec
type Users interface {

	// GetUser Gets a User from the database
	GetUser(ctx context.Context, ID int) (Body *User, err error)
}

// User as defined in swagger. http://localhost:8080/_spec#Users/User
type User struct {
	ID   *int    // id as defined in swagger. http://localhost:8080/_spec#Users/User/ID
	Name *string // name as defined in swagger. http://localhost:8080/_spec#Users/User/Name
}

// _getUserRequest is not to be used outside of this file.
// see https://gokit.io/examples/stringsvc.html#requests-and-responses for more detail
type _getUserRequest struct {
	ID int // Id of the user

}

// _getUserResponse is not to be used outside of this file.
// see https://gokit.io/examples/stringsvc.html#requests-and-responses for more detail
type _getUserResponse struct {
	Body *User // body as defined in swagger. http://localhost:8080/_spec#Users/GetUser/Body

}

// endpoints as used in https://gokit.io/examples/stringsvc.html#endpoints
func makeGetUserEndpoint(svc Users, opts *options.Options) endpoint.Endpoint {
	getUserEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(_getUserRequest)

		resp := _getUserResponse{}
		var err error

		resp.Body, err = svc.GetUser(ctx, req.ID)

		return resp, err
	}

	getUserMiddleware := opts.OpMiddlewares("GetUser")

	return getUserMiddleware(getUserEndpoint)
}
