package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ericaro/frontmatter"
	"github.com/fogleman/gg"
)

type Post struct {
	Title string   `yaml:"title"`
	Tags  []string `yaml:"tags"`
}

func main() {
	if err := filepath.Walk("./content/posts/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		body, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var out Post
		frontmatter.Unmarshal(body, &out)

		img, err := run(out.Title)
		if err != nil {
			return err
		}

		filename := "./static/images/" + strings.ReplaceAll(info.Name(), ".md", ".gif")
		fmt.Printf("Saving %s", filename)
		ioutil.WriteFile(filename, img, 0755)

		return nil
	}); err != nil {
		panic(err)
	}
}

const (
	width  = 1200
	height = 628
)

var (
	monoFont   = filepath.Join("/fonts", "Noto.ttf")
	headerFont = filepath.Join("/fonts", "Inter-SemiBold.ttf")
	fontColor  = color.RGBA{169, 169, 179, 0xff}
)

func run(title string) ([]byte, error) {
	dc := gg.NewContext(width, height)

	// Background
	dc.SetHexColor("#292a2d")
	dc.DrawRectangle(0, 0, width, height)
	dc.Fill()

	// Header
	dc.SetHexColor("#252627")
	dc.DrawRectangle(0, 0, width, 100)
	dc.Fill()

	if err := dc.LoadFontFace(monoFont, 45); err != nil {
		return nil, err
	}
	dc.SetColor(fontColor)
	dc.DrawString("#til", 45, 60)

	dc.LoadFontFace(monoFont, 20)
	dc.SetColor(fontColor)
	dc.DrawString("https://til.marcusnoble.co.uk", 895, 623)

	dc.LoadFontFace(headerFont, 80)
	textRightMargin := 50.0
	textTopMargin := 130.0
	x := textRightMargin
	y := textTopMargin
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin
	dc.SetColor(color.White)
	dc.DrawStringWrapped(title, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	dc.SetColor(fontColor)
	dc.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)

	frame1 := dc.Image()

	dc = gg.NewContextForImage(frame1)

	dc.LoadFontFace(headerFont, 80)
	title += " |"
	textRightMargin = 50.0
	textTopMargin = 130.0
	x = textRightMargin
	y = textTopMargin
	maxWidth = float64(dc.Width()) - textRightMargin - textRightMargin
	dc.SetColor(color.White)
	dc.DrawStringWrapped(title, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	dc.SetColor(fontColor)
	dc.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)

	frame2 := dc.Image()

	palettedImage1 := image.NewPaletted(frame1.Bounds(), palette.Plan9)
	draw.Over.Draw(palettedImage1, frame1.Bounds(), frame1, image.ZP)
	palettedImage2 := image.NewPaletted(frame2.Bounds(), palette.Plan9)
	draw.Over.Draw(palettedImage2, frame2.Bounds(), frame2, image.ZP)

	var output bytes.Buffer
	gif.EncodeAll(bufio.NewWriter(&output), &gif.GIF{
		Image: []*image.Paletted{
			palettedImage1,
			palettedImage2,
		},
		Delay: []int{50, 50},
	})

	return output.Bytes(), nil
}
