package store

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
