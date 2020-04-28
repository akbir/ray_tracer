package primatives

var verticalUp = Vector{X: 0, Y: 1, Z: 0}

type Camera struct {
	lowerLeft, horizontal, vertical, origin Vector
}

func NewCamera() *Camera {
	c := new(Camera)
	c.lowerLeft = Vector{-2.0, -1.0, -1.0}
	c.horizontal = Vector{4.0, 0.0, 0.0}
	c.vertical = Vector{0.0, 2.0, 0.0}
	c.origin = Vector{0.0, 0.0, 0.0}
	return c
}

func (c *Camera) RayAt(u, v float64) Ray {
	position := c.Position(u, v)
	direction := c.Direction(position)
	return Ray{Origin: c.origin, Direction: direction}
}

func (c *Camera) Position(u, v float64) Vector {
	horizontal := c.horizontal.MultiplyScalar(u)
	vertical := c.vertical.MultiplyScalar(v)
	return horizontal.Add(vertical)
}

func (c *Camera) Direction(position Vector) Vector {
	// still unsure of why this is the case
	//direction = lowerLeft + position
	return c.lowerLeft.Add(position)
}
