// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterStateChangedEvent] class.
var (
	MTRActionsClusterStateChangedEventClass     _MTRActionsClusterStateChangedEventClass
	MTRActionsClusterStateChangedEventClassOnce sync.Once
)

func getMTRActionsClusterStateChangedEventClass() _MTRActionsClusterStateChangedEventClass {
	MTRActionsClusterStateChangedEventClassOnce.Do(func() {
		MTRActionsClusterStateChangedEventClass = _MTRActionsClusterStateChangedEventClass{objc.GetClass("MTRActionsClusterStateChangedEvent")}
	})
	return MTRActionsClusterStateChangedEventClass
}

type _MTRActionsClusterStateChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterStateChangedEvent] class.
type IMTRActionsClusterStateChangedEvent interface {
	objectivec.IObject
	ActionID() foundation.Number
	SetActionID(value foundation.INumber)
	InvokeID() foundation.Number
	SetInvokeID(value foundation.INumber)
	NewState() foundation.Number
	SetNewState(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStateChangedEvent
type MTRActionsClusterStateChangedEvent struct {
	objectivec.Object
}

// MTRActionsClusterStateChangedEventFrom constructs a [MTRActionsClusterStateChangedEvent] from an unsafe.Pointer.
func MTRActionsClusterStateChangedEventFrom(ptr unsafe.Pointer) MTRActionsClusterStateChangedEvent {
	return MTRActionsClusterStateChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterStateChangedEventClass) Alloc() MTRActionsClusterStateChangedEvent {
	rv := objc.Send[MTRActionsClusterStateChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterStateChangedEventClass) New() MTRActionsClusterStateChangedEvent {
	rv := objc.Send[MTRActionsClusterStateChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterStateChangedEvent) Init() MTRActionsClusterStateChangedEvent {
	rv := objc.Send[MTRActionsClusterStateChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterStateChangedEvent) Autorelease() MTRActionsClusterStateChangedEvent {
	rv := objc.Send[MTRActionsClusterStateChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterStateChangedEvent creates a new MTRActionsClusterStateChangedEvent instance.
func NewMTRActionsClusterStateChangedEvent() MTRActionsClusterStateChangedEvent {
	return getMTRActionsClusterStateChangedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstatechangedevent/actionid
func (m_ MTRActionsClusterStateChangedEvent) ActionID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("actionID"))
	return rv
}


// SetActionID sets the value of the actionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstatechangedevent/actionid
func (m_ MTRActionsClusterStateChangedEvent) SetActionID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstatechangedevent/invokeid
func (m_ MTRActionsClusterStateChangedEvent) InvokeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("invokeID"))
	return rv
}


// SetInvokeID sets the value of the invokeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstatechangedevent/invokeid
func (m_ MTRActionsClusterStateChangedEvent) SetInvokeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstatechangedevent/newstate
func (m_ MTRActionsClusterStateChangedEvent) NewState() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("newState"))
	return rv
}


// SetNewState sets the value of the newState property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusterstatechangedevent/newstate
func (m_ MTRActionsClusterStateChangedEvent) SetNewState(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewState:"), value)
}



