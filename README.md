## Golang 环境配置

- 1, Mac OS X（也称为 Darwin） go安装

> https://golang.org/dl/

- 2, 配置GOROOT和PATH

```
vim ~/.bash_profile

//在~/.bash_profile文件中添加以下2条命令
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

- 3, 测试，在终端输入`go version`
  - 3.1 运行go程序文件
```
go run xx.go
```

### 接下来需要配置Go的一款开发工具GoLand

- 1, GoLand安装与配置
> https://www.jetbrains.com/go/



