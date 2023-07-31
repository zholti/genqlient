// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package queries

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// SimpleQueryResponse is returned by SimpleQuery on success.
type SimpleQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User SimpleQueryUser `json:"user"`
}

// GetUser returns SimpleQueryResponse.User, and is useful for accessing the field via an interface.
func (v *SimpleQueryResponse) GetUser() SimpleQueryUser { return v.User }

// SimpleQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetId returns SimpleQueryUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleQueryUser) GetId() testutil.ID { return v.Id }

// The query or mutation executed by SimpleQuery.
const SimpleQuery_Operation = `
query SimpleQuery {
	user {
		id
	}
}
`

func SimpleQuery(
	ctx context.Context,
	client graphql.Client,
) (*SimpleQueryResponse, error) {
	req := &graphql.Request{
		OpName: "SimpleQuery",
		Query:  SimpleQuery_Operation,
	}
	var err error

	var data_ SimpleQueryResponse
	resp := &graphql.Response{Data: &data_}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data_, err
}

