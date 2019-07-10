package procimg

import (
	"image"
)



func Preparation(inputImage image.Image) image.Image {

	// Flip upside down
	centerImage := reverse(inputImage, 0)
	// Flip horizontal
	sideImage := reverse(inputImage, 1)
	// Flip upside down & horizontal
	diagonalImage := reverse(inputImage, -1)

	size := inputImage.Bounds()
	inputImageX := inputImage.Bounds().Max.X
	inputImageY := inputImage.Bounds().Max.Y
	size.Max.X = inputImage.Bounds().Max.X * 3
	size.Max.Y = inputImage.Bounds().Max.Y * 3

	outputImage := image.NewRGBA(size)

	// (0, 0)
	insertionImage(diagonalImage, outputImage, 0, 0)
	// (1, 0)
	insertionImage(centerImage, outputImage, inputImageX, 0)
	// (2, 0)
	insertionImage(diagonalImage, outputImage, inputImageX * 2, 0)
	// (0, 1)
	insertionImage(sideImage, outputImage, 0, inputImageY)
	// (1, 1)
	insertionImage(inputImage, outputImage, inputImageX, inputImageY)
	// (2, 1)
	insertionImage(sideImage, outputImage, inputImageX * 2, inputImageY)
	// (0, 2)
	insertionImage(diagonalImage, outputImage, 0, inputImageY * 2)
	// (1, 2)
	insertionImage(centerImage, outputImage, inputImageX, inputImageY * 2)
	// (2, 2)
	insertionImage(diagonalImage, outputImage, inputImageX * 2, inputImageY * 2)


	return outputImage
}

func insertionImage(inputImage image.Image, outputImage *image.RGBA, pointX int, pointY int) {

	// Insert pixel
	for y := inputImage.Bounds().Min.Y; y < inputImage.Bounds().Max.Y; y++ {
		for x := inputImage.Bounds().Min.X; x < inputImage.Bounds().Max.Y; x++ {

			outputImage.Set(x + pointX, y + pointY, inputImage.At(x, y))

		}
	}
}
