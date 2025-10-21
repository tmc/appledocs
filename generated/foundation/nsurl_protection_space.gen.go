// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLProtectionSpace] class.
var (
	URLProtectionSpaceClass     _URLProtectionSpaceClass
	URLProtectionSpaceClassOnce sync.Once
)

func getURLProtectionSpaceClass() _URLProtectionSpaceClass {
	URLProtectionSpaceClassOnce.Do(func() {
		URLProtectionSpaceClass = _URLProtectionSpaceClass{objc.GetClass("NSURLProtectionSpace")}
	})
	return URLProtectionSpaceClass
}

type _URLProtectionSpaceClass struct {
	class objc.Class
}

// An interface definition for the [URLProtectionSpace] class.
type IURLProtectionSpace interface {
	objectivec.IObject
}

// A server or an area on a server, commonly referred to as a realm, that requires authentication.
//
// A protection space defines a series of matching constraints that determine which credential should be provided. For example, if a request provides your delegate with a object that requests a client username and password, your app should provide the correct username and password for the particular host, port, protocol, and realm, as specified in the challenge’s protection space.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace
type URLProtectionSpace struct {
	objectivec.Object
}

// URLProtectionSpaceFrom constructs a [URLProtectionSpace] from an unsafe.Pointer.
//
// A server or an area on a server, commonly referred to as a realm, that requires authentication.
func URLProtectionSpaceFrom(ptr unsafe.Pointer) URLProtectionSpace {
	return URLProtectionSpace{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLProtectionSpaceClass) Alloc() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLProtectionSpaceClass) New() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLProtectionSpace) Init() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLProtectionSpace) Autorelease() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLProtectionSpace creates a new URLProtectionSpace instance.
func NewURLProtectionSpace() URLProtectionSpace {
	return getURLProtectionSpaceClass().New()
}

// The authentication method used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/authenticationMethod
func (u_ URLProtectionSpace) AuthenticationMethod() string {
	rv := objc.Send[string](u_.ID, objc.Sel("authenticationMethod"))
	return rv
}

// The acceptable certificate-issuing authorities for client certificate authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/distinguishedNames
func (u_ URLProtectionSpace) DistinguishedNames() []Data {
	rv := objc.Send[[]Data](u_.ID, objc.Sel("distinguishedNames"))
	return rv
}

// The receiver’s proxy type.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/proxyType
func (u_ URLProtectionSpace) ProxyType() string {
	rv := objc.Send[string](u_.ID, objc.Sel("proxyType"))
	return rv
}
