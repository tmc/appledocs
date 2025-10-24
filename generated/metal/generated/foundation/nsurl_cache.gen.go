// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLCache] class.
var (
	URLCacheClass     _URLCacheClass
	URLCacheClassOnce sync.Once
)

func getURLCacheClass() _URLCacheClass {
	URLCacheClassOnce.Do(func() {
		URLCacheClass = _URLCacheClass{objc.GetClass("NSURLCache")}
	})
	return URLCacheClass
}

type _URLCacheClass struct {
	class objc.Class
}

// An interface definition for the [URLCache] class.
type IURLCache interface {
	objectivec.IObject
	// properties:
	CurrentDiskUsage() uint /* primitive/slice/pointer. */
	CurrentMemoryUsage() uint /* primitive/slice/pointer. */
	DiskCapacity() uint /* primitive/slice/pointer. */
	SetDiskCapacity(value uint /* primitive/slice/pointer. */)
	MemoryCapacity() uint /* primitive/slice/pointer. */
	SetMemoryCapacity(value uint /* primitive/slice/pointer. */)
	// methods:
	CachedResponseForRequest(request URLRequest /* not a class type */) ICachedURLResponse
	GetCachedResponseForDataTaskCompletionHandler(dataTask IURLSessionDataTask, completionHandler unsafe.Pointer)
	RemoveAllCachedResponses()
	RemoveCachedResponseForRequest(request URLRequest /* not a class type */)
	RemoveCachedResponseForDataTask(dataTask IURLSessionDataTask)
	RemoveCachedResponsesSinceDate(date IDate)
	StoreCachedResponseForRequest(cachedResponse ICachedURLResponse, request URLRequest /* not a class type */)
	StoreCachedResponseForDataTask(cachedResponse ICachedURLResponse, dataTask IURLSessionDataTask)
}

// An object that maps URL requests to cached response objects.
//
// The class implements the caching of responses to URL load requests, by mapping objects to objects. It provides a composite in-memory and on-disk cache, and lets you manipulate the sizes of both the in-memory and on-disk portions. You can also control the path where cache data is persistently stored.


// An object that maps URL requests to cached response objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache
type URLCache struct {
	objectivec.Object
}

