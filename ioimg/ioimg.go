package ioimg

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

const imgQt int = 60

func Input (filePath string) image.Image {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	pngImage, _, err := image.Decode(file)

	return pngImage

}

func Output (outputimage image.Image, filePath string, format string) {

	dst, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	switch format {
	case "png":
		// PNGの場合
		err = png.Encode(dst, outputimage)
		if err != nil {
			log.Fatal(err)
		}
		break
	case "jpg":
		// JPGの場合
		qt := jpeg.Options{
			Quality:imgQt,
		}
		err = jpeg.Encode(dst, outputimage, &qt)
		if err != nil {
			log.Fatal(err)
		}
		break
	case "gif":
		// GIFの場合
		log.Println("can't output gif")
		break
	default:
		// 標準で対応していないフォーマットの場合
		log.Fatal("Unsupported format.")
	}
}

func OutputJpeg(outputimage image.Image, filePath string) {
	dst, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	qt := jpeg.Options{
		Quality:imgQt,
	}
	err = jpeg.Encode(dst, outputimage, &qt)
	if err != nil {
		log.Fatal(err)
	}

}

func OutputGif (gifData []image.Image, filePath string) {

	outputGif := &gif.GIF{}

	for _, gifImg := range gifData {
		outputGif.Image = append(outputGif.Image, gifImg.(*image.Paletted))
		outputGif.Delay = append(outputGif.Delay, 0)
	}

	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)
	gif.EncodeAll(file, outputGif)
}