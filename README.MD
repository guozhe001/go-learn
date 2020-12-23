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


