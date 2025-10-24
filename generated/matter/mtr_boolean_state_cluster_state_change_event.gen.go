// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBooleanStateClusterStateChangeEvent] class.
var (
	MTRBooleanStateClusterStateChangeEventClass     _MTRBooleanStateClusterStateChangeEventClass
	MTRBooleanStateClusterStateChangeEventClassOnce sync.Once
)

func getMTRBooleanStateClusterStateChangeEventClass() _MTRBooleanStateClusterStateChangeEventClass {
	MTRBooleanStateClusterStateChangeEventClassOnce.Do(func() {
		MTRBooleanStateClusterStateChangeEventClass = _MTRBooleanStateClusterStateChangeEventClass{objc.GetClass("MTRBooleanStateClusterStateChangeEvent")}
	})
	return MTRBooleanStateClusterStateChangeEventClass
}

type _MTRBooleanStateClusterStateChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBooleanStateClusterStateChangeEvent] class.
type IMTRBooleanStateClusterStateChangeEvent interface {
	objectivec.IObject
	// properties:
	StateValue() objc.IObject /* cross-framework: NSNumber */
	SetStateValue(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBooleanStateClusterStateChangeEvent
type MTRBooleanStateClusterStateChangeEvent struct {
	objectivec.Object
}

// MTRBooleanStateClusterStateChangeEventFrom constructs a [MTRBooleanStateClusterStateChangeEvent] from an unsafe.Pointer.
func MTRBooleanStateClusterStateChangeEventFrom(ptr unsafe.Pointer) MTRBooleanStateClusterStateChangeEvent {
	return MTRBooleanStateClusterStateChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBooleanStateClusterStateChangeEventClass) Alloc() MTRBooleanStateClusterStateChangeEvent {
	rv := objc.Send[MTRBooleanStateClusterStateChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBooleanStateClusterStateChangeEventClass) New() MTRBooleanStateClusterStateChangeEvent {
	rv := objc.Send[MTRBooleanStateClusterStateChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBooleanStateClusterStateChangeEvent) Init() MTRBooleanStateClusterStateChangeEvent {
	rv := objc.Send[MTRBooleanStateClusterStateChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBooleanStateClusterStateChangeEvent) Autorelease() MTRBooleanStateClusterStateChangeEvent {
	rv := objc.Send[MTRBooleanStateClusterStateChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBooleanStateClusterStateChangeEvent creates a new MTRBooleanStateClusterStateChangeEvent instance.
func NewMTRBooleanStateClusterStateChangeEvent() MTRBooleanStateClusterStateChangeEvent {
	return getMTRBooleanStateClusterStateChangeEventClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateclusterstatechangeevent/statevalue
func (m_ MTRBooleanStateClusterStateChangeEvent) StateValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stateValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbooleanstateclusterstatechangeevent/statevalue
func (m_ MTRBooleanStateClusterStateChangeEvent) SetStateValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStateValue:"), value)
}
