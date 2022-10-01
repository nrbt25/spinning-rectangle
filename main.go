package main

import (
	"fmt"
	"math"
	"time"
)

type Rectangle struct {
    ax int;
    ay int;
    bx int;
    by int;
}

func (r Rectangle) draw(angle float64) {
    mi := min(r.ay, r.by)
    ma := max(r.ay, r.by)

    for y := mi; y <= ma; y++ {

	drawLine(r.ax, y, r.bx, y, angle)	
    }
}

func min(a, b int) int {
    if a < b {
	return a 
    }

    return b
}

func max(a, b int) int {
    if a > b {
	return a
    }

    return b 
}

func drawLine(ax, ay, bx, by int, angle float64) {
    // find the equation of the line
    // y = ax + b

    a := float64((by - ay)) / float64((bx - ax))
    b := float64(ay) - a * float64(ax)
    y := (func(x int) int { 
	return int(math.Round(a * float64(x) + b))
    })

    for x := ax; x < bx; x++ {
	xx := int(math.Round(float64(x) * math.Cos(angle) - float64(y(x)) * math.Sin(angle)))	
	yy := int(math.Round(float64(x) * math.Sin(angle) + float64(y(x)) * math.Cos(angle))) 
	fmt.Printf("\x1b[%d;%dH", yy + 20, xx * 2 + 20)
	fmt.Print("# ")
    }
}

func main() {

    r := Rectangle{
	-5,
	-5,
	5,
	5,
    }
    
    loop := 0

    for {

    fmt.Print("\x1b[2J")
	 
	r.draw((math.Pi / 16) * float64(loop))
	time.Sleep(time.Second / 20)
	loop += 1;
    }
   // drawLine(r.x, r.y, r.x + r.width, r.y + r.height) 
}
