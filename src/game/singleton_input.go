package game

/*
info
---
Singleton input handler
**/

import (
	"sync"

	"../gologger"
	"../vec2d"
	"github.com/veandco/go-sdl2/sdl"
)

// DeadZone  ... stick dead zone
const DeadZone int16 = 8000

// Mouse Buttons Enum
const (
	MouseLeft = iota
	MouseMiddle
	MouseRight
)

// sticks ... analogue sticks
type sticks struct {
	first  *vec2d.Vector2D
	second *vec2d.Vector2D
}

// Singelton ... turn inputHandler into a singleton
var (
	ih     *InHandler
	ihOnce sync.Once
)

// SInputHandler ... singleton
var SInputHandler = newInputHandler()

// InHandler ... input handler
type InHandler struct {
	isJSInitialised bool
	inSticks        []*sdl.Joystick
	inStickVal      []*sticks
	inButtons       [][]bool
	inMouseButtons  []bool
	inMousePos      *vec2d.Vector2D
	inKeyState      []uint8
}

// create new InputHandler ...
func newInputHandler() *InHandler {
	gologger.SLogger.Println("Init New Input Handler")

	ihOnce.Do(func() {
		ih = &InHandler{}
		ih.inSticks = make([]*sdl.Joystick, 0)

		ih.inStickVal = make([]*sticks, 0)

		ih.inButtons = make([][]bool, 0)

		ih.inMouseButtons = make([]bool, 3, 3)

		ih.inMousePos = vec2d.NewVector2d(0, 0)

		ih.inKeyState = make([]uint8, 0)
	})

	return ih
}

// Update ... update game inputs
func (input *InHandler) Update() {
	var event sdl.Event

	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		// catch events
		switch v := event.(type) {

		// quit app
		case *sdl.QuitEvent:
			STheGame.Quit()

		// mouse buttons
		case *sdl.MouseButtonEvent:
			if v.GetType() == sdl.MOUSEBUTTONUP {
				input.OnMouseButtonUp(v)
			}
			if v.GetType() == sdl.MOUSEBUTTONDOWN {
				input.OnMouseButtonDown(v)
			}

		// mouse move
		case *sdl.MouseMotionEvent:
			if v.GetType() == sdl.MOUSEMOTION {
				input.OnMouseMove(v)
			}

		// joypad button
		case *sdl.JoyButtonEvent:
			if v.GetType() == sdl.JOYBUTTONUP {
				input.OnJoyButtonUp(v)
			}
			if v.GetType() == sdl.JOYBUTTONDOWN {
				input.OnJoyButtonDown(v)
			}

		// joypad sticks
		case *sdl.JoyAxisEvent:
			input.OnJoyAxisMove(v)

		// keyboard
		case *sdl.KeyboardEvent:
			if v.State == sdl.PRESSED {
				input.OnKeyDown()
			}

			if v.State == sdl.RELEASED {
				input.OnKeyUp()
			}
		}
	}
}

// Clean ... close all joysticks within the slice *inSticks
func (input *InHandler) Clean() {
	if input.JoySticksInitialised() {
		for _, v := range input.inSticks {
			v.Close()
		}
	}
}

// InitialiseJoySticks ... initialised a joy stick if detected and add it to slice []*inSticks
func (input *InHandler) InitialiseJoySticks() {
	var err error

	// check if joystick subsystem has been initialised
	if sdl.WasInit(sdl.INIT_JOYSTICK) == 0 {
		err = sdl.InitSubSystem(sdl.INIT_JOYSTICK)
		checkError(err)
	}

	// if any  joysticks attached to system
	if sdl.NumJoysticks() > 0 {

		// add found to inSticks
		for joyNumber := 0; joyNumber < sdl.NumJoysticks(); joyNumber++ {

			joy := sdl.JoystickOpen(joyNumber)

			if joy != nil {

				input.inSticks = append(input.inSticks, joy)
				input.inStickVal = append(input.inStickVal, &sticks{vec2d.NewVector2d(0, 0), vec2d.NewVector2d(0, 0)})

				// add number of button found on joypad
				tmpButton := make([]bool, 0)

				for j := 0; j < joy.NumButtons(); j++ {
					tmpButton = append(tmpButton, false)
				}

				input.inButtons = append(input.inButtons, tmpButton)
			}
		}

		// listen for events
		sdl.JoystickEventState(sdl.ENABLE)

		// joy stick has been initialised
		input.isJSInitialised = true

		// log how many added
		gologger.SLogger.Println("Activated", len(input.inSticks), "joysticks")

	} else {
		gologger.SLogger.Println("No joysticks were activated")

		input.isJSInitialised = false
	}

}

// JoySticksInitialised ... return has joy stick been initialised
func (input *InHandler) JoySticksInitialised() bool {
	return input.isJSInitialised
}

// GetXvalue ... get current X value
// joy :: id of joypad
// - 0 joy stick 1
// - 1 joy stick 2
// stick - sticks on the controller, most have two a left and a right
// - 0 left
// - 1 right
func (input *InHandler) GetXvalue(joy, stick int) int {
	// if any sticks found
	if len(input.inStickVal) > 0 {

		if stick == 1 {
			// left
			v := input.inStickVal[joy].first.GetX()
			return int(v)
		} else if stick == 2 {
			// right
			v := input.inStickVal[joy].second.GetX()
			return int(v)
		}
	}

	// nothing
	return 0
}

