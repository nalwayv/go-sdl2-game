package game

// SINGLETON

import (
	"errors"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// TextureManager ...
type TextureManager struct {
	textureMap map[string]*sdl.Texture
}

var tm *TextureManager
var tOnce sync.Once

// STextureManager ...
var STextureManager = newTManager()

// newTManager ... convert into a singleton
func newTManager() *TextureManager {
	tOnce.Do(func() {
		tm = &TextureManager{
			textureMap: make(map[string]*sdl.Texture),
		}
	})
	return tm
}

// Load ...
func (t *TextureManager) Load(fileName, id string, render *sdl.Renderer) {
	var err error

	tmpSurface, err := img.Load(fileName)
	checkError(err)
	defer tmpSurface.Free()

	texture, err := render.CreateTextureFromSurface(tmpSurface)
	checkError(err)

	t.textureMap[id] = texture
}

// Draw ...
func (t *TextureManager) Draw(id string, x, y, width, height int32, ren *sdl.Renderer, flip sdl.RendererFlip) {
	// des Rect :: were on screen to draw
	// src Rect :: what part of the texture to render

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
	// des Rect :: were on screen to draw
	// src Rect :: what part of the texture to render

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

// ClearFromTextureMap ... remove texture from 'texture map' via its id if found
func (t *TextureManager) ClearFromTextureMap(id string) error {
	_, ok := t.textureMap[id]
	if !ok {
		return errors.New("texture id not found within texture map")
	}

	delete(t.textureMap, id)
	return nil
}
