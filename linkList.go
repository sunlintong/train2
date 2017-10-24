package main

import "fmt"

//链表结点
type Node struct{
	data interface{}
	next *Node
}

//链表
type LinkList struct{
	head *Node
	tail *Node
	size int
}

//新建空链表，即创建Node指针head，用来指向链表第一个结点，初始为空
func createLinkList() LinkList{
	l:=LinkList{}
	l.head=nil //head指向头部结点
	l.tail=nil //tail指向尾部结点
	l.size=0
	return l
}

//是否为空链表
func(l *LinkList) isEmpty() bool{
	return l.size==0
}

//获取链表长度
func(l *LinkList) getLength() int{
	return l.size
}

//是否含有指定结点
func(l *LinkList) exist(node *Node) bool{
	var p *Node=l.head
	for p!=nil{
		if(p==node){
			return true
		}else{
			p=p.next
		}
	}
	return false
}

//获取含有指定数据的第一个结点
func(l *LinkList) getNode(e interface{}) *Node{
	var p *Node=l.head
	for p!=nil{
		//找到该数据所在结点
		if e==p.data{
			return p
		}else{
			p=p.next
		}
	}
	return nil
}

//在链表尾部添加数据
func(l *LinkList) append(e interface{}){
	//为数据创建新结点
	newNode:=Node{}
	newNode.data=e
	newNode.next=nil	
	
	if(l.size==0){
		l.head=&newNode
		l.tail=&newNode
	}else{
		l.tail.next=&newNode
		l.tail=&newNode
	}
	l.size++
}

//在链表头部插入数据
func(l *LinkList) insertHead(e interface{}){
	newNode:=Node{}
	newNode.data=e
	newNode.next=l.head
	l.head=&newNode
	if(l.size==0){
	l.tail=&newNode	
	}
	l.size++
}

//在指定结点后面插入数据
func(l *LinkList) insertAfterNode(pre *Node,e interface{}){
	//如果链表中存在该结点，才进行插入
	if l.exist(pre){
		newNode:=Node{}
		newNode.data=e
		if(pre.next==nil){
			l.append(e)
		}else{
			newNode.next=pre.next
			pre.next=&newNode
		}
		l.size++
	}else{
		fmt.Println("链表中不存在该结点")
	}		
}

//在第一次出现指定数据的结点后插入数据,若链表中无该数据，返回false
func(l *LinkList) insertAfterData(preData interface{},e interface{}) bool{
	var p *Node=l.head
	for p!=nil{
		//找到该数据所在结点
		if p.data==preData{
			l.insertAfterNode(p,e)
			return true
		}else{
			p=p.next
		}
	}
	//没有找到该数据
	fmt.Println("链表中没有该数据，插入失败")
	return false
}

//在指定下标处插入数据
func(l *LinkList) insert(position int,e interface{}) bool{
	if position<0{
		fmt.Println("指定下标不合法")
		return false
	}else if position==0{
		//在头部插入
		l.insertHead(e)
		return true
	}else if position==l.size{
		//在尾部插入
		l.append(e)
		return true
	}else if position>l.size{
		fmt.Println("指定下标超出链表长度")
		return false
	}else{
		//在中间插入
		var index int
		var p *Node=l.head
		//逐个移动指针
		//position是插入后新结点的下标，position-1时需要定位到的其前一个结点的下标
		for index=0;index<position-1;index++{
			p=p.next
		}
		//找到
		l.insertAfterNode(p,e)
		return true
	}
	
}



//删除指定结点
func(l *LinkList) deleteNode(node *Node){
	//存在该结点
	if l.exist(node){
		//如果是头部结点
		if node==l.head{
			l.head=l.head.next
		//如果是尾部结点
		}else if node==l.tail{
			//寻找指向其前一个结点的指针
			var p *Node=l.head
			for p.next!=l.tail{
				p=p.next
			}
			p.next=nil
			l.tail=p
		//中间结点	
		}else{
			var p *Node=l.head
			for p.next!=node{
				p=p.next
			}
			p.next=node.next
		}
		l.size--
	}
}

//删除第一个含指定数据的结点
func(l *LinkList) delete(e interface{}){
	p:=l.getNode(e)
	if p==nil{
		fmt.Println("链表中无该数据，删除失败")
	}else{	
		l.deleteNode(p)
	}
}



//遍历链表
func(l *LinkList) traverse(){
	var p *Node=l.head
	if(l.isEmpty()){
		fmt.Println("LinkList is empty")
	}else{
		for p!=nil{
			fmt.Print(p.data," ")
			p=p.next
		}
		fmt.Println()
	}
}




//打印链表信息
func(l *LinkList) printInfo(){
	fmt.Println("###############################################")
	fmt.Println("链表长度为：",l.getLength())
	fmt.Println("链表是否为空:",l.isEmpty())
	fmt.Print("遍历链表：")
	l.traverse()
	fmt.Println("###############################################")
}


func main(){
	linkList:=createLinkList()
	fmt.Println("链表创建成功：")
	linkList.printInfo()	
	fmt.Println()
	fmt.Println("---------------------------现在开始测试插入数据-------------------------")
	fmt.Println("在空链表尾部加一个int 1")
	linkList.append(1)
	linkList.printInfo()

	fmt.Println()
	fmt.Println("在头部加一个int 0")
	linkList.insertHead(0)
	linkList.printInfo()

	fmt.Println()
	fmt.Println("在0后面插入一个int 2")
	linkList.insertAfterData(0,2)
	linkList.printInfo()

	fmt.Println()
	fmt.Println("在下标0插入一个int 100")
	flag:=linkList.insert(0,100)
	fmt.Println("插入成功？",flag)
	linkList.printInfo()

	fmt.Println()
	fmt.Println("在下标2插入一个int 200")
	flag=linkList.insert(2,200)
	fmt.Println("插入成功？",flag)
	linkList.printInfo()

	fmt.Println()
	fmt.Println("在下标5插入一个string sunlintong")
	flag=linkList.insert(5,"sunlintong")
	fmt.Println("插入成功？",flag)
	linkList.printInfo()

	fmt.Println()
	fmt.Println("在下标7插入一个int 400")
	flag=linkList.insert(7,400)
	fmt.Println("插入成功？",flag)
	linkList.printInfo()

	fmt.Println()
	fmt.Println("---------------------------现在开始测试删除数据-------------------------")
	fmt.Println("删除 0")
	linkList.delete(0)
	linkList.printInfo()

	fmt.Println()
	fmt.Println("删除 string sunlintong")
	linkList.delete("sunlintong")
	linkList.printInfo()

	fmt.Println()
	fmt.Println("删除 500")
	linkList.delete(500)
	linkList.printInfo()

}

