# Cmdexec
This is a Golang program for creating an alias of regularly run Linux commands, or commands you choose to save and don't want to rewrite every time.

## Features
- [x] Full support for actions with JSON file
- [x] Support for running commands
- [x] Support for running sudo commands
-  Support for custom sudo time-out
- [x] Temporarily saving sudo password for custom sudo time-out
-  CLI support (in progress)
-  Separate daemon project for running the commands and daemon process 

## Usage
To use this program, first build and install it using the following command:
```sh
go build
```
To add command
```sh
./cmdexec -a "command:sudo whoami,name: testa" -a "name: test1, command:sudo who"
```

check help 
```sh
./cmdexec
```