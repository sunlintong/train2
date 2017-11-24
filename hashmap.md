# golang实现hashmap数据结构
## 简介
哈希表是存放键值对的数据结构，在哈希表中进行删除、插入操作几乎能达到O(1)的时间效率。简单来说，哈希表是数组与链表的结合，为了更好理解哈希表，先看我测试hashmap时的一段输出，比较形象的描绘了哈希表的结构  
>0:----->&{20 atongmu <nil>}  
1:----->&{61 atongmu <nil>}  
2:  
3:  
4:  
5:  
6:  
7:----->&{67 atongmu <nil>}  
8:----->&{68 atongmu <nil>}  
9:  
10:  
11:  
12:  
13:  
14:  
15:----->&{35 atongmu <nil>}  
16:----->&{16 atongmu 0xc0420024e0}----->&{76 atongmu <nil>}  
17:----->&{57 atongmu <nil>}  
18:----->&{58 atongmu <nil>}  
19:----->&{79 atongmu <nil>}   

这一整块便是一个哈希表了，它由表头和数据链组成。最左侧的一列是表头，这一列也就是由数组构成的了，每个表头不存放数据，它只有数据链的指针，如果没有entry存放在该位置，其指针就为nil。数组的下标表示表头在bucket中额位置，也是每个key的hash值。右边的所有东西就是数据链了，为了让哈希表在空间占用轻量的同时，尽量达到O（1）的查找、插入、删除的时间效率，需要通过负载因子自动调节bucket的长度。

## 主要特点
1. 我的hashmap使用的key接收int，value接收string
2. 当负载因子（已有entry数/bucket容量)大于0.75时自动扩容
3. put时如果发现hashmap中已有该key，直接覆盖其value
4. put时如果发现key的hash值在表中已有结点，则将其附加在该结点所在链表尾部

## 关键类型、常量、变量、及函数介绍

### 1. Entry和HashMap类型
````golang
type Entry struct {
	key   int
	value string
	next  *Entry
}  
````  
Entry用来保存键值对及下一个Entry的指针，它描述hashmap中的每一个结点

````golang
type HashMap struct {
	size    int
	buckets []Entry
}
````  
HashMap类型就是整个hashmap了，它描述其大小，用Entry类型的切片实现

### 2. 全局常量、变量
````golang
const maxCapacity int = 100    
const loadFactor float32 = 0.75     
var nowCapacity int = 10   
````  
`maxCapacity`定义了hashmap中bucket的最大扩充容量，即，虽然hashmap可以自动扩充大小一提高使用效率，但不能无休止地扩充。  
`loadFactor`定义负载因子，当表中entry数量达到当前容量的负载因子时（此处定义为0.75），容量翻倍  
`nowCapacity`作为变量保存当前表的负载能力，也就是bucket的长度

### 3.CreateHashMap()新建hashmap
````golang
func CreateHashMap() *HashMap {
	hm := &HashMap{}
	hm.size = 0
	hm.buckets = make([]Entry, nowCapacity, maxCapacity)
	return hm
}
````  
用户只能通过调用此函数建立并初始化hashmap，它初始化了hashmap的size(初始Entry个数)、初始负载以及最大负载

### 4. getHash()获取key的散列值
````golang
func getHash(k int) int {
	return k % nowCapacity
}
````  
因为我定义的key为int类型，所以就通过简单的除留余数法获取key的散列值

### 5. DeleteEntry(k int)在hashmap中删除某个key对应的Entry
````golang
//删除指定key的Entry
func (hm *HashMap) Delete(k int) bool {
	e, ok := hm.GetEntry(k)
	if ok {
		fmt.Println("有该Entry,其hash为：", getHash(e.key), "值为：", e)
		hm.DeleteEntry(e)
		return true
	} else {
		fmt.Println("没有该Entry，删除失败")
		return false
	}
}
````  
调用了封装好的GetEntry(key int)和DeleteEntry(e *Entry),先找到该key所在的Entry，在执行删除

### 6. Put(k int,v value)向hashmap中插入数据，并检测负载程度自动调节容量
````golang
func (hm *HashMap) Put(k int, v string) {

	e := &Entry{k, v, nil}
	hm.insert(e)
	//达到负载因子且还能扩容时，扩容并迁移数据
	if float32(hm.size)/float32(nowCapacity) >= loadFactor && nowCapacity < maxCapacity {
		if 2*nowCapacity > maxCapacity {
			nowCapacity = maxCapacity
		} else {
			nowCapacity = 2 * nowCapacity
		}
		newHm := CreateHashMap()

		var index int
		for index = 0; index < len(hm.buckets); index++ {
			p := hm.buckets[index].next
			for p != nil {
				//临时保存p的next
				pNext := p.next
				p.next = nil
				newHm.insert(p)
				p = pNext
			}
		}

		var b1 int
		var b2 int = len(newHm.buckets)

		//把hm的切片扩容至newHm一样大
		for b1 = len(hm.buckets); b1 < b2; b1++ {
			hm.buckets = append(hm.buckets, Entry{})
		}
		//移动回数据
		for b1 = 0; b1 < b2; b1++ {
			hm.buckets[b1].next = newHm.buckets[b1].next
			newHm.buckets[b1].next = nil
		}
	}
}
````  
先调用`insert`函数直接插入键值对到hashmap中，再检测是否需要扩容，`float32(hm.size)/float32(nowCapacity) >= loadFactor`表示已超出负载因子，`nowCapacity < maxCapacity`表示还有扩容空间，满足这两个条件时，便新建容量翻倍（或直接值maxCapacity)的新的hashmap,再迁移其数据

