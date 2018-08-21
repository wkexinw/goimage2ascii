package main

import (
	"fmt"
	"strconv"
	"image/png"
	"log"
	"os"

)

func main(){

	if len(os.Args)<2 {
		log.Println("WARN: Please input an image.")
		return
	}
	image := os.Args[1]
	DrawAscii(image)
}

func DrawAscii(image string){
	base := "@#&$%*o!;."
	f, _ := os.Open(image)
	img, _ := png.Decode(f)
	bounds:=img.Bounds()
	ascii:=""
	for y:=0;y<bounds.Dy();y+=2{
		for x:=0;x<bounds.Dx();x++{
			pixel:=img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			r=r&0xFF
			g=g&0xFF
			b=b&0xFF
			gray := 0.299 * float64(r) + 0.578 * float64(g) + 0.114 * float64(b)
			temp:=fmt.Sprintf("%.0f",gray*float64(len(base)+1)/255)
			index,_:=strconv.Atoi(temp)
			if index>=len(base) {
				ascii+=" "
				//fmt.Print(" ")
			}else{
				ascii+=string(base[index])
				//fmt.Print(string(base[index]))
			}

		}
		
		ascii+="\n"

	}
	f.Close()

	fmt.Printf("\033[32;1m%s\033[0m\n",ascii)
}