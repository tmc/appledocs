// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROvenCavityOperationalStateClusterOperationCompletionEvent] class.
var (
	MTROvenCavityOperationalStateClusterOperationCompletionEventClass     _MTROvenCavityOperationalStateClusterOperationCompletionEventClass
	MTROvenCavityOperationalStateClusterOperationCompletionEventClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterOperationCompletionEventClass() _MTROvenCavityOperationalStateClusterOperationCompletionEventClass {
	MTROvenCavityOperationalStateClusterOperationCompletionEventClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterOperationCompletionEventClass = _MTROvenCavityOperationalStateClusterOperationCompletionEventClass{objc.GetClass("MTROvenCavityOperationalStateClusterOperationCompletionEvent")}
	})
	return MTROvenCavityOperationalStateClusterOperationCompletionEventClass
}

type _MTROvenCavityOperationalStateClusterOperationCompletionEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenCavityOperationalStateClusterOperationCompletionEvent] class.
type IMTROvenCavityOperationalStateClusterOperationCompletionEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationCompletionEvent
type MTROvenCavityOperationalStateClusterOperationCompletionEvent struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterOperationCompletionEventFrom constructs a [MTROvenCavityOperationalStateClusterOperationCompletionEvent] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterOperationCompletionEventFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	return MTROvenCavityOperationalStateClusterOperationCompletionEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterOperationCompletionEventClass) Alloc() MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationCompletionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenCavityOperationalStateClusterOperationCompletionEventClass) New() MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationCompletionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) Init() MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationCompletionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) Autorelease() MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationCompletionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterOperationCompletionEvent creates a new MTROvenCavityOperationalStateClusterOperationCompletionEvent instance.
func NewMTROvenCavityOperationalStateClusterOperationCompletionEvent() MTROvenCavityOperationalStateClusterOperationCompletionEvent {
	return getMTROvenCavityOperationalStateClusterOperationCompletionEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationCompletionEvent/completionErrorCode
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) CompletionErrorCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("completionErrorCode"))
	return rv
}


// SetCompletionErrorCode sets the value of the completionErrorCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationCompletionEvent/completionErrorCode
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) SetCompletionErrorCode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCompletionErrorCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationCompletionEvent/pausedTime
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) PausedTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pausedTime"))
	return rv
}


// SetPausedTime sets the value of the pausedTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationCompletionEvent/pausedTime
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) SetPausedTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPausedTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationCompletionEvent/totalOperationalTime
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) TotalOperationalTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalOperationalTime"))
	return rv
}


// SetTotalOperationalTime sets the value of the totalOperationalTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationCompletionEvent/totalOperationalTime
func (m_ MTROvenCavityOperationalStateClusterOperationCompletionEvent) SetTotalOperationalTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTotalOperationalTime:"), value)
}



