// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// __getUserInput is used internally by genqlient
type __getUserInput struct {
	Login string `json:"Login"`
}

// GetLogin returns __getUserInput.Login, and is useful for accessing the field via an interface.
func (v *__getUserInput) GetLogin() string { return v.Login }

// getUserResponse is returned by getUser on success.
type getUserResponse struct {
	// Lookup a user by login.
	User getUserUser `json:"user"`
}

// GetUser returns getUserResponse.User, and is useful for accessing the field via an interface.
func (v *getUserResponse) GetUser() getUserUser { return v.User }

// getUserUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A user is an individual's account on GitHub that owns repositories and can make new content.
type getUserUser struct {
	// The user's public profile name.
	TheirName string `json:"theirName"`
	// Identifies the date and time when the object was created.
	CreatedAt time.Time `json:"createdAt"`
}

// GetTheirName returns getUserUser.TheirName, and is useful for accessing the field via an interface.
func (v *getUserUser) GetTheirName() string { return v.TheirName }

// GetCreatedAt returns getUserUser.CreatedAt, and is useful for accessing the field via an interface.
func (v *getUserUser) GetCreatedAt() time.Time { return v.CreatedAt }

// getViewerResponse is returned by getViewer on success.
type getViewerResponse struct {
	// The currently authenticated user.
	Viewer getViewerViewerUser `json:"viewer"`
}

// GetViewer returns getViewerResponse.Viewer, and is useful for accessing the field via an interface.
func (v *getViewerResponse) GetViewer() getViewerViewerUser { return v.Viewer }

// getViewerViewerUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A user is an individual's account on GitHub that owns repositories and can make new content.
type getViewerViewerUser struct {
	// The user's public profile name.
	MyName string `json:"MyName"`
	// Identifies the date and time when the object was created.
	CreatedAt time.Time `json:"createdAt"`
}

// GetMyName returns getViewerViewerUser.MyName, and is useful for accessing the field via an interface.
func (v *getViewerViewerUser) GetMyName() string { return v.MyName }

// GetCreatedAt returns getViewerViewerUser.CreatedAt, and is useful for accessing the field via an interface.
func (v *getViewerViewerUser) GetCreatedAt() time.Time { return v.CreatedAt }

// The query or mutation executed by getUser.
const getUser_Operation = `
query getUser ($Login: String!) {
	user(login: $Login) {
		theirName: name
		createdAt
	}
}
`

// getUser gets the given user's name from their username.
func getUser(
	ctx_ context.Context,
	client_ graphql.Client,
	Login string,
) (*getUserResponse, error) {
	req_ := &graphql.Request{
		OpName: "getUser",
		Query:  getUser_Operation,
		Variables: &__getUserInput{
			Login: Login,
		},
	}
	var err_ error

	var data_ getUserResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by getViewer.
const getViewer_Operation = `
query getViewer {
	viewer {
		MyName: name
		createdAt
	}
}
`

func getViewer(
	ctx_ context.Context,
	client_ graphql.Client,
) (*getViewerResponse, error) {
	req_ := &graphql.Request{
		OpName: "getViewer",
		Query:  getViewer_Operation,
	}
	var err_ error

	var data_ getViewerResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
