// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLProtocol] class.
var (
	uRLProtocolClass     _URLProtocolClass
	uRLProtocolClassOnce sync.Once
)

func getURLProtocolClass() _URLProtocolClass {
	uRLProtocolClassOnce.Do(func() {
		uRLProtocolClass = _URLProtocolClass{objc.GetClass("NSURLProtocol")}
	})
	return uRLProtocolClass
}

type _URLProtocolClass struct {
	class objc.Class
}

// An interface definition for the [URLProtocol] class.
type IURLProtocol interface {
	objectivec.IObject
}

// An abstract class that handles the loading of protocol-specific URL data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol
type URLProtocol struct {
	objectivec.Object
}

// URLProtocolFrom constructs a [URLProtocol] from an unsafe.Pointer.
//
// An abstract class that handles the loading of protocol-specific URL data.
func URLProtocolFrom(ptr unsafe.Pointer) URLProtocol {
	return URLProtocol{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLProtocolClass) Alloc() URLProtocol {
	rv := objc.Send[URLProtocol](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLProtocolClass) New() URLProtocol {
	rv := objc.Send[URLProtocol](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLProtocol) Init() URLProtocol {
	rv := objc.Send[URLProtocol](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLProtocol) Autorelease() URLProtocol {
	rv := objc.Send[URLProtocol](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLProtocol creates a new URLProtocol instance.
func NewURLProtocol() URLProtocol {
	return getURLProtocolClass().New()
}




