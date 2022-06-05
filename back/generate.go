package mandelgo

import (
	"image"
	"image/color"
	"math/cmplx"
)

type ColorScheme int

const (
	Mono ColorScheme = iota
	GrayScale
	RedScale
	GreenScale
	BlueScale
)

func ColorSchemeFromValue(val int) ColorScheme {
	switch val {
	case 0:
		return Mono
	case 1:
		return GrayScale
	case 2:
		return RedScale
	case 3:
		return GreenScale
	case 4:
		return BlueScale
	}
	return GrayScale
}

func mandelbrot(z complex128, cs ColorScheme) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch cs {
			case Mono:
				return color.White
			case GrayScale:
				return color.Gray{255 - contrast*n}
			case RedScale:
				return color.RGBA{255 - contrast*n, 20, 20, 255}
			case GreenScale:
				return color.RGBA{20, 255 - contrast*n, 20, 255}
			case BlueScale:
				return color.RGBA{20, 20, 255 - contrast*n, 255}
			}
		}
	}
	return color.Black
}

func GenerateImage(width int, height int, xmin float64, ymin float64, xmax float64, ymax float64, cs ColorScheme) image.Image {
	img := image.NewNRGBA((image.Rect(0, 0, width, height)))
	yvar := ymax - ymin
	xvar := xmax - xmin
	for yc := 0; yc < height; yc++ {
		y := float64(yc)/float64(height)*yvar + ymin
		for xc := 0; xc < width; xc++ {
			x := float64(xc)/float64(width)*xvar + xmin
			z := complex(x, y)
			img.Set(xc, yc, mandelbrot(z, cs))
		}
	}
	return img
}
