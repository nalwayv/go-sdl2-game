package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Frame ...
type Frame struct {
	numcolumns     int
	numrows        int
	padding        int
	currentTick    int
	specifiedTicks int
	currentRow     int32
	currentCol     int32
	beginRow       int32 // current frame
	beginCol       int32
	endRow         int32 // current frame
	endCol         int32
}

// NewFrame ...
func NewFrame() *Frame {
	f := &Frame{}
	return f
}

// SetBeginColumn ...
func (f *Frame) SetBeginColumn(val int32) {
	f.beginCol = val
}

// SetBeginRow ...
func (f *Frame) SetBeginRow(val int32) {
	f.beginRow = val
}

// SetEndColumn ...
func (f *Frame) SetEndColumn(val int32) {
	f.endCol = val
}

// SetEndRow ...
func (f *Frame) SetEndRow(val int32) {
	f.endRow = val
}

// SetSpeed .. set animation speed
func (f *Frame) SetSpeed(ticks int) {
	f.specifiedTicks = ticks
}

// IncrementSpeed ... update animation speed
func (f *Frame) IncrementSpeed(delta int) {

	f.specifiedTicks += delta
}

// SetFrame ...
func (f *Frame) SetFrame(beginRow, beginCol, endRow, endCol int32, tickInterval, numframes int) {
	f.numcolumns = numframes

	f.specifiedTicks = tickInterval

	f.beginRow = beginRow
	f.beginCol = beginCol

	f.endRow = endRow
	f.endCol = endCol

	f.currentRow = f.beginRow
	f.currentCol = f.beginCol

	f.currentTick = 0
}

// UpdateFrame ...
func (f *Frame) UpdateFrame() {
	f.currentTick++
	if f.currentTick > f.specifiedTicks {

		f.currentTick = 0
		f.currentCol++
		if f.currentCol > f.endCol {

			f.currentCol = f.beginCol
			f.currentRow++
			if f.currentRow > f.endRow {
				f.currentRow = f.beginRow
			}
		}
	}
}

// DrawFrame ...
func (f *Frame) DrawFrame(id string, x, y, w, h int32, render *sdl.Renderer, flip sdl.RendererFlip) {
	STextureManager.DrawFrame(id, x, y, w, h, int32(f.currentRow), int32(f.currentCol), render, flip)
}
