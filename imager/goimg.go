package goimage

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	gocolor "github.com/egor-erm/goimager/manager"
	"github.com/go-gl/mathgl/mgl32"
)

type goimage struct {
	name  string
	image *image.RGBA
}

func New(name string, xmax int, ymax int) *goimage {
	b := image.Rect(0, 0, xmax, ymax)
	image := image.NewRGBA(b)

	return &goimage{name, image}
}

func NewWithCorners(name string, xmin int, ymin int, xmax int, ymax int) *goimage {
	b := image.Rect(xmin, ymin, xmax, ymax)
	image := image.NewRGBA(b)

	return &goimage{name, image}
}

func (gimg *goimage) Save() error {
	file, err := os.Create(gimg.name)

	if err != nil || file == nil {
		file, err = os.Open(gimg.name)
		if err != nil {
			return fmt.Errorf("error opening file: %s", err)
		}
	}

	err = png.Encode(file, gimg.image)
	if err != nil {
		return fmt.Errorf("error encoding image: %s", err)
	}

	file.Close()
	return nil
}

func (gimg *goimage) SetPixel(x int, y int, color color.RGBA) {
	gimg.image.SetRGBA(x, y, color)
}

func (gimg *goimage) SetPixelByVector(vec mgl32.Vec2, color color.RGBA) {
	gimg.image.SetRGBA(int(vec.X()), int(vec.Y()), color)
}

func (gimg *goimage) SetHEXPixel(x int, y int, c string) {
	color := gocolor.HEXtoRGBA(c)
	gimg.SetPixel(x, y, color)
}

func (gimg *goimage) SetHEXPixelByVector(vec mgl32.Vec2, c string) {
	color := gocolor.HEXtoRGBA(c)
	gimg.SetPixel(int(vec.X()), int(vec.Y()), color)
}

func (gimg *goimage) SetHEXAlphaPixel(x int, y int, c string, alpha uint8) {
	color := gocolor.HEXAlphatoRGBA(c, alpha)
	gimg.SetPixel(x, y, color)
}

func (gimg *goimage) SetHEXAlphaPixelByVector(vec mgl32.Vec2, c string, alpha uint8) {
	color := gocolor.HEXAlphatoRGBA(c, alpha)
	gimg.SetPixel(int(vec.X()), int(vec.Y()), color)
}

func (gimg *goimage) ClearPixel(x int, y int) {
	gimg.SetPixel(x, y, color.RGBA{0, 0, 0, 0})
}

func (gimg *goimage) ClearPixelByVectros(vec mgl32.Vec2) {
	gimg.SetPixel(int(vec.X()), int(vec.Y()), color.RGBA{0, 0, 0, 0})
}