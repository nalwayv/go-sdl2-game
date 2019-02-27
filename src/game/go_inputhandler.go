package game

import (
	"fmt"
	"sync"

	"../gologger"
	"../vec2d"
	"github.com/veandco/go-sdl2/sdl"
)

// DeadZone  ... stick dead zone
const DeadZone int16 = 8000

// Mouse BUttons Enum
const (
	MouseLeft = iota
	MouseMiddle
	MouseRight
)

// sticks ... analog sticks
type sticks struct {
	first  *vec2d.Vector2D
	second *vec2d.Vector2D
}

// InHandler ... input handler
type InHandler struct {
	isJSInitialised bool
	inSticks        []*sdl.Joystick
	inStickVal      []*sticks
	inButtons       [][]bool
	inMouseButtons  []bool
	inMousePos      *vec2d.Vector2D
}

// Singelton ... turn inputHandler into a singleton
var (
	ih    *InHandler
	ionce sync.Once
)

// SInputHandler ... singleton init and means of communicating with input functions
var SInputHandler = newInputHandler()

func newInputHandler() *InHandler {
	ionce.Do(func() {
		ih = &InHandler{}
		ih.inSticks = make([]*sdl.Joystick, 0)

		ih.inStickVal = make([]*sticks, 0)

		ih.inButtons = make([][]bool, 0)

		ih.inMouseButtons = make([]bool, 3, 3)

		ih.inMousePos = vec2d.NewVector2d(0, 0)
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

			if v.GetType() == sdl.MOUSEBUTTONDOWN {
				if v.Button == sdl.BUTTON_LEFT {
					fmt.Println("Mouse Left Down")
					input.inMouseButtons[MouseLeft] = true
				}

				if v.Button == sdl.BUTTON_MIDDLE {
					fmt.Println("Mouse Middle Down")
					input.inMouseButtons[MouseMiddle] = true
				}

				if v.Button == sdl.BUTTON_RIGHT {
					fmt.Println("Mouse Right Down")
					input.inMouseButtons[MouseRight] = true
				}
			}

			if v.GetType() == sdl.MOUSEBUTTONUP {
				if v.Button == sdl.BUTTON_LEFT {
					fmt.Println("Mouse Left Up")
					input.inMouseButtons[MouseLeft] = false
				}

				if v.Button == sdl.BUTTON_MIDDLE {
					fmt.Println("Mouse Middle Up")
					input.inMouseButtons[MouseMiddle] = false
				}

				if v.Button == sdl.BUTTON_RIGHT {
					fmt.Println("Mouse Right Up")
					input.inMouseButtons[MouseRight] = false
				}
			}

		// mouse move
		case *sdl.MouseMotionEvent:
			if v.GetType() == sdl.MOUSEMOTION {
				input.inMousePos.SetX(float64(v.X))
				input.inMousePos.SetY(float64(v.Y))
			}

		// joypad button
		case *sdl.JoyButtonEvent:
			// activate
			if v.GetType() == sdl.JOYBUTTONDOWN {
				whichOne := v.Which
				input.inButtons[whichOne][v.Button] = true
			}

			// de-activate
			if v.GetType() == sdl.JOYBUTTONUP {
				whichOne := v.Which
				input.inButtons[whichOne][v.Button] = false
			}

		// joypad sticks
		case *sdl.JoyAxisEvent:
			whichOne := v.Which

			// left analog stick moved --
			// right
			// left
			// down
			// up
			if v.Axis == 0 {
				if v.Value > DeadZone {
					fmt.Println("left stick 'X' right")
					input.inStickVal[whichOne].first.SetX(1)

				} else if v.Value < -DeadZone {
					fmt.Println("left stick 'X' left")
					input.inStickVal[whichOne].first.SetX(-1)

				} else {
					fmt.Println("left stick 'X' default")
					input.inStickVal[whichOne].first.SetX(0)
				}
			}
			if v.Axis == 1 {
				if v.Value > DeadZone {
					fmt.Println("left stick 'Y' down")
					input.inStickVal[whichOne].first.SetY(1)

				} else if v.Value < -DeadZone {
					fmt.Println("left stick 'Y' up")
					input.inStickVal[whichOne].first.SetY(-1)

				} else {
					fmt.Println("left stick 'Y' default")
					input.inStickVal[whichOne].first.SetY(0)
				}
			}

			// right analog stick moved --
			// right
			// left
			// down
			// up
			if v.Axis == 3 {
				if v.Value > DeadZone {
					fmt.Println("right stick 'X' right")
					input.inStickVal[whichOne].second.SetX(1)

				} else if v.Value < -DeadZone {
					fmt.Println("right stick 'X' left")
					input.inStickVal[whichOne].second.SetX(-1)

				} else {
					fmt.Println("right stick 'X' default")
					input.inStickVal[whichOne].second.SetX(0)
				}
			}
			if v.Axis == 4 {
				if v.Value > DeadZone {
					fmt.Println("right stick 'Y' down")
					input.inStickVal[whichOne].second.SetY(1)

				} else if v.Value < -DeadZone {
					fmt.Println("right stick 'Y' up")
					input.inStickVal[whichOne].second.SetY(-1)

				} else {
					fmt.Println("right stick 'Y' default")
					input.inStickVal[whichOne].second.SetY(0)
				}
			}
		}
	}
}

