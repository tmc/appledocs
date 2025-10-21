// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKRouteStep] class.
var (
	MKRouteStepClass     _MKRouteStepClass
	MKRouteStepClassOnce sync.Once
)

func getMKRouteStepClass() _MKRouteStepClass {
	MKRouteStepClassOnce.Do(func() {
		MKRouteStepClass = _MKRouteStepClass{objc.GetClass("MKRouteStep")}
	})
	return MKRouteStepClass
}

type _MKRouteStepClass struct {
	class objc.Class
}

// An interface definition for the [MKRouteStep] class.
type IMKRouteStep interface {
	objectivec.IObject
}

// One portion of an overall route.
//
// Each object corresponds to a single instruction that the person needs to follow when navigating between two points. For example, a step might involve following a single road until continuing along the route requires a turn. You don’t create instances of this class directly. An object contains the objects associated with a route. For more information about requesting directions, see .
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKRoute/Step
type MKRouteStep struct {
	objectivec.Object
}

// MKRouteStepFrom constructs a [MKRouteStep] from an unsafe.Pointer.
//
// One portion of an overall route.
func MKRouteStepFrom(ptr unsafe.Pointer) MKRouteStep {
	return MKRouteStep{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKRouteStepClass) Alloc() MKRouteStep {
	rv := objc.Send[MKRouteStep](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKRouteStepClass) New() MKRouteStep {
	rv := objc.Send[MKRouteStep](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKRouteStep) Init() MKRouteStep {
	rv := objc.Send[MKRouteStep](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKRouteStep) Autorelease() MKRouteStep {
	rv := objc.Send[MKRouteStep](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKRouteStep creates a new MKRouteStep instance.
func NewMKRouteStep() MKRouteStep {
	return getMKRouteStepClass().New()
}




