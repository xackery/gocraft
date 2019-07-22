package camera

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/icexin/gocraft/model"
)

type CameraMovement int

const (
	MoveForward CameraMovement = iota
	MoveBackward
	MoveLeft
	MoveRight
)

type Camera struct {
	pos    mgl32.Vec3
	up     mgl32.Vec3
	right  mgl32.Vec3
	front  mgl32.Vec3
	wfront mgl32.Vec3

	rotatex, rotatey float32

	Sens float32

	isFlying bool
}

// New creates a new camera
func New(pos mgl32.Vec3) *Camera {
	c := &Camera{
		pos:      pos,
		front:    mgl32.Vec3{0, 0, -1},
		rotatey:  0,
		rotatex:  -90,
		Sens:     0.14,
		isFlying: false,
	}
	c.updateAngles()
	return c
}

// Restore restores a camera's position
func (c *Camera) Restore(state model.PlayerState) {
	c.pos = mgl32.Vec3{state.X, state.Y, state.Z}
	c.rotatex = state.Rx
	c.rotatey = state.Ry
	c.updateAngles()
}

// State returns the current player state
func (c *Camera) State() model.PlayerState {
	return model.PlayerState{
		X:  c.pos.X(),
		Y:  c.pos.Y(),
		Z:  c.pos.Z(),
		Rx: c.rotatex,
		Ry: c.rotatey,
	}
}

// Matrix returns a camera's matrix
func (c *Camera) Matrix() mgl32.Mat4 {
	return mgl32.LookAtV(c.pos, c.pos.Add(c.front), c.up)
}

// SetPos will set a camera's position
func (c *Camera) SetPos(pos mgl32.Vec3) {
	c.pos = pos
}

// Pos returns a camera's position
func (c *Camera) Pos() mgl32.Vec3 {
	return c.pos
}

// Front returns the front vector of a camera
func (c *Camera) Front() mgl32.Vec3 {
	return c.front
}

// FlipFlying toggles the flying option
func (c *Camera) FlipFlying() {
	c.isFlying = !c.isFlying
}

// IsFlying returns if a camera is flying or not
func (c *Camera) IsFlying() bool {
	return c.isFlying
}

func (c *Camera) OnAngleChange(dx, dy float32) {
	if mgl32.Abs(dx) > 200 || mgl32.Abs(dy) > 200 {
		return
	}
	c.rotatex += dx * c.Sens
	c.rotatey += dy * c.Sens
	if c.rotatey > 89 {
		c.rotatey = 89
	}
	if c.rotatey < -89 {
		c.rotatey = -89
	}
	c.updateAngles()
}

func (c *Camera) OnMoveChange(dir CameraMovement, delta float32) {
	if c.isFlying {
		delta = 5 * delta
	}
	switch dir {
	case MoveForward:
		if c.isFlying {
			c.pos = c.pos.Add(c.front.Mul(delta))
		} else {
			c.pos = c.pos.Add(c.wfront.Mul(delta))
		}
	case MoveBackward:
		if c.isFlying {
			c.pos = c.pos.Sub(c.front.Mul(delta))
		} else {
			c.pos = c.pos.Sub(c.wfront.Mul(delta))
		}
	case MoveLeft:
		c.pos = c.pos.Sub(c.right.Mul(delta))
	case MoveRight:
		c.pos = c.pos.Add(c.right.Mul(delta))
	}
}
func (c *Camera) updateAngles() {
	front := mgl32.Vec3{
		model.Cos(model.Radian(c.rotatey)) * model.Cos(model.Radian(c.rotatex)),
		model.Sin(model.Radian(c.rotatey)),
		model.Cos(model.Radian(c.rotatey)) * model.Sin(model.Radian(c.rotatex)),
	}
	c.front = front.Normalize()
	c.right = c.front.Cross(mgl32.Vec3{0, 1, 0}).Normalize()
	c.up = c.right.Cross(c.front).Normalize()
	c.wfront = mgl32.Vec3{0, 1, 0}.Cross(c.right).Normalize()
}
