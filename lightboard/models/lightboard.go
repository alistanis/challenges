package models

import "github.com/alistanis/challenges/util"

type LightController struct {
	LightState  uint
	numSwitches int
}

func NewLightController(numSwitches int) *LightController {
	return &LightController{0, numSwitches}
}

func (s *LightController) Toggle(i int) {
	s.LightState ^= 1 << uint(i)
}

func (s *LightController) NumToggles() int {
	return s.numSwitches
}

// because we're using a bit set we know that there will always be a light on and this will never return -1,
// but the caller doesn't know that
func (s *LightController) Check() int {
	return int(s.LightState)
}

func (s *LightController) Reset() {
	s.LightState = 0
}

func (s *LightController) NumLights() int {
	return util.Pow(2, s.numSwitches)
}

type Toggler interface {
	Toggle(int)
	Check() int
	NumToggles() int
}
