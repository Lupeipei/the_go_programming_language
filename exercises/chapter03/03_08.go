package main

// for complex64
import (
  "image"
  "image/color"
  "image/png"
  "math/cmplx"
  "net/http"
  "log"
)

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrot(z complex64) color.Color {
  const iterations = 200
  const contrast = 15

  var v complex64
  for n := uint8(0); n < iterations ; n++ {
    v = v*v + z
    if cmplx.Abs(complex128(v)) > 2 {
      return color.Gray{255 - contrast*n}
    }
  }
  return color.Black
}

func handler(w http.ResponseWriter, r *http.Request) {
    const (
      xmin, ymin, xmax, ymax = -2, -2, +2, +2
      width, height = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0,0, width, height))
    for py := 0; py < height ; py++ {
      y := float32(py)/height*(ymax-ymin) + ymin
      for px := 0; px < width; px++ {
        x := float32(px)/width*(xmax-xmin) + xmin
        z := complex(x, y)
        img.Set(px, py, mandelbrot(z))
      }
    }
    png.Encode(w, img)
}
