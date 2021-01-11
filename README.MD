# 如果插件下载不下来可以使用蓝灯的代理

## 命令记录

### 格式化
```shell
go fmt hello.go
```

### 运行
```shell
go run hello.go
```

### 构建成可执行二进制文件
```shell
go build hello.go
./hello
```

### 构建成可执行二进制文件并安装到$GOPAT/bin目录下
```shell
go build hello.go
./$GOPATH/bin/hello
```

### 下载和安装包
```shell
 go get github.com/guozhe001/go-learn
```

### 阅读包文档
```shell
 go doc fmt
```

### 指定方法名阅读文档
```shell
 go doc strconv ParseFloat
```

### 在本地运行一个web的go文档
```shell
godoc -http=6060
```
如果没有此命令先安装`go get -v  golang.org/x/tools/cmd/godoc`
如果安装时因为网络原因导致报错，可以设置代理：

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```


## 注释的约定
* 包注释应该以`Package`开头，然后跟着包名
* 方法注释应该以方法名开头


## 切片
每一个切片都构建于一个底层的数组之上。实际上是底层的数组存储了切片的数据；切片仅仅是数组中的一部分（或者所有）元素的视图。

*注意：由于切片只是底层数组内容的视图，如果你修改底层数组，这些变化也会反映到切片*
*上面的定义反过来也成立：即如果修改了切片的值，底层的数组的值也被修改了*


## 方法
在函数的名称前面添加“接收器参数名称”和“接收器参数类型”，就变成了方法。
* **Go使用接收器参数来替代其他语言中的“self”或者“this”**
* **方法和类型必须定义在同一个包中**
* **接收器的类型也可以是指针类型**
* **当你使用非指针的变量调用一个需要指针的接收器的方法时，Go会自动为你将非指针类型转换为指针类型；反之亦然**
* **所有的方法的接收器类型都应该一致，所以一个类型函数要吗接收值类型的接收器参数，要吗接收指针类型的接收器参数；避免两者同时混用的情况**

如下面的方法，可以说（sayHi方法定义在MyType上），也可以说为类型MyType定义的方法sayHi：

```go
package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi!")
}

```
### 接收器参数（几乎）就是另一个参数
可以在方法内部使用接收器参数，如改造后的sayHi方法：
```go
package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}
```


*思考，为什么有了函数还需要方法？*

### 方法（几乎）就像一个函数
* 方法可以有其他的参数
* 方法可以有返回值
* 方法的首字母大写则认为是导出的，首字母小写则认为是不导出的。


## 封装
类型的属性不导出，方法导出，来实现封装，如：`headfirstgo/ch9/calendar/date.go`

## 嵌入
如：`headfirstgo/ch9/calendar/event.go`

## 接口


## 错误处理

### 延迟函数调用：defer
* 如果一个方法中有多个defer关键字，则这些被defer修饰的调用会按照倒序执行
* 在return之后的defer语句也不会被执行

### 抛出异常 panic
* panic很像抛出一个异常
* panic抛出的异常有堆栈信息

### 捕获异常 recover
* recover()方法会捕获panic

**不鼓励使用panic和recover，最好使用error来进行判断错误并且处理**

## 测试
测试代码的约定如下：
* 你不需要将测试代码和被测试代码放在同一个包，但是如果你想要从包中访问未导出的类型或函数，则需要这样做。
* 测试需要使用testing包中的类型，所以需要在每个测试文件的顶部导入该包
* 测试函数应该以Test开头。（名字的其余部分可以是任意内容，但它应该以大写字母开头）
* 测试函数应该接受单个参数：一个指向testing.T值的指针
* 你可以通过testing.T值调用方法（例如Error）来报告测试失败。大多数方法都接受一个字符串，其中包含解释测试失败原因的信息。
  
## 一级函数
Go语言支持一级函数，也就是说，函数在Go语言中被视为“一等公民”。
在具有一级函数的编程语言中，可以将函数分配给变量，然后从这些变量调用函数。
例子：`ch15/hello.go`