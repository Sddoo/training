package utils

import (
	"fmt"
	"math"
	"strconv"
	"net/http"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4       // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	up						= 0.2
	down    			= -0.2
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func surface(w http.ResponseWriter, r *http.Request) {
	var color string
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1 := corner(i+1, j)
			bx, by, z2 := corner(i, j)
			cx, cy, z3 := corner(i, j+1)
			dx, dy, z4 := corner(i+1, j+1)
			mz := (z1 + z2 + z3 + z4) / 4
			color = getColor(mz)
			fmt.Fprintf(w, "<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func getColor(mz float64) (string) {
	r, g, b := 0, 0, 0
	if mz > up {
		mz = up
	} else if mz < down {
		mz = down
	}
	percent := 100 * ((mz - down) / (up - down))
	r = int((percent / 100) * 255)
	b = 255 - r
	return fmt.Sprintf("#%.2x%.2x%.2x", r, g, b)
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

// var flag = false

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	r = math.Sin(r) / r 
	r, err := strconv.ParseFloat(fmt.Sprintf("%.8f", r), 64)
	if err != nil {
		return 0
	}
	return r
}

//!-