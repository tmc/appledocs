// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKHTTPCookieStore */

/* debug [class_header]: Header for WKHTTPCookieStore */
// The class instance for the [HTTPCookieStore] class.
var (
	HTTPCookieStoreClass     _HTTPCookieStoreClass
	HTTPCookieStoreClassOnce sync.Once
)

func getHTTPCookieStoreClass() _HTTPCookieStoreClass {
	HTTPCookieStoreClassOnce.Do(func() {
		HTTPCookieStoreClass = _HTTPCookieStoreClass{objc.GetClass("WKHTTPCookieStore")}
	})
	return HTTPCookieStoreClass
}

type _HTTPCookieStoreClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for HTTPCookieStore */
// An interface definition for the [HTTPCookieStore] class.
type IHTTPCookieStore interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for HTTPCookieStore */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for HTTPCookieStore */
	// methods:
	AddObserver(observer unsafe.Pointer)
	DeleteCookieCompletionHandler(cookie foundation.HTTPCookie, completionHandler func())
	GetAllCookies(completionHandler func([]unsafe.Pointer))
	GetCookiePolicy(completionHandler func(unsafe.Pointer))
	RemoveObserver(observer unsafe.Pointer)
	SetCookieCompletionHandler(cookie foundation.HTTPCookie, completionHandler func())
	SetCookiePolicyCompletionHandler(policy CookiePolicy, completionHandler func())
	SetCookiesCompletionHandler(cookies []foundation.HTTPCookie, completionHandler func())
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for HTTPCookieStore */
// Alloc allocates a new instance without initialization.
func (hc _HTTPCookieStoreClass) Alloc() HTTPCookieStore {
	rv := objc.Send[HTTPCookieStore](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HTTPCookieStoreClass) New() HTTPCookieStore {
	rv := objc.Send[HTTPCookieStore](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HTTPCookieStore) Init() HTTPCookieStore {
	rv := objc.Send[HTTPCookieStore](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HTTPCookieStore) Autorelease() HTTPCookieStore {
	rv := objc.Send[HTTPCookieStore](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHTTPCookieStore creates a new HTTPCookieStore instance.
func NewHTTPCookieStore() HTTPCookieStore {
	return getHTTPCookieStoreClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for HTTPCookieStore */
// An object that manages the HTTP cookies associated with a particular web view.
//
// Use a to specify the initial cookies for your webpages, and to manage cookies for your web content. For example, you might use this object to delete the cookie for the current session when the user logs out. To detect when the webpage changes a cookie, install a cookie observer using the method. You don’t create a object directly. Instead, retrieve this object from the object in your web view’s configuration object.

// An object that manages the HTTP cookies associated with a particular web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore
type HTTPCookieStore struct {
	objectivec.Object
}

// HTTPCookieStoreFrom constructs a [HTTPCookieStore] from an unsafe.Pointer.
//
// An object that manages the HTTP cookies associated with a particular web view.
func HTTPCookieStoreFrom(ptr unsafe.Pointer) HTTPCookieStore {
	return HTTPCookieStore{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for HTTPCookieStore */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for HTTPCookieStore */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for HTTPCookieStore */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for HTTPCookieStore */

// Adds an observer to the cookie store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/add(_:)
func (h_ HTTPCookieStore) AddObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addObserver:"), observer)
} /* debug [instance_methods/method]: AddObserver */

// Deletes the specified cookie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/delete(_:completionHandler:)
func (h_ HTTPCookieStore) DeleteCookieCompletionHandler(cookie foundation.HTTPCookie, completionHandler func()) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteCookie:completionHandler:"), cookie, completionHandler)
} /* debug [instance_methods/method]: DeleteCookieCompletionHandler */

// Fetches all stored cookies asynchronously and delivers them to the specified completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/getAllCookies(_:)
func (h_ HTTPCookieStore) GetAllCookies(completionHandler func([]unsafe.Pointer)) {
	objc.Send[objc.ID](h_.ID, objc.Sel("getAllCookies:"), completionHandler)
} /* debug [instance_methods/method]: GetAllCookies */

// Returns a cookie policy that indicates whether the cookie store allows cookie storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/getCookiePolicy(_:)
func (h_ HTTPCookieStore) GetCookiePolicy(completionHandler func(unsafe.Pointer)) {
	objc.Send[objc.ID](h_.ID, objc.Sel("getCookiePolicy:"), completionHandler)
} /* debug [instance_methods/method]: GetCookiePolicy */

// Removes an observer from the cookie store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/remove(_:)
func (h_ HTTPCookieStore) RemoveObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("removeObserver:"), observer)
} /* debug [instance_methods/method]: RemoveObserver */

// Adds a cookie to the cookie store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/setCookie(_:completionHandler:)
func (h_ HTTPCookieStore) SetCookieCompletionHandler(cookie foundation.HTTPCookie, completionHandler func()) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookie:completionHandler:"), cookie, completionHandler)
} /* debug [instance_methods/method]: SetCookieCompletionHandler */

// Sets a cookie policy that indicates whether the cookie store allows cookie storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/setCookiePolicy(_:completionHandler:)
func (h_ HTTPCookieStore) SetCookiePolicyCompletionHandler(policy CookiePolicy, completionHandler func()) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookiePolicy:completionHandler:"), policy, completionHandler)
} /* debug [instance_methods/method]: SetCookiePolicyCompletionHandler */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/setCookies(_:completionHandler:)
func (h_ HTTPCookieStore) SetCookiesCompletionHandler(cookies []foundation.HTTPCookie, completionHandler func()) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookies:completionHandler:"), cookies, completionHandler)
} /* debug [instance_methods/method]: SetCookiesCompletionHandler */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for HTTPCookieStore */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKHTTPCookieStore */
