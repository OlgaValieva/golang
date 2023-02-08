package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func printTree(prefix string, tree *TreeNode, isLeft bool) {
	if tree != nil {
		fmt.Print( prefix)
		if isLeft {
			fmt.Print("├────")
		} else {
			fmt.Print("└────")
		}
		if tree.HasToy {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		if isLeft {
			printTree(prefix+"│   ", tree.Left, true)
		} else {
			printTree(prefix+"    ", tree.Left, true)
		}
		if isLeft {
			printTree(prefix+"│   ", tree.Right, false)
		} else {
			printTree(prefix+"    ", tree.Right, false)
		}
	}
}

func throughLevel(level int, curLevel int, tree *TreeNode) ([]bool, error) {
	var result []bool
	var tmp []bool
	if level % 2 == 0 {
		if tree.Right != nil && curLevel != level {
			tmp, _ := throughLevel(level, curLevel+1, tree.Right)
			result = append(result, tmp...)
		}
		if tree.Left != nil && curLevel != level {
			tmp, _ := throughLevel(level, curLevel+1, tree.Left)
			result = append(result, tmp...)
		}
	} else {
		if tree.Left != nil && curLevel != level {
			tmp, _ := throughLevel(level, curLevel+1, tree.Left)
			result = append(result, tmp...)
		}
		if tree.Right != nil && curLevel != level {
			tmp, _ := throughLevel(level, curLevel+1, tree.Right)
			result = append(result, tmp...)
		}
	}
	if curLevel == level {
		result = append(result, tree.HasToy)
	}
	_ = tmp
	return result, nil
}

func unrollGarland(tree *TreeNode) []bool {
	var slice []bool
	var tmp []bool
	for i := 0; i < 100; i++ {
		tmp, _ = throughLevel(i, 0, tree)
		slice = append(slice, tmp...)
	}
	return slice
}

func main() {
	fmt.Println("unrollGarland:")
	var Xtree = &TreeNode{true,
		&TreeNode{true,
			&TreeNode{true, nil, nil},
			&TreeNode{false, nil, nil}},
		&TreeNode{false,
			&TreeNode{true, nil, nil},
			&TreeNode{true, nil, nil}}}
	printTree("", Xtree, false,)
	fmt.Println("This return: ", unrollGarland(Xtree))

	fmt.Println("\nOneUnrollGarland:")
	var XtreeOne = &TreeNode{false, nil, nil}
	printTree("", XtreeOne, false)
	fmt.Println("This return: ", unrollGarland(XtreeOne))

	fmt.Println("\nBigUnrollGarland:")
	var Bigtree = &TreeNode{true,
		&TreeNode{true,
			&TreeNode{true, nil, &TreeNode{true,
				&TreeNode{true, nil, nil},
				&TreeNode{false, nil, nil}}},
			&TreeNode{false, nil, nil}},
		&TreeNode{false,
			&TreeNode{true, nil, nil},
			&TreeNode{true, nil, nil}}}
	printTree("", Bigtree, false,)
	fmt.Println("This return: ", unrollGarland(Bigtree))
}