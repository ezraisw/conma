package condition

import "errors"

var (
	ErrEmptyCond             = errors.New("empty condition")
	ErrInvalidInterval       = errors.New("invalid interval")
	ErrInvalidMaxDist        = errors.New("invalid max distance")
	ErrInvalidStartDist      = errors.New("invalid start distance")
	ErrInvalidMaxOrStartDist = errors.New("invalid max or start distance")
)
