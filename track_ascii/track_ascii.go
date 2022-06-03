package track_ascii

//This code is largely lifted from https://github.com/stdupp/goasciiart
//There's no license on that repo, so I'm assuming public domain or some kind of share and share alike
//@stdupp if you ever have a problem with this or want me to credit differently I'm happy to

import (
	"strings"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/ipc"
	"github.com/adamrdrew/mosh/server"
	"github.com/nfnt/resize"

	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"reflect"
)

const ASCIISTR = "MND8OZ$7I?+=~:,.."

func MakeConverterForTrack(track ipc.ResponseItem) ImageConverter {
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

	buffer, w, h := i.scaleImage(image, 40)

	out := i.convert2Ascii(buffer, w, h)

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

func (i *ImageConverter) convert2Ascii(img image.Image, w, h int) []byte {
	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes()
}
