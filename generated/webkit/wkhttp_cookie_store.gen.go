// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HTTPCookieStore] class.
type IHTTPCookieStore interface {
	objectivec.IObject
	SetCookieCompletionHandler(cookie unsafe.Pointer, completionHandler func())
}

// An object that manages the HTTP cookies associated with a particular web view.
//
// Use a to specify the initial cookies for your webpages, and to manage cookies for your web content. For example, you might use this object to delete the cookie for the current session when the user logs out. To detect when the webpage changes a cookie, install a cookie observer using the method. You don’t create a object directly. Instead, retrieve this object from the object in your web view’s configuration object.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HTTPCookieStoreClass) Alloc() HTTPCookieStore {
	rv := objc.Send[HTTPCookieStore](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Adds a cookie to the cookie store.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/setCookie(_:completionHandler:)
func (h_ HTTPCookieStore) SetCookieCompletionHandler(cookie unsafe.Pointer, completionHandler func()) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookie:completionHandler:"), cookie, completionHandler)
}



