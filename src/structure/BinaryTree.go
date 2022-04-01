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
	LeftNode  *BinaryTreeNode
	RightNode *BinaryTreeNode
	Data      int
}

// 二叉树
type BinaryTree struct {
	Size int
	Root *BinaryTreeNode
}

func (b *BinaryTree) binaryTreeNodeChildrenNotEmpty(node *BinaryTreeNode) bool {
	return node.LeftNode != nil && node.RightNode != nil
}

// 注意判断：如果为空直接添加
// 从根节点开始比较：如果大于往右节点找，如果小于往左节点找
// 最后找到需要存放的根节点， 大于放右面，小于放左面
// 遇到值一样的覆盖或者不处理 直接跳出

// 如何找到最后添加的节点：的判断条件

func (b *BinaryTree) Add(n int) {

	// 判断根节点是空的
	if nil == b.Root {
		b.Root = &BinaryTreeNode{Data: n}
	} else {

		curr := b.Root

		for nil != curr.LeftNode || nil != curr.RightNode {

			// 相等不处理
			if curr.Data == n {
				return
			}

			// 大于往右找
			if curr.Data < n {
				curr = curr.RightNode
			}

			// 小于往左找
			if curr.Data > n {
				curr = curr.LeftNode
			}
		}

		node := &BinaryTreeNode{Data: n}

		if node.Data > curr.Data {
			curr.RightNode = node
		} else {
			curr.LeftNode = node
		}

	}

	b.Size++
}

// 删除度为1节点：
//1. 删除根节点直接删除, 如果下面有节点，重置根节点，原来的子节点变成根节点
//2. 删除子节点，用子节点代替原来的节点即可  left 或者是 right

// 删除度为2节点：
// 1.删除节点的时候需要判断，用左节点替代还是右节点替代 用左子树的<最大>或者右子树的<最小>
//  左子树最右 ||，右子树最左
//

// 如果该节点是叶子节点，则直接删除它：root = null。
// 如果该节点不是叶子节点且有右节点，则用它的后继节点的值替代 root.val = successor.val，然后删除后继节点。
// 如果该节点不是叶子节点且只有左节点，则用它的前驱节点的值替代 root.val = predecessor.val，然后删除前驱节点。

// 三种情况
// 如果目标节点没有子节点，我们可以直接移除该目标节点。
// 如果目标节只有一个子节点，我们可以用其子节点作为替换。
// 如果目标节点有两个子节点，我们需要用其中序后继节点或者前驱节点来替换，再删除该目标节点。

func (b *BinaryTree) Remove(node *BinaryTreeNode, n int) *BinaryTreeNode {

	if node == nil {
		return nil
	}

	if node.Data == n {

		// 删除节点为叶子节点, 下面没有节点了
		if node.LeftNode == nil && node.RightNode == nil {
			node = nil
			return node
		}

		//  删除节点有右子树
		if nil != node.RightNode && node.LeftNode == nil {
			node.Data = maxNode(node.RightNode)
			node.RightNode = b.Remove(node.RightNode, node.Data)
		}

		// 删除节点无右子树但存在左子树
		if nil != node.LeftNode && node.RightNode == nil {
			node.Data = minNode(node.RightNode)
			node.LeftNode = b.Remove(node.LeftNode, node.Data)
		}
	}

	if n > node.Data {
		node.RightNode = b.Remove(node.RightNode, n)
	} else {
		node.LeftNode = b.Remove(node.LeftNode, n)
	}

	return node
}

// 找到最右侧最小的节点
func minNode(node *BinaryTreeNode) int {
	for node.LeftNode != nil {
		node = node.LeftNode
	}
	return node.Data
}

//找到最左侧最大的节点
func maxNode(node *BinaryTreeNode) int {
	for node.RightNode != nil {
		node = node.RightNode
	}
	return node.Data
}

func (b *BinaryTree) SearchNode(n int) *BinaryTreeNode {

	node := b.Root

	for nil != node.LeftNode || nil != node.RightNode {

		// 相等不处理
		if node.Data == n {
			return node
		}

		// 大于往右找
		if node.Data < n {
			node = node.RightNode
		}

		// 小于往左找
		if node.Data > n {
			node = node.LeftNode
		}
	}

	return node
}

func (b *BinaryTree) lastFor() []int {
	var res []int

	lastLoop(b.Root, &res)

	return res
}

func lastLoop(node *BinaryTreeNode, res *[]int) {
	if nil != node {
		PreLoop(node.LeftNode, res)
		PreLoop(node.RightNode, res)
		*res = append(*res, node.Data)
	}
}

func (b *BinaryTree) PreFor() []int {
	var res []int

	PreLoop(b.Root, &res)

	return res
}

func PreLoop(node *BinaryTreeNode, res *[]int) {

	if nil != node {
		*res = append(*res, node.Data)
		PreLoop(node.LeftNode, res)
		PreLoop(node.RightNode, res)
	}
}

func (b *BinaryTree) MiddleFor() []int {
	var res []int

	middleLoop(b.Root, &res)

	return res
}

func middleLoop(node *BinaryTreeNode, res *[]int) {

	if nil != node {
		middleLoop(node.LeftNode, res)
		*res = append(*res, node.Data)
		middleLoop(node.RightNode, res)
	}
}

func (b *BinaryTree) MaxDepth() int {
	return MaxDepth(b.Root)
}

func MaxDepth(node *BinaryTreeNode) int {

	if nil == node {
		return 0
	}

	a := MaxDepth(node.LeftNode)
	b := MaxDepth(node.RightNode)
	return max(a, b) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
