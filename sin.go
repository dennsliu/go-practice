package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	//图⽚⼤⼩
	const size = 300
	//根据给定⼤⼩创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//遍历每⼀个元素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//填充为⽩⾊
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	for x := 0; x < size; x++ {
		//让的值的范围在之间 sin 0~2Pi
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为⼀半的像素。向下偏移⼀半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		//⽤⿊⾊绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	//创建⽂件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	//使⽤png格式将数据写⼊⽂件
	png.Encode(file, pic) //将image信息写⼊⽂件中
	//image关闭⽂件
	file.Close()
}
