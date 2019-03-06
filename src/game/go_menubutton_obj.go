package game

import "fmt"

/*
	*IGameObject*
	- Draw()
 	- Update()
 	- Clean()
*/

// enum mouse states
const (
	MouseOut  = iota // 0
	MouseOver        // 1
	Clicked          // 2
)

// MenuButton ...
type MenuButton struct {
	obj            *SdlGameObject
	buttonReleased bool
	callback       func()
}

// NewMenuButton ...
func NewMenuButton(params *LoadParams, callback func()) *MenuButton {
	mb := &MenuButton{}

	mb.obj = NewSdlGObj(params)

	mb.obj.CurrentFrame = MouseOut

	mb.callback = callback

	return mb
}

// Draw ...
func (mb *MenuButton) Draw() {
	mb.obj.Draw()
}

// Update ...
func (mb *MenuButton) Update() {
	// get current mouse pos
	mousePos := SInputHandler.GetMousePosition()

	mouseX := mousePos.GetX()
	mouseY := mousePos.GetY()
	posX := mb.obj.Position.GetX()
	posY := mb.obj.Position.GetY()
	width := float64(mb.obj.Width)
	height := float64(mb.obj.Height)

	// AABB collision check
	if mouseX < (posX+width) && mouseX > posX && mouseY < (posY+height) && mouseY > posY {

		mb.obj.CurrentFrame = MouseOver

		// change state on mouse click / mouse over
		if SInputHandler.GetMouseButtonState(MouseLeft) && mb.buttonReleased {
			fmt.Println("button clicked")

			mb.obj.CurrentFrame = Clicked
			mb.callback()
			mb.buttonReleased = false

		} else if !SInputHandler.GetMouseButtonState(MouseLeft) {

			mb.buttonReleased = true
			mb.obj.CurrentFrame = MouseOver

		}
	} else {
		mb.obj.CurrentFrame = MouseOut
	}

}

// Clean ...
func (mb *MenuButton) Clean() {
	mb.obj.Clean()
}
