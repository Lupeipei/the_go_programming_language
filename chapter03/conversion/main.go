package main
import (
  "fmt"
  "strconv"
)

func main() {
  x := 123
  y := fmt.Sprintf("%d", x)

  fmt.Println(y, strconv.Itoa(x))

  s := fmt.Sprintf("x=%b", x)
  fmt.Println(s)
  _, err := strconv.Atoi("123")
  if err != nil {
    fmt.Println(err)
  }
}
