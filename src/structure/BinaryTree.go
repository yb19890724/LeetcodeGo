package structure

// 二叉树
// 1.每个节点度为2 最多拥有两棵树
// 2.左右字数是有顺序的

// 1.非空二叉树的第i层,最多有2^(i-1) 个节点 (i>=1)
// 2.高度为h的二叉树上最有有 (2^h)-1个节点 (h>=1)
// 3.非空二叉树 如果叶子节点个数为n0 度为2的节点个数为n2 则有n0 = n2+1
// 	二叉树的总节点：n =n0 + n1 + n2
//

type BinaryTreeNode struct {
	LeftNode  *BinaryTree
	RightNode *BinaryTree
	Data      int
}

// 二叉树
type BinaryTree struct {
	Size int
	root *BinaryTreeNode
}

// 创建新树
func NewBinaryTree(data int) *BinaryTreeNode {
	return &BinaryTreeNode{Data: data}
}

// 注意判断：如果为空直接添加
// 从根节点开始比较：如果大于往右节点找，如果小于往左节点找
func (d *BinaryTree) Add(n int) {

	node := NewBinaryTree(n)

	// 判断根节点是空的
	if nil == d.root{
		d.root = node
	} else {

	}

	d.Size++
}

func main() {

}
