// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineStateSerialization] class.
var (
	CKSyncEngineStateSerializationClass     _CKSyncEngineStateSerializationClass
	CKSyncEngineStateSerializationClassOnce sync.Once
)

func getCKSyncEngineStateSerializationClass() _CKSyncEngineStateSerializationClass {
	CKSyncEngineStateSerializationClassOnce.Do(func() {
		CKSyncEngineStateSerializationClass = _CKSyncEngineStateSerializationClass{objc.GetClass("CKSyncEngineStateSerialization")}
	})
	return CKSyncEngineStateSerializationClass
}

type _CKSyncEngineStateSerializationClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineStateSerialization] class.
type ICKSyncEngineStateSerialization interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An opaque object that contains the serialized representation of a sync engine’s current state.


// An opaque object that contains the serialized representation of a sync engine’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineStateSerialization
type CKSyncEngineStateSerialization struct {
	objectivec.Object
}

// CKSyncEngineStateSerializationFrom constructs a [CKSyncEngineStateSerialization] from an unsafe.Pointer.
//
// An opaque object that contains the serialized representation of a sync engine’s current state.
func CKSyncEngineStateSerializationFrom(ptr unsafe.Pointer) CKSyncEngineStateSerialization {
	return CKSyncEngineStateSerialization{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineStateSerializationClass) Alloc() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineStateSerializationClass) New() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineStateSerialization) Init() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineStateSerialization) Autorelease() CKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineStateSerialization creates a new CKSyncEngineStateSerialization instance.
func NewCKSyncEngineStateSerialization() CKSyncEngineStateSerialization {
	return getCKSyncEngineStateSerializationClass().New()
}




