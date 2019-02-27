package game

/*
	Singleton
*/

import (
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// TextureManager ...
type TextureManager struct {
	textureMap map[string]*sdl.Texture
}

var tm *TextureManager
var tonce sync.Once

// STextureManager ...
var STextureManager = newTManager()

// newTManager ... convert into a singleton
func newTManager() *TextureManager {
	tonce.Do(func() {
		tm = &TextureManager{
			textureMap: make(map[string]*sdl.Texture),
		}
	})
	return tm
}

// Load ...
func (t *TextureManager) Load(fname, id string, render *sdl.Renderer) {
	var err error

	tmpSurface, err := img.Load(fname)
	checkError(err)
	defer tmpSurface.Free()

	texure, err := render.CreateTextureFromSurface(tmpSurface)
	checkError(err)

	t.textureMap[id] = texure
}

// Draw ...
func (t *TextureManager) Draw(id string, x, y, width, height int32, ren *sdl.Renderer, flip sdl.RendererFlip) {
	var err error

	desRect := sdl.Rect{
		X: x,
		Y: y,
		W: width,
		H: height}

	srcRect := sdl.Rect{
		X: 0,
		Y: 0,
		W: desRect.W,
		H: desRect.H}

	err = ren.CopyEx(t.textureMap[id], &srcRect, &desRect, 0, nil, flip)
	checkError(err)
}

// DrawFrame ...
func (t *TextureManager) DrawFrame(id string, x, y, width, height, currentRow, currentFrame int32, render *sdl.Renderer, flip sdl.RendererFlip) {
	var err error

	desRect := sdl.Rect{
		X: x,
		Y: y,
		W: width,
		H: height,
	}

	srcRect := sdl.Rect{
		X: width * currentFrame,
		Y: height * (currentRow - 1),
		W: desRect.W,
		H: desRect.H,
	}

	err = render.CopyEx(t.textureMap[id], &srcRect, &desRect, 0, nil, flip)
	checkError(err)
}
