package mandelgo

import (
	"image"
)

func GenerateImage(width int, height int, xmin int, ymin int, xmax int, ymax int) image.Image {
	// TODO add real body
	return image.NewNRGBA((image.Rect(0, 0, width, height)))
}