// GetYvalue ... get current Y value of stick moved on joypad
// joy :: id of joypad
// - 0 joy stick 1
// - 1 joy stick 2
// stick - sticks on the controller, most have two a left and a right
// - 0 left
// - 1 right
func (input *InHandler) GetYvalue(joy, stick int) int {
	// if any sticks found
	if len(input.inStickVal) > 0 {
		if stick == 1 {
			// left
			v := input.inStickVal[joy].first.GetY()
			return int(v)

		} else if stick == 2 {
			// right
			v := input.inStickVal[joy].second.GetY()
			return int(v)
		}
	}

	// nothing
	return 0
}

// GetButtonState ... return if current button state is active or not
// joy - joypad to act upon
// buttonNum - button on joypad to act upon
//
//   * Xbox Buttons
//   -----------------
//  | Button | Number |
//  | ------ + ------ |
//	| A      |      0 |
//	| B      |      1 |
//	| X      |      2 |
//	| Y      |      3 |
//	| L      |      4 |
//	| R      |      5 |
//    -----------------
func (input *InHandler) GetButtonState(joy, buttonNum int) bool {

	return input.inButtons[joy][buttonNum]
}

// GetMouseButtonState ... returns current state of mouse button click
//   * Mouse Buttons
//   -----------------
//  | Button | Number |
//  | ------ + ------ |
//	| Left   |      0 |
//	| Middle |      1 |
//	| Right  |      3 |
//   -----------------
func (input *InHandler) GetMouseButtonState(button int) bool {
	return input.inMouseButtons[button]
}

// GetMousePosition ... return mouse (X, Y) position
func (input *InHandler) GetMousePosition() *vec2d.Vector2D {
	return input.inMousePos
}

// --- KEYBOARD EVENTS

// OnKeyUp ...
func (input *InHandler) OnKeyUp() {
	// get current keys that are active or not
	input.inKeyState = sdl.GetKeyboardState()
}

// OnKeyDown ...
func (input *InHandler) OnKeyDown() {
	// get current keys that are active or not
	input.inKeyState = sdl.GetKeyboardState()
}

// IsKeyDown ...
func (input *InHandler) IsKeyDown(key sdl.Scancode) bool {
	if len(input.inKeyState) == 0 {
		return false
	}

	// is key active
	if input.inKeyState[key] == 1 {
		return true
	}

	return false
}

// --- MOUSE EVENTS

// OnMouseMove ...
func (input *InHandler) OnMouseMove(v *sdl.MouseMotionEvent) {
	input.inMousePos.SetX(float64(v.X))
	input.inMousePos.SetY(float64(v.Y))
}

// OnMouseButtonUp ...
func (input *InHandler) OnMouseButtonUp(v *sdl.MouseButtonEvent) {
	if v.Button == sdl.BUTTON_LEFT {
		input.inMouseButtons[MouseLeft] = false
	}

	if v.Button == sdl.BUTTON_MIDDLE {
		input.inMouseButtons[MouseMiddle] = false
	}

	if v.Button == sdl.BUTTON_RIGHT {
		input.inMouseButtons[MouseRight] = false
	}
}

// OnMouseButtonDown ...
func (input *InHandler) OnMouseButtonDown(v *sdl.MouseButtonEvent) {
	if v.Button == sdl.BUTTON_LEFT {
		input.inMouseButtons[MouseLeft] = true
	}

	if v.Button == sdl.BUTTON_MIDDLE {
		input.inMouseButtons[MouseMiddle] = true
	}

	if v.Button == sdl.BUTTON_RIGHT {
		input.inMouseButtons[MouseRight] = true
	}
}

// Reset ... reset mouse states to false
func (input *InHandler) Reset() {
	input.inMouseButtons[MouseLeft] = false
	input.inMouseButtons[MouseMiddle] = false
	input.inMouseButtons[MouseRight] = false
}

// --- JOYSTICK EVENTS

// OnJoyAxisMove ...
func (input *InHandler) OnJoyAxisMove(v *sdl.JoyAxisEvent) {
	whichOne := v.Which

	// left analogue stick moved --
	if v.Axis == 0 {
		if v.Value > DeadZone {
			input.inStickVal[whichOne].first.SetX(1)

		} else if v.Value < -DeadZone {
			input.inStickVal[whichOne].first.SetX(-1)

		} else {
			input.inStickVal[whichOne].first.SetX(0)
		}
	}
	if v.Axis == 1 {
		if v.Value > DeadZone {
			input.inStickVal[whichOne].first.SetY(1)

		} else if v.Value < -DeadZone {
			input.inStickVal[whichOne].first.SetY(-1)

		} else {
			input.inStickVal[whichOne].first.SetY(0)
		}
	}

	// right analogue stick moved --
	if v.Axis == 3 {
		if v.Value > DeadZone {
			input.inStickVal[whichOne].second.SetX(1)

		} else if v.Value < -DeadZone {
			input.inStickVal[whichOne].second.SetX(-1)

		} else {
			input.inStickVal[whichOne].second.SetX(0)
		}
	}
	if v.Axis == 4 {
		if v.Value > DeadZone {
			input.inStickVal[whichOne].second.SetY(1)

		} else if v.Value < -DeadZone {
			input.inStickVal[whichOne].second.SetY(-1)

		} else {
			input.inStickVal[whichOne].second.SetY(0)
		}
	}
}

// OnJoyButtonUp ...
func (input *InHandler) OnJoyButtonUp(v *sdl.JoyButtonEvent) {
	input.inButtons[v.Which][v.Button] = false
}

// OnJoyButtonDown ...
func (input *InHandler) OnJoyButtonDown(v *sdl.JoyButtonEvent) {
	input.inButtons[v.Which][v.Button] = true
}
