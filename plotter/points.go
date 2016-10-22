package plotter

import(
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/math/fixed"
)

var palette = []color.Color{color.White, color.Black}

func AddLabel(img *image.Paletted, x, y int, label string) {
    col := color.RGBA{200, 100, 0, 255}
    point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

    d := &font.Drawer{
        Dst:  img,
        Src:  image.NewUniform(col),
        Face: basicfont.Face7x13,
        Dot:  point,
    }
    d.DrawString(label)
}

func PlotPoints(out io.Writer, points []Point, w int, h int, nframes int, delay int, label string) {
    steps := make([]int, nframes, nframes)
    for k := 0; k < nframes; k++ {
        steps[k] = int(math.Min(float64(len(points)-1), math.Pow(math.SqrtPhi, float64(k))))
    }

    anim := gif.GIF{LoopCount: nframes}
    ymargin := 10
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, w, h)
        img := image.NewPaletted(rect, palette)
        for j := 0; j <= steps[i]; j++ {
            img.Set(int(points[j].X), int(points[j].Y) + ymargin, color.Black)
        }

        AddLabel(img, w/3+50, h-45, label)

        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }

    gif.EncodeAll(out, &anim)
}


