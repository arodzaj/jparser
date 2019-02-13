package store

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingGetNoder(t *testing.T) {
	assert := assert.New(t)
	js := `{"obj": {"a":1, "b":2, "c":[3,4,5,6,7]}}`
	var buff interface{}
	json.Unmarshal([]byte(js), &buff)

	root := Parse(buff)
	node := getNode("obj.a", root)
	assert.Equal(node.String(), "1")

	node = getNode("", root)
	assert.Equal(node, root)

	node = getNode("obj.c.3", root)
	assert.Equal(node.String(), "6")

}
