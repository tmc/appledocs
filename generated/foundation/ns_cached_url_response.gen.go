// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CachedURLResponse] class.
var (
	CachedURLResponseClass     _CachedURLResponseClass
	CachedURLResponseClassOnce sync.Once
)

func getCachedURLResponseClass() _CachedURLResponseClass {
	CachedURLResponseClassOnce.Do(func() {
		CachedURLResponseClass = _CachedURLResponseClass{objc.GetClass("NSCachedURLResponse")}
	})
	return CachedURLResponseClass
}

type _CachedURLResponseClass struct {
	class objc.Class
}

// An interface definition for the [CachedURLResponse] class.
type ICachedURLResponse interface {
	objectivec.IObject
}

// A cached response to a URL request.
//
// A object provides the server’s response metadata in the form of a object, along with an object containing the actual cached content data. Its storage policy determines whether the response should be cached on disk, in memory, or not at all. Cached responses also contain a user info dictionary where you can store app-specific information about the cached item. The class stores and retrieves instances of .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getCachedURLResponseClass().New()
}
