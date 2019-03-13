package game

// LoadParams ...
type LoadParams struct {
	x              int32
	y              int32
	width          int32
	height         int32
	textureID      string
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
	param.textureID = id
	param.numframes = numframes
	param.callbackID = callbackID
	param.animationSpeed = animationSpeed

	return param
}

// X ... get x
func (p LoadParams) GetX() int32 {
	return p.x
}

// Y ... get y
func (p LoadParams) GetY() int32 {
	return p.y
}

// Width ... get width
func (p LoadParams) GetWidth() int32 {
	return p.width
}

// Height ... get height
func (p LoadParams) GetHeight() int32 {
	return p.height
}

// ID ... get id
func (p LoadParams) GetTextureID() string {
	return p.textureID
}

// NumFrames ... get number of frames
func (p LoadParams) GetNumFrames() int {
	return p.numframes
}

// CallBackID ... get callback ID
func (p LoadParams) GetCallBackID() int {
	return p.callbackID
}

// AnimationSpeed ... get animation speed
func (p LoadParams) GetAnimationSpeed() int {
	return p.animationSpeed
}
