package main

import "fmt"

func main() {

	l1 := CreateList()

	els := []ElType{{name: "xxb1", tel: 111}, {name: "xxb2", tel: 222}, {name: "xxb3", tel: 333}}

	//l1.addTailElem(els)
	l1.addTailElem(els)
	//l1.Insert(2,ElType{name:"xxb4",tel:444})
	fmt.Println(l1.Head.Next.Data)
	fmt.Println(l1.Head.Next.Next.Data)
	fmt.Println(l1.Head.Next.Next.Next.Data)
	fmt.Println(l1.Length)
	l1.Delete(3)
	fmt.Println(l1.Head.Next.Data)
	fmt.Println(l1.Head.Next.Next.Data)
	fmt.Println(l1.Length)


}

type ElType struct {
	name string
	tel  int
}

type Node struct {
	Data ElType
	Next *Node
}

type List struct {
	Head   *Node
	Length int
}

func CreateList() *List {
	head := new(Node)
	return &List{Head: head, Length: 0}
}

//头插法 初始化插入联系人
func (l *List) addHeadElem(el []ElType) bool {
	for i := 0; i < len(el); i++ {
		node := &Node{Data: el[i], Next: nil}
		//把头 节点 l.Head.Next 指向 给新节点的NEXT
		node.Next = l.Head.Next
		l.Head.Next = node
		l.Length++
	}
	return true
}

//尾插法 初始化插入联系人
func (l *List) addTailElem(el []ElType) bool {
	r := l.Head
	for i := 0; i < len(el); i++ {
		node := &Node{Data: el[i], Next: nil}
		r.Next = node
		r = node
		l.Length++
	}
	return true
}

func (l *List) addTableBefore(el ElType) bool {
	node := &Node{Data: el, Next: nil}
	//把头 节点 l.Head.Next 指向 给新节点的NEXT
	node.Next = l.Head.Next
	l.Head.Next = node
	l.Length++
	return true
}

func (l *List) addTableAfter(el ElType) {
	node := &Node{Data: el, Next: nil}
	p := l.Head
	if p.Next != nil {
		for p.Next != nil {
			p = p.Next
		}
	}
	p.Next = node
	l.Length++

}

//根据位置插入
func (l *List) Insert(index int, el ElType) bool {
	if index < 1 {
		return false
	}

	if index > l.Length {
		l.addTableAfter(el)
		return true
	}

	if l.Head.Next == nil {
		l.addTableBefore(el)
	}
	p := l.Head
	for i := 1; i < index; i++ {
		p = p.Next
	}
	node := &Node{Data: el}
	node.Next = p.Next
	p.Next = node
	l.Length++

	return true
}

//删除指定位置的链表元素。
func (l *List) Delete(index int) bool {
	if index < 1 || index > l.Length {
		return false
	}
	p := l.Head
	for i := 1; i < index; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	l.Length--
	return true

}
