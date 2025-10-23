// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineEvent] class.
var (
	CKSyncEngineEventClass     _CKSyncEngineEventClass
	CKSyncEngineEventClassOnce sync.Once
)

func getCKSyncEngineEventClass() _CKSyncEngineEventClass {
	CKSyncEngineEventClassOnce.Do(func() {
		CKSyncEngineEventClass = _CKSyncEngineEventClass{objc.GetClass("CKSyncEngineEvent")}
	})
	return CKSyncEngineEventClass
}

type _CKSyncEngineEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineEvent] class.
type ICKSyncEngineEvent interface {
	objectivec.IObject
}

// A parent class referenced by other CloudKit classes.


// A parent class referenced by other CloudKit classes. [Full Topic]
type CKSyncEngineEvent struct {
	objectivec.Object
}

// CKSyncEngineEventFrom constructs a [CKSyncEngineEvent] from an unsafe.Pointer.
//
// A parent class referenced by other CloudKit classes.
func CKSyncEngineEventFrom(ptr unsafe.Pointer) CKSyncEngineEvent {
	return CKSyncEngineEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineEventClass) Alloc() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineEventClass) New() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineEvent) Init() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineEvent) Autorelease() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineEvent creates a new CKSyncEngineEvent instance.
func NewCKSyncEngineEvent() CKSyncEngineEvent {
	return getCKSyncEngineEventClass().New()
}




