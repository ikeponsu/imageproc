package procimg

import "image"

var filterWeight = map[int]map[int]float64{}

// 重み値の合計
const weightSum int = 16

func initgaussian() {
	// フィルタ処理で使用する重み値
	filterWeight[-1] = make(map[int]float64)
	filterWeight[0] = make(map[int]float64)
	filterWeight[1] = make(map[int]float64)
	filterWeight[-1][-1] = 1
	filterWeight[-1][0] = 2
	filterWeight[-1][1] = 1
	filterWeight[0][-1] = 2
	filterWeight[0][0] = 4
	filterWeight[0][1] = 2
	filterWeight[1][-1] = 1
	filterWeight[1][0] = 2
	filterWeight[1][1] = 1
}

func gaussian(img image.Image) {

}
