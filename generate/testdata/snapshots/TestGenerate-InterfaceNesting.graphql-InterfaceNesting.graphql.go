// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceNestingResponse is returned by InterfaceNesting on success.
type InterfaceNestingResponse struct {
	Root InterfaceNestingRootTopic `json:"root"`
}

// GetRoot returns InterfaceNestingResponse.Root, and is useful for accessing the field via an interface.
func (v *InterfaceNestingResponse) GetRoot() InterfaceNestingRootTopic { return v.Root }

// InterfaceNestingRootTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                `json:"id"`
	Children []InterfaceNestingRootTopicChildrenContent `json:"-"`
}

// GetId returns InterfaceNestingRootTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopic) GetId() testutil.ID { return v.Id }

// GetChildren returns InterfaceNestingRootTopic.Children, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopic) GetChildren() []InterfaceNestingRootTopicChildrenContent {
	return v.Children
}

func (v *InterfaceNestingRootTopic) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*InterfaceNestingRootTopic
		Children []json.RawMessage `json:"children"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InterfaceNestingRootTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Children
		src := firstPass.Children
		*dst = make(
			[]InterfaceNestingRootTopicChildrenContent,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			if len(src) != 0 && string(src) != "null" {
				err = __unmarshalInterfaceNestingRootTopicChildrenContent(
					src, dst)
				if err != nil {
					return fmt.Errorf(
						"unable to unmarshal InterfaceNestingRootTopic.Children: %w", err)
				}
			}
		}
	}
	return nil
}

type __premarshalInterfaceNestingRootTopic struct {
	Id testutil.ID `json:"id"`

	Children []json.RawMessage `json:"children"`
}

func (v *InterfaceNestingRootTopic) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *InterfaceNestingRootTopic) __premarshalJSON() (*__premarshalInterfaceNestingRootTopic, error) {
	var retval __premarshalInterfaceNestingRootTopic

	retval.Id = v.Id
	{

		dst := &retval.Children
		src := v.Children
		*dst = make(
			[]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			var err error
			*dst, err = __marshalInterfaceNestingRootTopicChildrenContent(
				&src)
			if err != nil {
				return nil, fmt.Errorf(
					"unable to marshal InterfaceNestingRootTopic.Children: %w", err)
			}
		}
	}
	return &retval, nil
}

// InterfaceNestingRootTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceNestingRootTopicChildrenArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id     testutil.ID                                         `json:"id"`
	Parent InterfaceNestingRootTopicChildrenContentParentTopic `json:"parent"`
}

// GetTypename returns InterfaceNestingRootTopicChildrenArticle.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenArticle) GetTypename() string { return v.Typename }

// GetId returns InterfaceNestingRootTopicChildrenArticle.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenArticle) GetId() testutil.ID { return v.Id }

// GetParent returns InterfaceNestingRootTopicChildrenArticle.Parent, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenArticle) GetParent() InterfaceNestingRootTopicChildrenContentParentTopic {
	return v.Parent
}

// InterfaceNestingRootTopicChildrenContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceNestingRootTopicChildrenContent is implemented by the following types:
// InterfaceNestingRootTopicChildrenArticle
// InterfaceNestingRootTopicChildrenTopic
// InterfaceNestingRootTopicChildrenVideo
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNestingRootTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetParent returns the interface-field "parent" from its implementation.
	GetParent() InterfaceNestingRootTopicChildrenContentParentTopic
}

func (v *InterfaceNestingRootTopicChildrenArticle) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent() {
}
func (v *InterfaceNestingRootTopicChildrenTopic) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent() {
}
func (v *InterfaceNestingRootTopicChildrenVideo) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent() {
}

func __unmarshalInterfaceNestingRootTopicChildrenContent(b []byte, v *InterfaceNestingRootTopicChildrenContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceNestingRootTopicChildrenArticle)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(InterfaceNestingRootTopicChildrenTopic)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(InterfaceNestingRootTopicChildrenVideo)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for InterfaceNestingRootTopicChildrenContent: "%v"`, tn.TypeName)
	}
}

