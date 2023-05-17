# 1. 通用基础
## 1.1 go struct 能不能比较
**答案**  
（1）不同类型的 struct 之间不能进行比较，编译期就会报错（GoLand 会直接提示）  
（2）同类型的 struct 也分为两种情况  
  - struct 的所有成员都是可以比较的，则该 strcut 的不同实例可以比较  
  - struct 中含有不可比较的成员（如 Slice，map，函数类型），则该 struct 不可以比较  

**案例分析**  
(1) 相同类型的struct做比较  
```go
import "fmt"

type A struct {
	age  int
	name string
}

func StructCompare1() {
	aObj1 := A{
		age:  13,
		name: "张三",
	}
	aObj2 := A{
		age:  13,
		name: "张三",
	}
	fmt.Println(aObj1 == aObj2) // true

	aObj3 := &A{
		age:  13,
		name: "张三",
	}
	aObj4 := &A{
		age:  13,
		name: "张三",
	}

	fmt.Println(aObj3 == aObj4) // false

	var aObj5 A
	fmt.Println(aObj5) //{0 } ，未明确初始化时，struct 实例的成员取各自的零值
	//fmt.Println( aObj5 == nil)  // 报错，无法将 nil 转换为类型 A

	var aObj6 *A
	fmt.Println(aObj6)        // <nil> ，指针类型数据的零值为 nil
	fmt.Println(aObj6 == nil) //  true，指针类型的数据可以和 nil 比较
}
```
假如 struct 中包含了不可比较的类型，Goland会直接报错   
![](assets/2023-05-17-23-25-24.png)

# 2. 常见数据结构
## 2.1 数组和切片                    
## 2.2 Map
## 2.3 Chancel
# 3. 并发编程
## 3.1 Mutex
## 3.2 WaitGroup
# 4. 调度机制
# 5. 内存分配
# 6. 垃圾回收