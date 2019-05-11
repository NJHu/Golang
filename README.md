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

- 3, 测试，在终端输入go version
  - 3.1 运行go程序文件
```
go run xx.go
```


- 4, GoLand安装与配置
> https://www.jetbrains.com/go/

- 5, 配置GOPATH
> 这一步非常重要，否则代码中的依赖包将无法识别。一般GOPATH都配置到项目名这一层