// Clean ... close all joysticks within the slice *inSticks
func (input *InHandler) Clean() {
	for _, v := range input.inSticks {
		v.Close()
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

	if sdl.NumJoysticks() > 0 {

		// add joypads found to inSticks
		for i := 0; i < sdl.NumJoysticks(); i++ {
			joy := sdl.JoystickOpen(i)

			if joy != nil {

				input.inSticks = append(input.inSticks, joy)

				input.inStickVal = append(input.inStickVal,
					&sticks{
						vec2d.NewVector2d(0, 0),
						vec2d.NewVector2d(0, 0)},
				)

				// add number of button found on joypad to inButtons
				tmpButton := make([]bool, 0)

				for j := 0; j < joy.NumButtons(); j++ {
					tmpButton = append(tmpButton, false)
				}

				input.inButtons = append(input.inButtons, tmpButton)

				// log added buttons
				for _, bval := range input.inButtons {
					gologger.SLogger.Println("added to inButtons", bval)
				}
			}
		}

		// listen for events
		sdl.JoystickEventState(sdl.ENABLE)

		// joy stick has been initialised
		input.isJSInitialised = true

		// log how many
		gologger.SLogger.Println("initialised", len(input.inSticks), "joysticks")

	} else {
		gologger.SLogger.Println("no joysticks were initialised")

		input.isJSInitialised = false
	}

}

// JoySticksInitialised ... return has joy stick been initialised
func (input InHandler) JoySticksInitialised() bool {
	return input.isJSInitialised
}

// Xvalue ... get current X value
// joy   - id of joypad for example 0 == joypad 1
// stick - current stick on controller most have 2 left and a right
func (input InHandler) Xvalue(joy, stick int) int {
	if len(input.inStickVal) > 0 {
		if stick == 1 {
			v := input.inStickVal[joy].first.GetX()
			return int(v)

		} else if stick == 2 {
			v := input.inStickVal[joy].second.GetX()
			return int(v)
		}
	}
	return 0
}

// Yvalue ... get current Y value
// joy   - id of joypad for example 0 == joypad 1
// stick - current stick on controller most have 2 left and a right
func (input InHandler) Yvalue(joy, stick int) int {
	if len(input.inStickVal) > 0 {
		if stick == 1 {
			v := input.inStickVal[joy].first.GetY()
			return int(v)

		} else if stick == 2 {
			v := input.inStickVal[joy].second.GetY()
			return int(v)
		}
	}
	return 0
}

// GetButtonState ... return if current button state is active or not
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
func (input InHandler) GetButtonState(joy, buttonNum int) bool {
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
func (input InHandler) GetMouseButtonState(button int) bool {
	return input.inMouseButtons[button]
}

// GetMousePosition ... return mouse (X, Y) position
func (input InHandler) GetMousePosition() *vec2d.Vector2D {
	return input.inMousePos
}
