package ui

import (
	"TextEdit/base/buffer"
	"TextEdit/base/settings"

	"github.com/veandco/go-sdl2/sdl"
)

func BufferToSurface(buffer buffer.Buffer) sdl.Surface {
	font := settings.ProgSettings.Font
	textColor := settings.ProgSettings.TextColor
	surface, _ := font.RenderUTF8Blended("test", textColor)

	for key, value := range buffer.Content {

	}
}
