// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CXCallUpdate] class.
var (
	CXCallUpdateClass     _CXCallUpdateClass
	CXCallUpdateClassOnce sync.Once
)

func getCXCallUpdateClass() _CXCallUpdateClass {
	CXCallUpdateClassOnce.Do(func() {
		CXCallUpdateClass = _CXCallUpdateClass{objc.GetClass("CXCallUpdate")}
	})
	return CXCallUpdateClass
}

type _CXCallUpdateClass struct {
	class objc.Class
}

// An interface definition for the [CXCallUpdate] class.
type ICXCallUpdate interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An encapsulation of new and changed information about a call.
//
// objects are used by the system to communicate changes to calls over time. Not every property on a object must be set each time, as each object includes only new and changed information. For example, when a call is started, only some properties may be known and included in the first object sent to the system, such as . Later in the same call, other properties may change; for example, a call may be upgraded from audio only to audio and video, which would be reflected by a new object with its property set to . When an incoming call is received, you construct a object specifying a and pass that to the method to notify the telephony provider. When an active call is updated, you construct a object specifying any updated information and pass that to the method. For example, if a user changes their contact information during a call, you could notify the telephony provider of this change using a new object with the new value set to its property.


// An encapsulation of new and changed information about a call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate
type CXCallUpdate struct {
	objectivec.Object
}

// CXCallUpdateFrom constructs a [CXCallUpdate] from an unsafe.Pointer.
//
// An encapsulation of new and changed information about a call.
func CXCallUpdateFrom(ptr unsafe.Pointer) CXCallUpdate {
	return CXCallUpdate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CXCallUpdateClass) Alloc() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXCallUpdateClass) New() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallUpdate) Init() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallUpdate) Autorelease() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallUpdate creates a new CXCallUpdate instance.
func NewCXCallUpdate() CXCallUpdate {
	return getCXCallUpdateClass().New()
}



