# goTrain
关于go的学习

1 Go语言结构
2 Go语言类型
3 Go语言变量
4 Go语言常量
5 Go语言函数 https://www.runoob.com/go/go-functions.html
6 数组
7 指针
8 结构体
9 切片(Slice)https://www.runoob.com/go/go-slice.html
10 范围range
11 集合map
12 递归
13 类型转换
14 接口
### Go语言结构
第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
下一行 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
下一行 /*...*/ 是注释，在程序执行时将被忽略。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。
下一行 fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
使用 fmt.Print("hello, world\n") 可以得到相同的结果。
Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。
当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
### 2.Go语言类型
* 数据类型：
* 布尔、
* 字符串、
* 派生类型
    * 指针Pointer
    * 数组
    * 结构化类型struct
    * Channel
    * 函数
    * 切片
    * 接口interface
    * Map
### 3.Go语言变量
```
var identifier type
var a string = "Runoob"
var b, c int = 1, 2
```
不赋值，默认零值
```
var b int   默认为0
var c bool  默认为false
var a string  默认为""
```
#### 以下几种类型为 nil
```var a *int
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error 是接口
var d = true 可根据赋值的类型判断变量的类型
f := "Runoob"  （省略var）
```
#### 相同类型多变量
```
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3
var vname1, vname2, vname3 = v1, v2, v3
vname1, vname2, vname3 := v1, v2, v3
```
#### 值类型和引用类型
当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝：
可以通过 &i 来获取变量 i 的内存地址
同一个引用类型的指针指向的多个字可以是在连续的内存地址中（内存布局是连续的），这也是计算效率最高的一种存储形式；也可以将这些字分散存放在内存中，每个字都指示了下一个字所在的内存地址。当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响。
空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)。
a , b , c := 1 , 2 , "str"
### 4.Go语言常量
显式类型定义： const b string = "abc"
隐式类型定义： const b = "abc"
多个相同类型声明 const c_name1, c_name2 = value1, value2
```
inta
const (
    a = iota
    b = iota
    c = iota
)第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：const (
    a = iota
    b
    c
)

const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)--------0 1 2 ha ha 100 100 7 8
```
### 5.Go语言函数 https://www.runoob.com/go/go-functions.html
Go 语言函数定义格式如下：
func function_name( [parameter list] ) [return_types] {
   函数体
}

闭包：
```
func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i  
   }
}
```
函数内包含一个匿名函数，可以直接使用函数内的变量而不必声明。
### 6.数组
var variable_name [SIZE] variable_type
多维数组：
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
初始化数组：
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
### 7.指针
取地址 &a   如： 20818a220
#### 什么是指针
一个指针变量指向了一个值的内存地址。
类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下
```
var var_name *var-type
var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：
var ip *int        /* 指向整型*/
var fp *float32    /* 指向浮点型 */
```


示例
```
package main
import "fmt"
func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */
   ip = &a  /* 指针变量的存储地址 */
   fmt.Printf("a 变量的地址是: %x\n", &a  )//a 变量的地址是: 20818a220
   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )//ip 变量储存的指针地址: 20818a220
   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )//*ip 变量的值: 20
}
```
空指针判断：
```
if(ptr != nil)     /* ptr 不是空指针 */
if(ptr == nil)    /* ptr 是空指针 */

整型指针数组的声明：
var ptr [MAX]*int;
指向指针的指针变量声明格式如下：（可以不断叠加*的个数表示指向指针的指针的指针...）
var ptr **int;
a = 3000
   /* 指针 ptr 地址 */
   ptr = &a

   /* 指向指针 ptr 地址 */
   pptr = &ptr
--------------------------------------------------------------------------------
```
### 8.结构体
结构体格式：
```
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```
一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
示例：
```
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}


func main() {

    // 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

    // 忽略的字段为 0 或 空
   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}
```
访问结构体成员
如果要访问结构体成员，需要使用点号 . 操作符，格式为：
结构体.成员名"
赋值： Book1.title = "Go 语言"
打印： fmt.Printf( "Book 1 title : %s\n", Book1.title)    //Book 1 title : Go 语言

***

### 9.切片(Slice)
https://www.runoob.com/go/go-slice.html
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
定义切片
```
1.var identifier []type   //未指定大小的数组
2.var slice1 []type = make([]type, len)
```
也可以简写为
slice1 := make([]type, len)
切片初始化
```
s :=[] int {1,2,3 } 
s := arr[:]   //s是arr的引用
s := arr[startIndex:endIndex]  //arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
s := arr[startIndex:] 
s := arr[:endIndex] 
s1 := s[startIndex:endIndex] 
s :=make([]int,len,cap)  //通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
相关函数：
x = make([]int,3,5)
fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)  //len=3 cap=5 slice=[0 0 0]
```
空切片
```
var numbers []int
fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
len=0 cap=0 slice=[]
numbers == nil
```
切片截取
```
numbers := []int{0,1,2,3,4,5,6,7,8} //len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8] 
numbers[1:4] == [1 2 3]
numbers[:3] == [0 1 2]
numbers[4:] == [4 5 6 7 8]
numbers1 := make([]int,0,5) //len=0 cap=5 slice=[]
number2 := numbers[:2] //len=2 cap=9 slice=[0 1]
number3 := numbers[2:5] //len=3 cap=7 slice=[2 3 4]
```
切片的扩容
```
var numbers []int  //len=0 cap=0 slice=[]
numbers = append(numbers, 0)  //len=1 cap=1 slice=[0]
numbers = append(numbers, 1) //len=2 cap=2 slice=[0 1]
numbers = append(numbers, 2,3,4)   //len=5 cap=6 slice=[0 1 2 3 4]
 /* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)
  /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)
   ```
### 10 范围range
数组：
```
nums := []int{2, 3, 4}
sum := 0
for _, num := range nums {
    sum += num
}
```
需要index：
```
for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
map：
 kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
```
### 11.集合map
定义map
```
/* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type    //var countryCapitalMap map[string]string /*创建集合 */
/* 使用 make 函数 */
map_variable := make(map[key_data_type]value_data_type)  //countryCapitalMap = make(map[string]string)
 /* map插入key - value对,各个国家对应的首都 */
    countryCapitalMap [ "France" ] = "巴黎"
    countryCapitalMap [ "Italy" ] = "罗马"
    countryCapitalMap [ "Japan" ] = "东京"
    countryCapitalMap [ "India " ] = "新德里"

    /*使用键输出地图值 */
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }

 /*删除元素*/ delete(countryCapitalMap, "France")
```
### 12 递归
阶乘
```
package main

import "fmt"

func Factorial(n uint64)(result uint64) {
    if (n > 0) {
        result = n * Factorial(n-1)
        return result
    }
    return 1
}

func main() {  
    var i int = 15
    fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}
```
### 13.类型转换
type_name(expression)
实例
以下实例中将整型转化为浮点型，并计算结果，将结果赋值给浮点型变量：
实例
```
package main

import "fmt"

func main() {
   var sum int = 17
   var count int = 5
   var mean float32
   
   mean = float32(sum)/float32(count)
   fmt.Printf("mean 的值为: %f\n",mean)
}
```
以上实例执行输出结果为：
mean 的值为: 3.400000
### 14.接口
interface
### 15 并发：
go 函数名( 参数列表 )<br>
ch := make(chan int)// 通道<br>
ch := make(chan int, 100) //通道缓冲区
