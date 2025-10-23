// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCRacingWheel] class.
var (
	GCRacingWheelClass     _GCRacingWheelClass
	GCRacingWheelClassOnce sync.Once
)

func getGCRacingWheelClass() _GCRacingWheelClass {
	GCRacingWheelClassOnce.Do(func() {
		GCRacingWheelClass = _GCRacingWheelClass{objc.GetClass("GCRacingWheel")}
	})
	return GCRacingWheelClass
}

type _GCRacingWheelClass struct {
	class objc.Class
}

// An interface definition for the [GCRacingWheel] class.
type IGCRacingWheel interface {
	objectivec.IObject
	Capture() GCRacingWheel
	RelinquishDevice()
	Acquired() bool
	Snapshot() bool
	IsAcquired() bool
	SetIsAcquired(value bool)
	IsSnapshot() bool
	SetIsSnapshot(value bool)
	WheelInput() unsafe.Pointer
	SetWheelInput(value unsafe.Pointer)
}

// An object that represents a physical racing wheel controller connected to a device.


// An object that represents a physical racing wheel controller connected to a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel
type GCRacingWheel struct {
	objectivec.Object
}

// GCRacingWheelFrom constructs a [GCRacingWheel] from an unsafe.Pointer.
//
// An object that represents a physical racing wheel controller connected to a device.
func GCRacingWheelFrom(ptr unsafe.Pointer) GCRacingWheel {
	return GCRacingWheel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCRacingWheelClass) Alloc() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCRacingWheelClass) New() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCRacingWheel) Init() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCRacingWheel) Autorelease() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCRacingWheel creates a new GCRacingWheel instance.
func NewGCRacingWheel() GCRacingWheel {
	return getGCRacingWheelClass().New()
}



// Returns a snapshot of the racing wheel with its current element values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/capture()
func (g_ GCRacingWheel) Capture() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](g_.ID, objc.Sel("capture"))
	return rv
}


// Stops receiving events from the racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/relinquishDevice()
func (g_ GCRacingWheel) RelinquishDevice() {
	objc.Send[objc.ID](g_.ID, objc.Sel("relinquishDevice"))
}


// A Boolean value that indicates whether the racing wheel sends events to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/isAcquired
func (g_ GCRacingWheel) Acquired() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("acquired"))
	return rv
}


// A Boolean value that indicates whether the object is a snapshot of a racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheel/isSnapshot
func (g_ GCRacingWheel) Snapshot() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("snapshot"))
	return rv
}


// A Boolean value that indicates whether the racing wheel sends events to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcracingwheel/isacquired
func (g_ GCRacingWheel) IsAcquired() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isAcquired"))
	return rv
}


// A Boolean value that indicates whether the racing wheel sends events to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcracingwheel/isacquired
func (g_ GCRacingWheel) SetIsAcquired(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsAcquired:"), value)
}


// A Boolean value that indicates whether the object is a snapshot of a racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcracingwheel/issnapshot
func (g_ GCRacingWheel) IsSnapshot() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isSnapshot"))
	return rv
}


// A Boolean value that indicates whether the object is a snapshot of a racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcracingwheel/issnapshot
func (g_ GCRacingWheel) SetIsSnapshot(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsSnapshot:"), value)
}


// The physical input profile for the racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcracingwheel/wheelinput
func (g_ GCRacingWheel) WheelInput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("wheelInput"))
	return rv
}


// The physical input profile for the racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcracingwheel/wheelinput
func (g_ GCRacingWheel) SetWheelInput(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setWheelInput:"), value)
}



