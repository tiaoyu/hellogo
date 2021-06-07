package main

import (
    "fmt"
    "log"
)

type a int32

const (
    a_1 a = 1
    a_2 a = 2
)

func main() {
    fmt.Println("这个是主干用的merge")
    fmt.Println("I'll merge this.")
    fmt.Println("Hello, Go! This is issue001 update~")
    fmt.Println("commit 001")
    fmt.Println("commit 002")
    fmt.Println("commit 003")
    fmt.Println("Hello, Go! This is issue001 update2~")
    fmt.Println("I'll merge in master~")
    var tmp = 2147483647
    fmt.Println(a(tmp))
    defer HandlerPanic()
    log.Fatal("asdf")
    panic("panic main")
}

func HandlerPanic() {
   if err := recover(); err != nil {
       fmt.Println(err)
   } 
}
