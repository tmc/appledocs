// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBasicInformationClusterReachableChangedEvent] class.
var (
	MTRBasicInformationClusterReachableChangedEventClass     _MTRBasicInformationClusterReachableChangedEventClass
	MTRBasicInformationClusterReachableChangedEventClassOnce sync.Once
)

func getMTRBasicInformationClusterReachableChangedEventClass() _MTRBasicInformationClusterReachableChangedEventClass {
	MTRBasicInformationClusterReachableChangedEventClassOnce.Do(func() {
		MTRBasicInformationClusterReachableChangedEventClass = _MTRBasicInformationClusterReachableChangedEventClass{objc.GetClass("MTRBasicInformationClusterReachableChangedEvent")}
	})
	return MTRBasicInformationClusterReachableChangedEventClass
}

type _MTRBasicInformationClusterReachableChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicInformationClusterReachableChangedEvent] class.
type IMTRBasicInformationClusterReachableChangedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicInformationClusterReachableChangedEvent
type MTRBasicInformationClusterReachableChangedEvent struct {
	objectivec.Object
}

// MTRBasicInformationClusterReachableChangedEventFrom constructs a [MTRBasicInformationClusterReachableChangedEvent] from an unsafe.Pointer.
func MTRBasicInformationClusterReachableChangedEventFrom(ptr unsafe.Pointer) MTRBasicInformationClusterReachableChangedEvent {
	return MTRBasicInformationClusterReachableChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicInformationClusterReachableChangedEventClass) Alloc() MTRBasicInformationClusterReachableChangedEvent {
	rv := objc.Send[MTRBasicInformationClusterReachableChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicInformationClusterReachableChangedEventClass) New() MTRBasicInformationClusterReachableChangedEvent {
	rv := objc.Send[MTRBasicInformationClusterReachableChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicInformationClusterReachableChangedEvent) Init() MTRBasicInformationClusterReachableChangedEvent {
	rv := objc.Send[MTRBasicInformationClusterReachableChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicInformationClusterReachableChangedEvent) Autorelease() MTRBasicInformationClusterReachableChangedEvent {
	rv := objc.Send[MTRBasicInformationClusterReachableChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicInformationClusterReachableChangedEvent creates a new MTRBasicInformationClusterReachableChangedEvent instance.
func NewMTRBasicInformationClusterReachableChangedEvent() MTRBasicInformationClusterReachableChangedEvent {
	return getMTRBasicInformationClusterReachableChangedEventClass().New()
}




