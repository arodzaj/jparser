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

// List -
type List struct {
}

// Type -
func (l *List) Type() string {
	return "list"
}

// Value -
func (l *List) Value() string {
	return ""
}

// Branch -
type Branch struct {
}

// Type -
func (b *Branch) Type() string {
	return "branch"
}

// Value -
func (b *Branch) Value() string {
	return ""
}
