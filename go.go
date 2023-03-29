package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Height int
}

func main() {
	var n int
	fmt.Println("Ingrese la cantidad de números a agregar al árbol:")
	fmt.Scan(&n)
	var root *Node
	for i := 0; i < n; i++ {
		var val int
		fmt.Printf("Ingrese el número %d: ", i+1)
		fmt.Scan(&val)
		root = insert(root, val)
	}
	fmt.Println("Árbol Equilibrado AVL.")
	printTree(root, 0)
	fmt.Println("Recorrido en preorden:")
	preOrder(root)
	fmt.Println("\nRecorrido en inorden:")
	inOrder(root)
	fmt.Println("\nRecorrido en postorden:")
	postOrder(root)
	fmt.Printf("\nAltura del árbol AVL: %d\n", root.Height)
	fmt.Println("Nodos Hoja del árbol:")
	printLeaves(root)
}

func insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{Value: val, Height: 1}
	}
	if val < root.Value {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	root.Height = int(math.Max(float64(height(root.Left)), float64(height(root.Right)))) + 1
	balance := getBalance(root)
	if balance > 1 && val < root.Left.Value {
		return rotateRight(root)
	}
	if balance < -1 && val > root.Right.Value {
		return rotateLeft(root)
	}
	if balance > 1 && val > root.Left.Value {
		root.Left = rotateLeft(root.Left)
		return rotateRight(root)
	}
	if balance < -1 && val < root.Right.Value {
		root.Right = rotateRight(root.Right)
		return rotateLeft(root)
	}
	return root
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func rotateRight(node *Node) *Node {
	left := node.Left
	right := left.Right
	left.Right = node
	node.Left = right
	node.Height = int(math.Max(float64(height(node.Left)), float64(height(node.Right)))) + 1
	left.Height = int(math.Max(float64(height(left.Left)), float64(height(left.Right)))) + 1
	return left
}

func rotateLeft(node *Node) *Node {
	right := node.Right
	left := right.Left
	right.Left = node
	node.Right = left
	node.Height = int(math.Max(float64(height(node.Left)), float64(height(node.Right)))) + 1
	right.Height = int(math.Max(float64(height(right.Left)), float64(height(right.Right)))) + 1
	return right
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.Value)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.Left)
		fmt.Printf("%d ", node.Value)
		inOrder(node.Right)
	}
}

func postOrder(node *Node) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		fmt.Printf("%d ", node.Value)
	}
}

func printLeaves(node *Node) {
	if node != nil {
		if node.Left == nil && node.Right == nil {
			fmt.Printf("%d ", node.Value)
		}
		printLeaves(node.Left)
		printLeaves(node.Right)
	}
}

func printTree(node *Node, level int) {
	if node != nil {
		printTree(node.Right, level+1)
		for i := 0; i < level; i++ {
			fmt.Print("\t")
		}
		fmt.Printf("%d\n", node.Value)
		printTree(node.Left, level+1)
	}
}
