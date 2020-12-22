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


