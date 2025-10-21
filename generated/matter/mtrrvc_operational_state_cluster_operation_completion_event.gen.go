// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCOperationalStateClusterOperationCompletionEvent] class.
var (
	MTRRVCOperationalStateClusterOperationCompletionEventClass     _MTRRVCOperationalStateClusterOperationCompletionEventClass
	MTRRVCOperationalStateClusterOperationCompletionEventClassOnce sync.Once
)

func getMTRRVCOperationalStateClusterOperationCompletionEventClass() _MTRRVCOperationalStateClusterOperationCompletionEventClass {
	MTRRVCOperationalStateClusterOperationCompletionEventClassOnce.Do(func() {
		MTRRVCOperationalStateClusterOperationCompletionEventClass = _MTRRVCOperationalStateClusterOperationCompletionEventClass{objc.GetClass("MTRRVCOperationalStateClusterOperationCompletionEvent")}
	})
	return MTRRVCOperationalStateClusterOperationCompletionEventClass
}

type _MTRRVCOperationalStateClusterOperationCompletionEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCOperationalStateClusterOperationCompletionEvent] class.
type IMTRRVCOperationalStateClusterOperationCompletionEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterOperationCompletionEvent
type MTRRVCOperationalStateClusterOperationCompletionEvent struct {
	objectivec.Object
}

// MTRRVCOperationalStateClusterOperationCompletionEventFrom constructs a [MTRRVCOperationalStateClusterOperationCompletionEvent] from an unsafe.Pointer.
func MTRRVCOperationalStateClusterOperationCompletionEventFrom(ptr unsafe.Pointer) MTRRVCOperationalStateClusterOperationCompletionEvent {
	return MTRRVCOperationalStateClusterOperationCompletionEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCOperationalStateClusterOperationCompletionEventClass) Alloc() MTRRVCOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationCompletionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCOperationalStateClusterOperationCompletionEventClass) New() MTRRVCOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationCompletionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCOperationalStateClusterOperationCompletionEvent) Init() MTRRVCOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationCompletionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCOperationalStateClusterOperationCompletionEvent) Autorelease() MTRRVCOperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationCompletionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCOperationalStateClusterOperationCompletionEvent creates a new MTRRVCOperationalStateClusterOperationCompletionEvent instance.
func NewMTRRVCOperationalStateClusterOperationCompletionEvent() MTRRVCOperationalStateClusterOperationCompletionEvent {
	return getMTRRVCOperationalStateClusterOperationCompletionEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationcompletionevent/totaloperationaltime
func (m_ MTRRVCOperationalStateClusterOperationCompletionEvent) TotalOperationalTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("totalOperationalTime"))
	return rv
}


// SetTotalOperationalTime sets the value of the totalOperationalTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationcompletionevent/totaloperationaltime
func (m_ MTRRVCOperationalStateClusterOperationCompletionEvent) SetTotalOperationalTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTotalOperationalTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationcompletionevent/pausedtime
func (m_ MTRRVCOperationalStateClusterOperationCompletionEvent) PausedTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("pausedTime"))
	return rv
}


// SetPausedTime sets the value of the pausedTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationcompletionevent/pausedtime
func (m_ MTRRVCOperationalStateClusterOperationCompletionEvent) SetPausedTime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPausedTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationcompletionevent/completionerrorcode
func (m_ MTRRVCOperationalStateClusterOperationCompletionEvent) CompletionErrorCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("completionErrorCode"))
	return rv
}


// SetCompletionErrorCode sets the value of the completionErrorCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationcompletionevent/completionerrorcode
func (m_ MTRRVCOperationalStateClusterOperationCompletionEvent) SetCompletionErrorCode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCompletionErrorCode:"), value)
}



