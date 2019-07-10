package procimg

import (
	"image"
	"image/color"
	"math"
)

func Bilinear(inputImage image.Image, f float64) image.Image {
	// 重み値を定義
	var x float64
	var y float64

	// リサイズ後
	size := inputImage.Bounds()
	size.Max.X = int(float64(inputImage.Bounds().Max.X) * f)
	size.Max.Y = int(float64(inputImage.Bounds().Max.Y) * f)

	// 逆数
	reciprocalScalingRows := 1 / f
	reciprocalScalingCols := 1 / f

	// アウトプット画像を定義
	outputImage := image.NewRGBA(size)



	var outputColor color.RGBA64

	// 画像の左上から順に画素を読み込む
	for imgRows := 0; imgRows < size.Max.Y; imgRows++ {
		for imgCols := 0; imgCols < size.Max.X; imgCols++ {

			// 双一次補完式

			// 元画像の座標定義
			// 元画像の縦の座標
			inputRows := int(float64(imgRows) * reciprocalScalingRows)
			// 元画像の横の座標
			inputCols := int(float64(imgCols) * reciprocalScalingCols)

			// 補完式で使う元画像のpixel
			// point(0, 0)
			src00 := inputImage.At(inputCols, inputRows)
			// point(0, 1)
			src01 := inputImage.At(inputCols + 1, inputRows)
			// point(1, 0)
			src10 := inputImage.At(inputCols, inputRows + 1)
			// point(1, 1)
			src11 := inputImage.At(inputCols + 1, inputRows + 1)

			// 重み値を算出
			x = float64(imgCols) * reciprocalScalingCols
			y = float64(imgRows) * reciprocalScalingRows

			// 小数点以下を抽出
			x = x - float64(math.Trunc(x))
			y = y - float64(math.Trunc(y))

			r00, g00, b00, a00 := src00.RGBA()
			r01, g01, b01, _ := src01.RGBA()
			r10, g10, b10, _ := src10.RGBA()
			r11, g11, b11, _ := src11.RGBA()

			// 拡大後の画素を算出
			outputColor.R = uint16((1 - x) * (1 - y) * float64(r00))
			outputColor.G = uint16((1 - x) * (1 - y) * float64(g00))
			outputColor.B = uint16((1 - x) * (1 - y) * float64(b00))


			outputColor.R += uint16(x * (1 - y) * float64(r01))
			outputColor.G += uint16(x * (1 - y) * float64(g01))
			outputColor.B += uint16(x * (1 - y) * float64(b01))

			outputColor.R += uint16((1 - x) * y * float64(r10))
			outputColor.G += uint16((1 - x) * y * float64(g10))
			outputColor.B += uint16((1 - x) * y * float64(b10))

			outputColor.R += uint16(x * y * float64(r11))
			outputColor.G += uint16(x * y * float64(g11))
			outputColor.B += uint16(x * y * float64(b11))

			outputColor.A = uint16(a00)


			outputImage.Set(imgCols, imgRows, outputColor)
		}
	}

	return outputImage
}
