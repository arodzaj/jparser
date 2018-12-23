package store

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingTypes(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input   string
		expType string
	}{
		{`"string"`, "leaf"},
		{`true`, "leaf"},
		{`23`, "leaf"},
		{`0.34`, "leaf"},
		{`{"a":1, "b":2}`, "branch"},
		{`[1,2,3,4]`, "list"},
	}

	for _, test := range tests {
		var buff interface{}
		json.Unmarshal([]byte(test.input), &buff)
		parsed := Parse(buff)
		assert.Equal(parsed.Type(), test.expType)
	}
}

func TestParsingValue(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input string
		value string
	}{
		{`"string"`, "string"},
		{`true`, "true"},
		{`23`, "23"},
		{`0.34`, "0.34"},
		{`{"a":1, "b":2}`, ""},
		{`[1,2,3,4]`, ""},
	}

	for _, test := range tests {
		var buff interface{}
		json.Unmarshal([]byte(test.input), &buff)
		parsed := Parse(buff)
		assert.Equal(parsed.Value(), test.value)
	}
}

func TestParsingReccuring(t *testing.T) {
	assert := assert.New(t)
	js := `{"obj": {"a":1, "b":2}}`
	var buff interface{}
	json.Unmarshal([]byte(js), &buff)

	node := Parse(buff)
	assert.Equal(node.Type(), "branch")

}
