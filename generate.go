package mandelgo

import (
	"image"
	"image/color"
	"math/cmplx"
)

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func GenerateImage(width int, height int, xmin int, ymin int, xmax int, ymax int) image.Image {
	img := image.NewNRGBA((image.Rect(0, 0, width, height)))
	yvar := float64(ymax - ymin)
	xvar := float64(xmax - xmin)
	for yc := 0; yc < height; yc++ {
		y := float64(yc)/float64(height)*yvar + float64(ymin)
		for xc := 0; xc < width; xc++ {
			x := float64(xc)/float64(width)*xvar + float64(xmin)
			z := complex(x, y)
			img.Set(xc, yc, mandelbrot(z))
		}
	}
	return img
}
