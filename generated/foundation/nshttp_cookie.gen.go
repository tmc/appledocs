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
}

// A representation of an HTTP cookie.
//
// An object is immutable, initialized from a dictionary that contains the attributes of the cookie. This class supports two different cookie versions: Version 0: The original cookie format defined by Netscape. Most cookies are in this format. Version 1: The cookie format defined in , HTTP State Management Mechanism.
//
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


// The cookie’s port list.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/portlist
func (h_ HTTPCookie) PortList() Number {
	rv := objc.Send[Number](h_.ID, objc.Sel("portList"))
	return rv
}


// SetPortList sets the value of the portList property.
// The cookie’s port list.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/portlist
func (h_ HTTPCookie) SetPortList(value Number) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPortList:"), value)
}

// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h_ HTTPCookie) IsSessionOnly() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isSessionOnly"))
	return rv
}


// SetIsSessionOnly sets the value of the isSessionOnly property.
// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h_ HTTPCookie) SetIsSessionOnly(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsSessionOnly:"), value)
}

// A Boolean value that indicates whether the cookie should only be sent to HTTP servers.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/ishttponly
func (h_ HTTPCookie) IsHTTPOnly() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isHTTPOnly"))
	return rv
}


// SetIsHTTPOnly sets the value of the isHTTPOnly property.
// A Boolean value that indicates whether the cookie should only be sent to HTTP servers.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/ishttponly
func (h_ HTTPCookie) SetIsHTTPOnly(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsHTTPOnly:"), value)
}

// A Boolean value that indicates whether the cookie may only be sent over secure channels.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issecure
func (h_ HTTPCookie) IsSecure() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isSecure"))
	return rv
}


// SetIsSecure sets the value of the isSecure property.
// A Boolean value that indicates whether the cookie may only be sent over secure channels.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issecure
func (h_ HTTPCookie) SetIsSecure(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsSecure:"), value)
}

// The domain of the cookie.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/domain
func (h_ HTTPCookie) Domain() string {
	rv := objc.Send[string](h_.ID, objc.Sel("domain"))
	return rv
}


// SetDomain sets the value of the domain property.
// The domain of the cookie.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/domain
func (h_ HTTPCookie) SetDomain(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDomain:"), objc.String(value))
}

// A Boolean value that indicates whether to restrict the cookie to requests sent back to the same site that created it.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/samesitepolicy
func (h_ HTTPCookie) SameSitePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("sameSitePolicy"))
	return rv
}


// SetSameSitePolicy sets the value of the sameSitePolicy property.
// A Boolean value that indicates whether to restrict the cookie to requests sent back to the same site that created it.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/samesitepolicy
func (h_ HTTPCookie) SetSameSitePolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSameSitePolicy:"), value)
}

// The cookie’s string value.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/value
func (h_ HTTPCookie) Value() string {
	rv := objc.Send[string](h_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
// The cookie’s string value.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/value
func (h_ HTTPCookie) SetValue(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setValue:"), objc.String(value))
}

// The cookie’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/name
func (h_ HTTPCookie) Name() string {
	rv := objc.Send[string](h_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The cookie’s name.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/name
func (h_ HTTPCookie) SetName(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setName:"), objc.String(value))
}

// The cookie’s expiration date.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/expiresdate
func (h_ HTTPCookie) ExpiresDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("expiresDate"))
	return rv
}


// SetExpiresDate sets the value of the expiresDate property.
// The cookie’s expiration date.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/expiresdate
func (h_ HTTPCookie) SetExpiresDate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setExpiresDate:"), value)
}

// The cookie’s path.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/path
func (h_ HTTPCookie) Path() string {
	rv := objc.Send[string](h_.ID, objc.Sel("path"))
	return rv
}


// SetPath sets the value of the path property.
// The cookie’s path.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/path
func (h_ HTTPCookie) SetPath(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPath:"), objc.String(value))
}

// The cookie’s comment string.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/comment
func (h_ HTTPCookie) Comment() string {
	rv := objc.Send[string](h_.ID, objc.Sel("comment"))
	return rv
}


// SetComment sets the value of the comment property.
// The cookie’s comment string.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/comment
func (h_ HTTPCookie) SetComment(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setComment:"), objc.String(value))
}

// The cookie’s version.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/version
func (h_ HTTPCookie) Version() int {
	rv := objc.Send[int](h_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
// The cookie’s version.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/version
func (h_ HTTPCookie) SetVersion(value int) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setVersion:"), value)
}

// The cookie’s comment URL.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/commenturl
func (h_ HTTPCookie) CommentURL() URL {
	rv := objc.Send[URL](h_.ID, objc.Sel("commentURL"))
	return rv
}


// SetCommentURL sets the value of the commentURL property.
// The cookie’s comment URL.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/commenturl
func (h_ HTTPCookie) SetCommentURL(value URL) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCommentURL:"), value)
}

// The cookie’s properties.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/properties
func (h_ HTTPCookie) Properties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("properties"))
	return rv
}



