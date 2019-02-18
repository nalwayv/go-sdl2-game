package main

import (
	"fmt"

	"./gologger"
	"github.com/veandco/go-sdl2/sdl"
)

// Window ...
var (
	Window  *sdl.Window
	Surface *sdl.Surface
	golog   = gologger.GetInstance("src/gologger/golog.log")
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

func checkError(err error) {
	if err != nil {
		golog.Fatalln(err)
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
		return nil, err
	}

	return window, nil
}

func loadMedia(fname string, format *sdl.PixelFormat) (*sdl.Surface, error) {
	surface, err := sdl.LoadBMP(fname)
	if err != nil {
		return nil, err
	}

	optimizedSurface, err := surface.Convert(format, 0)
	if err != nil {
		return nil, err
	}

	return optimizedSurface, nil
}

func setUpKeyPressSurface(face *sdl.Surface) ([]*sdl.Surface, error) {

	keySurface := make([]*sdl.Surface, 5)
	var err error
	format := face.Format

	def, err := loadMedia("assets/hello_world.bmp", format)
	if err != nil {
		return nil, err
	}
	keySurface[KEY_PRESS_DEFAULT] = def

	up, err := loadMedia("assets/up.bmp", format)
	if err != nil {
		return nil, err
	}
	keySurface[KEY_PRESS_UP] = up

	down, err := loadMedia("assets/down.bmp", format)
	if err != nil {
		return nil, err
	}
	keySurface[KEY_PRESS_DOWN] = down

	left, err := loadMedia("assets/left.bmp", format)
	if err != nil {
		return nil, err
	}
	keySurface[KEY_PRESS_LEFT] = left

	right, err := loadMedia("assets/right.bmp", format)
	if err != nil {
		return nil, err
	}
	keySurface[KEY_PRESS_RIGHT] = right

	return keySurface, nil
}

func cleanSurfaces(imgs []*sdl.Surface) {
	golog.Println("clean keypress surface")

	for _, v := range imgs {
		v.Free()
	}
}

func close() {
	golog.Println("cleanup")

	Window.Destroy()

	sdl.Quit()
}

func main() {
	// SETUP --
	var err error

	Window, err = InitSDL()
	checkError(err)
	defer close()

	Surface, err := Window.GetSurface()
	checkError(err)
	defer Surface.Free()

	KeyPressSurfaces, err := setUpKeyPressSurface(Surface)
	checkError(err)
	defer cleanSurfaces(KeyPressSurfaces)

	rect := sdl.Rect{X: 0, Y: 0, W: WINDOW_WIDTH, H: WINDOW_HEIGHT}
	KeyPressSurfaces[KEY_PRESS_DEFAULT].BlitScaled(nil, Surface, &rect)
	Window.UpdateSurface()

	// EVENT-LOOP --
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
					fmt.Println("Up")
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
			}
		}

		rect := sdl.Rect{X: 0, Y: 0, W: WINDOW_WIDTH, H: WINDOW_HEIGHT}
		nextSurface.BlitScaled(nil, Surface, &rect)
		Window.UpdateSurface()
	}
}
