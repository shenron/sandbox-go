# sandbox-go
Start with golang

# Install
## There are two compilers (go, gcc)
https://wiki.archlinux.org/index.php/Go

```shell
sudo pacman -S go
```

# Configuration
## Add in you .bash_profile or .bashrc
```shell
# gopath
export GOPATH=$HOME/dev/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

# Run
## In each folder run with program
```shell
go run myFile.go
```
