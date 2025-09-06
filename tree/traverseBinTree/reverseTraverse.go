package main

import "fmt"

// Задание: необхожимо реализовать обратный обход дерева.

// Обратный обход - обход, аналогичный прямому обходу дерева, но теперь
// сначала рассматриваются все поддеревья, и только потом корень поддерева.
// То есть в этом слуае узел будет обрабатываться, когда мы уходим из него навсегда.

func traverseReverse(root *Node) {
	if root == nil {
		return
	}
	
	traverseReverse(root.left)
	traverseReverse(root.right)
	fmt.Println(root.elem)
}
