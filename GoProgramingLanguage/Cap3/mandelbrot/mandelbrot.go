package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math/rand"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		samples                = 2 // superamostragem 2x2
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var rSum, gSum, bSum, aSum float64
			for sy := 0; sy < samples; sy++ {
				for sx := 0; sx < samples; sx++ {
					// Subpixel offsets
					offsetX := (float64(sx) + 0.5) / float64(samples)
					offsetY := (float64(sy) + 0.5) / float64(samples)
					x := (float64(px)+offsetX)/float64(width)*(xmax-xmin) + xmin
					yy := (float64(py)+offsetY)/float64(height)*(ymax-ymin) + ymin
					z := complex(x, yy)
					c := mandelbrot(z)
					r, g, b, a := c.RGBA()
					rSum += float64(r)
					gSum += float64(g)
					bSum += float64(b)
					aSum += float64(a)
				}
			}
			// Média das amostras (RGBA retorna valores de 0-65535)
			amostrasTotais := float64(samples * samples)
			img.Set(px, py, color.RGBA{
				uint8(rSum / amostrasTotais / 257),
				uint8(gSum / amostrasTotais / 257),
				uint8(bSum / amostrasTotais / 257),
				uint8(aSum / amostrasTotais / 257),
			})
		}
	}
	png.Encode(os.Stdout, img)
}

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
	var randColor1 = uint8(rand.Intn(256))
	var randColor2 = uint8(rand.Intn(256))
	var randColor3 = uint8(rand.Intn(256))
	return color.RGBA{randColor1, randColor2, randColor3, 255}
}
