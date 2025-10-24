// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCachedURLResponse */


/* debug [class_header]: Header for NSCachedURLResponse */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CachedURLResponse */
// An interface definition for the [CachedURLResponse] class.
type ICachedURLResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CachedURLResponse */
	// properties:
	StoragePolicy() URLCacheStoragePolicy /* not a class type */
	Data() IData
	SetData(value IData)
	Response() IURLResponse
	SetResponse(value IURLResponse)
	UserInfo() objectivec.IObject
	SetUserInfo(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CachedURLResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CachedURLResponse */
// Alloc allocates a new instance without initialization.
func (cc _CachedURLResponseClass) Alloc() CachedURLResponse {
	rv := objc.Send[CachedURLResponse](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CachedURLResponse */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CachedURLResponse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CachedURLResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CachedURLResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CachedURLResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CachedURLResponse */

// The cached response’s storage policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/CachedURLResponse/storagePolicy
func (c_ CachedURLResponse) StoragePolicy() URLCacheStoragePolicy /* not a class type */ {
	rv := objc.Send[URLCacheStoragePolicy](c_.ID, objc.Sel("storagePolicy"))
	return rv
}/* debug [instance_properties/getter]: storagePolicy */


// The cached response’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/cachedurlresponse/data
func (c_ CachedURLResponse) Data() IData {
	rv := objc.Send[Data](c_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The cached response’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/cachedurlresponse/data
func (c_ CachedURLResponse) SetData(value IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// The URL response object associated with the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/cachedurlresponse/response
func (c_ CachedURLResponse) Response() IURLResponse {
	rv := objc.Send[URLResponse](c_.ID, objc.Sel("response"))
	return rv
}/* debug [instance_properties/getter]: response */


// The URL response object associated with the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/cachedurlresponse/response
func (c_ CachedURLResponse) SetResponse(value IURLResponse) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResponse:"), value)
}/* debug [instance_properties/setter]: response */


// The cached response’s user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/cachedurlresponse/userinfo
func (c_ CachedURLResponse) UserInfo() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// The cached response’s user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/cachedurlresponse/userinfo
func (c_ CachedURLResponse) SetUserInfo(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCachedURLResponse */



