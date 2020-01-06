package main

import (
	"fmt"
	"strings"

	"github.com/fogleman/gg"
	"github.com/rivo/uniseg"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	jpFontPath := os.Args[1]
	jpFont, err := gg.LoadFontFace(jpFontPath, 96)
	if err != nil {
		panic(err)
	}
	emojiFontPath := os.Args[2]
	emojiFont, err := gg.LoadFontFace(emojiFontPath, 96)
	if err != nil {
		panic(err)
	}

	dc.SetFontFace(jpFont)
	w, _ := dc.MeasureString("こんにちは")
	dc.DrawString("こんにちは", 0, S/2)
	//dc.DrawString("🤛", w, S/2)
	dc.SetFontFace(emojiFont)

	str := strings.Join([]string{"👁", "👁︎", "👁️"}, "")
	gr := uniseg.NewGraphemes(str)
	for gr.Next() {
		rs := gr.Runes()
		if len(rs) > 1 {
			fmt.Println("skip", string(rs))
			dc.DrawString(string(rs), w, S/2)
			dw, _ := dc.MeasureString(string(rs))
			w += dw
			continue
		}

		_, _, ok := emojiFont.GlyphBounds(rs[0])
		if ok {
			dc.DrawString(string(rs), w, S/2)
			dw, _ := dc.MeasureString(string(rs))
			w += dw
			fmt.Println(string(rs))
		} else {
			fmt.Println("skip", string(rs))
		}
	}
	dc.SavePNG("out.png")
}
