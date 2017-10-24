package main
import (
	"fmt"
	"time"
	"math/rand"
)

const initCapacity int=20

type Entry struct{
	key int
	value string
	next *Entry
}

type HashMap struct{
	size int
	buckets []Entry
}

func createHashMap() HashMap{
	hm:=HashMap{}
	hm.size=0
	hm.buckets=make([]Entry,initCapacity)
	return hm
}

//获取key的hash
func getHash(k int) int{
	return k%initCapacity
}

//由key在hashMap中找到其Entry指针
func(hm *HashMap) getEntry(k int) *Entry{
	p:=hm.buckets[getHash(k)].next
	for p!=nil{
		//找到,返回
		if p.key==k{
			return p
		}else{
			p=p.next
		}	
	}
	//此时p==nil，没找到
	return nil
}

func(hm *HashMap) put(k int,v string){
	newEntry:=Entry{}
	newEntry.key=k
	newEntry.value=v
	hash:=getHash(k)
	//从表头中的指针开始遍历
	var p *Entry=&hm.buckets[hash]
	for p.next!=nil{
		//如果找到相同的key，则覆盖其中的值,完成put，return
		if p.next.key==k{
			p.next.value=v
			return
		}else{
			p=p.next
		}		
	}
	//此时p指向尾部结点，p.next==nil
	//将新结点的next赋为nil，填至尾部
	newEntry.next=nil
	p.next=&newEntry
}

//删除指定Entry
func(hm *HashMap) deleteEntry(e *Entry){
	//注意这里不能是p:=hm.buckets[getHash(e.key)].next,因为表头存的指针不一定为nil
	p:=&hm.buckets[getHash(e.key)]
	for p!=nil{
		if p.next==e{
			fmt.Println("正在删除")
			p.next=p.next.next
			fmt.Println("删除成功")
			return
		}else{
			p=p.next
		}
	}
	fmt.Println("删除失败")
}

//删除指定key的Entry
func(hm *HashMap) delete(k int) bool{
	e:=hm.getEntry(k)
	if e!=nil{
		fmt.Println("有该Entry,其hash为：",getHash(e.key),"值为：",e)
		hm.deleteEntry(e)
		return true
	}else{
		fmt.Println("没有该Entry，删除失败")
		return false
	}
}

//遍历hashMap
func(hm *HashMap) traverse(){
	var index int
	for index=0;index<initCapacity;index++{
		p:=hm.buckets[index].next
		if p!=nil{
			fmt.Print(index,":")
		}
		for p!=nil{			
			fmt.Print("----->",p)
			p=p.next
			if p==nil{
				fmt.Println()
			}
		}
	}
	fmt.Println()
}




func main(){
	hm:=createHashMap()
	//插入40个随机key，value对
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<100;i++{
		hm.put(r.Intn(100),"atongmu")
	}
	
	fmt.Println("插入100个Entry后的HashMap:")
	hm.traverse()

	fmt.Println()
	ke:=r.Intn(100)
	fmt.Println("删除key:",ke)
	hm.delete(ke)
	hm.traverse()
}



