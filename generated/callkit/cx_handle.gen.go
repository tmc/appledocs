// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CXHandle] class.
var (
	CXHandleClass     _CXHandleClass
	CXHandleClassOnce sync.Once
)

func getCXHandleClass() _CXHandleClass {
	CXHandleClassOnce.Do(func() {
		CXHandleClass = _CXHandleClass{objc.GetClass("CXHandle")}
	})
	return CXHandleClass
}

type _CXHandleClass struct {
	class objc.Class
}

// An interface definition for the [CXHandle] class.
type ICXHandle interface {
	objectivec.IObject
	IsEqualToHandle(handle unsafe.Pointer) bool
}

// A way to reach a call recipient, such as a phone number or email address.
//
// When the telephony provider receives an incoming call or the user starts an outgoing call, the other caller is identified by a object. For a caller identified by a phone number, the handle type is and the value is a sequence of digits. For a caller identified by an email address, the handle type is and the value is an email address. For a caller identified in any other way, the handle type is and the value typically follows some domain-specific format, such as a username, numeric ID, or URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle
type CXHandle struct {
	objectivec.Object
}

// CXHandleFrom constructs a [CXHandle] from an unsafe.Pointer.
//
// A way to reach a call recipient, such as a phone number or email address.
func CXHandleFrom(ptr unsafe.Pointer) CXHandle {
	return CXHandle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CXHandleClass) Alloc() CXHandle {
	rv := objc.Send[CXHandle](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXHandleClass) New() CXHandle {
	rv := objc.Send[CXHandle](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXHandle) Init() CXHandle {
	rv := objc.Send[CXHandle](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXHandle) Autorelease() CXHandle {
	rv := objc.Send[CXHandle](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXHandle creates a new CXHandle instance.
func NewCXHandle() CXHandle {
	return getCXHandleClass().New()
}




// Initializes a new handle of a given type with the specified value.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/init(type:value:)
func NewCXHandleWithTypeValue(type_ unsafe.Pointer, value string) CXHandle {
	instance := getCXHandleClass().Alloc()
	rv := objc.Send[CXHandle](instance.ID, objc.Sel("initWithType:value:"), type_, objc.String(value))
	rv.Autorelease()
	return rv
}


// Returns a Boolean value that indicates whether a given handle is equal to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/isEqualToHandle:
func (c_ CXHandle) IsEqualToHandle(handle unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToHandle:"), handle)
	return rv
}

// The type of the handle.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/type
func (c_ CXHandle) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("type"))
	return rv
}

// The value of the handle.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/value
func (c_ CXHandle) Value() string {
	rv := objc.Send[string](c_.ID, objc.Sel("value"))
	return rv
}


