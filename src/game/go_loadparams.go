package game

// LoadParams ...
type LoadParams struct {
	x      int32
	y      int32
	width  int32
	height int32
	id     string
}

// NewParams ...
func NewParams(x, y, w, h int32, id string) *LoadParams {
	param := &LoadParams{}

	param.x = x
	param.y = y

	param.width = w
	param.height = h

	param.id = id

	return param
}

// Getters --

// X ...
func (p LoadParams) X() int32 {
	return p.x
}

// Y ...
func (p LoadParams) Y() int32 {
	return p.y
}

// Width ...
func (p LoadParams) Width() int32 {
	return p.width
}

// Height ...
func (p LoadParams) Height() int32 {
	return p.height
}

// ID ...
func (p LoadParams) ID() string {
	return p.id
}
