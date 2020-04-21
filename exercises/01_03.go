package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    echo1()
    echo2()
    echo3()
}

func echo1() {
    start := time.Now()
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Printf("%.8fs %s\n", time.Since(start).Seconds(), s)
}

func echo2() {
    start := time.Now()
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Printf("%.8fs %s\n", time.Since(start).Seconds(), s)
}

func echo3() {
    start := time.Now()
    s := strings.Join(os.Args[1:], " ")

    fmt.Printf("%.8fs %s\n", time.Since(start).Seconds(), s)
}

