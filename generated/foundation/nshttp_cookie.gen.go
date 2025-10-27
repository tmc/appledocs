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
	ExpiresDate() IDate
	Comment() IString
	SetComment(value IString)
	CommentURL() IURL
	SetCommentURL(value IURL)
	Domain() IString
	SetDomain(value IString)
	IsHTTPOnly() bool
	SetIsHTTPOnly(value bool)
	IsSecure() bool
	SetIsSecure(value bool)
	IsSessionOnly() bool
	SetIsSessionOnly(value bool)
	Name() IString
	SetName(value IString)
	Path() IString
	SetPath(value IString)
	PortList() INumber
	SetPortList(value INumber)
	Properties() objectivec.IObject
	SetProperties(value objectivec.IObject)
	SameSitePolicy() objectivec.IObject
	SetSameSitePolicy(value objectivec.IObject)
	Value() IString
	SetValue(value IString)
	Version() int
	SetVersion(value int)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (hc _HTTPCookieClass) Alloc() HTTPCookie {
	rv := objc.Send[HTTPCookie](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

























// The cookie’s expiration date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPCookie/expiresDate
func (h_ HTTPCookie) ExpiresDate() IDate {
	rv := objc.Send[Date](h_.ID, objc.Sel("expiresDate"))
	return rv
}


// The cookie’s comment string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/comment
func (h_ HTTPCookie) Comment() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("comment"))
	return rv
}


// The cookie’s comment string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/comment
func (h_ HTTPCookie) SetComment(value IString) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setComment:"), value)
}


// The cookie’s comment URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/commenturl
func (h_ HTTPCookie) CommentURL() IURL {
	rv := objc.Send[URL](h_.ID, objc.Sel("commentURL"))
	return rv
}


// The cookie’s comment URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/commenturl
func (h_ HTTPCookie) SetCommentURL(value IURL) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCommentURL:"), value)
}


// The domain of the cookie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/domain
func (h_ HTTPCookie) Domain() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("domain"))
	return rv
}


// The domain of the cookie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/domain
func (h_ HTTPCookie) SetDomain(value IString) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDomain:"), value)
}


// A Boolean value that indicates whether the cookie should only be sent to HTTP servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/ishttponly
func (h_ HTTPCookie) IsHTTPOnly() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isHTTPOnly"))
	return rv
}


// A Boolean value that indicates whether the cookie should only be sent to HTTP servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/ishttponly
func (h_ HTTPCookie) SetIsHTTPOnly(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsHTTPOnly:"), value)
}


// A Boolean value that indicates whether the cookie may only be sent over secure channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issecure
func (h_ HTTPCookie) IsSecure() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isSecure"))
	return rv
}


// A Boolean value that indicates whether the cookie may only be sent over secure channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issecure
func (h_ HTTPCookie) SetIsSecure(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsSecure:"), value)
}


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h_ HTTPCookie) IsSessionOnly() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isSessionOnly"))
	return rv
}


// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h_ HTTPCookie) SetIsSessionOnly(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIsSessionOnly:"), value)
}


// The cookie’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/name
func (h_ HTTPCookie) Name() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("name"))
	return rv
}


// The cookie’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/name
func (h_ HTTPCookie) SetName(value IString) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setName:"), value)
}


// The cookie’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/path
func (h_ HTTPCookie) Path() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("path"))
	return rv
}


// The cookie’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/path
func (h_ HTTPCookie) SetPath(value IString) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPath:"), value)
}


// The cookie’s port list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/portlist
func (h_ HTTPCookie) PortList() INumber {
	rv := objc.Send[Number](h_.ID, objc.Sel("portList"))
	return rv
}


// The cookie’s port list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/portlist
func (h_ HTTPCookie) SetPortList(value INumber) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPortList:"), value)
}


// The cookie’s properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/properties
func (h_ HTTPCookie) Properties() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](h_.ID, objc.Sel("properties"))
	return rv
}


// The cookie’s properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/properties
func (h_ HTTPCookie) SetProperties(value objectivec.IObject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setProperties:"), value)
}


// A Boolean value that indicates whether to restrict the cookie to requests sent back to the same site that created it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/samesitepolicy
func (h_ HTTPCookie) SameSitePolicy() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](h_.ID, objc.Sel("sameSitePolicy"))
	return rv
}


// A Boolean value that indicates whether to restrict the cookie to requests sent back to the same site that created it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/samesitepolicy
func (h_ HTTPCookie) SetSameSitePolicy(value objectivec.IObject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSameSitePolicy:"), value)
}


// The cookie’s string value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/value
func (h_ HTTPCookie) Value() IString {
	rv := objc.Send[String](h_.ID, objc.Sel("value"))
	return rv
}


// The cookie’s string value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/value
func (h_ HTTPCookie) SetValue(value IString) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setValue:"), value)
}


// The cookie’s version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/version
func (h_ HTTPCookie) Version() int {
	rv := objc.Send[int](h_.ID, objc.Sel("version"))
	return rv
}


// The cookie’s version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpcookie/version
func (h_ HTTPCookie) SetVersion(value int) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setVersion:"), value)
}








