// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTargetNavigatorClusterTargetUpdatedEvent] class.
var (
	MTRTargetNavigatorClusterTargetUpdatedEventClass     _MTRTargetNavigatorClusterTargetUpdatedEventClass
	MTRTargetNavigatorClusterTargetUpdatedEventClassOnce sync.Once
)

func getMTRTargetNavigatorClusterTargetUpdatedEventClass() _MTRTargetNavigatorClusterTargetUpdatedEventClass {
	MTRTargetNavigatorClusterTargetUpdatedEventClassOnce.Do(func() {
		MTRTargetNavigatorClusterTargetUpdatedEventClass = _MTRTargetNavigatorClusterTargetUpdatedEventClass{objc.GetClass("MTRTargetNavigatorClusterTargetUpdatedEvent")}
	})
	return MTRTargetNavigatorClusterTargetUpdatedEventClass
}

type _MTRTargetNavigatorClusterTargetUpdatedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRTargetNavigatorClusterTargetUpdatedEvent] class.
type IMTRTargetNavigatorClusterTargetUpdatedEvent interface {
	objectivec.IObject
	CurrentTarget() foundation.Number
	SetCurrentTarget(value foundation.INumber)
	Data() foundation.NSData
	SetData(value foundation.IData)
	TargetList() objc.ID
	SetTargetList(value objc.ID)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent
type MTRTargetNavigatorClusterTargetUpdatedEvent struct {
	objectivec.Object
}

// MTRTargetNavigatorClusterTargetUpdatedEventFrom constructs a [MTRTargetNavigatorClusterTargetUpdatedEvent] from an unsafe.Pointer.
func MTRTargetNavigatorClusterTargetUpdatedEventFrom(ptr unsafe.Pointer) MTRTargetNavigatorClusterTargetUpdatedEvent {
	return MTRTargetNavigatorClusterTargetUpdatedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTargetNavigatorClusterTargetUpdatedEventClass) Alloc() MTRTargetNavigatorClusterTargetUpdatedEvent {
	rv := objc.Send[MTRTargetNavigatorClusterTargetUpdatedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTargetNavigatorClusterTargetUpdatedEventClass) New() MTRTargetNavigatorClusterTargetUpdatedEvent {
	rv := objc.Send[MTRTargetNavigatorClusterTargetUpdatedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) Init() MTRTargetNavigatorClusterTargetUpdatedEvent {
	rv := objc.Send[MTRTargetNavigatorClusterTargetUpdatedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) Autorelease() MTRTargetNavigatorClusterTargetUpdatedEvent {
	rv := objc.Send[MTRTargetNavigatorClusterTargetUpdatedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTargetNavigatorClusterTargetUpdatedEvent creates a new MTRTargetNavigatorClusterTargetUpdatedEvent instance.
func NewMTRTargetNavigatorClusterTargetUpdatedEvent() MTRTargetNavigatorClusterTargetUpdatedEvent {
	return getMTRTargetNavigatorClusterTargetUpdatedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent/currentTarget
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) CurrentTarget() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("currentTarget"))
	return rv
}


// SetCurrentTarget sets the value of the currentTarget property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent/currentTarget
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) SetCurrentTarget(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentTarget:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent/data
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) Data() foundation.NSData {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent/data
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) SetData(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent/targetList
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) TargetList() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("targetList"))
	return rv
}


// SetTargetList sets the value of the targetList property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent/targetList
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) SetTargetList(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetList:"), value)
}



