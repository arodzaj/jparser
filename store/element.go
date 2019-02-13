package store

import (
	"fmt"
	"strconv"
)

func parse(data interface{}) Node {

	switch data.(type) {
	case int:
		l := new(Leaf)
		l.value = strconv.Itoa(data.(int))
		return l

	case float64:
		l := new(Leaf)
		l.value = strconv.FormatFloat(data.(float64), 'f', -1, 64)
		return l

	case string:
		l := new(Leaf)
		l.value = "\"" + data.(string) + "\""
		return l

	case bool:
		l := new(Leaf)
		l.value = strconv.FormatBool(data.(bool))
		return l

	case []interface{}:
		a := new(Array)
		a.childs = []Node{}

		for _, v := range data.([]interface{}) {
			a.childs = append(a.childs, parse(v))
		}

		return nil

	case map[string]interface{}:
		o := new(Branch)
		o.childs = map[string]Node{}

		cnt := 0
		for k, v := range data.(map[string]interface{}) {
			o.childs[k] = parse(v)
			cnt++
		}
		o.length = cnt + 1

		return o
	}

	return nil
}

// Node
type Node interface {
	String() string
	Value() string
	Type() string
	Childs() map[string]Node
}

// Branch -
type Branch struct {
	length int
	childs map[string]Node
}

func (b *Branch) String() string {
	return fmt.Sprintf("<type:branch,childs:%d>", b.length)
}

func (b *Branch) Value() string {
	return ""
}

func (b *Branch) Type() string {
	return "branch"
}

func (b *Branch) Childs() map[string]Node {
	return b.childs
}

type Leaf struct {
	value string
}

func (l *Leaf) String() string {
	return fmt.Sprintf("<type:Leaf,value:%s>", l.value)
}

func (l *Leaf) Value() string {
	return l.value
}

func (l *Leaf) Type() string {
	return "leaf"
}

func (l *Leaf) Childs() map[string]Node {
	return map[string]Node{}
}

type Array struct {
	childs []Node
}

func (a *Array) String() string {
	return fmt.Sprintf("<type:array,childs:%d>", len(a.childs))
}

func (a *Array) Value() string {
	return ""
}

func (a *Array) Type() string {
	return "array"
}

func (a *Array) Childs() map[string]Node {
	els := map[string]Node{}

	for k, v := range a.childs {
		key := strconv.Itoa(k)
		els[key] = v
	}

	return els
}
