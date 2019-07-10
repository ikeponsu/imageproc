package main

import (
	"fmt"
	"image/jpeg"
	"img-test/ioimg"
	"img-test/procimg"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

func main() {

	count := 0

	datapath := "Data/sample/*.jpg"

	file, _ := filepath.Glob(datapath)

	timeLayout := time.Now()
	timeString := timeLayout.Format("20060102150405")
	d_path := filepath.Join("data", "result", timeString)
	if err := os.Mkdir(d_path, 0777); err != nil {
		fmt.Println(err)
	}

	start := time.Now()
	var wg sync.WaitGroup

	for _, item := range file{

		wg.Add(1)
		go func(item string, timeString string, count int) {
			defer wg.Done()
			exc(item, timeString, count)

		}(item, timeString, count)
		count++
	}

	wg.Wait()

	end := time.Now()
	fmt.Println("B.Total time: ", end.Sub(start))

}

func exc(item string, timeString string, count int) {

	img := ioimg.Input(item)

	rimg := procimg.Bilinear(img, 0.5)

	path := filepath.Join("Data", "result", timeString, strconv.Itoa(count) + ".jpg")

	qt := jpeg.Options{
		Quality:60,
	}
	file, _ := os.Create(path)
	jpeg.Encode(file, rimg, &qt)
}
