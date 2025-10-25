// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLCache */


/* debug [class_header]: Header for NSURLCache */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLCache */
// An interface definition for the [URLCache] class.
type IURLCache interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLCache */
	// properties:
	CurrentDiskUsage() int
	SetCurrentDiskUsage(value int)
	CurrentMemoryUsage() int
	SetCurrentMemoryUsage(value int)
	DiskCapacity() int
	SetDiskCapacity(value int)
	MemoryCapacity() int
	SetMemoryCapacity(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLCache */
	// methods:
	CachedResponseForRequest(request IURLRequest) ICachedURLResponse
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLCache */
// Alloc allocates a new instance without initialization.
func (uc _URLCacheClass) Alloc() URLCache {
	rv := objc.Send[URLCache](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLCache */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLCache *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLCache */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLCache */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLCache */

// Returns the cached URL response in the cache for the specified URL request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCache/cachedResponse(for:)
func (u_ URLCache) CachedResponseForRequest(request IURLRequest) ICachedURLResponse {
	rv := objc.Send[CachedURLResponse](u_.ID, objc.Sel("cachedResponseForRequest:"), request)
	return rv
}/* debug [instance_methods/method]: CachedResponseForRequest */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLCache */

// The current size of the on-disk cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcache/currentdiskusage
func (u_ URLCache) CurrentDiskUsage() int {
	rv := objc.Send[int](u_.ID, objc.Sel("currentDiskUsage"))
	return rv
}/* debug [instance_properties/getter]: currentDiskUsage */


// The current size of the on-disk cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcache/currentdiskusage
func (u_ URLCache) SetCurrentDiskUsage(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCurrentDiskUsage:"), value)
}/* debug [instance_properties/setter]: currentDiskUsage */


// The current size of the in-memory cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcache/currentmemoryusage
func (u_ URLCache) CurrentMemoryUsage() int {
	rv := objc.Send[int](u_.ID, objc.Sel("currentMemoryUsage"))
	return rv
}/* debug [instance_properties/getter]: currentMemoryUsage */


// The current size of the in-memory cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcache/currentmemoryusage
func (u_ URLCache) SetCurrentMemoryUsage(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCurrentMemoryUsage:"), value)
}/* debug [instance_properties/setter]: currentMemoryUsage */


// The capacity of the on-disk cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcache/diskcapacity
func (u_ URLCache) DiskCapacity() int {
	rv := objc.Send[int](u_.ID, objc.Sel("diskCapacity"))
	return rv
}/* debug [instance_properties/getter]: diskCapacity */


// The capacity of the on-disk cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcache/diskcapacity
func (u_ URLCache) SetDiskCapacity(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDiskCapacity:"), value)
}/* debug [instance_properties/setter]: diskCapacity */


// The capacity of the in-memory cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcache/memorycapacity
func (u_ URLCache) MemoryCapacity() int {
	rv := objc.Send[int](u_.ID, objc.Sel("memoryCapacity"))
	return rv
}/* debug [instance_properties/getter]: memoryCapacity */


// The capacity of the in-memory cache, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcache/memorycapacity
func (u_ URLCache) SetMemoryCapacity(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMemoryCapacity:"), value)
}/* debug [instance_properties/setter]: memoryCapacity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLCache */



