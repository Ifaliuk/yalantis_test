package internal

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strings"
	"testing"
)

type testExpectedResultParams struct {
	Field string
	Message string
}

func (params testExpectedResultParams) String() string {
	return fmt.Sprintf("validation.Errors: %s: %s", params.Field, params.Message)
}

type testPair struct {
	name string
	email string
	expect interface{}
}

func (pair testPair) InputToString() string {
	return fmt.Sprintf("{name: '%s', email: '%s'}", pair.name, pair.email)
}

func parseResult(res interface{}, params testExpectedResultParams) bool  {
	err, ok := res.(validation.Errors)
	if !ok {
		return false
	}
	message, ok := err[params.Field]
	if !ok {
		return false
	}
	return strings.Contains(message.Error(), params.Message)
}

var testPairs = []testPair {
	{
	"",
	"john@gmail.com",
		testExpectedResultParams{"Name", "cannot be blank"},
	},
	{
	"John",
	"",
		testExpectedResultParams{"Email", "cannot be blank"},
	},
	{
	"John",
	"john",
		testExpectedResultParams{"Email", "must be a valid email address"},
	},
	{
	"John",
	"john@gmail.com",
	nil,
	},
	{
	"Mike",
	"john@gmail.com",
		testExpectedResultParams{"Email", "user with this email already exists"},
	},

}

func TestAddMember(t *testing.T)  {
	members := Members{}
	for _, testPair :=  range testPairs {
		m := NewMember(testPair.name, testPair.email)
		res := members.AddMember(&m)
		if expectedResultParams, ok := testPair.expect.(testExpectedResultParams); ok {
			if expect := parseResult(res, expectedResultParams); !expect {
				t.Error(
					"For", testPair.InputToString(),
					"expected", expectedResultParams,
					"got", res,
				)
			}
		} else {
			if res != nil {
				t.Error(
					"For", testPair.InputToString(),
					"expected", nil,
					"got", res,
				)
			}
		}
	}
}
