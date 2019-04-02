package game

/*
info
---
Singleton texture manager that uses a map to store sld.Texture values

**/

import (
	"errors"
	"sync"

	"../gologger"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	tm     *TextureManager
	tmOnce sync.Once
)

// STextureManager ...
var STextureManager = newTManager()

// TextureManager ...
type TextureManager struct {
	textureMap map[string]*sdl.Texture
}

// newTManager ... convert into a singleton
func newTManager() *TextureManager {
	gologger.SLogger.Println("Init New Texture Manager")

	tmOnce.Do(func() {
		tm = &TextureManager{
			textureMap: make(map[string]*sdl.Texture),
		}
	})
	return tm
}

// Load ... load a texture and store it within textureMap
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
func (t *TextureManager) Draw(id string, x, y, width, height int32, render *sdl.Renderer, flip sdl.RendererFlip) {
	// des Rect :: were on screen to draw
	// src Rect :: what part of the texture to render

	var err error

	desRect := sdl.Rect{x, y, width, height}

	srcRect := sdl.Rect{0, 0, desRect.W, desRect.H}

	err = render.CopyEx(t.textureMap[id], &srcRect, &desRect, 0, nil, flip)

	checkError(err)
}

// DrawFrame ...
func (t *TextureManager) DrawFrame(id string, x, y, width, height, currentRow, currentFrame int32, render *sdl.Renderer, flip sdl.RendererFlip) {
	// des Rect :: were on screen to draw
	// src Rect :: what part of the texture to render

	var err error

	desRect := sdl.Rect{x, y, width, height}

	srcRect := sdl.Rect{
		X: width * currentFrame,
		Y: height * (currentRow - 1),
		W: desRect.W,
		H: desRect.H,
	}

	err = render.CopyEx(t.textureMap[id], &srcRect, &desRect, 0, nil, flip)

	checkError(err)
}

// DrawTile ... draw tile with included margin and spacing values
func (t *TextureManager) DrawTile(id string, margin, spacing, x, y, width, height, currentRow, currentFrame int32, render *sdl.Renderer) {
	// des Rect :: were on screen to draw
	// src Rect :: what part of the texture to render

	var err error

	desRect := sdl.Rect{x, y, width, height}

	srcRect := sdl.Rect{
		X: margin + (spacing+width)*currentFrame,
		Y: margin + (spacing+height)*currentRow,
		W: desRect.W,
		H: desRect.H,
	}

	err = render.CopyEx(t.textureMap[id], &srcRect, &desRect, 0, nil, sdl.FLIP_NONE)

	checkError(err)
}

// ClearFromTextureMap ... remove texture from 'texture map' by its id if found
func (t *TextureManager) ClearFromTextureMap(id string) error {
	_, ok := t.textureMap[id]
	if ok {
		delete(t.textureMap, id)
	} else {
		return errors.New("texture id not found in texture map")
	}
	return nil
}
