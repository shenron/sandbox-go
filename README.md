# sandbox-go
Start with golang

# Install
## There are two compilers (go, gcc-go)
https://wiki.archlinux.org/index.php/Go

```shell
sudo pacman -S go

# construct workspace
mkdir -p ~/dev/go/src
mkdir ~/dev/go/bin
mkdir ~/dev/go/pkg

cd ~/dev/go/src
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
## In each folder launch program with 'go run'
```shell
go run ~/dev/go/src/sandbox-go/hello/hello.go
```
