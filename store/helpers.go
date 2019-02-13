package store

import "strings"

func getNode(path string, root Node) Node {
	p := strings.Split(path, ".")
	cur := root
	var result Node

	for {
		if len(p) > 0 {
			break
		}
		head, tail := p[0], p[1:]

		// switch cur.Type == '' {
		// }

	}

	return result
}
