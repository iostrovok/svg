package style

import (
	"fmt"
	"math"
)

const (
	lineWidth    = "%d.%d"
	lineWidthInt = "%d"
)

type WIDTH struct {
	ONESTYLE
	value float64
}

// Constructor
func Width(value float64) WIDTH {
	return WIDTH{
		value: value,
	}
}

// Source() returns svg implementation of WIDTH style
func (width WIDTH) Source() string {
	if math.Floor(width.value) == width.value {
		return fmt.Sprintf(lineWidthInt, int(width.value))
	}

	return fmt.Sprintf(lineWidth, int(math.Floor(width.value)), int((width.value-math.Floor(width.value))*100))
}
