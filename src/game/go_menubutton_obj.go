package game

import "fmt"

/*
*IGameObject*

- Draw()
- Update()
- Clean()
- Load()
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
	callbackID     int
}

// NewMenuButton ...
func NewMenuButton(params *LoadParams, callback func()) *MenuButton {
	mb := &MenuButton{}

	//mb.obj = NewSdlGObj(params)

	mb.obj.CurrentFrame = MouseOut

	mb.callback = callback

	return mb
}

// Load ...
func (mb *MenuButton) Load(params *LoadParams) {
	mb.obj.Load(params)

	mb.obj.CurrentFrame = MouseOut

	mb.callbackID = params.CallBackID()
}

// Draw ...
func (mb *MenuButton) Draw() {
	mb.obj.Draw()
}

// Update ...
func (mb *MenuButton) Update() {
	// get current mouse pos
	mousePos := SInputHandler.GetMousePosition()

	// AABB collision check
	if mousePos.GetX() < mb.obj.Position.GetX()+float64(mb.obj.Width) &&
		mousePos.GetX() > mb.obj.Position.GetX() &&
		mousePos.GetY() < mb.obj.Position.GetY()+float64(mb.obj.Height) &&
		mousePos.GetY() > mb.obj.Position.GetY() {

		mb.obj.CurrentFrame = MouseOver

		// change state on mouse click / mouse over
		if SInputHandler.GetMouseButtonState(MouseLeft) && mb.buttonReleased {
			fmt.Println("button clicked")

			mb.obj.CurrentFrame = Clicked

			mb.callback() // run callback function

			mb.buttonReleased = false

		} else if !SInputHandler.GetMouseButtonState(MouseLeft) {
			fmt.Println("button over")
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

// SetCallBackID ...
func (mb *MenuButton) SetCallBackID(val int) {
	mb.callbackID = val
}

// GetCallBackID ...
func (mb MenuButton) GetCallBackID() int {
	return mb.callbackID
}
