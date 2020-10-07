package main

import "fmt"

func  main()  {
  var test []int
  test = append(test, 1)
  test = append(test, 2, 3)
  test = append(test, 4, 5, 6)
  test = append(test, test...)
  fmt.Println(test)

  var x, y []int
  for i := 0; i < 10; i++ {
    y = appendInt(x, i)
    fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
    x = y
  }
}

func appendInt(x []int, y int) []int {
  var z []int
  zlen := len(x) + 1
  if zlen <= cap(x) {
    z = x[:zlen]
  } else {
    zcap := zlen
    if zcap < 2*len(x) {
      zcap = 2*len(x)
    }
    z = make([]int, zlen, zcap)
    copy(z, x)
  }
  z[len(x)] = y
  return z
}
