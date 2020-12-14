# 《The Go Programming Language》

## 入门

名为main的包用来定义是一个独立的可执行程序，main包里面的main函数时一个程序开始执行的地方，如果执行一个不在main包的方法，会报错：

```shell
(base) bogon:the-go-pl guozhe$ go run hello.go
go run: cannot run non-main package
```

运行一个.go文件

```shell
go run helloworld.go
```

编译成一个可执行文件：

```shell
go build helloworld.go
./helloworld
```



