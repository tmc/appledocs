// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWaterHeaterManagementClusterBoostStartedEvent] class.
var (
	MTRWaterHeaterManagementClusterBoostStartedEventClass     _MTRWaterHeaterManagementClusterBoostStartedEventClass
	MTRWaterHeaterManagementClusterBoostStartedEventClassOnce sync.Once
)

func getMTRWaterHeaterManagementClusterBoostStartedEventClass() _MTRWaterHeaterManagementClusterBoostStartedEventClass {
	MTRWaterHeaterManagementClusterBoostStartedEventClassOnce.Do(func() {
		MTRWaterHeaterManagementClusterBoostStartedEventClass = _MTRWaterHeaterManagementClusterBoostStartedEventClass{objc.GetClass("MTRWaterHeaterManagementClusterBoostStartedEvent")}
	})
	return MTRWaterHeaterManagementClusterBoostStartedEventClass
}

type _MTRWaterHeaterManagementClusterBoostStartedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRWaterHeaterManagementClusterBoostStartedEvent] class.
type IMTRWaterHeaterManagementClusterBoostStartedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostStartedEvent
type MTRWaterHeaterManagementClusterBoostStartedEvent struct {
	objectivec.Object
}

// MTRWaterHeaterManagementClusterBoostStartedEventFrom constructs a [MTRWaterHeaterManagementClusterBoostStartedEvent] from an unsafe.Pointer.
func MTRWaterHeaterManagementClusterBoostStartedEventFrom(ptr unsafe.Pointer) MTRWaterHeaterManagementClusterBoostStartedEvent {
	return MTRWaterHeaterManagementClusterBoostStartedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterManagementClusterBoostStartedEventClass) Alloc() MTRWaterHeaterManagementClusterBoostStartedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostStartedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWaterHeaterManagementClusterBoostStartedEventClass) New() MTRWaterHeaterManagementClusterBoostStartedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostStartedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterManagementClusterBoostStartedEvent) Init() MTRWaterHeaterManagementClusterBoostStartedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostStartedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterManagementClusterBoostStartedEvent) Autorelease() MTRWaterHeaterManagementClusterBoostStartedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostStartedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterManagementClusterBoostStartedEvent creates a new MTRWaterHeaterManagementClusterBoostStartedEvent instance.
func NewMTRWaterHeaterManagementClusterBoostStartedEvent() MTRWaterHeaterManagementClusterBoostStartedEvent {
	return getMTRWaterHeaterManagementClusterBoostStartedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostStartedEvent/boostInfo
func (m_ MTRWaterHeaterManagementClusterBoostStartedEvent) BoostInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("boostInfo"))
	return rv
}


// SetBoostInfo sets the value of the boostInfo property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostStartedEvent/boostInfo
func (m_ MTRWaterHeaterManagementClusterBoostStartedEvent) SetBoostInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoostInfo:"), value)
}


