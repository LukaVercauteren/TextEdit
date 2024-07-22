package settings

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Settings struct {
	Font      ttf.Font
	FontName  string
	FontPath  string
	TextSize  uint8
	TextColor sdl.Color
}

var ProgSettings Settings
