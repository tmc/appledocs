// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	CompletionErrorCode() objc.IObject /* cross-framework: NSNumber */
	SetCompletionErrorCode(value objc.IObject /* cross-framework: NSNumber */)
	PausedTime() objc.IObject /* cross-framework: NSNumber */
	SetPausedTime(value objc.IObject /* cross-framework: NSNumber */)
	TotalOperationalTime() objc.IObject /* cross-framework: NSNumber */
	SetTotalOperationalTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusteroperationcompletionevent/completionerrorcode
func (m_ MTROperationalStateClusterOperationCompletionEvent) CompletionErrorCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("completionErrorCode"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusteroperationcompletionevent/completionerrorcode
func (m_ MTROperationalStateClusterOperationCompletionEvent) SetCompletionErrorCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCompletionErrorCode:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusteroperationcompletionevent/pausedtime
func (m_ MTROperationalStateClusterOperationCompletionEvent) PausedTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("pausedTime"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusteroperationcompletionevent/pausedtime
func (m_ MTROperationalStateClusterOperationCompletionEvent) SetPausedTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPausedTime:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusteroperationcompletionevent/totaloperationaltime
func (m_ MTROperationalStateClusterOperationCompletionEvent) TotalOperationalTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("totalOperationalTime"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalstateclusteroperationcompletionevent/totaloperationaltime
func (m_ MTROperationalStateClusterOperationCompletionEvent) SetTotalOperationalTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTotalOperationalTime:"), value)
}
