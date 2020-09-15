package main

import "fmt"

func main () {
  s := "hello, go language"
  fmt.Println(comma(s))
}

func comma(s string) string {
  n := len(s)
  if n <= 3 {
    return s
  }
  return comma(s[:n-3]) + ", " + s[n-3:]
}
