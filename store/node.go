// Package store provides primitives and logic of loaded JSON files
package store

import (
	"fmt"
	"strconv"
)

// Parse function transforms json into structure
// If passed object is not a leaf, the function calls itself
// for each branch.
// It returns a branch object that implementing Node interface
func Parse(data interface{}) Node {

	switch data.(type) {
	case string, bool, float64:
		l := new(Leaf)
		l.value = data
		return l

	case []interface{}:
		a := new(List)

		a.childs = []Node{}
		for _, value := range data.([]interface{}) {
			a.childs = append(a.childs, Parse(value))
		}

		return a

	case map[string]interface{}:
		a := new(Branch)

		a.childs = map[string]Node{}
		for k, value := range data.(map[string]interface{}) {
			a.childs[k] = Parse(value)
		}

		return a
	}
	return nil
}

// Node interface respresents abstraction a node in JSON tree
type Node interface {
	Type() string
	String() string
	Child(key interface{}) Node
	// Iter() <-chan Node
}

// Leaf -
type Leaf struct {
	value interface{}
}

// Type -
func (l *Leaf) Type() string {
	return "leaf"
}

// String
func (l *Leaf) String() string {
	switch l.value.(type) {
	case string:
		return l.value.(string)

	case float64:
		return strconv.FormatFloat(l.value.(float64), 'f', -1, 64)

	case bool:
		return strconv.FormatBool(l.value.(bool))
	}

	return ""
}

// Child -
func (l *Leaf) Child(key interface{}) Node {
	return nil
}

// List -
type List struct {
	childs []Node
}

// Type -
func (l *List) Type() string {
	return "list"
}

// String -
func (l *List) String() string {
	return fmt.Sprintf("<type:list, childs:%d>", len(l.childs))
}

// Child -
func (l *List) Child(key interface{}) Node {
	k := key.(int)
	return l.childs[k]
}

// Branch -
type Branch struct {
	childs map[string]Node
}

// Type -
func (b *Branch) Type() string {
	return "branch"
}

// String
func (b *Branch) String() string {
	n := 0
	for range b.childs {
		n++
	}
	return fmt.Sprintf("<type:branch, childs:%d>", n)
}

// Child -
func (b *Branch) Child(key interface{}) Node {
	k := key.(string)
	return b.childs[k].(Node)
}
