# Hello World

## 安装和配置

安装方法可以查看[官网](https://go.dev/dl/)。

在安装完成后，可以查看一下 `GOPATH` 的路径：

```bash
go env | grep GOPATH
```

> 如果使用 Zsh，可能会无法执行该命令。可以**临时**（到时候记得删除）在 `~/.zshrc` 中添加 `export PATH=$PATH:/usr/local/go/bin`，然后使用 `source ~/.zshrc` 重新载入。

上面提到的 `GOPATH` 路径是 workspace，是可以修改的，这个下面会说。

现在，需要到 shell 的配置文件（Zsh 是 `~/.zshrc`，Bash 是 `~/.bashrc`）中，自己添加环境变量：

```
export GOROOT=/usr/local/go
// 可以自定义一个路径
export GOPATH=/Users/username/go/workspace
export PATH=$PATH:$GOPATH/bin
```

添加完成后，记得使用 `source ~/.zshrc` 重新载入。

然后，在 `GOPATH` 路径下，创建三个路径即可：

```bash
mkdir -p bin pkg src
```

## Hello World

[main.go](./main/main.go)
