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
	// properties:
	CookieAcceptPolicy() HTTPCookieAcceptPolicy
	SetCookieAcceptPolicy(value HTTPCookieAcceptPolicy)
	Cookies() []HTTPCookie /* primitive/slice/pointer. */
	IsSessionOnly() bool /* primitive/slice/pointer. */
	SetIsSessionOnly(value bool /* primitive/slice/pointer. */)
	// methods:
	CookiesForURL(URL IURL) []HTTPCookie /* primitive/slice/pointer. */
	DeleteCookie(cookie IHTTPCookie)
	GetCookiesForTaskCompletionHandler(task IURLSessionTask, completionHandler unsafe.Pointer)
	RemoveCookiesSinceDate(date IDate)
	SetCookie(cookie IHTTPCookie)
	SetCookiesForURLMainDocumentURL(cookies []HTTPCookie /* primitive/slice/pointer. */, URL IURL, mainDocumentURL IURL)
	SortedCookiesUsingDescriptors(sortOrder []SortDescriptor /* primitive/slice/pointer. */) []HTTPCookie /* primitive/slice/pointer. */
	StoreCookiesForTask(cookies []HTTPCookie /* primitive/slice/pointer. */, task IURLSessionTask)
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



// Returns the cookie storage instance for the container associated with the specified app group identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/sharedCookieStorage(forGroupContainerIdentifier:)
func (hc _HTTPCookieStorageClass) SharedCookieStorageForGroupContainerIdentifier(identifier IString) IHTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](objc.ID(hc.class), objc.Sel("sharedCookieStorageForGroupContainerIdentifier:"), identifier)
	return rv
}


// The shared cookie storage instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/shared
func (hc _HTTPCookieStorageClass) SharedHTTPCookieStorage() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](objc.ID(hc.class), objc.Sel("sharedHTTPCookieStorage"))
	return rv
}

// Returns all the cookie storage’s cookies that are sent to a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/cookies(for:)
func (h_ HTTPCookieStorage) CookiesForURL(URL IURL) []HTTPCookie /* primitive/slice/pointer. */ {
	rv := objc.Send[[]HTTPCookie](h_.ID, objc.Sel("cookiesForURL:"), URL)
	return rv
}


// Deletes the specified cookie from the cookie storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/deleteCookie(_:)
func (h_ HTTPCookieStorage) DeleteCookie(cookie IHTTPCookie) {
	objc.Send[objc.ID](h_.ID, objc.Sel("deleteCookie:"), cookie)
}


// Fetches cookies relevant to the specified task and passes them to the completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/getCookiesFor(_:completionHandler:)
func (h_ HTTPCookieStorage) GetCookiesForTaskCompletionHandler(task IURLSessionTask, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("getCookiesForTask:completionHandler:"), task, completionHandler)
}


// Removes cookies that were stored after a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/removeCookies(since:)
func (h_ HTTPCookieStorage) RemoveCookiesSinceDate(date IDate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("removeCookiesSinceDate:"), date)
}


// Stores a specified cookie in the cookie storage if the cookie accept policy permits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/setCookie(_:)
func (h_ HTTPCookieStorage) SetCookie(cookie IHTTPCookie) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookie:"), cookie)
}


// Adds an array of cookies to the cookie storage if the storage’s cookie acceptance policy permits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/setCookies(_:for:mainDocumentURL:)
func (h_ HTTPCookieStorage) SetCookiesForURLMainDocumentURL(cookies []HTTPCookie /* primitive/slice/pointer. */, URL IURL, mainDocumentURL IURL) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookies:forURL:mainDocumentURL:"), cookies, URL, mainDocumentURL)
}


// Returns all of the cookie storage’s cookies, sorted according to a given set of sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/sortedCookies(using:)
func (h_ HTTPCookieStorage) SortedCookiesUsingDescriptors(sortOrder []SortDescriptor /* primitive/slice/pointer. */) []HTTPCookie /* primitive/slice/pointer. */ {
	rv := objc.Send[[]HTTPCookie](h_.ID, objc.Sel("sortedCookiesUsingDescriptors:"), sortOrder)
	return rv
}


// Stores an array of cookies in the cookie storage, on behalf of the provided task, if the cookie accept policy permits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/storeCookies(_:for:)
func (h_ HTTPCookieStorage) StoreCookiesForTask(cookies []HTTPCookie /* primitive/slice/pointer. */, task IURLSessionTask) {
	objc.Send[objc.ID](h_.ID, objc.Sel("storeCookies:forTask:"), cookies, task)
}


// The cookie storage’s cookie accept policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/cookieAcceptPolicy
func (h_ HTTPCookieStorage) CookieAcceptPolicy() HTTPCookieAcceptPolicy {
	rv := objc.Send[HTTPCookieAcceptPolicy](h_.ID, objc.Sel("cookieAcceptPolicy"))
	return rv
}


// The cookie storage’s cookie accept policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/cookieAcceptPolicy
func (h_ HTTPCookieStorage) SetCookieAcceptPolicy(value HTTPCookieAcceptPolicy) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCookieAcceptPolicy:"), value)
}


// The cookie storage’s cookies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/cookies
func (h_ HTTPCookieStorage) Cookies() []HTTPCookie /* primitive/slice/pointer. */ {
	rv := objc.Send[[]HTTPCookie](h_.ID, objc.Sel("cookies"))
	return rv
}


// The shared cookie storage instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/shared
func (h_ HTTPCookieStorage) SharedHTTPCookieStorage() IHTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](h_.ID, objc.Sel("sharedHTTPCookieStorage"))
	return rv
}


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h_ HTTPCookieStorage) IsSessionOnly() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("isSessionOnly"))
	return rv
}


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h_ HTTPCookieStorage) SetIsSessionOnly(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsSessionOnly:"), value)
}



