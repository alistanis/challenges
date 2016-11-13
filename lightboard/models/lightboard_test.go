package models

import (
	"errors"
	"testing"

	"github.com/alistanis/challenges/util"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSwitchController(t *testing.T) {
	Convey("We can test a switch controller's functions", t, func() {
		s := NewLightController(8)
		// start a 0
		So(s.NumLights(), ShouldEqual, 256)

		So(s.LightState, ShouldEqual, 0)
		s.Toggle(0)
		So(s.LightState, ShouldEqual, 1)
		// reset to original state
		s.Toggle(0)
		So(s.LightState, ShouldEqual, 0)
		s.Toggle(1)
		So(s.LightState, ShouldEqual, 2)
		s.Toggle(0)
		So(s.LightState, ShouldEqual, 3)

		s.Toggle(5)
		So(s.LightState, ShouldEqual, 35)
		s.Reset()
		So(s.LightState, ShouldEqual, 0)

		s.Reset()

		Convey("We can test that every permutation turns on exactly one light", func() {
			m := uint(0)
			mask := &m
			checkPermutations(s, s.NumToggles(), mask)
			So(*mask, ShouldEqual, util.Pow(2, s.NumToggles())-1)
		})
	})
}

// toggles every switch to ensure that the light controller works properly
func checkPermutations(t Toggler, i int, mask *uint) error {

	for j := 0; j < i; j++ {
		if t.Check() == -1 {
			errors.New("No light is on")
		}
		*mask |= uint(t.Check())
		t.Toggle(j)
		err := checkPermutations(t, i-1, mask)
		if err != nil {
			return err
		}
	}
	return nil
}
