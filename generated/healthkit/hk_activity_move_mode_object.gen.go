// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKActivityMoveModeObject] class.
var (
	HKActivityMoveModeObjectClass     _HKActivityMoveModeObjectClass
	HKActivityMoveModeObjectClassOnce sync.Once
)

func getHKActivityMoveModeObjectClass() _HKActivityMoveModeObjectClass {
	HKActivityMoveModeObjectClassOnce.Do(func() {
		HKActivityMoveModeObjectClass = _HKActivityMoveModeObjectClass{objc.GetClass("HKActivityMoveModeObject")}
	})
	return HKActivityMoveModeObjectClass
}

type _HKActivityMoveModeObjectClass struct {
	class objc.Class
}

// An interface definition for the [HKActivityMoveModeObject] class.
type IHKActivityMoveModeObject interface {
	objectivec.IObject
	// properties:
	ActivityMoveMode() unsafe.Pointer
	SetActivityMoveMode(value unsafe.Pointer)
	// methods:
}

// An object that contains a movement mode value.


// An object that contains a movement mode value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivityMoveModeObject
type HKActivityMoveModeObject struct {
	objectivec.Object
}

// HKActivityMoveModeObjectFrom constructs a [HKActivityMoveModeObject] from an unsafe.Pointer.
//
// An object that contains a movement mode value.
func HKActivityMoveModeObjectFrom(ptr unsafe.Pointer) HKActivityMoveModeObject {
	return HKActivityMoveModeObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKActivityMoveModeObjectClass) Alloc() HKActivityMoveModeObject {
	rv := objc.Send[HKActivityMoveModeObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKActivityMoveModeObjectClass) New() HKActivityMoveModeObject {
	rv := objc.Send[HKActivityMoveModeObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKActivityMoveModeObject) Init() HKActivityMoveModeObject {
	rv := objc.Send[HKActivityMoveModeObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKActivityMoveModeObject) Autorelease() HKActivityMoveModeObject {
	rv := objc.Send[HKActivityMoveModeObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKActivityMoveModeObject creates a new HKActivityMoveModeObject instance.
func NewHKActivityMoveModeObject() HKActivityMoveModeObject {
	return getHKActivityMoveModeObjectClass().New()
}



// A property that contains the movement mode value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitymovemodeobject/activitymovemode
func (h_ HKActivityMoveModeObject) ActivityMoveMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("activityMoveMode"))
	return rv
}


// A property that contains the movement mode value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkactivitymovemodeobject/activitymovemode
func (h_ HKActivityMoveModeObject) SetActivityMoveMode(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setActivityMoveMode:"), value)
}



