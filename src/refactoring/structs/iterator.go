package _struct

import (
	"strings"
)

// 迭代器操作

type Goods struct {
	Id    int
	Name  string
	Price int
	Desc  string
}


type Goodss []*Goods


func (g Goodss) Each(f func(goods *Goods)) {
	
	for _, d := range g {
		f(d)
	}
	
}

func (g *Goods) GoodsDesc() {
	g.Desc = strings.Join([]string{g.Name, g.Name}, "test")
}

func TestEach()  {


}