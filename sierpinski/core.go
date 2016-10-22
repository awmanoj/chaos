package sierpinski

import(
    "math"
    "math/rand"
    "github.com/awmanoj/chaos/plotter"
)

const(
    π = math.Pi  // pi
    φ = π / 12   // rotation offset for axes
    δ = 0.5      // fraction of distance between current point and the chosen vertex (new point)
)

// number of vertices
var n int32 = 3

// dimension of canvas 
var w float64 = 600
var h float64 = 600
var iter int = 50000

func GeneratePoints(w float64, h float64, iter int) []plotter.Point {
    sin := math.Sin
    cos := math.Cos

    θ := 2*π/float64(n) + φ

    xr := w/2
    yr := h/2

    cx := xr*(cos(θ) + 1)
    cy := yr*(sin(θ) + 1)
  
    points := []plotter.Point{}
    points = append(points, plotter.Point{X: cx, Y: cy})

    for i := 0; i < iter; i++ {
        vn := rand.Int31() % n + 1 // current vertex number chosen randomly
        vx := xr*(cos(θ*float64(vn)) + 1)     // current vertex x coordinate 
        vy := yr*(sin(θ*float64(vn)) + 1)     // current vertex y coordinate
        cx += δ * (vx - cx)          // translate to new x coordinate 
        cy += δ * (vy - cy)          // translate to new y coordinate
        points = append(points, plotter.Point{X: cx, Y: cy})
    }

    return points
}
