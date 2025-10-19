// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLProtectionSpace] class.
var uRLProtectionSpaceClass = _URLProtectionSpaceClass{objc.GetClass("NSURLProtectionSpace")}

type _URLProtectionSpaceClass struct {
	class objc.Class
}

// An interface definition for the [URLProtectionSpace] class.
type IURLProtectionSpace interface {
	objectivec.IObject
}

// A server or an area on a server, commonly referred to as a realm, that requires authentication. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return uRLProtectionSpaceClass.New()
}




