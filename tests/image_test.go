package tests

import (
	"fmt"
	"testing"

	"avenger/tools/image"
)

func TestImage(t *testing.T) {
	width, height, err := image.Size("./sample/images/test.png")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(width, height)
}

func TestImageScale(t *testing.T) {
	err := image.Scale("./sample/images/test.png", "./sample/images/scale.png", 200, 200, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = image.ScaleByWidth("./sample/images/test.png", "./sample/images/ScaleByWidth.png", 200, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = image.ScaleByHeight("./sample/images/test.png", "./sample/images/ScaleByHeight.png", 200, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
}

func TestImageClip(t *testing.T) {
	err := image.Thumbnail("./sample/images/test.png", "./sample/images/Clip_100_100.png", 200, 200, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = image.Thumbnail("./sample/images/test.png", "./sample/images/Clip_100_300.png", 100, 300, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = image.Thumbnail("./sample/images/test.png", "./sample/images/Clip_300_100.png", 300, 100, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
}

func TestImageRotate(t *testing.T) {
	err := image.Rotate("./sample/images/test.png", "./sample/images/Rotate90.png", image.Rotate90, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = image.Rotate("./sample/images/test.png", "./sample/images/Rotate180.png", image.Rotate180, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = image.Rotate("./sample/images/test.png", "./sample/images/Rotate270.png", image.Rotate270, image.PNG)
	if err != nil {
		t.Fatal(err)
	}
}
