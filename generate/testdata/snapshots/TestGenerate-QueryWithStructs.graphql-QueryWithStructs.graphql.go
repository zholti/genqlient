// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
)

// QueryWithStructsResponse is returned by QueryWithStructs on success.
type QueryWithStructsResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithStructsUser `json:"user"`
}

// GetUser returns QueryWithStructsResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithStructsResponse) GetUser() QueryWithStructsUser { return v.User }

// QueryWithStructsUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithStructsUser struct {
	AuthMethods []QueryWithStructsUserAuthMethodsAuthMethod `json:"authMethods"`
}

// GetAuthMethods returns QueryWithStructsUser.AuthMethods, and is useful for accessing the field via an interface.
func (v *QueryWithStructsUser) GetAuthMethods() []QueryWithStructsUserAuthMethodsAuthMethod {
	return v.AuthMethods
}

// QueryWithStructsUserAuthMethodsAuthMethod includes the requested fields of the GraphQL type AuthMethod.
type QueryWithStructsUserAuthMethodsAuthMethod struct {
	Provider string `json:"provider"`
	Email    string `json:"email"`
}

// GetProvider returns QueryWithStructsUserAuthMethodsAuthMethod.Provider, and is useful for accessing the field via an interface.
func (v *QueryWithStructsUserAuthMethodsAuthMethod) GetProvider() string { return v.Provider }

// GetEmail returns QueryWithStructsUserAuthMethodsAuthMethod.Email, and is useful for accessing the field via an interface.
func (v *QueryWithStructsUserAuthMethodsAuthMethod) GetEmail() string { return v.Email }

// The query or mutation executed by QueryWithStructs.
const QueryWithStructs_Operation = `
query QueryWithStructs {
	user {
		authMethods {
			provider
			email
		}
	}
}
`

func QueryWithStructs(
	client graphql.Client,
) (*QueryWithStructsResponse, error) {
	req := &graphql.Request{
		OpName: "QueryWithStructs",
		Query:  QueryWithStructs_Operation,
	}
	var err error

	var data_ QueryWithStructsResponse
	resp := &graphql.Response{Data: &data_}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data_, err
}

