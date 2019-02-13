package store

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingTypes(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected string
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
		assert.Equal(parsed.Type(), test.expected)
	}
}

func TestParsingString(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected string
	}{
		{`"string"`, "string"},
		{`true`, "true"},
		{`23`, "23"},
		{`0.34`, "0.34"},
		{`{"a":1, "b":2}`, "<type:branch, childs:2>"},
		{`[1,2,3,4]`, "<type:list, childs:4>"},
	}

	for _, test := range tests {
		var buff interface{}
		json.Unmarshal([]byte(test.input), &buff)
		node := Parse(buff)
		assert.Equal(node.String(), test.expected)
	}
}

func TestChild(t *testing.T) {
	assert := assert.New(t)

	js := `{"a":1, "b":2}`
	var buff interface{}
	json.Unmarshal([]byte(js), &buff)
	node := Parse(buff)
	assert.Equal("1", node.Child("a").String())

	js = `["a","b","c","d"]`
	json.Unmarshal([]byte(js), &buff)
	node = Parse(buff)
	assert.Equal("c", node.Child(2).String())
}

func TestParsingChildKeys(t *testing.T) {
	assert := assert.New(t)
	js := `{"a":1, "b":2, "c":3, "d":4}`
	var buff interface{}
	json.Unmarshal([]byte(js), &buff)

	root := Parse(buff)
	for _, key := range root.ChildKeys() {
		assert.NotNil(root.Child(key))
	}

	js = `[1,2,3,4,5,6,7,8,9]`
	json.Unmarshal([]byte(js), &buff)

	root = Parse(buff)
	for _, key := range root.ChildKeys() {
		assert.NotNil(root.Child(key))
	}

}
