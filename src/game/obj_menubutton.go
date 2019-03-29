package game

/*
IGameObject interface
---

- Draw()
- Update()
- Clean()
- Load(*params)
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
	buttonReleased bool   // pressed or not
	callback       func() // function execusted when button is clicked
	callbackID     int    // num indicating what function to call within data slice
}

// NewMenuButton ...
func NewMenuButton() *MenuButton {
	mb := &MenuButton{}
	mb.obj = NewSdlGObj()
	return mb
}

// Load ...
func (mb *MenuButton) Load(params *LoadParams) {
	mb.obj.Load(params)

	mb.callbackID = params.GetCallBackID()

	mb.obj.CurrentFrame = MouseOut
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
			mb.obj.CurrentFrame = Clicked

			mb.callback() // run callback function

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

// SetCallBack ... function called when clicked
func (mb *MenuButton) SetCallBack(callback Callback) {
	mb.callback = callback
}

// GetCallBackID ... id of the function to call
func (mb MenuButton) GetCallBackID() int {
	return mb.callbackID
}
