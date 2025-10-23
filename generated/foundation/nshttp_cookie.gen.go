// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HTTPCookie] class.
var (
	HTTPCookieClass     _HTTPCookieClass
	HTTPCookieClassOnce sync.Once
)

func getHTTPCookieClass() _HTTPCookieClass {
	HTTPCookieClassOnce.Do(func() {
		HTTPCookieClass = _HTTPCookieClass{objc.GetClass("NSHTTPCookie")}
	})
	return HTTPCookieClass
}

type _HTTPCookieClass struct {
	class objc.Class
}

// An interface definition for the [HTTPCookie] class.
type IHTTPCookie interface {
	objectivec.IObject
	// properties:
	Comment() IString
	CommentURL() IURL
	Domain() IString
	ExpiresDate() IDate
	HTTPOnly() bool /* primitive/slice/pointer. */
	Secure() bool /* primitive/slice/pointer. */
	SessionOnly() bool /* primitive/slice/pointer. */
	Name() IString
	Path() IString
	PortList() []Number /* primitive/slice/pointer. */
	Properties() IDictionary /* already interface */
	SameSitePolicy() objc.IObject /* cross-framework: HTTPCookieStringPolicy */
	Value() IString
	Version() uint /* primitive/slice/pointer. */
	IsHTTPOnly() bool /* primitive/slice/pointer. */
	SetIsHTTPOnly(value bool /* primitive/slice/pointer. */)
	IsSecure() bool /* primitive/slice/pointer. */
	SetIsSecure(value bool /* primitive/slice/pointer. */)
	IsSessionOnly() bool /* primitive/slice/pointer. */
	SetIsSessionOnly(value bool /* primitive/slice/pointer. */)
	// methods:
}

// A representation of an HTTP cookie.
//
// An object is immutable, initialized from a dictionary that contains the attributes of the cookie. This class supports two different cookie versions: Version 0: The original cookie format defined by Netscape. Most cookies are in this format. Version 1: The cookie format defined in , HTTP State Management Mechanism.


// A representation of an HTTP cookie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie
type HTTPCookie struct {
	objectivec.Object
}

// HTTPCookieFrom constructs a [HTTPCookie] from an unsafe.Pointer.
//
// A representation of an HTTP cookie.
func HTTPCookieFrom(ptr unsafe.Pointer) HTTPCookie {
	return HTTPCookie{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HTTPCookieClass) Alloc() HTTPCookie {
	rv := objc.Send[HTTPCookie](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HTTPCookieClass) New() HTTPCookie {
	rv := objc.Send[HTTPCookie](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HTTPCookie) Init() HTTPCookie {
	rv := objc.Send[HTTPCookie](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HTTPCookie) Autorelease() HTTPCookie {
	rv := objc.Send[HTTPCookie](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHTTPCookie creates a new HTTPCookie instance.
func NewHTTPCookie() HTTPCookie {
	return getHTTPCookieClass().New()
}



// Initializes an HTTP cookie object with the given cookie properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/init(properties:)
func NewHTTPCookieWithProperties(properties IDictionary /* already interface */) HTTPCookie {
	instance := getHTTPCookieClass().Alloc()
	rv := objc.Send[HTTPCookie](instance.ID, objc.Sel("initWithProperties:"), properties)
	rv.Autorelease()
	return rv
}



// Creates an array of HTTP cookies that corresponds to the provided response header fields for the provided URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/cookies(withResponseHeaderFields:for:)
func (hc _HTTPCookieClass) CookiesWithResponseHeaderFieldsForURL(headerFields IDictionary /* already interface */, URL IURL) []HTTPCookie /* primitive/slice/pointer. */ {
	rv := objc.Send[[]HTTPCookie](objc.ID(hc.class), objc.Sel("cookiesWithResponseHeaderFields:forURL:"), headerFields, URL)
	return rv
}


// Converts an array of cookies to a dictionary of header fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/requestHeaderFields(with:)
func (hc _HTTPCookieClass) RequestHeaderFieldsWithCookies(cookies []HTTPCookie /* primitive/slice/pointer. */) IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](objc.ID(hc.class), objc.Sel("requestHeaderFieldsWithCookies:"), cookies)
	return rv
}


// Creates and initializes an HTTP cookie object using the provided properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHTTPCookie/cookieWithProperties:
func (hc _HTTPCookieClass) CookieWithProperties(properties IDictionary /* already interface */) IHTTPCookie {
	rv := objc.Send[HTTPCookie](objc.ID(hc.class), objc.Sel("cookieWithProperties:"), properties)
	return rv
}


// The cookie’s comment string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/comment
func (h_ HTTPCookie) Comment() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("comment"))
	return rv
}


// The cookie’s comment URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/commentURL
func (h_ HTTPCookie) CommentURL() IURL {
	rv := objc.Send[URL](h_.ID, objc.Sel("commentURL"))
	return rv
}


// The domain of the cookie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/domain
func (h_ HTTPCookie) Domain() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("domain"))
	return rv
}


