package procimg

import (
	"image"
	"image/draw"
	"log"
)


// Process to rotate the image
func reverse(inputImage image.Image, angleCode int) image.Image {

	var outputImage image.Image

	switch angleCode {
	case -1:
		// Flip upside down & horizontal
		outputImage = all(inputImage)
		break
	case 0:
		// Flip upside down
		outputImage = upDown(inputImage)
		break
	case 1:
		// Flip horizontal
		outputImage = horizon(inputImage)
		break
	default:
		log.Fatal("angle code does not exist")
		break
	}

	return outputImage

}

// Flip upside down
func upDown(inputImage image.Image) image.Image {

	size := inputImage.Bounds()

	outputImage := image.NewRGBA(size)
	
	for y := size.Min.Y; y < size.Max.Y; y++ {
		for x := size.Min.X; x < size.Max.Y; x++ {
			outputImage.Set(x, y, inputImage.At(size.Max.X - x, y))
		}
	}

	draw.FloydSteinberg.Draw(outputImage, inputImage.Bounds(), inputImage, image.ZP)

	return outputImage
}

// Flip horizontal
func horizon(inputImage image.Image) image.Image {

	size := inputImage.Bounds()

	outputImage := image.NewRGBA(size)

	for y := size.Min.Y; y < size.Max.Y; y++ {
		for x := size.Min.X; x < size.Max.Y; x++ {
			outputImage.Set(x, y, inputImage.At(x, size.Max.Y - y))
		}
	}
	
	return outputImage
}

// Flip upside down & horizontal
func all(inputImage image.Image) image.Image {

	outputImage := upDown(inputImage)
	outputImage = horizon(outputImage)
	
	return outputImage
}

