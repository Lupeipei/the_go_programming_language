package main

import (
  "fmt"
  "bytes"
)

func main() {
  fmt.Println(floatsToString([]float64{0.1, 0.2, 0.3}))
}

func floatsToString(values []float64) string {
  var buf bytes.Buffer
  buf.WriteByte('[')
  for i, v := range values {
    if i > 0 {
      buf.WriteString(", ")
    }
    fmt.Fprintf(&buf, "%3.1f", v)
  }
  buf.WriteByte(']')
  return buf.String()
}
