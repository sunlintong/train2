package main

import "fmt"

const initSize int=20 //定义常量栈的初始大小initSize为20

//栈结构体Stack
type Stack struct{
	size int //容量
	top int	//栈顶
	data []interface{} //用切片data作容器，定义为interface{}类型的切片以接收任意类型
}

//创建并初始化栈方法createStack，返回Stack
func createStack() Stack{
	s:=Stack{} //声明Stack变量s，注意：声明结构体变量得加{}
	s.size=initSize
	s.top=-1
	s.data=make([]interface{},initSize)
	return s
}

//操作Stack的方法isEmpty判断栈是否为空
func(s *Stack) isEmpty() bool{
	return s.top==-1
}
func(s *Stack) isFull() bool{
	return s.top==s.size-1
}

//入栈
func(s *Stack) push(e interface{}) bool{
	if(s.isFull()){
		fmt.Println("Stack is full,push failed")
		return false
	}
	s.top++
	s.data[s.top]=e
	return true
}
//出栈
func(s *Stack) pop() interface{}{
	if(s.isEmpty()){
		fmt.Println("Stack is empty,pop failed")
		return nil
	}
	e:=s.data[s.top]
	s.top--
	return e
}

//栈已有元素的长度
func(s *Stack) getLength() int{
	length:=s.top+1
	return length
}

//清空栈
func(s *Stack) clear(){
	s.top=-1
}

//遍历栈
func(s *Stack) traverse(){
	if(s.isEmpty()){
		fmt.Println("Stack is empty")
	}else{	//注意这里的else不能换到下一行
		for i:=0;i<=s.top;i++{
			fmt.Print(s.data[i]," ")
		}
		fmt.Println()
	}
}

//打印栈的当前信息
func(s *Stack) printInfo(){
	fmt.Println("###############################################")
	fmt.Println("栈容量为：",s.size)
	fmt.Println("栈顶指数为：",s.top)
	fmt.Println("长度为：",s.getLength())
	fmt.Println("栈是否为空:",s.isEmpty())
	fmt.Println("栈是否为满：",s.isFull())
	fmt.Print("遍历栈：")
	s.traverse()
	fmt.Println("###############################################")
}


func main(){
	//初始情况测试
	s:=createStack()
	fmt.Println("栈创建成功：")
	s.printInfo()	
	fmt.Println()
	fmt.Println("---------------------------现在开始测试入栈-------------------------")
	fmt.Println("用10个int测试入栈：")
	for i:=0;i<10;i++{
		s.push(i)
	}
	s.printInfo()
	fmt.Println()

	fmt.Println("再加10个string入栈：")
	var str string="a"
	for i:=10;i<20;i++{
		s.push(str)
		str+="a"
	}
	s.printInfo()

	fmt.Println()
	fmt.Println("此时栈已满，再加一个int入栈试试：")
	s.push(21)
	s.printInfo()

	fmt.Println("-----------------------现在开始测试出栈----------------------------")
	fmt.Println("先pop一个出栈顶试试")
	fmt.Println("pop出的元素是：",s.pop())
	s.printInfo()
	
	fmt.Println()
	fmt.Println("清空栈")
	s.clear()
	s.printInfo()

	fmt.Println()
	fmt.Println("再pop一个试试")
	s.pop()
	s.printInfo()
}
