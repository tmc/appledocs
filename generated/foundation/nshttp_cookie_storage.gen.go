// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSHTTPCookieStorage */


/* debug [class_header]: Header for NSHTTPCookieStorage */
// The class instance for the [HTTPCookieStorage] class.
var (
	HTTPCookieStorageClass     _HTTPCookieStorageClass
	HTTPCookieStorageClassOnce sync.Once
)

func getHTTPCookieStorageClass() _HTTPCookieStorageClass {
	HTTPCookieStorageClassOnce.Do(func() {
		HTTPCookieStorageClass = _HTTPCookieStorageClass{objc.GetClass("NSHTTPCookieStorage")}
	})
	return HTTPCookieStorageClass
}

type _HTTPCookieStorageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HTTPCookieStorage */
// An interface definition for the [HTTPCookieStorage] class.
type IHTTPCookieStorage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HTTPCookieStorage */
	// properties:
	IsSessionOnly() bool
	SetIsSessionOnly(value bool)
	CookieAcceptPolicy() objectivec.IObject
	SetCookieAcceptPolicy(value objectivec.IObject)
	Cookies() IHTTPCookie
	SetCookies(value IHTTPCookie)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HTTPCookieStorage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HTTPCookieStorage */
// Alloc allocates a new instance without initialization.
func (hc _HTTPCookieStorageClass) Alloc() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HTTPCookieStorageClass) New() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HTTPCookieStorage) Init() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HTTPCookieStorage) Autorelease() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHTTPCookieStorage creates a new HTTPCookieStorage instance.
func NewHTTPCookieStorage() HTTPCookieStorage {
	return getHTTPCookieStorageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HTTPCookieStorage */
// A container that manages the storage of cookies.
//
// Each stored cookie is represented by an instance of the class.


// A container that manages the storage of cookies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage
type HTTPCookieStorage struct {
	objectivec.Object
}

// HTTPCookieStorageFrom constructs a [HTTPCookieStorage] from an unsafe.Pointer.
//
// A container that manages the storage of cookies.
func HTTPCookieStorageFrom(ptr unsafe.Pointer) HTTPCookieStorage {
	return HTTPCookieStorage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HTTPCookieStorage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HTTPCookieStorage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HTTPCookieStorage */

// The shared cookie storage instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/shared
func (hc _HTTPCookieStorageClass) SharedHTTPCookieStorage() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](objc.ID(hc.class), objc.Sel("sharedHTTPCookieStorage"))
	return rv
}/* debug [class_properties_class/property]: sharedHTTPCookieStorage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HTTPCookieStorage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HTTPCookieStorage */

// The shared cookie storage instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/shared
func (h_ HTTPCookieStorage) SharedHTTPCookieStorage() IHTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](h_.ID, objc.Sel("sharedHTTPCookieStorage"))
	return rv
}/* debug [instance_properties/getter]: sharedHTTPCookieStorage */


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h_ HTTPCookieStorage) IsSessionOnly() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isSessionOnly"))
	return rv
}/* debug [instance_properties/getter]: isSessionOnly */


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h_ HTTPCookieStorage) SetIsSessionOnly(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsSessionOnly:"), value)
}/* debug [instance_properties/setter]: isSessionOnly */


// The cookie storage’s cookie accept policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookiestorage/cookieacceptpolicy
func (h_ HTTPCookieStorage) CookieAcceptPolicy() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](h_.ID, objc.Sel("cookieAcceptPolicy"))
	return rv
}/* debug [instance_properties/getter]: cookieAcceptPolicy */


// The cookie storage’s cookie accept policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookiestorage/cookieacceptpolicy
func (h_ HTTPCookieStorage) SetCookieAcceptPolicy(value objectivec.IObject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookieAcceptPolicy:"), value)
}/* debug [instance_properties/setter]: cookieAcceptPolicy */


// The cookie storage’s cookies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookiestorage/cookies
func (h_ HTTPCookieStorage) Cookies() IHTTPCookie {
	rv := objc.Send[HTTPCookie](h_.ID, objc.Sel("cookies"))
	return rv
}/* debug [instance_properties/getter]: cookies */


// The cookie storage’s cookies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookiestorage/cookies
func (h_ HTTPCookieStorage) SetCookies(value IHTTPCookie) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookies:"), value)
}/* debug [instance_properties/setter]: cookies */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSHTTPCookieStorage */



