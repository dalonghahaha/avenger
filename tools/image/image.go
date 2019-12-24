package image

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
)

type Format string

type Degree int

//支持的输出类型
const (
	PNG  Format = ".png"
	JPEG Format = ".jpg"
	GIF  Format = ".gif"
)

//顺时针旋转角度
const (
	Rotate90  Degree = 90
	Rotate180 Degree = 190
	Rotate270 Degree = 270
)

//输出jpeg默认参数
var jpegOption = &jpeg.Options{
	Quality: jpeg.DefaultQuality,
}

//输出gif默认参数
var gifOption = &gif.Options{
	NumColors: 256,
}

//Size 获取图片尺寸
func Size(path string) (int, int, error) {
	src, err := Load(path)
	if err != nil {
		return 0, 0, err
	}
	bound := src.Bounds()
	return bound.Dx(), bound.Dy(), nil
}

//Load 加载图片
func Load(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, err
}

//Save 保存图片
func Save(path string, img image.Image, format Format) error {
	imgfile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		imgfile.Close()
	}()
	switch format {
	case PNG:
		return png.Encode(imgfile, img)
	case JPEG:
		return jpeg.Encode(imgfile, img, jpegOption)
	case GIF:
		return gif.Encode(imgfile, img, gifOption)
	default:
		return fmt.Errorf("save format error")
	}
}

//Scale 缩放图片
func Scale(srcPath string, dstPath string, width, height int, format Format) error {
	src, err := Load(srcPath)
	if err != nil {
		return err
	}
	dst := imaging.Resize(src, width, height, imaging.Lanczos)
	return Save(dstPath, dst, format)
}

//ScaleByWidth 按宽度等比例缩放图片
func ScaleByWidth(srcPath string, dstPath string, width int, format Format) error {
	src, err := Load(srcPath)
	if err != nil {
		return err
	}
	dst := imaging.Resize(src, width, 0, imaging.Lanczos)
	return Save(dstPath, dst, format)
}

//ScaleByWidth 按高度等比例缩放图片
func ScaleByHeight(srcPath string, dstPath string, height int, format Format) error {
	src, err := Load(srcPath)
	if err != nil {
		return err
	}
	dst := imaging.Resize(src, 0, height, imaging.Lanczos)
	return Save(dstPath, dst, format)
}

//Clip 按比例自动剪裁(截取中间部分)并缩放图片
func Thumbnail(srcPath string, dstPath string, width, height int, format Format) error {
	src, err := Load(srcPath)
	if err != nil {
		return err
	}
	dst := imaging.Fill(src, width, height, imaging.Center, imaging.Lanczos)
	return Save(dstPath, dst, format)
}

//Rotate 旋转图片
func Rotate(srcPath string, dstPath string, degree Degree, format Format) error {
	src, err := Load(srcPath)
	if err != nil {
		return err
	}
	var dst *image.NRGBA
	switch degree {
	case Rotate90:
		dst = imaging.Rotate90(src)
	case Rotate180:
		dst = imaging.Rotate180(src)
	case Rotate270:
		dst = imaging.Rotate270(src)
	default:
		return fmt.Errorf("wrong degree")
	}
	return Save(dstPath, dst, format)
}
