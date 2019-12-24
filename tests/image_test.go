package tests

import (
	"fmt"
	"testing"

	"avenger/tools/image"
)

func TestImage(t *testing.T) {
	width, height, err := image.Size("./sample/test.png")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(width, height)
	//Scale
	err = image.Scale("./sample/test.png", "./sample/scale.png", 200, 200, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = image.ScaleByWidth("./sample/test.png", "./sample/ScaleByWidth.png", 200, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = image.ScaleByHeight("./sample/test.png", "./sample/ScaleByHeight.png", 200, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	//Clip
	err = image.Thumbnail("./sample/test.png", "./sample/Clip_200_200.png", 200, 200, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	//Rotate
	err = image.Rotate("./sample/test.png", "./sample/Rotate90.png", image.Rotate90, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = image.Rotate("./sample/test.png", "./sample/Rotate180.png", image.Rotate180, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = image.Rotate("./sample/test.png", "./sample/Rotate270.png", image.Rotate270, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
}
