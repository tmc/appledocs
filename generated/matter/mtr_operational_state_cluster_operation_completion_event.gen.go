// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROperationalStateClusterOperationCompletionEvent] class.
var (
	MTROperationalStateClusterOperationCompletionEventClass     _MTROperationalStateClusterOperationCompletionEventClass
	MTROperationalStateClusterOperationCompletionEventClassOnce sync.Once
)

func getMTROperationalStateClusterOperationCompletionEventClass() _MTROperationalStateClusterOperationCompletionEventClass {
	MTROperationalStateClusterOperationCompletionEventClassOnce.Do(func() {
		MTROperationalStateClusterOperationCompletionEventClass = _MTROperationalStateClusterOperationCompletionEventClass{objc.GetClass("MTROperationalStateClusterOperationCompletionEvent")}
	})
	return MTROperationalStateClusterOperationCompletionEventClass
}

type _MTROperationalStateClusterOperationCompletionEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalStateClusterOperationCompletionEvent] class.
type IMTROperationalStateClusterOperationCompletionEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalStateClusterOperationCompletionEvent
type MTROperationalStateClusterOperationCompletionEvent struct {
	objectivec.Object
}

// MTROperationalStateClusterOperationCompletionEventFrom constructs a [MTROperationalStateClusterOperationCompletionEvent] from an unsafe.Pointer.
func MTROperationalStateClusterOperationCompletionEventFrom(ptr unsafe.Pointer) MTROperationalStateClusterOperationCompletionEvent {
	return MTROperationalStateClusterOperationCompletionEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalStateClusterOperationCompletionEventClass) Alloc() MTROperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROperationalStateClusterOperationCompletionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalStateClusterOperationCompletionEventClass) New() MTROperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROperationalStateClusterOperationCompletionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalStateClusterOperationCompletionEvent) Init() MTROperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROperationalStateClusterOperationCompletionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalStateClusterOperationCompletionEvent) Autorelease() MTROperationalStateClusterOperationCompletionEvent {
	rv := objc.Send[MTROperationalStateClusterOperationCompletionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalStateClusterOperationCompletionEvent creates a new MTROperationalStateClusterOperationCompletionEvent instance.
func NewMTROperationalStateClusterOperationCompletionEvent() MTROperationalStateClusterOperationCompletionEvent {
	return getMTROperationalStateClusterOperationCompletionEventClass().New()
}




