# Chaos Game 

Sierpinski Triangle visualization using golang image library.

### Introduction

`Sierpinski Triangle` were first described in 1915 by Waclaw Sierpinski, a polish mathematician. Sierpinski Gasket can be drawn using pencil and paper. Follow these steps:

Step 1: Draw a triangle (equilateral).

![1.png](/public/imgs/1.png?raw=true)

Step 2: Draw a point in the middle of each of the three sides of the triangle and then connect these points to form a new triangle. 

![2.png](/public/imgs/2.png?raw=true)

Step 3: Repeat `Step2` for the outer three triangles. 

![3.png](/public/imgs/3.png?raw=true)

If you repeat this `iter` number of times you generate a sierpinski gasket or triangle or pattern. It is one of the most basic and popular fractals out there. 

### Chaos Game

The Chaos Game builds the gasket by using a roll of a dice (or equivalent). 

Step 1: Draw the three vertices of a triangle and pick a random point somewhere inside it.

![4.png](/public/imgs/4.png?raw=true)

Step 2: Generate a random number (integer) between 1 and 3 (both inclusive). If 1, then go towards vertex at the bottom left. If 2, then go towards the vertex at the bottom right and if 3, then go towards the vertex at top. For `sierpinski gasket` we move exactly half the distance between current point and the chosen vertex (towards the vertex).

![5.png](/public/imgs/5.png?raw=true)

Step 3: Repeat `Step 2` for the new point generated and go on.

![6.png](/public/imgs/6.png?raw=true)

Following is chaos game after 100 iterations: 

![7.png](/public/imgs/7.png?raw=true)

After 50,000 iterations: 

![8.png](/public/imgs/8.png?raw=true)

### Program 

This program implements a web service which generates a GIF image that illustrates the chaos game - GIF is a dynamic image which beautifully demonstrates the beauty of order in chaos. 

It is written in `golang`. 

It starts with one point on bottom left and follows the `roll of a dice` pattern to generate more points. For making the example not too slow and still keep the idea clear - number of points plotted in each iteration grow exponentially 

```
...
      for k := 0; k < nframes; k++ {
          steps[k] = int(math.Min(float64(len(points)-1), math.Pow(math.SqrtPhi, float64(k))))
      }
...
```
So you see more points in later iterations and lesser in initial iterations.

### Usage: 

o Build and start the server 

```
$ go build 
$ nohup ./chaos & 
```

o Play! 

```
$ curl localhost:8080/sierpinski > out.gif
$ open out.gif # open in default program
```

### Credits

* Fractals: A programmer's approach by Ben Trube (contains a C++ implementation)

### Feedback 

Drop a line on [@awmanoj](https://twitter.com/awmanoj). 