// The cookie’s expiration date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/expiresDate
func (h_ HTTPCookie) ExpiresDate() IDate {
	rv := objc.Send[Date](h_.ID, objc.Sel("expiresDate"))
	return rv
}


// A Boolean value that indicates whether the cookie should only be sent to HTTP servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/isHTTPOnly
func (h_ HTTPCookie) HTTPOnly() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("HTTPOnly"))
	return rv
}


// A Boolean value that indicates whether the cookie may only be sent over secure channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/isSecure
func (h_ HTTPCookie) Secure() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("secure"))
	return rv
}


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/isSessionOnly
func (h_ HTTPCookie) SessionOnly() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("sessionOnly"))
	return rv
}


// The cookie’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/name
func (h_ HTTPCookie) Name() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("name"))
	return rv
}


// The cookie’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/path
func (h_ HTTPCookie) Path() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("path"))
	return rv
}


// The cookie’s port list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/portList
func (h_ HTTPCookie) PortList() []Number /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Number](h_.ID, objc.Sel("portList"))
	return rv
}


// The cookie’s properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/properties
func (h_ HTTPCookie) Properties() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](h_.ID, objc.Sel("properties"))
	return rv
}


// A Boolean value that indicates whether to restrict the cookie to requests sent back to the same site that created it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/sameSitePolicy
func (h_ HTTPCookie) SameSitePolicy() objc.IObject /* cross-framework: HTTPCookieStringPolicy */ {
	rv := objc.Send[HTTPCookieStringPolicy](h_.ID, objc.Sel("sameSitePolicy"))
	return rv
}


// The cookie’s string value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/value
func (h_ HTTPCookie) Value() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("value"))
	return rv
}


// The cookie’s version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/version
func (h_ HTTPCookie) Version() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](h_.ID, objc.Sel("version"))
	return rv
}


// A Boolean value that indicates whether the cookie should only be sent to HTTP servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/ishttponly
func (h_ HTTPCookie) IsHTTPOnly() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("isHTTPOnly"))
	return rv
}


// A Boolean value that indicates whether the cookie should only be sent to HTTP servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/ishttponly
func (h_ HTTPCookie) SetIsHTTPOnly(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsHTTPOnly:"), value)
}


// A Boolean value that indicates whether the cookie may only be sent over secure channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issecure
func (h_ HTTPCookie) IsSecure() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("isSecure"))
	return rv
}


// A Boolean value that indicates whether the cookie may only be sent over secure channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issecure
func (h_ HTTPCookie) SetIsSecure(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsSecure:"), value)
}


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h_ HTTPCookie) IsSessionOnly() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](h_.ID, objc.Sel("isSessionOnly"))
	return rv
}


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h_ HTTPCookie) SetIsSessionOnly(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsSessionOnly:"), value)
}


