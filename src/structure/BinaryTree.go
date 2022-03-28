package structure

// 二叉树
type BinaryTree struct {
	LeftNode  *BinaryTree
	RightNode *BinaryTree
	Data      int
}

// 创建新树
func NewBinaryTree(data int) *BinaryTree {
	return &BinaryTree{Data: data}
}

