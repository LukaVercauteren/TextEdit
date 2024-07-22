package main

import (
	"TextEdit/base/cursor"
	"TextEdit/base/settings"
	"TextEdit/ui"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	fontName = "Iosevka Medium"
	fontPath = "/usr/local/share/fonts/i/Iosevka_Medium.ttc"
	fontSize = 12
)

func main() {
	// Init and get font
	ttf.Init()

	// Init settings struct
	font, err := ttf.OpenFont(fontPath, fontSize)
	if err != nil {
		return
	}
	defer font.Close()

	settings.ProgSettings = settings.Settings{
		Font:      *font,
		FontName:  fontName,
		FontPath:  fontPath,
		TextSize:  fontSize,
		TextColor: sdl.Color{R: 255, G: 255, B: 255, A: 255},
	}

	ui.BufferToSurface()
	cursor := cursor.NewCursor()
	fmt.Printf("cursor: %v\n", cursor)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("TextEdit Development Version", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	surface.FillRect(nil, 0)

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	color := sdl.Color{R: 255, G: 0, B: 255, A: 255}
	pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)

	surface.FillRect(&rect, pixel)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
				println("Quit")
				running = false
				break
			}
		}

		sdl.Delay(33)
	}
}
