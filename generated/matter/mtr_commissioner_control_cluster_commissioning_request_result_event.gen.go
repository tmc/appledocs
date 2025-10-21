// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRCommissionerControlClusterCommissioningRequestResultEvent] class.
var (
	MTRCommissionerControlClusterCommissioningRequestResultEventClass     _MTRCommissionerControlClusterCommissioningRequestResultEventClass
	MTRCommissionerControlClusterCommissioningRequestResultEventClassOnce sync.Once
)

func getMTRCommissionerControlClusterCommissioningRequestResultEventClass() _MTRCommissionerControlClusterCommissioningRequestResultEventClass {
	MTRCommissionerControlClusterCommissioningRequestResultEventClassOnce.Do(func() {
		MTRCommissionerControlClusterCommissioningRequestResultEventClass = _MTRCommissionerControlClusterCommissioningRequestResultEventClass{objc.GetClass("MTRCommissionerControlClusterCommissioningRequestResultEvent")}
	})
	return MTRCommissionerControlClusterCommissioningRequestResultEventClass
}

type _MTRCommissionerControlClusterCommissioningRequestResultEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRCommissionerControlClusterCommissioningRequestResultEvent] class.
type IMTRCommissionerControlClusterCommissioningRequestResultEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent
type MTRCommissionerControlClusterCommissioningRequestResultEvent struct {
	objectivec.Object
}

// MTRCommissionerControlClusterCommissioningRequestResultEventFrom constructs a [MTRCommissionerControlClusterCommissioningRequestResultEvent] from an unsafe.Pointer.
func MTRCommissionerControlClusterCommissioningRequestResultEventFrom(ptr unsafe.Pointer) MTRCommissionerControlClusterCommissioningRequestResultEvent {
	return MTRCommissionerControlClusterCommissioningRequestResultEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCommissionerControlClusterCommissioningRequestResultEventClass) Alloc() MTRCommissionerControlClusterCommissioningRequestResultEvent {
	rv := objc.Send[MTRCommissionerControlClusterCommissioningRequestResultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCommissionerControlClusterCommissioningRequestResultEventClass) New() MTRCommissionerControlClusterCommissioningRequestResultEvent {
	rv := objc.Send[MTRCommissionerControlClusterCommissioningRequestResultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) Init() MTRCommissionerControlClusterCommissioningRequestResultEvent {
	rv := objc.Send[MTRCommissionerControlClusterCommissioningRequestResultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) Autorelease() MTRCommissionerControlClusterCommissioningRequestResultEvent {
	rv := objc.Send[MTRCommissionerControlClusterCommissioningRequestResultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissionerControlClusterCommissioningRequestResultEvent creates a new MTRCommissionerControlClusterCommissioningRequestResultEvent instance.
func NewMTRCommissionerControlClusterCommissioningRequestResultEvent() MTRCommissionerControlClusterCommissioningRequestResultEvent {
	return getMTRCommissionerControlClusterCommissioningRequestResultEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent/clientNodeID
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) ClientNodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("clientNodeID"))
	return rv
}


// SetClientNodeID sets the value of the clientNodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent/clientNodeID
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) SetClientNodeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClientNodeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent/fabricIndex
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent/fabricIndex
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent/requestID
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) RequestID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("requestID"))
	return rv
}


// SetRequestID sets the value of the requestID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent/requestID
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) SetRequestID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent/statusCode
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) StatusCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("statusCode"))
	return rv
}


// SetStatusCode sets the value of the statusCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterCommissioningRequestResultEvent/statusCode
func (m_ MTRCommissionerControlClusterCommissioningRequestResultEvent) SetStatusCode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusCode:"), value)
}



