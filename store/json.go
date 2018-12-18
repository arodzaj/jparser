package store

type json struct {
}

func (j *json) String() string {
	return ""
}

type Element interface {
	String() string
	Parent() *Element
	Value() string
	Type() string // leaf, array, object

	Childs() map[string]Element
}

// type Element struct {
// 	elType int
// 	value  string
// 	level  int
// }
