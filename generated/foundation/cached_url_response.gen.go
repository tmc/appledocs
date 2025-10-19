// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CachedURLResponse] class.
var cachedURLResponseClass = _CachedURLResponseClass{objc.GetClass("NSCachedURLResponse")}

type _CachedURLResponseClass struct {
	class objc.Class
}

// An interface definition for the [CachedURLResponse] class.
type ICachedURLResponse interface {
	objectivec.IObject
}

// A cached response to a URL request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/CachedURLResponse

type CachedURLResponse struct {
	objectivec.Object
}

// CachedURLResponseFrom constructs a [CachedURLResponse] from an unsafe.Pointer.
//
// A cached response to a URL request.
func CachedURLResponseFrom(ptr unsafe.Pointer) CachedURLResponse {
	return CachedURLResponse{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (cc _CachedURLResponseClass) Alloc() CachedURLResponse {
	rv := objc.Send[CachedURLResponse](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _CachedURLResponseClass) New() CachedURLResponse {
	rv := objc.Send[CachedURLResponse](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CachedURLResponse) Init() CachedURLResponse {
	rv := objc.Send[CachedURLResponse](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CachedURLResponse) Autorelease() CachedURLResponse {
	rv := objc.Send[CachedURLResponse](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCachedURLResponse creates a new CachedURLResponse instance.
func NewCachedURLResponse() CachedURLResponse {
	return cachedURLResponseClass.New()
}