## 完整代码实现
````golang
package hashmap

import (
	"fmt"
)

const maxCapacity int = 100     //bucket的最大容量
const loadFactor float32 = 0.75 //负载因子

var nowCapacity int = 10 //当前容量

type Entry struct {
	key   int
	value string
	next  *Entry
}

type HashMap struct {
	size    int
	buckets []Entry
}

func CreateHashMap() *HashMap {
	hm := &HashMap{}
	hm.size = 0
	hm.buckets = make([]Entry, nowCapacity, maxCapacity)
	return hm
}

//获取key的hash
func getHash(k int) int {
	return k % nowCapacity
}

func (hm *HashMap) GetSize() int {
	return hm.size
}

//由key在hashMap中找到其Entry指针
func (hm *HashMap) GetEntry(k int) (*Entry, bool) {
	p := hm.buckets[getHash(k)].next
	for p != nil {
		//找到,返回
		if p.key == k {
			return p, true
		} else {
			p = p.next
		}
	}
	//此时p==nil，没找到
	return nil, false
}

//不判断是否达到负载因子,直接往hashmap中插entry
func (hm *HashMap) insert(e *Entry) {

	hash := getHash(e.key)
	//从表头中的指针开始遍历
	var p *Entry = &hm.buckets[hash]
	for p.next != nil {
		//如果找到相同的key，则覆盖其中的值,完成insert，return
		if p.next.key == e.key {
			p.next.value = e.value
			return
		} else {
			p = p.next
		}
	}
	//此时p指向尾部结点
	p.next = e
	hm.size++
}

//判断是否达到负载因子再insert
func (hm *HashMap) Put(k int, v string) {
	e := &Entry{k, v, nil}
	hm.insert(e)
	//达到负载因子且还能扩容时，扩容并迁移数据
	if float32(hm.size)/float32(nowCapacity) >= loadFactor && nowCapacity < maxCapacity {
		if 2*nowCapacity > maxCapacity {
			nowCapacity = maxCapacity
		} else {
			nowCapacity = 2 * nowCapacity
		}

		newHm := CreateHashMap()
		var index int
		for index = 0; index < len(hm.buckets); index++ {
			p := hm.buckets[index].next
			for p != nil {
				//临时保存p的next
				pNext := p.next
				p.next = nil
				newHm.insert(p)
				p = pNext
			}
		}

		var b1 int
		var b2 int = len(newHm.buckets)

		//把hm的切片扩容至newHm一样大
		for b1 = len(hm.buckets); b1 < b2; b1++ {
			hm.buckets = append(hm.buckets, Entry{})
		}
		//移回数据
		for b1 = 0; b1 < b2; b1++ {
			hm.buckets[b1].next = newHm.buckets[b1].next
			newHm.buckets[b1].next = nil
		}
	}
}

//删除指定Entry
func (hm *HashMap) DeleteEntry(e *Entry) {
	//注意这里不能是p:=hm.buckets[getHash(e.key)].next,因为表头存的指针不一定为nil
	p := &hm.buckets[getHash(e.key)]
	for p != nil {
		if p.next == e {
			fmt.Println("正在删除")
			p.next = p.next.next
			fmt.Println("删除成功")

			hm.size--
			return
		} else {
			p = p.next
		}
	}
	fmt.Println("删除失败")
}

//删除指定key的Entry
func (hm *HashMap) Delete(k int) bool {
	e, ok := hm.GetEntry(k)
	if ok {
		fmt.Println("有该Entry,其hash为：", getHash(e.key), "值为：", e)
		hm.DeleteEntry(e)
		return true
	} else {
		fmt.Println("没有该Entry，删除失败")
		return false
	}
}

//遍历hashMap
func (hm *HashMap) Traverse() {
	var index int
	for index = 0; index < nowCapacity; index++ {
		p := hm.buckets[index].next
		if p == nil {
			fmt.Println(index, ":")
		} else {
			fmt.Print(index, ":")
		}
		for p != nil {
			fmt.Print("----->", p)
			p = p.next
			if p == nil {
				fmt.Println()
			}
		}
	}
	fmt.Println()
}
````

