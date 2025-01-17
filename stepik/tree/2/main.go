/*
https://stepik.org/lesson/45970/step/2?unit=24123
Проверка свойства дерева поиска
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	if n == 0 {
		fmt.Println("CORRECT")
		return
	}

	root := treeBuild(n, os.Stdin)
	var vl validatorV1
	if vl.isValidBST(root) {
		fmt.Println("CORRECT")
	} else {
		fmt.Println("INCORRECT")

	}
}

// InOrder Traversal
type validatorV1 struct {
	prev *node
}

func (v *validatorV1) isValidBST(root *node) bool {
	if root == nil {
		return true
	}

	if !v.isValidBST(root.left) {
		return false
	}
	if v.prev != nil && v.prev.key >= root.key {
		return false
	}
	v.prev = root
	return v.isValidBST(root.rigth)
}

// Recursive Traversal
type validatorV2 struct{}

func (v *validatorV2) isValidBST(root, low, high *node) bool {
	if root == nil {
		return true
	}

	if low != nil && root.key <= low.key || high != nil && root.key >= high.key {
		return false
	}

	return v.isValidBST(root.left, low, root) && v.isValidBST(root.rigth, root, high)
}

// Iterative Traversal
type validatorV3 struct{}

func (v *validatorV3) isValidBST(root *node) bool {
	if root == nil {
		return true
	}
	lower := new(stack)
	upper := new(stack)
	stk := new(stack)
	stk.push(root)

	for !stk.isEmpty() {
		root := stk.pop()
		low := lower.pop()
		high := upper.pop()

		if low != nil && root.key <= low.key || high != nil && root.key >= high.key {
			return false
		}

		if root.rigth != nil {
			stk.push(root.rigth)
			lower.push(root)
			upper.push(high)
		}
		if root.left != nil {
			stk.push(root.left)
			lower.push(low)
			upper.push(root)
		}
	}

	return true
}

type stack struct {
	data []*node
}

func (st *stack) pop() *node {
	if len(st.data) == 0 {
		return nil
	}
	v := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return v
}
func (st *stack) push(v *node) {
	st.data = append(st.data, v)
}
func (st *stack) isEmpty() bool {
	return len(st.data) == 0
}

func treeBuild(n int, r io.Reader) *node {
	reader := bufio.NewReaderSize(r, 1024)

	var key, left, right int
	nodes := make([]*node, n)
	root := new(node)
	nodes[0] = root

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &key, &left, &right)

		if nodes[i] == nil {
			nodes[i] = new(node)
		}
		nodes[i].key = key

		if left != -1 {
			if nodes[left] == nil {
				nodes[left] = new(node)
			}
			nodes[i].left = nodes[left]
		}

		if right != -1 {
			if nodes[right] == nil {
				nodes[right] = new(node)
			}
			nodes[i].rigth = nodes[right]
		}
	}

	return root
}

type node struct {
	key         int
	left, rigth *node
}
