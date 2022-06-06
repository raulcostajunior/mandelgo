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
	MultiHue
)

func ColorSchemeFromValue(val int) ColorScheme {
	switch val {
	case 0:
		return Mono
	case 1:
		return GrayScale
	case 2:
		return MultiHue
	}
	return MultiHue
}

func HSV_2_RGBA(h uint16, s uint8, v uint8) (r, g, b, a uint32) {
	// Converts a color given in (H)ue, (S)aturation, (V)alue model into its
	// equivalent in the RGBA model.
	// Direct implementation of the graph in this image:
	// https://en.wikipedia.org/wiki/HSL_and_HSV#/media/File:HSV-RGB-comparison.svg
	max := uint32(v) * 255
	min := uint32(v) * uint32(255-s)

	h %= 360
	segment := h / 60
	offset := uint32(h % 60)
	mid := ((max - min) * offset) / 60

	switch segment {
	case 0:
		return max, min + mid, min, 0xffff
	case 1:
		return max - mid, max, min, 0xffff
	case 2:
		return min, max, min + mid, 0xffff
	case 3:
		return min, max - mid, max, 0xffff
	case 4:
		return min + mid, min, max, 0xffff
	case 5:
		return max, min, max - mid, 0xffff
	}
	return 0, 0, 0, 0xffff
}

func mandelbrot(z complex128, cs ColorScheme) color.Color {
	// TODO add smooth coloring option (avoid banding). Details at:
	// https://en.wikipedia.org/wiki/Plotting_algorithms_for_the_Mandelbrot_set
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
			case MultiHue:
				r, g, b, a := HSV_2_RGBA(uint16(n+15), 180, 160)
				return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			}
		}
	}
	return color.Black
}

func GenerateImage(width int, height int, xmin float64, ymin float64, xmax float64,
	ymax float64, cs ColorScheme) image.Image {

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
