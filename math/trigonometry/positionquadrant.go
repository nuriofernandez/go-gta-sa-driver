package trigonometry

type PositionQuadrant byte

const (
	CENTER      PositionQuadrant = 0 // No quadrant, on an axis line.
	PLUS_PLUS                    = 1 // Top right
	PLUS_MINUS                   = 2 // Down right
	MINUS_MINUS                  = 3 // Down left
	MINUS_PLUS                   = 4 // Top left
)
