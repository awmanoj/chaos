package main

import(
    //"fmt"
    "log"
    "net/http"
    "time"
    "github.com/awmanoj/chaos/sierpinski"
    "github.com/awmanoj/chaos/plotter"
)

func SierpinskiHandler(w http.ResponseWriter, r *http.Request) {
    var width float64 = 600 
    var height float64 = 600 
    var iterations int = 50000 
    var nframes = 64
    var delay = 10
    plotter.PlotPoints(w, sierpinski.GeneratePoints(width, height, iterations),
        int(width), int(height), nframes, delay, "SIERPINSKI TRIANGLE")
}

func main() {
    http.HandleFunc("/sierpinski", SierpinskiHandler)
    s := &http.Server{
        Addr:           ":8080",
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    log.Println("listening on", s.Addr)    
    log.Fatal(s.ListenAndServe())    
}
