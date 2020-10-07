package main

import (
  "fmt"
  "crypto/sha256"
)

func main() {
  c1 := sha256.Sum256([]byte("go"))
  c2 := sha256.Sum256([]byte("GO"))

  fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1==c2, c1)
}
