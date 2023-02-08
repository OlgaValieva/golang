package main

import "fmt"

type TreeNode struct {
	hasToy bool
	left   *TreeNode
	right  *TreeNode
	
}

func aretoysBalanced(tree *TreeNode) bool {
	if tree.right == nil && tree.left == nil {
		return true
	} else if tree.right == nil {
		return toys(tree.left) == 0
	} else if tree.left == nil {
		return toys(tree.right) == 0
	}
	return toys(tree.left) == toys(tree.right)
}

func toys (tree *TreeNode) int {
	var sum int
	var toy int

	if tree == nil {
		return 0
	}
	if tree.hasToy == true {
		toy = 1
	}
	sum = toy + toys (tree.left) + toys (tree.right)
	return sum
}

func printTree(prefix string, tree *TreeNode, isLeft bool) {
	if tree != nil {
		fmt.Print(prefix)
		if isLeft {
			fmt.Print("├────")
		} else {
			fmt.Print("└────")
		}
		if tree.hasToy {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		if isLeft {
			printTree(prefix + "│   ", tree.left, true)
		} else {
			printTree(prefix + "    ", tree.left, true)
		}
		if isLeft {
			printTree(prefix + "│   ", tree.right, false)
		} else {
			printTree(prefix + "    ", tree.right, false)
		}
	}
}

func main() {
	fmt.Println("\nTrue_1:")
	var example1 = &TreeNode{false,
		&TreeNode{false,
			&TreeNode{false, nil, nil},&TreeNode{true, nil, nil}},
		&TreeNode{true, nil, nil}}
	printTree("", example1, false)
	fmt.Println( "This return: ", aretoysBalanced(example1))

	fmt.Println("\nTrue_2:")
	var example2 = &TreeNode{true,
		&TreeNode{true,
			&TreeNode{true, nil, nil},
			&TreeNode{false, nil, nil}},
		&TreeNode{false,
			&TreeNode{true, nil, nil},
			&TreeNode{true, nil, nil}}}
	printTree("", example2, false)
	fmt.Println("This return: ",  aretoysBalanced(example2))
		
	fmt.Println("\nFalse_1:")
	var example3 = &TreeNode{true,
		&TreeNode{true, nil, nil},
		&TreeNode{false, nil, nil}}
	printTree("", example3, false)
	fmt.Println("This return: ",  aretoysBalanced (example3))

	fmt.Println("\nFalse_2:")
	var example4 = &TreeNode{false,
		&TreeNode{true, nil,
			&TreeNode{true, nil, nil}},
		&TreeNode{false, nil,
			&TreeNode{true, nil, nil}}}
	printTree("", example4, false)
	fmt.Println("This return: ",aretoysBalanced(example4))

	fmt.Println("\nOne:")
	var tree5 = &TreeNode{false, nil, nil}
	printTree("", tree5, false)
	fmt.Println("This return: ", aretoysBalanced(tree5))
}