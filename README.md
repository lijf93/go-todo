```text
   ______          ______          __    
  / ____/___      /_  __/___  ____/ /___ 
 / / __/ __ \______/ / / __ \/ __  / __ \
/ /_/ / /_/ /_____/ / / /_/ / /_/ / /_/ /
\____/\____/     /_/  \____/\__,_/\____/ 

A cli todo list app in Golang.
```
[![](https://img.shields.io/github/license/lijf93/go-todo.svg)](https://github.com/lijf93/go-todo/blob/master/LICENSE)

My first Golang application, in order to learn Golang syntax.

## Installation
```bash
cd ~/go/src/go-todo
go build
./go-todo h
```

### ScreenShot
![gotodo](https://github.com/lijf93/go-todo/blob/master/screenshot/gotodo-screenshot.png)

## Usage
### add
```bash
go-todo add test add
go-todo a test add
```

### done
```bash
go-todo done 1
go-todo do 1
```

### delete
```bash
go-todo delete 1
go-todo del 1
```

### list
```bash
go-todo list
go-todo l
```

## Thanks
* [urfave/cli](https://github.com/urfave/cli)
* [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)

## License
[MIT License](https://github.com/lijf93/go-todo/blob/master/LICENSE)