// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBasicInformationClusterStartUpEvent] class.
var (
	MTRBasicInformationClusterStartUpEventClass     _MTRBasicInformationClusterStartUpEventClass
	MTRBasicInformationClusterStartUpEventClassOnce sync.Once
)

func getMTRBasicInformationClusterStartUpEventClass() _MTRBasicInformationClusterStartUpEventClass {
	MTRBasicInformationClusterStartUpEventClassOnce.Do(func() {
		MTRBasicInformationClusterStartUpEventClass = _MTRBasicInformationClusterStartUpEventClass{objc.GetClass("MTRBasicInformationClusterStartUpEvent")}
	})
	return MTRBasicInformationClusterStartUpEventClass
}

type _MTRBasicInformationClusterStartUpEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicInformationClusterStartUpEvent] class.
type IMTRBasicInformationClusterStartUpEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicInformationClusterStartUpEvent
type MTRBasicInformationClusterStartUpEvent struct {
	objectivec.Object
}

// MTRBasicInformationClusterStartUpEventFrom constructs a [MTRBasicInformationClusterStartUpEvent] from an unsafe.Pointer.
func MTRBasicInformationClusterStartUpEventFrom(ptr unsafe.Pointer) MTRBasicInformationClusterStartUpEvent {
	return MTRBasicInformationClusterStartUpEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicInformationClusterStartUpEventClass) Alloc() MTRBasicInformationClusterStartUpEvent {
	rv := objc.Send[MTRBasicInformationClusterStartUpEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicInformationClusterStartUpEventClass) New() MTRBasicInformationClusterStartUpEvent {
	rv := objc.Send[MTRBasicInformationClusterStartUpEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicInformationClusterStartUpEvent) Init() MTRBasicInformationClusterStartUpEvent {
	rv := objc.Send[MTRBasicInformationClusterStartUpEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicInformationClusterStartUpEvent) Autorelease() MTRBasicInformationClusterStartUpEvent {
	rv := objc.Send[MTRBasicInformationClusterStartUpEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicInformationClusterStartUpEvent creates a new MTRBasicInformationClusterStartUpEvent instance.
func NewMTRBasicInformationClusterStartUpEvent() MTRBasicInformationClusterStartUpEvent {
	return getMTRBasicInformationClusterStartUpEventClass().New()
}




