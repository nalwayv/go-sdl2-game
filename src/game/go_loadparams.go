package game

// LoadParams ...
type LoadParams struct {
	x              int32
	y              int32
	width          int32
	height         int32
	id             string
	numframes      int
	callbackID     int
	animationSpeed int
}

// NewParams ...
func NewParams(x, y, w, h int32, id string, numframes, callbackID, animationSpeed int) *LoadParams {
	param := &LoadParams{}

	param.x = x
	param.y = y
	param.width = w
	param.height = h
	param.id = id
	param.numframes = numframes
	param.callbackID = callbackID
	param.animationSpeed = animationSpeed

	return param
}

// X ... get x
func (p LoadParams) X() int32 {
	return p.x
}

// Y ... get y
func (p LoadParams) Y() int32 {
	return p.y
}

// Width ... get width
func (p LoadParams) Width() int32 {
	return p.width
}

// Height ... get height
func (p LoadParams) Height() int32 {
	return p.height
}

// ID ... get id
func (p LoadParams) ID() string {
	return p.id
}

// NumFrames ... get number of frames
func (p LoadParams) NumFrames() int {
	return p.numframes
}

// CallBackID ... get callback ID
func (p LoadParams) CallBackID() int {
	return p.callbackID
}

// AnimationSpeed ... get animation speed
func (p LoadParams) AnimationSpeed() int {
	return p.animationSpeed
}
