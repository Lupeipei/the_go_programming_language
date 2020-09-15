package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    counts := make(map[string]map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    // for line := range counts {
    //     for file, n := range counts[line] {
    //         if n > 1 {
    //             fmt.Printf("%d\t%s\t%s\n", n, file, line)
    //         }
    //     }
    // }

    // need refactoring
    for file := range counts {
        for line, n := range counts[file] {
            if n > 1 {
                fmt.Printf("%d\t%s\t%s\n", n, file, line)
            }
        }
    }
}

func countLines(f *os.File, counts map[string]map[string]int) {
    input := bufio.NewScanner(f)
    counts[f.Name()] = make(map[string]int)
    for input.Scan() {
        counts[f.Name()][input.Text()]++
        // counts[input.Text()] = make(map[string]int)
        // counts[input.Text()][f.Name()]++
    }
}
