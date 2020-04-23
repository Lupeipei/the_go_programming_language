package main

import (
    "log"
    "net/http"
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0
    blackIndex = 1
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    query := make(map[string]string)
    for k, v := range r.Form {
        query[k] = v[0]
    }
    lissajous(w, query)
}

func lissajous(out io.Writer, query map[string]string) {
    var  cycles float64 = 5
    var res float64 = 0.001
    var size int64 = 100
    var nframes int64 = 64
    var delay int64 = 8

    if val, ok := query["cycles"]; ok {
        data, err := strconv.ParseFloat(val, 64)
        if err == nil {
            cycles = data
        }
    }

    if val, ok := query["res"]; ok {
        data, err := strconv.ParseFloat(val, 64)
        if err == nil {
            res = data
        }
    }
    
    if val, ok := query["size"]; ok {
        data, err := strconv.ParseInt(val, 10, 64)
        if err == nil {
            size = data
        }
    }

    if val, ok := query["nframes"]; ok {
        data, err := strconv.ParseInt(val, 10, 64)
        if err == nil {
            nframes = data
        }
    }

    if val, ok := query["delay"]; ok {
        data, err := strconv.ParseInt(val, 10, 64)
        if err == nil {
            delay = data
        }
    }

    freq := rand.Float64()*3.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette) 
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size + int(x*size + 0.5), size + int(y*size + 0.5), blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
