package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Window ...
var (
	Window *sdl.Window
)

// width / height
const (
	WINDOW_WIDTH  int32 = 800
	WINDOW_HEIGHT int32 = 600
)

// KEY PRESS ENUMS ...
const (
	KEY_PRESS_DEFAULT = iota
	KEY_PRESS_UP
	KEY_PRESS_DOWN
	KEY_PRESS_LEFT
	KEY_PRESS_RIGHT
)

func checkError(err error, msg string) {
	if err != nil {
		panic(err)
	}
}

// InitSDL ...
func InitSDL() (*sdl.Window, error) {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		sdl.GetError()
		return nil, err
	}

	window, err := sdl.CreateWindow("testing", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.GetError()
		return nil, err
	}

	return window, nil
}

func loadMedia(fname string) (*sdl.Surface, error) {
	surface, err := sdl.LoadBMP(fname)
	if err != nil {
		sdl.GetError()
		return nil, err
	}
	return surface, nil
}

func setUpKeyPressSurface() ([]*sdl.Surface, error) {

	keySurface := make([]*sdl.Surface, 5)
	var err error

	def, err := loadMedia("assets/hello_world.bmp")
	if err != nil {
		return nil, err
	}
	keySurface[KEY_PRESS_DEFAULT] = def

	up, err := loadMedia("assets/up.bmp")
	if err != nil {
		return nil, err
	}
	keySurface[KEY_PRESS_UP] = up

	down, err := loadMedia("assets/down.bmp")
	if err != nil {
		return nil, err
	}
	keySurface[KEY_PRESS_DOWN] = down

	left, err := loadMedia("assets/left.bmp")
	if err != nil {
		return nil, err
	}
	keySurface[KEY_PRESS_LEFT] = left

	right, err := loadMedia("assets/right.bmp")
	if err != nil {
		return nil, err
	}
	keySurface[KEY_PRESS_RIGHT] = right

	return keySurface, nil
}

func cleanSurfaces(imgs []*sdl.Surface) {
	fmt.Println("clean keypress surface")

	for _, v := range imgs {
		v.Free()
	}
}

func close() {
	fmt.Println("cleanup")

	Window.Destroy()

	sdl.Quit()
}

func main() {
	fmt.Print("SDL")
	var err error

	Window, err = InitSDL()
	checkError(err, "failed to init")
	defer close()

	surface, err := Window.GetSurface()
	checkError(err, "window surface fail")
	defer surface.Free()

	KeyPressSurfaces, err := setUpKeyPressSurface()
	checkError(err, "failed to get img")
	defer cleanSurfaces(KeyPressSurfaces)

	KeyPressSurfaces[KEY_PRESS_DEFAULT].Blit(nil, surface, nil)
	Window.UpdateSurface()

	var event sdl.Event
	quit := false
	for !quit {
		var nextSurface *sdl.Surface
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch t := event.(type) {
			case *sdl.QuitEvent:
				quit = true

			case *sdl.KeyboardEvent:
				switch t.Keysym.Scancode {
				case sdl.SCANCODE_UP:
					fmt.Println("UP")
					nextSurface = KeyPressSurfaces[KEY_PRESS_UP]

				case sdl.SCANCODE_DOWN:
					fmt.Println("DOWN")
					nextSurface = KeyPressSurfaces[KEY_PRESS_DOWN]

				case sdl.SCANCODE_LEFT:
					fmt.Println("LEFT")
					nextSurface = KeyPressSurfaces[KEY_PRESS_LEFT]

				case sdl.SCANCODE_RIGHT:
					fmt.Println("RIGHT")
					nextSurface = KeyPressSurfaces[KEY_PRESS_RIGHT]

				default:
					fmt.Println("DEFAULT")
					nextSurface = KeyPressSurfaces[KEY_PRESS_DEFAULT]
				}
				nextSurface.Blit(nil, surface, nil)
				Window.UpdateSurface()
			}
		}
	}
}
