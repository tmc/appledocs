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
	// properties:
	ActionID() objc.IObject /* cross-framework: NSNumber */
	SetActionID(value objc.IObject /* cross-framework: NSNumber */)
	Error() objc.IObject /* cross-framework: NSNumber */
	SetError(value objc.IObject /* cross-framework: NSNumber */)
	InvokeID() objc.IObject /* cross-framework: NSNumber */
	SetInvokeID(value objc.IObject /* cross-framework: NSNumber */)
	NewState() objc.IObject /* cross-framework: NSNumber */
	SetNewState(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/actionid
func (m_ MTRActionsClusterActionFailedEvent) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/actionid
func (m_ MTRActionsClusterActionFailedEvent) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/error
func (m_ MTRActionsClusterActionFailedEvent) Error() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("error"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/error
func (m_ MTRActionsClusterActionFailedEvent) SetError(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setError:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/invokeid
func (m_ MTRActionsClusterActionFailedEvent) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/invokeid
func (m_ MTRActionsClusterActionFailedEvent) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/newstate
func (m_ MTRActionsClusterActionFailedEvent) NewState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractionsclusteractionfailedevent/newstate
func (m_ MTRActionsClusterActionFailedEvent) SetNewState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewState:"), value)
}



