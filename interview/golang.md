# 1. 通用基础
## 1.1 go struct 能不能用等号做比较
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
(2) 不同类型的 struct 做比较   
Goland会直接报错  
![](assets/2023-05-18-06-52-41.png)   

## 1.2 new 和 make 有什么区别
**答案**  
(1) make 和 new 都是用来分配内存的。make 既分配内存也初始化内存；而 new 只是将内存清零，并没有初始化内存   
(2) make 只能用来分配及初始化 slice，map，channel 等数据类型；而new 可以分配任意类型的数据    
(3) make 返回的是引用类型本身；而 new 返回的是指向类型的指针   
   
**使用 new 分配内存**  
- 基础类型 int    
  ```go
  var v * int
  fmt.Println(*v) // 这步会因为空指针而panic
  fmt.Println(v) // <nil>

  v = new(int)
  fmt.Println(*v) // 0
  fmt.Println(v) // 0xc00004c088
  ```
  由此可以看出，初始化一个指针变量，其值为 nil， nil 值是不能直接赋值的。通过 new 返回一个指向新分配类型为 int 的指针，指针值为 0xc00004c088，这个指针值指向的内容的值为0 (zero value)。
- 复合类型-array   
  ```go
  var v * int
  fmt.Println(*v) // 这步会因为空指针而panic
  fmt.Println(v) // <nil>

  v = new(int)
  fmt.Println(*v) // 0
  fmt.Println(v) // 0xc00004c088
  ``` 
- 

**使用 make 分配内存**  

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