// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineAccountChangeEvent] class.
var (
	CKSyncEngineAccountChangeEventClass     _CKSyncEngineAccountChangeEventClass
	CKSyncEngineAccountChangeEventClassOnce sync.Once
)

func getCKSyncEngineAccountChangeEventClass() _CKSyncEngineAccountChangeEventClass {
	CKSyncEngineAccountChangeEventClassOnce.Do(func() {
		CKSyncEngineAccountChangeEventClass = _CKSyncEngineAccountChangeEventClass{objc.GetClass("CKSyncEngineAccountChangeEvent")}
	})
	return CKSyncEngineAccountChangeEventClass
}

type _CKSyncEngineAccountChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineAccountChangeEvent] class.
type ICKSyncEngineAccountChangeEvent interface {
	ICKSyncEngineEvent
}

// An event that provides information about a change to the device’s iCloud account.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeEvent
type CKSyncEngineAccountChangeEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineAccountChangeEventFrom constructs a [CKSyncEngineAccountChangeEvent] from an unsafe.Pointer.
//
// An event that provides information about a change to the device’s iCloud account.
func CKSyncEngineAccountChangeEventFrom(ptr unsafe.Pointer) CKSyncEngineAccountChangeEvent {
	return CKSyncEngineAccountChangeEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineAccountChangeEventClass) Alloc() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineAccountChangeEventClass) New() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineAccountChangeEvent) Init() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineAccountChangeEvent) Autorelease() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineAccountChangeEvent creates a new CKSyncEngineAccountChangeEvent instance.
func NewCKSyncEngineAccountChangeEvent() CKSyncEngineAccountChangeEvent {
	return getCKSyncEngineAccountChangeEventClass().New()
}




