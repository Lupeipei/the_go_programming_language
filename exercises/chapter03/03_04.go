package main

import (
  "log"
  "net/http"
  "math"
  "io"
  "github.com/ajstarks/svgo"
)

const (
  width, height = 600, 320
  cells = 100
  xyrange = 30.0
  xyscale = width / 2 / xyrange
  zscale = height * 0.4
  angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func corner(i, j int) (int, int) {

  x := xyrange * (float64(i)/cells - 0.5)
  y := xyrange * (float64(j)/cells - 0.5)

  z := f(x, y)
  sx := width/2 + (x-y)*cos30*xyscale
  sy := height/2 + (x+y)*sin30*xyscale - z*zscale
  return int(sx), int(sy)
}

func f(x, y float64) float64 {
  r := math.Hypot(x, y)
  return math.Sin(r) / r
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(out io.Writer) {
  canvas := svg.New(out)
  canvas.Start(width, height)
  var x []int
  var y []int
  for i := 0; i < cells; i ++ {
    for j := 0; j < cells; j ++ {
      ax, ay := corner(i+1, j)
      bx, by := corner(i, j)
      cx, cy := corner(i, j+1)
      dx, dy := corner(i+1, j+1)
      x = nil
      y = nil
      x = append(x, ax, bx, cx, dx)
      y = append(y, ay, by, cy, dy)
      canvas.Polygon(x, y)
    }
  }
  canvas.End()
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "image/svg+xml")
    surface(w)
}
