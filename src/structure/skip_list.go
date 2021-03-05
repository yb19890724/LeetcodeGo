package structure

import (
	"math/rand"
	"time"
)

const MAX_LEVEL=10

// 跳表
type SkipList struct {
	maxLevel int // 长度
	root []*SkipNode
}

// 节点
type SkipNode struct {
	level int
	prev  []*SkipNode
	next  []*SkipNode
	value int // 值
	count int // 重复数量
}

// 创建新节点
func NewNode(level, value int) *SkipNode {
	return &SkipNode{
		level: level,
		value: value,
		next:  make([]*SkipNode, level+1),
		prev:  make([]*SkipNode, level+1),
		count: 1,
	}
}

// 创建新调表
func NewSkipList() SkipList{
	return SkipList{
		maxLevel: -1,
	}
}

// 添加节点
func (n *SkipNode) AddNodeByLevel(level int, nn *SkipNode) {
	
	if n == nil || n.prev == nil || n.next == nil || nn == nil || nn.prev == nil || nn.next == nil {
		return
	}
	
	// 如果当前节点 大于添加节点
	if n.value > nn.value {
		
		// 取出当前节点的前一个节点
		prev := n.prev[level]
		
		n.prev[level] = nn
		nn.next[level] = n
		
		nn.prev[level] = prev
		
		if prev != nil {
			prev.next[level] = nn
		}
	} else {
		next := n.next[level]
		
		n.next[level] = nn
		nn.prev[level] = n
		
		nn.next[level] = next
		if next != nil {
			next.prev[level] = nn
		}
	}
}



func (s *SkipList) randomUpgrade() bool {
	rand.Seed(time.Now().UnixNano())
	
	r := rand.Intn(7) % 2
	
	if s.maxLevel > MAX_LEVEL {
		return false
	}
	
	// fmt.Println("rand:---",r)
	return r == 1
}

// 添加节点
func (s *SkipList) Add(num int)  {
	
	// 查找当前节点
	n:=s.SearchPos(num)
	
	if n==nil{
		s.maxLevel=0
		n=NewNode(0,num)
		s.root=append(s.root,n)
		return
	}
	
	if n.value==num{
		n.count++
		return
	}
	
	
	if s.randomUpgrade(){
		s.maxLevel++
		nn:=NewNode(s.maxLevel,num)
		s.root=append(s.root,nn)
		
		for i:=1;i<=s.maxLevel-1;i++{
			in:=s.root[i]
			for in!=nil && in.value>num  && in.prev!=nil && in.prev[i]!=nil{
				in=in.prev[i]
			}
			for in!=nil && in.value<num && in.next!=nil && in.next[i]!=nil{
				in=in.next[i]
			}
			in.AddNodeByLevel(i,nn)
		}
		
		n.AddNodeByLevel(0,nn)
	}else{
		nn:=NewNode(0,num)
		n.AddNodeByLevel(0,nn)
	}
}








// 搜索值
func (s *SkipList) Search(v int) bool {
	n := s.SearchPos(v)
	
	if n == nil {
		return false
	}
	return n.value == v
}

// 匹配值与节点值
func (s *SkipList) SearchPos(v int) *SkipNode {
	
	if s.maxLevel == -1 {
		return nil
	}
	
	next :=s.root[s.maxLevel]
	
	// 从最后一个节点开始 查找节点
	for l := next.level; l >= 0; l-- {
		
		// 找到对应value
		if v == next.value {
			return next
		}
		
		// 匹配值小于当前节点值
		if v < next.value {
			for next.prev[l]!=nil && v<=next.prev[l].value{
				next=next.prev[l]
			}
		}else{
			for next.next[l]!=nil && v<=next.next[l].value{
				next=next.next[l]
			}
		}
		
		
	}
	
	return next
}

