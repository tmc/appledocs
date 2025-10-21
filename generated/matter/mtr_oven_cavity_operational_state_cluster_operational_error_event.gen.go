// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROvenCavityOperationalStateClusterOperationalErrorEvent] class.
var (
	MTROvenCavityOperationalStateClusterOperationalErrorEventClass     _MTROvenCavityOperationalStateClusterOperationalErrorEventClass
	MTROvenCavityOperationalStateClusterOperationalErrorEventClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterOperationalErrorEventClass() _MTROvenCavityOperationalStateClusterOperationalErrorEventClass {
	MTROvenCavityOperationalStateClusterOperationalErrorEventClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterOperationalErrorEventClass = _MTROvenCavityOperationalStateClusterOperationalErrorEventClass{objc.GetClass("MTROvenCavityOperationalStateClusterOperationalErrorEvent")}
	})
	return MTROvenCavityOperationalStateClusterOperationalErrorEventClass
}

type _MTROvenCavityOperationalStateClusterOperationalErrorEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenCavityOperationalStateClusterOperationalErrorEvent] class.
type IMTROvenCavityOperationalStateClusterOperationalErrorEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalErrorEvent
type MTROvenCavityOperationalStateClusterOperationalErrorEvent struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterOperationalErrorEventFrom constructs a [MTROvenCavityOperationalStateClusterOperationalErrorEvent] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterOperationalErrorEventFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	return MTROvenCavityOperationalStateClusterOperationalErrorEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterOperationalErrorEventClass) Alloc() MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenCavityOperationalStateClusterOperationalErrorEventClass) New() MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterOperationalErrorEvent) Init() MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterOperationalErrorEvent) Autorelease() MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterOperationalErrorEvent creates a new MTROvenCavityOperationalStateClusterOperationalErrorEvent instance.
func NewMTROvenCavityOperationalStateClusterOperationalErrorEvent() MTROvenCavityOperationalStateClusterOperationalErrorEvent {
	return getMTROvenCavityOperationalStateClusterOperationalErrorEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalErrorEvent/errorState
func (m_ MTROvenCavityOperationalStateClusterOperationalErrorEvent) ErrorState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("errorState"))
	return rv
}


// SetErrorState sets the value of the errorState property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalErrorEvent/errorState
func (m_ MTROvenCavityOperationalStateClusterOperationalErrorEvent) SetErrorState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorState:"), value)
}


