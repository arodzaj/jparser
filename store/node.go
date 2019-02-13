// Package store provides primitives and logic of loaded JSON files
package store

import (
	"fmt"
	"strconv"
)

// Node interface respresents abstraction a node in JSON tree
type Node interface {
	Type() string
	String() string
	Child(key interface{}) Node
	ChildKeys() []string
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

func (l *Leaf) ChildKeys() []string {
	return []string{}
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
	var k int
	switch key.(type) {
	case int:
		k = key.(int)
		break
	case string:
		k, _ = strconv.Atoi(key.(string))
		break

	}

	if k < len(l.childs) {
		return l.childs[k].(Node)
	}

	return nil

}

func (l *List) ChildKeys() []string {
	res := []string{}

	for i := range l.childs {
		res = append(res, strconv.Itoa(i))
	}

	return res
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

	if val, ok := b.childs[k]; ok {
		return val.(Node)
	}

	return nil

}

func (b *Branch) ChildKeys() []string {
	res := []string{}

	for k := range b.childs {
		res = append(res, k)
	}

	return res
}
