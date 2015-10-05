# sandbox-go
Start with golang

# Install
## There are two compilers (go, gcc-go)
https://wiki.archlinux.org/index.php/Go

```shell
sudo pacman -S go

mkdir -p ~/dev/go

cd ~/dev/go
git clone https://github.com/shenron/sandbox-go
```

# Configuration
## Add in you .bash_profile or .bashrc
```shell
# gopath
export GOPATH=$HOME/dev/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

# Run
## In each folder launch programm with 'go run'
```shell
go run ~/dev/go/sandbox-go/hello/hello.go
```
