// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLConnection] class.
var uRLConnectionClass = _URLConnectionClass{objc.GetClass("NSURLConnection")}

type _URLConnectionClass struct {
	class objc.Class
}

// An interface definition for the [URLConnection] class.
type IURLConnection interface {
	objectivec.IObject
}

// An object that enables you to start and stop URL requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection

type URLConnection struct {
	objectivec.Object
}

// URLConnectionFrom constructs a [URLConnection] from an unsafe.Pointer.
//
// An object that enables you to start and stop URL requests.
func URLConnectionFrom(ptr unsafe.Pointer) URLConnection {
	return URLConnection{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (uc _URLConnectionClass) Alloc() URLConnection {
	rv := objc.Send[URLConnection](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (uc _URLConnectionClass) New() URLConnection {
	rv := objc.Send[URLConnection](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLConnection) Init() URLConnection {
	rv := objc.Send[URLConnection](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLConnection) Autorelease() URLConnection {
	rv := objc.Send[URLConnection](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLConnection creates a new URLConnection instance.
func NewURLConnection() URLConnection {
	return uRLConnectionClass.New()
}




