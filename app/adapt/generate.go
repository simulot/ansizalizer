package adapt

import (
	"bufio"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strings"

	"github.com/mccutchen/palettor"

	"github.com/Zebbeni/ansizalizer/controls/settings/palettes/adaptive"
)

func GeneratePalette(m adaptive.Model, imgFilePath string) (color.Palette, string) {
	if imgFilePath == "" {
		return nil, ""
	}

	var img image.Image
	imgFile, err := os.Open(imgFilePath)
	if err != nil {
		return nil, ""
	}
	defer imgFile.Close()
	imageReader := bufio.NewReader(imgFile)
	img, _, err = image.Decode(imageReader)
	if err != nil {
		return nil, ""
	}

	count, iterations := m.Info()
	palette, err := palettor.Extract(count, iterations, img)

	name := strings.Split(filepath.Base(imgFilePath), ".")[0]

	return palette.Colors(), name
}
