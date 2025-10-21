// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterActionFailedEvent] class.
var (
	MTRActionsClusterActionFailedEventClass     _MTRActionsClusterActionFailedEventClass
	MTRActionsClusterActionFailedEventClassOnce sync.Once
)

func getMTRActionsClusterActionFailedEventClass() _MTRActionsClusterActionFailedEventClass {
	MTRActionsClusterActionFailedEventClassOnce.Do(func() {
		MTRActionsClusterActionFailedEventClass = _MTRActionsClusterActionFailedEventClass{objc.GetClass("MTRActionsClusterActionFailedEvent")}
	})
	return MTRActionsClusterActionFailedEventClass
}

type _MTRActionsClusterActionFailedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterActionFailedEvent] class.
type IMTRActionsClusterActionFailedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionFailedEvent
type MTRActionsClusterActionFailedEvent struct {
	objectivec.Object
}

// MTRActionsClusterActionFailedEventFrom constructs a [MTRActionsClusterActionFailedEvent] from an unsafe.Pointer.
func MTRActionsClusterActionFailedEventFrom(ptr unsafe.Pointer) MTRActionsClusterActionFailedEvent {
	return MTRActionsClusterActionFailedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterActionFailedEventClass) Alloc() MTRActionsClusterActionFailedEvent {
	rv := objc.Send[MTRActionsClusterActionFailedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterActionFailedEventClass) New() MTRActionsClusterActionFailedEvent {
	rv := objc.Send[MTRActionsClusterActionFailedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterActionFailedEvent) Init() MTRActionsClusterActionFailedEvent {
	rv := objc.Send[MTRActionsClusterActionFailedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterActionFailedEvent) Autorelease() MTRActionsClusterActionFailedEvent {
	rv := objc.Send[MTRActionsClusterActionFailedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterActionFailedEvent creates a new MTRActionsClusterActionFailedEvent instance.
func NewMTRActionsClusterActionFailedEvent() MTRActionsClusterActionFailedEvent {
	return getMTRActionsClusterActionFailedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/actionid
func (m_ MTRActionsClusterActionFailedEvent) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/actionid
func (m_ MTRActionsClusterActionFailedEvent) SetActionID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/error
func (m_ MTRActionsClusterActionFailedEvent) Error() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("error"))
	return rv
}


// SetError sets the value of the error property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/error
func (m_ MTRActionsClusterActionFailedEvent) SetError(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setError:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/invokeid
func (m_ MTRActionsClusterActionFailedEvent) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/invokeid
func (m_ MTRActionsClusterActionFailedEvent) SetInvokeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/newstate
func (m_ MTRActionsClusterActionFailedEvent) NewState() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("newState"))
	return rv
}


// SetNewState sets the value of the newState property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/newstate
func (m_ MTRActionsClusterActionFailedEvent) SetNewState(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewState:"), value)
}



