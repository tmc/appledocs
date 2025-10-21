// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Floor] class.
var (
	FloorClass     _FloorClass
	FloorClassOnce sync.Once
)

func getFloorClass() _FloorClass {
	FloorClassOnce.Do(func() {
		FloorClass = _FloorClass{objc.GetClass("CLFloor")}
	})
	return FloorClass
}

type _FloorClass struct {
	class objc.Class
}

// An interface definition for the [Floor] class.
type IFloor interface {
	objectivec.IObject
}

// The floor of a building on which the user’s device is located.
//
// A object specifies the floor of the building on which the device is located. In places where floor information can be determined, a object may include a floor object along with the regular location data. You do not create instances of this class directly, nor should you subclass it.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLFloor
type Floor struct {
	objectivec.Object
}

// FloorFrom constructs a [Floor] from an unsafe.Pointer.
//
// The floor of a building on which the user’s device is located.
func FloorFrom(ptr unsafe.Pointer) Floor {
	return Floor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FloorClass) Alloc() Floor {
	rv := objc.Send[Floor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FloorClass) New() Floor {
	rv := objc.Send[Floor](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ Floor) Init() Floor {
	rv := objc.Send[Floor](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ Floor) Autorelease() Floor {
	rv := objc.Send[Floor](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFloor creates a new Floor instance.
func NewFloor() Floor {
	return getFloorClass().New()
}


// The logical floor of the building.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLFloor/level
func (f_ Floor) Level() int {
	rv := objc.Send[int](f_.ID, objc.Sel("level"))
	return rv
}


// SetLevel sets the value of the level property.
// The logical floor of the building.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLFloor/level
func (f_ Floor) SetLevel(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLevel:"), value)
}