// URLCacheFrom constructs a [URLCache] from an unsafe.Pointer.
//
// An object that maps URL requests to cached response objects.
func URLCacheFrom(ptr unsafe.Pointer) URLCache {
	return URLCache{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLCacheClass) Alloc() URLCache {
	rv := objc.Send[URLCache](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLCacheClass) New() URLCache {
	rv := objc.Send[URLCache](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLCache) Init() URLCache {
	rv := objc.Send[URLCache](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLCache) Autorelease() URLCache {
	rv := objc.Send[URLCache](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLCache creates a new URLCache instance.
func NewURLCache() URLCache {
	return getURLCacheClass().New()
}



// Creates a URL cache object with the specified memory and disk capacities, in the specified directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLCache/initWithMemoryCapacity:diskCapacity:directoryURL:
func NewURLCacheWithMemoryCapacityDiskCapacityDirectoryURL(memoryCapacity uint /* primitive/slice/pointer. */, diskCapacity uint /* primitive/slice/pointer. */, directoryURL IURL) URLCache {
	instance := getURLCacheClass().Alloc()
	rv := objc.Send[URLCache](instance.ID, objc.Sel("initWithMemoryCapacity:diskCapacity:directoryURL:"), memoryCapacity, diskCapacity, directoryURL)
	rv.Autorelease()
	return rv
}


// Creates a URL cache object with the specified values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/init(memoryCapacity:diskCapacity:diskPath:)
func NewURLCacheWithMemoryCapacityDiskCapacityDiskPath(memoryCapacity uint /* primitive/slice/pointer. */, diskCapacity uint /* primitive/slice/pointer. */, path IString) URLCache {
	instance := getURLCacheClass().Alloc()
	rv := objc.Send[URLCache](instance.ID, objc.Sel("initWithMemoryCapacity:diskCapacity:diskPath:"), memoryCapacity, diskCapacity, path)
	rv.Autorelease()
	return rv
}



// The shared URL cache instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/shared
func (uc _URLCacheClass) SharedURLCache() URLCache {
	rv := objc.Send[URLCache](objc.ID(uc.class), objc.Sel("sharedURLCache"))
	return rv
}

// Returns the cached URL response in the cache for the specified URL request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/cachedResponse(for:)
func (u_ URLCache) CachedResponseForRequest(request URLRequest /* not a class type */) ICachedURLResponse {
	rv := objc.Send[CachedURLResponse](u_.ID, objc.Sel("cachedResponseForRequest:"), request)
	return rv
}


// Gets the cached URL response for a data task, passing it to the provided completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/getCachedResponse(for:completionHandler:)
func (u_ URLCache) GetCachedResponseForDataTaskCompletionHandler(dataTask IURLSessionDataTask, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getCachedResponseForDataTask:completionHandler:"), dataTask, completionHandler)
}


// Clears the receiver’s cache, removing all stored cached URL responses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/removeAllCachedResponses()
func (u_ URLCache) RemoveAllCachedResponses() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllCachedResponses"))
}


// Removes the cached URL response for a specified URL request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/removeCachedResponse(for:)-1dh89
func (u_ URLCache) RemoveCachedResponseForRequest(request URLRequest /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCachedResponseForRequest:"), request)
}


// Removes the cached URL response for a specified data task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/removeCachedResponse(for:)-1zwp6
func (u_ URLCache) RemoveCachedResponseForDataTask(dataTask IURLSessionDataTask) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCachedResponseForDataTask:"), dataTask)
}


// Clears the given cache of any cached responses since the provided date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/removeCachedResponses(since:)
func (u_ URLCache) RemoveCachedResponsesSinceDate(date IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeCachedResponsesSinceDate:"), date)
}


// Stores a cached URL response for a specified request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/storeCachedResponse(_:for:)-7p7bl
func (u_ URLCache) StoreCachedResponseForRequest(cachedResponse ICachedURLResponse, request URLRequest /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("storeCachedResponse:forRequest:"), cachedResponse, request)
}


// Stores a cached URL response for a specified data task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/storeCachedResponse(_:for:)-8uq91
func (u_ URLCache) StoreCachedResponseForDataTask(cachedResponse ICachedURLResponse, dataTask IURLSessionDataTask) {
	objc.Send[objc.ID](u_.ID, objc.Sel("storeCachedResponse:forDataTask:"), cachedResponse, dataTask)
}


// The current size of the on-disk cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/currentDiskUsage
func (u_ URLCache) CurrentDiskUsage() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](u_.ID, objc.Sel("currentDiskUsage"))
	return rv
}


// The current size of the in-memory cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/currentMemoryUsage
func (u_ URLCache) CurrentMemoryUsage() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](u_.ID, objc.Sel("currentMemoryUsage"))
	return rv
}


// The capacity of the on-disk cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/diskCapacity
func (u_ URLCache) DiskCapacity() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](u_.ID, objc.Sel("diskCapacity"))
	return rv
}


// The capacity of the on-disk cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/diskCapacity
func (u_ URLCache) SetDiskCapacity(value uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDiskCapacity:"), value)
}


// The capacity of the in-memory cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/memoryCapacity
func (u_ URLCache) MemoryCapacity() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](u_.ID, objc.Sel("memoryCapacity"))
	return rv
}


// The capacity of the in-memory cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/memoryCapacity
func (u_ URLCache) SetMemoryCapacity(value uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMemoryCapacity:"), value)
}


// The shared URL cache instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/shared
func (u_ URLCache) SharedURLCache() IURLCache {
	rv := objc.Send[URLCache](u_.ID, objc.Sel("sharedURLCache"))
	return rv
}


// The shared URL cache instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/shared
func (u_ URLCache) SetSharedURLCache(value IURLCache) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSharedURLCache:"), value)
}


