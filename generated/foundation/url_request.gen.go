// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLRequest] class.
var (
	uRLRequestClass     _URLRequestClass
	uRLRequestClassOnce sync.Once
)

func getURLRequestClass() _URLRequestClass {
	uRLRequestClassOnce.Do(func() {
		uRLRequestClass = _URLRequestClass{objc.GetClass("NSURLRequest")}
	})
	return uRLRequestClass
}

type _URLRequestClass struct {
	class objc.Class
}

// An interface definition for the [URLRequest] class.
type IURLRequest interface {
	objectivec.IObject
}

// A URL load request that is independent of protocol or URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest
type URLRequest struct {
	objectivec.Object
}

// URLRequestFrom constructs a [URLRequest] from an unsafe.Pointer.
//
// A URL load request that is independent of protocol or URL scheme.
func URLRequestFrom(ptr unsafe.Pointer) URLRequest {
	return URLRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLRequestClass) Alloc() URLRequest {
	rv := objc.Send[URLRequest](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLRequestClass) New() URLRequest {
	rv := objc.Send[URLRequest](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLRequest) Init() URLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLRequest) Autorelease() URLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLRequest creates a new URLRequest instance.
func NewURLRequest() URLRequest {
	return getURLRequestClass().New()
}




