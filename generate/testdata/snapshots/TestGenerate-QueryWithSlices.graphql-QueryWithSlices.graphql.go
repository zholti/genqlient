// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
)

// QueryWithSlicesResponse is returned by QueryWithSlices on success.
type QueryWithSlicesResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithSlicesUser `json:"user"`
}

// GetUser returns QueryWithSlicesResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesResponse) GetUser() QueryWithSlicesUser { return v.User }

// QueryWithSlicesUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithSlicesUser struct {
	Emails                []string `json:"emails"`
	EmailsOrNull          []string `json:"emailsOrNull"`
	EmailsWithNulls       []string `json:"emailsWithNulls"`
	EmailsWithNullsOrNull []string `json:"emailsWithNullsOrNull"`
}

// GetEmails returns QueryWithSlicesUser.Emails, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmails() []string { return v.Emails }

// GetEmailsOrNull returns QueryWithSlicesUser.EmailsOrNull, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsOrNull() []string { return v.EmailsOrNull }

// GetEmailsWithNulls returns QueryWithSlicesUser.EmailsWithNulls, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsWithNulls() []string { return v.EmailsWithNulls }

// GetEmailsWithNullsOrNull returns QueryWithSlicesUser.EmailsWithNullsOrNull, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsWithNullsOrNull() []string { return v.EmailsWithNullsOrNull }

// The query or mutation executed by QueryWithSlices.
const QueryWithSlices_Operation = `
query QueryWithSlices {
	user {
		emails
		emailsOrNull
		emailsWithNulls
		emailsWithNullsOrNull
	}
}
`

func QueryWithSlices(
	client graphql.Client,
) (*QueryWithSlicesResponse, error) {
	req := &graphql.Request{
		OpName: "QueryWithSlices",
		Query:  QueryWithSlices_Operation,
	}
	var err_ error

	var data_ QueryWithSlicesResponse
	resp := &graphql.Response{Data: &data_}

	err_ = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data_, err_
}

