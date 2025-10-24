// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCOperationalStateClusterOperationalErrorEvent] class.
var (
	MTRRVCOperationalStateClusterOperationalErrorEventClass     _MTRRVCOperationalStateClusterOperationalErrorEventClass
	MTRRVCOperationalStateClusterOperationalErrorEventClassOnce sync.Once
)

func getMTRRVCOperationalStateClusterOperationalErrorEventClass() _MTRRVCOperationalStateClusterOperationalErrorEventClass {
	MTRRVCOperationalStateClusterOperationalErrorEventClassOnce.Do(func() {
		MTRRVCOperationalStateClusterOperationalErrorEventClass = _MTRRVCOperationalStateClusterOperationalErrorEventClass{objc.GetClass("MTRRVCOperationalStateClusterOperationalErrorEvent")}
	})
	return MTRRVCOperationalStateClusterOperationalErrorEventClass
}

type _MTRRVCOperationalStateClusterOperationalErrorEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCOperationalStateClusterOperationalErrorEvent] class.
type IMTRRVCOperationalStateClusterOperationalErrorEvent interface {
	objectivec.IObject
	// properties:
	ErrorState() IMTRRVCOperationalStateClusterErrorStateStruct
	SetErrorState(value IMTRRVCOperationalStateClusterErrorStateStruct)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterOperationalErrorEvent
type MTRRVCOperationalStateClusterOperationalErrorEvent struct {
	objectivec.Object
}

// MTRRVCOperationalStateClusterOperationalErrorEventFrom constructs a [MTRRVCOperationalStateClusterOperationalErrorEvent] from an unsafe.Pointer.
func MTRRVCOperationalStateClusterOperationalErrorEventFrom(ptr unsafe.Pointer) MTRRVCOperationalStateClusterOperationalErrorEvent {
	return MTRRVCOperationalStateClusterOperationalErrorEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCOperationalStateClusterOperationalErrorEventClass) Alloc() MTRRVCOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCOperationalStateClusterOperationalErrorEventClass) New() MTRRVCOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCOperationalStateClusterOperationalErrorEvent) Init() MTRRVCOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCOperationalStateClusterOperationalErrorEvent) Autorelease() MTRRVCOperationalStateClusterOperationalErrorEvent {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCOperationalStateClusterOperationalErrorEvent creates a new MTRRVCOperationalStateClusterOperationalErrorEvent instance.
func NewMTRRVCOperationalStateClusterOperationalErrorEvent() MTRRVCOperationalStateClusterOperationalErrorEvent {
	return getMTRRVCOperationalStateClusterOperationalErrorEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalerrorevent/errorstate
func (m_ MTRRVCOperationalStateClusterOperationalErrorEvent) ErrorState() IMTRRVCOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("errorState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalerrorevent/errorstate
func (m_ MTRRVCOperationalStateClusterOperationalErrorEvent) SetErrorState(value IMTRRVCOperationalStateClusterErrorStateStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorState:"), value)
}



