// Code generated from Apple documentation for PushKit. DO NOT EDIT.

package pushkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PushPayload] class.
var (
	PushPayloadClass     _PushPayloadClass
	PushPayloadClassOnce sync.Once
)

func getPushPayloadClass() _PushPayloadClass {
	PushPayloadClassOnce.Do(func() {
		PushPayloadClass = _PushPayloadClass{objc.GetClass("PKPushPayload")}
	})
	return PushPayloadClass
}

type _PushPayloadClass struct {
	class objc.Class
}

// An interface definition for the [PushPayload] class.
type IPushPayload interface {
	objectivec.IObject
	// properties:
	DictionaryPayload() objc.IObject /* cross-framework: NSDictionary */
	Type() objc.IObject /* cross-framework: PushType */
	SetType(value objc.IObject /* cross-framework: PushType */)
	// methods:
}

// An object that contains information about a received PushKit notification.


// An object that contains information about a received PushKit notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushPayload
type PushPayload struct {
	objectivec.Object
}

// PushPayloadFrom constructs a [PushPayload] from an unsafe.Pointer.
//
// An object that contains information about a received PushKit notification.
func PushPayloadFrom(ptr unsafe.Pointer) PushPayload {
	return PushPayload{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PushPayloadClass) Alloc() PushPayload {
	rv := objc.Send[PushPayload](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PushPayloadClass) New() PushPayload {
	rv := objc.Send[PushPayload](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PushPayload) Init() PushPayload {
	rv := objc.Send[PushPayload](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PushPayload) Autorelease() PushPayload {
	rv := objc.Send[PushPayload](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPushPayload creates a new PushPayload instance.
func NewPushPayload() PushPayload {
	return getPushPayloadClass().New()
}



// The contents of the received payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushPayload/dictionaryPayload
func (p_ PushPayload) DictionaryPayload() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("dictionaryPayload"))
	return rv
}


// The type value indicating how to interpret the payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pushkit/pkpushpayload/type
func (p_ PushPayload) Type() objc.IObject /* cross-framework: PushType */ {
	rv := objc.Send[PushType](p_.ID, objc.Sel("type"))
	return rv
}


// The type value indicating how to interpret the payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pushkit/pkpushpayload/type
func (p_ PushPayload) SetType(value objc.IObject /* cross-framework: PushType */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), value)
}



