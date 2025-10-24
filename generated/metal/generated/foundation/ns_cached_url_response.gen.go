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
	// properties:
	Data() IData
	Response() IURLResponse
	StoragePolicy() URLCacheStoragePolicy
	UserInfo() IDictionary
	// methods:
}

// A cached response to a URL request.
//
// A object provides the server’s response metadata in the form of a object, along with an object containing the actual cached content data. Its storage policy determines whether the response should be cached on disk, in memory, or not at all. Cached responses also contain a user info dictionary where you can store app-specific information about the cached item. The class stores and retrieves instances of .


// A cached response to a URL request.
//
// [Full Topic]
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



// Creates a cached URL response instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/CachedURLResponse/init(response:data:)
func NewCachedURLResponseWithResponseData(response IURLResponse, data IData) CachedURLResponse {
	instance := getCachedURLResponseClass().Alloc()
	rv := objc.Send[CachedURLResponse](instance.ID, objc.Sel("initWithResponse:data:"), response, data)
	rv.Autorelease()
	return rv
}


// Creates a cached URL response with a given server response, data, user-info dictionary, and storage policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/CachedURLResponse/init(response:data:userInfo:storagePolicy:)
func NewCachedURLResponseWithResponseDataUserInfoStoragePolicy(response IURLResponse, data IData, userInfo IDictionary, storagePolicy URLCacheStoragePolicy) CachedURLResponse {
	instance := getCachedURLResponseClass().Alloc()
	rv := objc.Send[CachedURLResponse](instance.ID, objc.Sel("initWithResponse:data:userInfo:storagePolicy:"), response, data, userInfo, storagePolicy)
	rv.Autorelease()
	return rv
}



// The cached response’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/CachedURLResponse/data
func (c_ CachedURLResponse) Data() IData {
	rv := objc.Send[Data](c_.ID, objc.Sel("data"))
	return rv
}


// The URL response object associated with the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/CachedURLResponse/response
func (c_ CachedURLResponse) Response() IURLResponse {
	rv := objc.Send[URLResponse](c_.ID, objc.Sel("response"))
	return rv
}


// The cached response’s storage policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/CachedURLResponse/storagePolicy
func (c_ CachedURLResponse) StoragePolicy() URLCacheStoragePolicy {
	rv := objc.Send[URLCacheStoragePolicy](c_.ID, objc.Sel("storagePolicy"))
	return rv
}


// The cached response’s user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/CachedURLResponse/userInfo
func (c_ CachedURLResponse) UserInfo() IDictionary {
	rv := objc.Send[Dictionary](c_.ID, objc.Sel("userInfo"))
	return rv
}


