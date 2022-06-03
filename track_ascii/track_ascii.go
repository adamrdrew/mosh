package track_ascii

//This code is largely lifted from https://github.com/ajmalsiddiqui/termpic
//That code isn't a lib, its an executable, so I copied the code out
//The original is MIT so lifting it should be cool
//Shout out to @ajmalsiddiqui for the code - great work!

import (
	"fmt"
	"strings"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/ipc"
	"github.com/adamrdrew/mosh/responses"
	"github.com/adamrdrew/mosh/server"
	"github.com/nfnt/resize"

	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

const UPPER_HALF_BLOCK = "▀"

func MakeConverterForResponseTrack(source responses.ResponseTrack) ImageConverter {
	track := ipc.ResponseItem{
		Image: source.Image,
	}
	i := ImageConverter{
		Track: track,
	}
	return i
}

func MakeConverterForResponseItem(track ipc.ResponseItem) ImageConverter {
	i := ImageConverter{
		Track: track,
	}
	return i
}

type ImageConverter struct {
	Track ipc.ResponseItem
}

func (i *ImageConverter) GetAscii() string {

	image := i.getImage()

	out := i.convert2Ascii(image, 24)

	return string(out)
}

func (i *ImageConverter) getImage() image.Image {
	conf := config.GetConfig()
	server := server.GetServer(&conf)
	encodedImage := server.GetArtForSong(i.Track)
	reader := strings.NewReader(encodedImage)
	img, _, _ := image.Decode(reader)
	return img
}

func (i *ImageConverter) scaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}

func (i *ImageConverter) convert2Ascii(img image.Image, skip int) string {
	// We'll just reuse this to increment the loop counters
	skip += 1
	ansi := i.resetColorSequence()
	yMax := img.Bounds().Max.Y
	xMax := img.Bounds().Max.X

	sequences := make([]string, yMax)

	for y := img.Bounds().Min.Y; y < yMax; y += 2 * skip {
		sequence := ""
		for x := img.Bounds().Min.X; x < xMax; x += skip {
			upperPix := img.At(x, y)
			lowerPix := img.At(x, y+skip)

			ur, ug, ub := i.convertColorToRGB(upperPix)
			lr, lg, lb := i.convertColorToRGB(lowerPix)

			if y+skip >= yMax {
				sequence += i.resetColorSequence()
			} else {
				sequence += i.rgbBackgroundSequence(lr, lg, lb)
			}

			sequence += i.rgbTextSequence(ur, ug, ub)
			sequence += UPPER_HALF_BLOCK

			sequences[y] = sequence
		}
	}

	for y := img.Bounds().Min.Y; y < yMax; y += 2 * skip {
		ansi += sequences[y] + i.resetColorSequence() + "\n"
	}

	return ansi
}

func (i *ImageConverter) rgbBackgroundSequence(r, g, b uint8) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b)
}

// 38;2;r;g;bm - set text colour to rgb
func (i *ImageConverter) rgbTextSequence(r, g, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func (i *ImageConverter) resetColorSequence() string {
	return "\x1b[0m"
}

func (i *ImageConverter) convertColorToRGB(col color.Color) (uint8, uint8, uint8) {
	rgbaColor := color.RGBAModel.Convert(col)
	_r, _g, _b, _ := rgbaColor.RGBA()
	// rgb values are uint8s, I cannot comprehend why the stdlib would return
	// int32s :facepalm:
	r := uint8(_r & 0xFF)
	g := uint8(_g & 0xFF)
	b := uint8(_b & 0xFF)
	return r, g, b
}
