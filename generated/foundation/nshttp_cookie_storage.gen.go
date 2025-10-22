// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HTTPCookieStorage] class.
type IHTTPCookieStorage interface {
	objectivec.IObject
	IsSessionOnly() bool
	SetIsSessionOnly(value bool)
	CookieAcceptPolicy() unsafe.Pointer
	SetCookieAcceptPolicy(value unsafe.Pointer)
	Cookies() NSHTTPCookie
	SetCookies(value IHTTPCookie)
}

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

// Alloc allocates a new instance without initialization.
func (hc _HTTPCookieStorageClass) Alloc() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The shared cookie storage instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/shared

func (hc _HTTPCookieStorageClass) SharedHTTPCookieStorage() HTTPCookieStorage {
	rv := objc.Send[NSHTTPCookieStorage](objc.ID(hc.class), objc.Sel("sharedHTTPCookieStorage"))
	return rv
}

// The shared cookie storage instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/shared

func (h_ HTTPCookieStorage) SharedHTTPCookieStorage() NSHTTPCookieStorage {
	rv := objc.Send[NSHTTPCookieStorage](h_.ID, objc.Sel("sharedHTTPCookieStorage"))
	return rv
}


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly

func (h_ HTTPCookieStorage) IsSessionOnly() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isSessionOnly"))
	return rv
}


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly

func (h_ HTTPCookieStorage) SetIsSessionOnly(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsSessionOnly:"), value)
}


// The cookie storage’s cookie accept policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookiestorage/cookieacceptpolicy

func (h_ HTTPCookieStorage) CookieAcceptPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("cookieAcceptPolicy"))
	return rv
}


// The cookie storage’s cookie accept policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookiestorage/cookieacceptpolicy

func (h_ HTTPCookieStorage) SetCookieAcceptPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookieAcceptPolicy:"), value)
}


// The cookie storage’s cookies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookiestorage/cookies

func (h_ HTTPCookieStorage) Cookies() NSHTTPCookie {
	rv := objc.Send[NSHTTPCookie](h_.ID, objc.Sel("cookies"))
	return rv
}


// The cookie storage’s cookies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookiestorage/cookies

func (h_ HTTPCookieStorage) SetCookies(value IHTTPCookie) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookies:"), value)
}



