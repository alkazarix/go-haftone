package halftone

import (
	"image"
	"image/color"
	"math"
)

var black = color.Gray{Y: 0}
var white = color.Gray{Y: 255}

func Halftone(img image.Image, filter Filter) image.Image {

	grayimg := grayscale(img)
	bounds := grayimg.Bounds()

	dx, dy := bounds.Dx(), bounds.Dy()
	dimx, dimy := len(filter.Matrix[0]), len(filter.Matrix)
	nx, ny := dimx*int(math.Floor(float64(dx))/float64(dimx)), dimy*int(math.Floor(float64(dy)/float64(dimy)))

	n := filter.Max() - filter.Min() + 2
	output := image.NewGray(image.Rect(0, 0, nx-1, ny-1))

	for x := 0; x < nx; x++ {
		for y := 0; y < ny; y++ {
			quantify := (float64(n-1) / 255) * float64(grayimg.GrayAt(bounds.Min.X+x, bounds.Min.Y+y).Y)
			f := filter.Matrix[x%dimx][y%dimy]
			if quantify < float64(f) {
				output.SetGray(x, y, black)
			} else {
				output.SetGray(x, y, white)
			}
		}
	}

	return output
}

func grayscale(img image.Image) *image.Gray {

	bounds := img.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()
	grayimg := image.NewGray(bounds)

	for x := bounds.Min.X; x < dx; x++ {
		for y := bounds.Min.Y; y < dy; y++ {
			grayimg.Set(x, y, img.At(x, y))
		}
	}

	return grayimg
}