func __marshalInterfaceNestingRootTopicChildrenContent(v *InterfaceNestingRootTopicChildrenContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *InterfaceNestingRootTopicChildrenArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNestingRootTopicChildrenArticle
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceNestingRootTopicChildrenTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNestingRootTopicChildrenTopic
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceNestingRootTopicChildrenVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNestingRootTopicChildrenVideo
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for InterfaceNestingRootTopicChildrenContent: "%T"`, v)
	}
}

// InterfaceNestingRootTopicChildrenContentParentTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenContentParentTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                                          `json:"id"`
	Children []InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent `json:"-"`
}

// GetId returns InterfaceNestingRootTopicChildrenContentParentTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenContentParentTopic) GetId() testutil.ID { return v.Id }

// GetChildren returns InterfaceNestingRootTopicChildrenContentParentTopic.Children, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenContentParentTopic) GetChildren() []InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent {
	return v.Children
}

func (v *InterfaceNestingRootTopicChildrenContentParentTopic) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*InterfaceNestingRootTopicChildrenContentParentTopic
		Children []json.RawMessage `json:"children"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InterfaceNestingRootTopicChildrenContentParentTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Children
		src := firstPass.Children
		*dst = make(
			[]InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			if len(src) != 0 && string(src) != "null" {
				err = __unmarshalInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent(
					src, dst)
				if err != nil {
					return fmt.Errorf(
						"unable to unmarshal InterfaceNestingRootTopicChildrenContentParentTopic.Children: %w", err)
				}
			}
		}
	}
	return nil
}

type __premarshalInterfaceNestingRootTopicChildrenContentParentTopic struct {
	Id testutil.ID `json:"id"`

	Children []json.RawMessage `json:"children"`
}

func (v *InterfaceNestingRootTopicChildrenContentParentTopic) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *InterfaceNestingRootTopicChildrenContentParentTopic) __premarshalJSON() (*__premarshalInterfaceNestingRootTopicChildrenContentParentTopic, error) {
	var retval __premarshalInterfaceNestingRootTopicChildrenContentParentTopic

	retval.Id = v.Id
	{

		dst := &retval.Children
		src := v.Children
		*dst = make(
			[]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			var err error
			*dst, err = __marshalInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent(
				&src)
			if err != nil {
				return nil, fmt.Errorf(
					"unable to marshal InterfaceNestingRootTopicChildrenContentParentTopic.Children: %w", err)
			}
		}
	}
	return &retval, nil
}

// InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// GetTypename returns InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle) GetTypename() string {
	return v.Typename
}

// GetId returns InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle) GetId() testutil.ID {
	return v.Id
}

// InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent is implemented by the following types:
// InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle
// InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic
// InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
}

func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent() {
}
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent() {
}
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent() {
}

func __unmarshalInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent(b []byte, v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent: "%v"`, tn.TypeName)
	}
}

func __marshalInterfaceNestingRootTopicChildrenContentParentTopicChildrenContent(v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNestingRootTopicChildrenContentParentTopicChildrenArticle
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for InterfaceNestingRootTopicChildrenContentParentTopicChildrenContent: "%T"`, v)
	}
}

// InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// GetTypename returns InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic) GetTypename() string {
	return v.Typename
}

// GetId returns InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenTopic) GetId() testutil.ID {
	return v.Id
}

// InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

// GetTypename returns InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo) GetTypename() string {
	return v.Typename
}

// GetId returns InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenContentParentTopicChildrenVideo) GetId() testutil.ID {
	return v.Id
}

// InterfaceNestingRootTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id     testutil.ID                                         `json:"id"`
	Parent InterfaceNestingRootTopicChildrenContentParentTopic `json:"parent"`
}

// GetTypename returns InterfaceNestingRootTopicChildrenTopic.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenTopic) GetTypename() string { return v.Typename }

// GetId returns InterfaceNestingRootTopicChildrenTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenTopic) GetId() testutil.ID { return v.Id }

// GetParent returns InterfaceNestingRootTopicChildrenTopic.Parent, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenTopic) GetParent() InterfaceNestingRootTopicChildrenContentParentTopic {
	return v.Parent
}

// InterfaceNestingRootTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceNestingRootTopicChildrenVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id     testutil.ID                                         `json:"id"`
	Parent InterfaceNestingRootTopicChildrenContentParentTopic `json:"parent"`
}

// GetTypename returns InterfaceNestingRootTopicChildrenVideo.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenVideo) GetTypename() string { return v.Typename }

// GetId returns InterfaceNestingRootTopicChildrenVideo.Id, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenVideo) GetId() testutil.ID { return v.Id }

// GetParent returns InterfaceNestingRootTopicChildrenVideo.Parent, and is useful for accessing the field via an interface.
func (v *InterfaceNestingRootTopicChildrenVideo) GetParent() InterfaceNestingRootTopicChildrenContentParentTopic {
	return v.Parent
}

// The query or mutation executed by InterfaceNesting.
const InterfaceNesting_Operation = `
query InterfaceNesting {
	root {
		id
		children {
			__typename
			id
			parent {
				id
				children {
					__typename
					id
				}
			}
		}
	}
}
`

func InterfaceNesting(
	client graphql.Client,
) (*InterfaceNestingResponse, error) {
	req := &graphql.Request{
		OpName: "InterfaceNesting",
		Query:  InterfaceNesting_Operation,
	}
	var err_ error

	var data_ InterfaceNestingResponse
	resp := &graphql.Response{Data: &data_}

	err_ = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data_, err_
}

