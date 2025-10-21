// Code generated from Apple documentation for PushKit. DO NOT EDIT.

package pushkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PushCredentials] class.
var (
	PushCredentialsClass     _PushCredentialsClass
	PushCredentialsClassOnce sync.Once
)

func getPushCredentialsClass() _PushCredentialsClass {
	PushCredentialsClassOnce.Do(func() {
		PushCredentialsClass = _PushCredentialsClass{objc.GetClass("PKPushCredentials")}
	})
	return PushCredentialsClass
}

type _PushCredentialsClass struct {
	class objc.Class
}

// An interface definition for the [PushCredentials] class.
type IPushCredentials interface {
	objectivec.IObject
}

// An object that encapsulates the device token you use to deliver push notifications to your app.
//
// When registering your app’s push types, PushKit creates a object for each type your app supports and delivers it to your delegate’s method. Don’t create objects yourself.
//
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushCredentials
type PushCredentials struct {
	objectivec.Object
}

// PushCredentialsFrom constructs a [PushCredentials] from an unsafe.Pointer.
//
// An object that encapsulates the device token you use to deliver push notifications to your app.
func PushCredentialsFrom(ptr unsafe.Pointer) PushCredentials {
	return PushCredentials{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PushCredentialsClass) Alloc() PushCredentials {
	rv := objc.Send[PushCredentials](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PushCredentialsClass) New() PushCredentials {
	rv := objc.Send[PushCredentials](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PushCredentials) Init() PushCredentials {
	rv := objc.Send[PushCredentials](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PushCredentials) Autorelease() PushCredentials {
	rv := objc.Send[PushCredentials](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPushCredentials creates a new PushCredentials instance.
func NewPushCredentials() PushCredentials {
	return getPushCredentialsClass().New()
}


// A unique device token to use when sending push notifications to the current device.
//
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushCredentials/token
func (p_ PushCredentials) Token() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("token"))
	return rv
}

// The push type constant associated with the token.
//
// [Full Topic]: https://developer.apple.com/documentation/pushkit/pkpushcredentials/type
func (p_ PushCredentials) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The push type constant associated with the token.

//
// [Full Topic]: https://developer.apple.com/documentation/pushkit/pkpushcredentials/type
func (p_ PushCredentials) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), value)
}



