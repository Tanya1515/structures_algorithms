package main

// Сложность поиска элемента в бинарном дереве по времени: O(H), где H - высота дерева.

// Задание: необходимо найти элемент в бинарном дереве (рекурсивным и обычным способом).

func findElemRecursive(root *Node, elem int) bool {

	if root == nil {
		return false
	}

	if root.elem == elem {
		return true
	}

	if root.elem > elem {
		return findElemRecursive(root.left, elem)
	}

	return findElemRecursive(root.right, elem)
}

func finElem(root *Node, elem int) bool {

	rootNode := root

	for rootNode != nil {
		if rootNode.elem == elem {
			return true
		}
		if rootNode.elem > elem {
			rootNode = rootNode.left
			continue
		}
		rootNode = rootNode.right
	}
	return false
}
