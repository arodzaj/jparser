package store

import "strconv"

// Parse -
func Parse(data interface{}) Node {

	switch data.(type) {
	case string:
		l := new(Leaf)
		l.value = data.(string)
		return l

	case bool:
		l := new(Leaf)
		l.value = strconv.FormatBool(data.(bool))
		return l

	case float64:
		l := new(Leaf)
		l.value = strconv.FormatFloat(data.(float64), 'f', -1, 64)
		return l

	case []interface{}:
		a := new(List)
		return a

	case map[string]interface{}:
		a := new(Branch)
		return a

	}
	return nil
}

// Node -
type Node interface {
	Type() string
	Value() string
	Child(key interface{}) Node
}

// Leaf -
type Leaf struct {
	value string
}

// Type -
func (l *Leaf) Type() string {
	return "leaf"
}

// Value -
func (l *Leaf) Value() string {
	return l.value
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

// Value -
func (l *List) Value() string {
	return ""
}

// Child -
func (l *List) Child(key interface{}) Node {
	k := key.(int)
	return l.childs[k]
}

// Branch -
type Branch struct {
	childs map[string]interface{}
}

// Type -
func (b *Branch) Type() string {
	return "branch"
}

// Value -
func (b *Branch) Value() string {
	return ""
}

// Child -
func (b *Branch) Child(key interface{}) Node {
	k := key.(string)
	return b.childs[k].(Node)
}
