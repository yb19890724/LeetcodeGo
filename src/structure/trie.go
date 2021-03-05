package structure

import (
	"unicode/utf8"
)

//  trie 存储 单词  cat,add,dog,doggy,cast,does 结构
//          [ 根节点 ]
//          /   |   \
//         a    c    d
//        /     |     \
//       d      a      o
//       |     / \   /   \
//       d    s  t   g   e
//            |      |   |
//            t      g   s
//                   |
//                   y


// 节点
type Node struct {
	Parent     *Node          // 父节点
	ParentName rune           // 父节点名字
	Children   map[rune]*Node // key是字符 v就是key节点的子节点
	Word       string         // 节点值
	Ending     bool           // 是都是一个单词的结尾
}

//
type Trie struct {
	Size int
	Root *Node // 根节点
}

// 清理所有
func (t *Trie) Clear() {
	t.Size = 0
	t.Root = nil
}

// 是否是空
func (t *Trie) isEmpty() bool {
	return t.Size == 0
}

// 清理所有
func (t *Trie) Add(k string, v string) string {
	
	if false == t.CheckKey(k) {
		return ""
	}
	
	p := t.Root
	
	for _, v := range k {
		
		// 节点不存在
		if p.Children[v] == nil {
			
			n := new(Node)
			
			// 父节点
			n.Parent = p
			n.ParentName = v
			
			// 初始化map
			n.Children = make(map[rune]*Node)
			
			// 初始化节点单词标志为假
			n.Ending = false
			
			p.Children[v] = n
		}
		p = p.Children[v]
	}
	
	// 之前存储在 覆盖 返回原始值
	if p.Ending {
		
		ov := p.Word
		
		p.Word = v
		
		return ov
	}
	
	// 新添加 当前节点是一个词的尾节点
	p.Ending = true
	// 单词
	p.Word = v
	
	return ""
}

// 获取子节点值
func (t *Trie) children(c map[rune]*Node) map[rune]*Node {
	
	if c == nil {
		c = make(map[rune]*Node)
	}
	
	return c
}

// 获取节点
func (t *Trie) Node(k string) *Node {
	
	// 节点是否有数据
	if t.Root == nil {
		return nil
	}
	
	// key是否存储在
	if t.CheckKey(k) == false {
		return nil
	}
	
	p := t.Root
	
	var n *Node
	
	// 按字符切串
	for _, v := range k {
		
		// 从节点一直向下找，如果没找到 单词不存在
		n = t.children(p.Children)[v]
		
		// 空节点
		if n == nil {
			return nil
		}
		
	}
	
	// @todo 节点存在不代表单词存在
	if true == n.Ending {
		return n
	}
	
	return nil
}

// 1.先检测是否存在，在进行删除

func (t *Trie) Remove(k string) string {
	
	n := t.Node(k)
	
	// 1.单词不存在
	// 2.不是单词尾节点
	if n == nil || n.Ending == false {
		return ""
	}
	
	t.Size --
	
	var w = n.Word // 最后一个节点的value
	
	// 如果还有子节点,把他的标记取消
	if n.Children != nil {
		
		n.Ending = false
		n.Word = ""
		return w
	}
	
	// 如果没有子节点，从后往回删除
	// 拿到父节点(parent)
	
	for p := n.Parent; p != nil; {
		
		// 删除当前节点
		delete(p.Children, p.ParentName)
		
		// @todo 如果删除的时候发现父节点是一个单词结尾或者父节点还有其他子节点，直接返回
		if true == p.Ending || p.Children != nil {
			break
		}
		
		// 继续往上走
		n = p
	}
	
	return w
}



// 校验key是否为空
func (t *Trie) CheckKey(k string) bool {
	if utf8.RuneCountInString(k) < 1 {
		return false
	}
	return true
}