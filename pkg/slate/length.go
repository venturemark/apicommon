package slate

func nodeLength(node Node) int {
	if node.Type == "" {
		return len(node.Text)
	}

	var totalLength int
	for _, n := range node.Children {
		totalLength += nodeLength(n)
	}

	return totalLength
}

func Length(nodes Nodes) int {
	var totalLength int
	for _, n := range nodes {
		totalLength += nodeLength(n)
	}

	return totalLength
}
