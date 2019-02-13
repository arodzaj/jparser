package store

import "strings"

func getNode(path string, root Node) Node {
	p := strings.Split(path, ".")
	if len(p) == 1 && p[0] == "" {
		return root
	}

	cur := root
	var key string

	for {
		key, p = p[0], p[1:]
		cur = cur.Child(key)

		if len(p) == 0 || cur == nil {
			break
		}
	}

	return cur
}
