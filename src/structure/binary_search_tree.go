package structure

// 存储数据必须有可比较性
type BstCompareer interface {
	// 获取可比较的元素值 用来进行比较
	Value(e BstCompareer) int
}

// 数据
type Data struct {
	Age int
}

func (d *Data) Value() int {
	return d.Age
}

// 数据2
type Data2 struct {
	Age int
}

func (d *Data2) Value() int {
	return d.Age
}


// 比较大的比较小
func (d *Data) Compare(Data Data) int {
	return d.Value() - Data.Value()
}

// 比较小的比较大
func (d *Data2) Compare(Data Data) int {
	return  Data.Value() -d.Value()
}

// 比较器
type Comparator interface {
	// 元素比较函数
	/**
	 *  e1 == e2 return = 0
	 *  e1 > e2 return  > 0
	 *  e1 < e2 return  < 0
	 */
	
	Compare (e1 BstCompareer,e2 BstCompareer)int
}

//
type Bster interface {
    Size() // 元素数量
    isEmpty() // 是否为空
    Clear() // 清空树
    Add(e string) // 添加节点
    Remove() // 移除节点
    Contains() // 是否包含元素
}

// 二叉搜索树
type Bst struct {
	Num int // 元素数量
	Root *BstNode // 根节点
	Comparator Comparator // 比较器
}

// 节点
type BstNode struct {
	Value  BstCompareer
	Left   *BstNode
	Right  *BstNode
	Parent *BstNode
}

func (b *Bst) Size()  {
	
}

func (b *Bst) Node()  {

}


// 添加节点先找到父节点，创建新节点 parent.left = node 或者 parent.right = node
// 如果值相等： 1.直接return 2.覆盖节点
func (b *Bst) Add(v BstCompareer)  {
	
	// 第一个节点
	if b.Root == nil { // 空节点
		b.Root = &BstNode{
			Value: v,
		}
		b.Num ++
		return
	}
	
	// 找到父节点
	
	n:=b.Root
	
	
	var (
		parent *BstNode
		cmp int
	)
	
	for n.Value != nil {
		
		cmp = b.Comparator.Compare(v, n.Value)
		
		parent = n.Parent // 每次保存父节点
		
		// 如果添加元素大于当前节点的元素往右找
		if cmp > 0 {
			n =n.Right
		}
		
		// 如果添加元素大于当前节点的元素往左找
		if cmp < 0 {
			n =n.Left
		}
		
		// 如果添加元素等于当前节点的元素 覆盖或者直接返回
		if cmp == 0 {
			return
		}
	}
	
	// 找到父节点，看放到哪个位置
	nn := &BstNode{
		Value: v,
	}
	
	if cmp > 0 {
		parent.Right = nn
	}
	
	if cmp < 0  {
		parent.Left = nn
	}
	
	b.Num++
}

