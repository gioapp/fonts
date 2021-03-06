package fonts

import (
	"fmt"

	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/font/gofont/gomonobolditalic"
	"golang.org/x/image/font/gofont/gomonoitalic"

	"github.com/gioapp/fonts/bariolbold"
	"github.com/gioapp/fonts/bariolbolditalic"
	"github.com/gioapp/fonts/bariollight"
	"github.com/gioapp/fonts/bariollightitalic"
	"github.com/gioapp/fonts/bariolregular"
	"github.com/gioapp/fonts/bariolregularitalic"
	"github.com/gioapp/fonts/plan9"
)

func Register() {
	register(text.Font{Typeface: "bariol"}, bariolregular.TTF)
	register(text.Font{Typeface: "plan9"}, plan9.TTF)
	register(text.Font{Typeface: "bariol", Style: text.Italic}, bariolregularitalic.TTF)
	register(text.Font{Typeface: "bariol", Weight: text.Bold}, bariolbold.TTF)
	register(text.Font{Typeface: "bariol", Style: text.Italic, Weight: text.Bold}, bariolbolditalic.TTF)
	register(text.Font{Typeface: "bariol", Weight: text.Medium}, bariollight.TTF)
	register(text.Font{Typeface: "bariol", Weight: text.Medium, Style: text.Italic}, bariollightitalic.TTF)
	register(text.Font{Typeface: "go"}, gomono.TTF)
	register(text.Font{Typeface: "go", Weight: text.Bold}, gomonobold.TTF)
	register(text.Font{Typeface: "go", Weight: text.Bold, Style: text.Italic}, gomonobolditalic.TTF)
	register(text.Font{Typeface: "go", Style: text.Italic}, gomonoitalic.TTF)
	register(text.Font{Typeface: "go", Style: text.Italic}, gomonoitalic.TTF)
}

func register(fnt text.Font, ttf []byte) {
	face, err := opentype.Parse(ttf)
	if err != nil {
		panic(fmt.Sprintf("failed to parse font: %v", err))
	}
	font.Register(fnt, face)
}
